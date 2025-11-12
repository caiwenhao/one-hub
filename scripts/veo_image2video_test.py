#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
使用 OpenAI 风格视频 API 测试 Veo 图生视频（veo-3.1-fast-generate-preview）：
- 创建任务：multipart 方式上传 1–3 张参考图（字段名统一为 input_reference，可重复）
- 轮询查询直到完成
- 下载视频到本地

说明
- Apimart/Sutui 通道会把多图作为“参考图”使用；Google 官方通道若恰为 2 张图，可能按首尾帧插值处理（具体以通道实际能力为准）。
- seconds 字段后端要求以“字符串”提交（,string），脚本已处理。

示例
python3 scripts/veo_image2video_test.py \
  --base-url http://localhost:3000/v1 \
  --api-key sk-xxx \
  --prompt "A silver sports car in neon rain" \
  --image ./ref1.jpg --image ./ref2.jpg \
  --seconds 6 --size 1280x720 --output out.mp4
"""

import argparse
import mimetypes
import os
import sys
import time
import requests
from typing import List

DEFAULT_MODEL = "veo-3.1-fast-generate-preview"
DEFAULT_SECONDS = 6
DEFAULT_SIZE = "1280x720"


def _auth_hdr(api_key: str) -> dict:
    return {"Authorization": f"Bearer {api_key}"}


def create_with_images(base_url: str, api_key: str, prompt: str, seconds: int, size: str, model: str, images: List[str]) -> str:
    if not images:
        raise ValueError("至少提供一张图片 --image")

    url = base_url.rstrip("/") + "/videos"

    # 文本字段（multipart）
    data = {
        "model": model,
        "prompt": prompt or "",
        "seconds": str(seconds),  # 后端使用 ,string 反序列化
        "size": size,
    }

    files = []  # [(fieldname, (filename, fileobj, mimetype))]
    extra_data = []  # 处理 http(s)/data: 这类“非文件”的参考图

    for it in images:
        it = it.strip()
        if it.lower().startswith("http://") or it.lower().startswith("https://") or it.lower().startswith("data:"):
            # 以纯文本形式追加一个 input_reference 字段
            extra_data.append(("input_reference", it))
            continue
        if not os.path.isfile(it):
            raise FileNotFoundError(f"图片不存在: {it}")
        mime, _ = mimetypes.guess_type(it)
        if not mime:
            mime = "image/jpeg"
        files.append(("input_reference", (os.path.basename(it), open(it, "rb"), mime)))

    # requests 同时支持 data=dict 与 files；若有重复字段名，使用 list of tuples 形式
    # 先把 data 转换为 list，再拼接 extra_data
    data_items = list(data.items()) + extra_data

    resp = requests.post(url, headers=_auth_hdr(api_key), data=data_items, files=files, timeout=60)
    # 关闭文件句柄
    for _, f in files:
        try:
            f[1].close()
        except Exception:
            pass
    if resp.status_code >= 400:
        raise RuntimeError(f"create failed: {resp.status_code} {resp.text}")
    data = resp.json()
    vid = data.get("id")
    if not vid:
        raise RuntimeError(f"create response missing id: {data}")
    print(f"[CREATE] id={vid} status={data.get('status')}")
    return vid


def retrieve(base_url: str, api_key: str, video_id: str) -> dict:
    url = base_url.rstrip("/") + f"/videos/{video_id}"
    resp = requests.get(url, headers=_auth_hdr(api_key), timeout=30)
    if resp.status_code >= 400:
        raise RuntimeError(f"retrieve failed: {resp.status_code} {resp.text}")
    return resp.json()


def wait_done(base_url: str, api_key: str, video_id: str, max_wait: int = 30 * 60, interval: int = 8) -> dict:
    start = time.time()
    last = None
    while True:
        if time.time() - start > max_wait:
            raise TimeoutError("wait timeout")
        try:
            cur = retrieve(base_url, api_key, video_id)
        except Exception as e:
            print(f"[POLL] error: {e}; sleep {interval}s")
            time.sleep(interval)
            continue
        st = (cur.get("status") or "").lower()
        pr = cur.get("progress")
        if st != last:
            print(f"[POLL] id={video_id} status={st} progress={pr}")
            last = st
        if st == "completed":
            vu = cur.get("video_url") or (cur.get("result") or {}).get("video_url")
            if vu:
                print(f"[POLL] video_url: {vu}")
            return cur
        if st in ("failed", "cancelled"):
            raise RuntimeError(f"task ended with status={st}, body={cur}")
        time.sleep(interval)


def download(base_url: str, api_key: str, video_id: str, outfile: str):
    url = base_url.rstrip("/") + f"/videos/{video_id}/content"
    with requests.get(url, headers=_auth_hdr(api_key), stream=True, timeout=60) as r:
        if r.status_code >= 400:
            raise RuntimeError(f"download failed: {r.status_code} {r.text}")
        with open(outfile, "wb") as f:
            for chunk in r.iter_content(chunk_size=8192):
                if chunk:
                    f.write(chunk)
    print(f"[DOWN] saved to {outfile}")


def main():
    p = argparse.ArgumentParser(description="Veo 图生视频测试（OpenAI 风格）")
    p.add_argument("--base-url", default=os.getenv("OPENAI_BASE_URL", "http://localhost:3000/v1"))
    p.add_argument("--api-key", default=os.getenv("OPENAI_API_KEY", ""))
    p.add_argument("--model", default=DEFAULT_MODEL)
    p.add_argument("--prompt", default="")
    p.add_argument("--seconds", type=int, default=DEFAULT_SECONDS, choices=[4, 6, 8])
    p.add_argument("--size", default=DEFAULT_SIZE, choices=["1280x720", "720x1280", "1920x1080", "1080x1920"])
    p.add_argument("--image", action="append", dest="images", help="参考图路径或 URL，可重复 1–3 次", required=True)
    p.add_argument("--output", default="output.mp4")
    args = p.parse_args()

    if not args.api_key:
        print("ERROR: 未提供 API Key（--api-key 或 OPENAI_API_KEY）", file=sys.stderr)
        sys.exit(2)

    print("=== Veo 图生视频（OpenAI 风格）===")
    print(f"Base URL : {args.base_url}")
    print(f"Model    : {args.model}")
    print(f"Prompt   : {args.prompt}")
    print(f"Seconds  : {args.seconds}")
    print(f"Size     : {args.size}")
    print(f"Images   : {args.images}")

    vid = create_with_images(args.base_url, args.api_key, args.prompt, args.seconds, args.size, args.model, args.images)
    _ = wait_done(args.base_url, args.api_key, vid)
    download(args.base_url, args.api_key, vid, args.output)
    print("=== 完成 ===")


if __name__ == "__main__":
    main()

