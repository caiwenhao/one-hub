#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
基于 openai 官方 Python SDK 的“参考图生视频”验证脚本。

用法示例（本地 one-hub 网关）：
  export OPENAI_BASE_URL="http://localhost:3000/v1"
  export OPENAI_API_KEY="sk-xxxx"
  python3 scripts/sora2_image_to_video_sdk.py \
    --prompt "百事可乐宣传片" \
    --seconds 4 \
    --size 720x1280 \
    --image ./ref.jpg

说明：
- 当传入文件句柄（input_reference）时，SDK 会自动以 multipart/form-data 方式提交。
- ezlinkai 严格要求 seconds ∈ {4,8,12} 且必须出现。
- 若未提供 --image，将自动下载一张竖版示例图到临时目录。
"""

import argparse
import os
import sys
import tempfile
import time
from typing import Optional

import requests
from openai import OpenAI


def ensure_image(image_path: Optional[str]) -> str:
    if image_path and os.path.isfile(image_path):
        return image_path
    # 下载一张示例图（竖版）
    fd, tmp_path = tempfile.mkstemp(prefix="sora_ref_", suffix=".jpg")
    os.close(fd)
    url = "https://picsum.photos/seed/pepsi/720/1280"
    try:
        r = requests.get(url, timeout=20)
        r.raise_for_status()
        with open(tmp_path, "wb") as f:
            f.write(r.content)
        return tmp_path
    except Exception as e:
        print(f"下载示例参考图失败: {e}", file=sys.stderr)
        raise


def main():
    parser = argparse.ArgumentParser()
    parser.add_argument("--prompt", default=os.getenv("SORA_PROMPT", "百事可乐宣传片"))
    parser.add_argument("--model", default=os.getenv("SORA_MODEL", "sora-2"))
    parser.add_argument("--seconds", default=os.getenv("SORA_SECONDS", "4"))
    parser.add_argument("--size", default=os.getenv("SORA_SIZE", "720x1280"))
    parser.add_argument("--image", default=os.getenv("SORA_IMAGE"), help="本地参考图路径；用于官方/支持 multipart 的通道")
    parser.add_argument("--image-url", dest="image_url", default=os.getenv("SORA_IMAGE_URL"), help="参考图直链 URL；用于 apimart 等仅支持 URL 的通道")
    parser.add_argument("--output", default=os.getenv("SORA_OUTPUT", "sora_image2video.mp4"))
    args = parser.parse_args()

    base_url = os.getenv("OPENAI_BASE_URL", "http://localhost:3000/v1")
    api_key = os.getenv("OPENAI_API_KEY")
    if not api_key:
        print("请设置 OPENAI_API_KEY 环境变量。", file=sys.stderr)
        sys.exit(1)

    # SDK 客户端
    client = OpenAI(base_url=base_url, api_key=api_key)

    # 处理 seconds：保持字符串，避免 SDK/后端的 ,string 解析差异
    sec = str(args.seconds).strip()
    if sec not in {"4", "8", "12"}:
        print(f"警告：seconds={sec} 可能不被上游接受（建议 4/8/12）", file=sys.stderr)

    image_path = None
    if not args.image_url:
        image_path = ensure_image(args.image)

    print("🎬 提交图生视频任务…")
    print(f"📝 prompt: {args.prompt}")
    print(f"🎯 model:  {args.model}")
    print(f"⏱️ seconds: {sec}")
    print(f"📐 size:    {args.size}")
    if args.image_url:
        print(f"🖼️ image_url: {args.image_url}")
    else:
        print(f"🖼️ image:   {image_path}")

    # 1) 创建任务（携带 input_reference 文件句柄 -> multipart 提交）
    try:
        if args.image_url:
            # 适用于 apimart：通过 input_image / input_images 传 URL，由服务端适配为 image_urls
            job = client.videos.create(
                prompt=args.prompt,
                model=args.model,
                seconds=sec,
                size=args.size,
                input_image=args.image_url,
            )
        else:
            with open(image_path, "rb") as f:
                job = client.videos.create(
                    prompt=args.prompt,
                    model=args.model,
                    seconds=sec,   # 字符串形式
                    size=args.size,
                    input_reference=f,  # 关键：参考图文件
                )
    except Exception as e:
        print(f"❌ 创建任务失败: {e}", file=sys.stderr)
        sys.exit(2)

    print(f"📦 任务已创建: {job.id} | 状态: {job.status}")

    # 2) 轮询状态
    last_status = None
    while True:
        try:
            cur = client.videos.retrieve(job.id)
        except Exception as e:
            print(f"⚠️  轮询失败，重试中: {e}")
            time.sleep(2.0)
            continue

        status = cur.status
        progress = getattr(cur, "progress", None)
        if status != last_status or progress is not None:
            if progress is not None:
                print(f"⏳ 状态: {status} | 进度: {progress}%")
            else:
                print(f"⏳ 状态: {status}")
            last_status = status

        if status in ("completed", "failed"):
            job = cur
            break

        time.sleep(2.0)

    # 3) 下载视频内容
    if job.status == "completed":
        try:
            print("📥 开始下载视频…")
            content = client.videos.download_content(job.id)
            content.write_to_file(args.output)
            print(f"✅ 已保存到: {args.output}")
        except Exception as e:
            print(f"❌ 下载失败: {e}", file=sys.stderr)
            sys.exit(3)
    else:
        err = getattr(job, "error", None)
        print(f"❌ 任务失败: {err}", file=sys.stderr)
        sys.exit(4)


if __name__ == "__main__":
    main()
