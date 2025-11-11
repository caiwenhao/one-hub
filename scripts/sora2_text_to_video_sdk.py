#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
使用官方 openai-python SDK 调用 /v1/videos：创建任务 + 轮询状态 + 下载视频

注意：
- 建议通过环境变量传入密钥/端点，避免把密钥写进代码：
  - OPENAI_API_KEY
  - OPENAI_BASE_URL（默认 https://models.kapon.cloud/v1）
- 对接本项目的 Sora 通道时，sora-2 常见合法秒数：10 或 15（按上游供应商要求）。
"""

import os
import sys
import time
from openai import OpenAI


def main():
    base_url = os.getenv("OPENAI_BASE_URL", "https://models.kapon.cloud/v1")
    api_key = os.getenv("OPENAI_API_KEY")
    if not api_key:
        print("请设置 OPENAI_API_KEY 环境变量。", file=sys.stderr)
        sys.exit(1)

    client = OpenAI(base_url=base_url, api_key=api_key)

    # 业务参数（可通过环境变量覆盖）
    prompt = os.getenv("SORA_PROMPT", "A calico cat playing a piano on stage")
    model = os.getenv("SORA_MODEL", "sora-2")
    # sora-2 常用 10 或 15；服务端 JSON 反序列化要求该字段为字符串
    seconds = os.getenv("SORA_SECONDS", "10")
    try:
        # 兼容传入数字，最终以字符串传给 SDK
        seconds = str(int(seconds))
    except Exception:
        # 若无法解析则直接按字符串透传
        pass
    size = os.getenv("SORA_SIZE", "720x1280")
    output = os.getenv("SORA_OUTPUT", "sora_video.mp4")

    print("🎬 提交生成任务…")
    print(f"📝 prompt: {prompt}")
    print(f"🎯 model:  {model}")
    print(f"⏱️ seconds: {seconds}")
    print(f"📐 size:    {size}")

    # 1) 创建任务
    try:
        job = client.videos.create(
            prompt=prompt,
            model=model,
            seconds=seconds,  # 以字符串形式传递，满足后端 `,string` 反序列化
            size=size,
        )
    except Exception as e:
        print(f"❌ 创建任务失败: {e}", file=sys.stderr)
        sys.exit(2)

    print(f"📦 任务已创建: {job.id} | 状态: {job.status}")

    # 2) 轮询状态（打印进度）
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
            content.write_to_file(output)
            print(f"✅ 已保存到: {output}")
        except Exception as e:
            print(f"❌ 下载失败: {e}", file=sys.stderr)
            sys.exit(3)
    else:
        err = getattr(job, "error", None)
        print(f"❌ 任务失败: {err}", file=sys.stderr)
        sys.exit(4)


if __name__ == "__main__":
    main()
