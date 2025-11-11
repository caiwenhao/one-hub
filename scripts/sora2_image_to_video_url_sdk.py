#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
使用 openai Python SDK，基于参考图 URL（image_url）发起图生视频任务，轮询直至完成并下载成片。

环境变量（可覆盖命令行参数）:
- OPENAI_BASE_URL  默认 http://localhost:3000/v1
- OPENAI_API_KEY   必填
- SORA_PROMPT      默认 "百事可乐宣传片"
- SORA_MODEL       默认 "sora-2"
- SORA_SECONDS     默认 "10"
- SORA_SIZE        默认 "720x1280"
- SORA_IMAGE_URL   若未通过 --image-url 指定，将尝试从此处读取
- SORA_OUTPUT      默认 "sora_image2video.mp4"

用法示例:
  export OPENAI_BASE_URL="http://localhost:3000/v1"
  export OPENAI_API_KEY="sk-xxxx"
  python3 scripts/sora2_image_to_video_url_sdk.py \
    --image-url "https://example.com/ref.jpg" \
    --prompt "百事可乐宣传片" \
    --seconds 10 \
    --size 720x1280 \
    --output sora_image2video.mp4
"""

import argparse
import os
import sys
import time
from openai import OpenAI


def main() -> None:
    parser = argparse.ArgumentParser()
    parser.add_argument("--prompt", default=os.getenv("SORA_PROMPT", "百事可乐宣传片"))
    parser.add_argument("--model", default=os.getenv("SORA_MODEL", "sora-2"))
    parser.add_argument("--seconds", default=os.getenv("SORA_SECONDS", "10"))
    parser.add_argument("--size", default=os.getenv("SORA_SIZE", "720x1280"))
    parser.add_argument("--image-url", dest="image_url", default=os.getenv("SORA_IMAGE_URL"), help="参考图直链 URL (http/https/base64)")
    parser.add_argument("--output", default=os.getenv("SORA_OUTPUT", "sora_image2video.mp4"))
    args = parser.parse_args()

    base_url = os.getenv("OPENAI_BASE_URL", "http://localhost:3000/v1")
    api_key = os.getenv("OPENAI_API_KEY")
    if not api_key:
        print("请设置 OPENAI_API_KEY 环境变量。", file=sys.stderr)
        sys.exit(1)

    if not args.image_url:
        print("请通过 --image-url 或 SORA_IMAGE_URL 指定参考图 URL。", file=sys.stderr)
        sys.exit(2)

    client = OpenAI(base_url=base_url, api_key=api_key)

    # seconds 以字符串传递，兼容后端 ,string 解析
    sec = str(args.seconds).strip()
    if sec not in {"4", "8", "10", "12", "15", "25"}:
        print(f"提示：seconds={sec} 可能不被部分上游接受。", file=sys.stderr)

    print("🎬 提交图生视频任务 (URL)…")
    print(f"📝 prompt: {args.prompt}")
    print(f"🎯 model:  {args.model}")
    print(f"⏱️ seconds: {sec}")
    print(f"📐 size:    {args.size}")
    print(f"🖼️ image_url: {args.image_url}")

    # 1) 创建任务（使用 input_image 传 URL）
    try:
        # 使用 extra_body 传递未在 SDK 显式声明的字段（input_image）
        job = client.videos.create(
            prompt=args.prompt,
            model=args.model,
            seconds=sec,
            size=args.size,
            extra_body={"input_image": args.image_url},
        )
    except Exception as e:
        print(f"❌ 创建任务失败: {e}", file=sys.stderr)
        sys.exit(3)

    print(f"📦 任务已创建: {job.id} | 状态: {job.status}")

    # 2) 轮询
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

    # 3) 下载
    if job.status == "completed":
        try:
            print("📥 开始下载视频…")
            content = client.videos.download_content(job.id)
            content.write_to_file(args.output)
            print(f"✅ 已保存到: {args.output}")
        except Exception as e:
            print(f"❌ 下载失败: {e}", file=sys.stderr)
            sys.exit(4)
    else:
        err = getattr(job, "error", None)
        print(f"❌ 任务失败: {err}", file=sys.stderr)
        sys.exit(5)


if __name__ == "__main__":
    main()
