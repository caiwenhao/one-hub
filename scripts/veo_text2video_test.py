#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
使用 OpenAI 风格视频 API 测试 Veo 文生视频（veo-3.1-fast-generate-preview）：
- 创建任务（POST /v1/videos）
- 轮询查询（GET /v1/videos/{id}）
- 下载内容（GET /v1/videos/{id}/content）

环境变量（可选）：
- OPENAI_BASE_URL：默认 https://models.kapon.cloud/v1 或传入 --base-url 覆盖
- OPENAI_API_KEY：默认从此读取或传入 --api-key 覆盖

示例：
python3 scripts/veo_text2video_test.py \
  --base-url https://models.kapon.cloud/v1 \
  --api-key sk-xxx \
  --prompt "A lion on the savannah at sunset, cinematic" \
  --seconds 6 \
  --size 1280x720 \
  --output out.mp4
"""

import argparse
import json
import os
import sys
import time
from typing import Optional

import requests

DEFAULT_MODEL = "veo-3.1-fast-generate-preview"
DEFAULT_SECONDS = 6
DEFAULT_SIZE = "1280x720"  # 720p；如选 1080p，请注意 Veo 官方仅支持 8 秒


def _headers(api_key: str) -> dict:
    return {
        "Authorization": f"Bearer {api_key}",
        "Content-Type": "application/json",
    }


essential_fields = ["id", "status"]


def create_video(base_url: str, api_key: str, prompt: str, seconds: int, size: str, model: str) -> str:
    url = base_url.rstrip("/") + "/videos"
    # 注意：后端 seconds 字段使用`,string`标签，必须以字符串传递
    payload = {
        "model": model,
        "prompt": prompt,
        "seconds": str(seconds),
        "size": size,
    }
    resp = requests.post(url, headers=_headers(api_key), data=json.dumps(payload), timeout=30)
    if resp.status_code >= 400:
        raise RuntimeError(f"create failed: {resp.status_code} {resp.text}")
    data = resp.json()
    for f in essential_fields:
        if f not in data:
            raise RuntimeError(f"create response missing field: {f}, body={data}")
    print(f"[CREATE] id={data['id']} status={data['status']}")
    return data["id"]


def retrieve_video(base_url: str, api_key: str, video_id: str) -> dict:
    url = base_url.rstrip("/") + f"/videos/{video_id}"
    resp = requests.get(url, headers={"Authorization": f"Bearer {api_key}"}, timeout=30)
    if resp.status_code >= 400:
        raise RuntimeError(f"retrieve failed: {resp.status_code} {resp.text}")
    return resp.json()


def download_video(base_url: str, api_key: str, video_id: str, outfile: str) -> None:
    url = base_url.rstrip("/") + f"/videos/{video_id}/content"
    with requests.get(url, headers={"Authorization": f"Bearer {api_key}"}, stream=True, timeout=60) as r:
        if r.status_code >= 400:
            raise RuntimeError(f"download failed: {r.status_code} {r.text}")
        r.raise_for_status()
        total = int(r.headers.get("Content-Length", 0))
        downloaded = 0
        with open(outfile, "wb") as f:
            for chunk in r.iter_content(chunk_size=8192):
                if not chunk:
                    continue
                f.write(chunk)
                downloaded += len(chunk)
                if total:
                    pct = downloaded * 100.0 / total
                    sys.stdout.write(f"\r[DOWN] {downloaded}/{total} bytes ({pct:.1f}%)")
                    sys.stdout.flush()
        sys.stdout.write("\n")
    print(f"[DOWN] saved to {outfile}")


def wait_until_complete(base_url: str, api_key: str, video_id: str, max_wait: int = 30 * 60, interval: int = 8) -> dict:
    """轮询直至完成或失败。max_wait 秒超时。"""
    start = time.time()
    last_status = None
    backoff = interval
    while True:
        if time.time() - start > max_wait:
            raise TimeoutError(f"wait timeout after {max_wait}s, last_status={last_status}")
        try:
            cur = retrieve_video(base_url, api_key, video_id)
        except Exception as e:
            print(f"[POLL] error: {e}; sleep {backoff}s")
            time.sleep(backoff)
            backoff = min(backoff * 2, 30)
            continue

        status = (cur.get("status") or "").lower()
        progress = cur.get("progress")
        if status != last_status:
            print(f"[POLL] id={video_id} status={status} progress={progress}")
            last_status = status

        if status in ("completed",):
            # 顶层 video_url 为主，兼容历史 result.video_url
            video_url = cur.get("video_url") or (cur.get("result") or {}).get("video_url")
            if video_url:
                print(f"[POLL] video_url: {video_url}")
            return cur
        if status in ("failed", "cancelled"):
            raise RuntimeError(f"task ended with status={status}, body={cur}")
        time.sleep(interval)


def main():
    parser = argparse.ArgumentParser(description="Veo 文生视频测试（OpenAI 风格视频 API）")
    parser.add_argument("--base-url", default=os.getenv("OPENAI_BASE_URL", "https://models.kapon.cloud/v1"), help="网关 Base URL（结尾不要带 /v1/videos）")
    parser.add_argument("--api-key", default=os.getenv("OPENAI_API_KEY", ""), help="API Key（Bearer）")
    parser.add_argument("--model", default=DEFAULT_MODEL, help="模型名，默认 veo-3.1-fast-generate-preview")
    parser.add_argument("--prompt", required=True, help="文本提示词")
    parser.add_argument("--seconds", type=int, default=DEFAULT_SECONDS, choices=[4, 6, 8], help="视频时长（秒）。1080p 场景仅支持 8 秒")
    parser.add_argument("--size", default=DEFAULT_SIZE, choices=["1280x720", "720x1280", "1920x1080", "1080x1920"], help="分辨率")
    parser.add_argument("--output", default="output.mp4", help="输出文件名")
    parser.add_argument("--timeout", type=int, default=30 * 60, help="最大等待时长（秒），默认 30 分钟")

    args = parser.parse_args()

    if not args.api_key:
        print("ERROR: 未提供 API Key（--api-key 或环境变量 OPENAI_API_KEY）", file=sys.stderr)
        sys.exit(2)

    print("=== Veo 文生视频（OpenAI 风格）===")
    print(f"Base URL : {args.base_url}")
    print(f"Model    : {args.model}")
    print(f"Prompt   : {args.prompt}")
    print(f"Seconds  : {args.seconds}")
    print(f"Size     : {args.size}")

    try:
        vid = create_video(args.base_url, args.api_key, args.prompt, args.seconds, args.size, args.model)
        cur = wait_until_complete(args.base_url, args.api_key, vid, max_wait=args.timeout)
        # 下载
        download_video(args.base_url, args.api_key, vid, args.output)
        print("=== 完成 ===")
    except Exception as e:
        print(f"FAILED: {e}", file=sys.stderr)
        sys.exit(1)


if __name__ == "__main__":
    main()
