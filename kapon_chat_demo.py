#!/usr/bin/env python3
"""
使用 openai 官方 SDK 调用 Kapon Chat Completions
运行前请先:
    pip install --upgrade openai
    export KAPON_API_KEY='sk-...'
"""

import os
from openai import OpenAI

def main() -> None:
    api_key = os.environ.get("KAPON_API_KEY")
    if not api_key:
        raise RuntimeError("请在环境变量 KAPON_API_KEY 中配置密钥")

    client = OpenAI(
        api_key=api_key,
        base_url="https://models.kapon.cloud/v1",
        default_headers={
            "User-Agent": "Mozilla/5.0 (compatible; KaponDemo/1.0)",
            "Accept": "application/json",
        },
    )

    response = client.chat.completions.create(
        model="gpt-4o",
        messages=[
            {"role": "system", "content": "你是一名专业的中文助手。"},
            {"role": "user", "content": "请用一句话欢迎 Kapon AI 的新用户。"},
        ],
        temperature=0.6,
        max_tokens=200,
    )

    print(response.choices[0].message.content.strip())

if __name__ == "__main__":
    main()
