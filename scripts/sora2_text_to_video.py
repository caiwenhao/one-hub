#!/usr/bin/env python3
# -*- coding: utf-8 -*-
"""
OpenAI Sora 2 项目对齐版 - Python 脚本
严格对齐本项目 /v1/videos 接口：multipart/form-data 提交，字段为 model/prompt/seconds/size；
创建后轮询 /v1/videos/{id}，依据 status=queued|in_progress|completed|failed，完成后读取 video_url。

依赖安装:
    pip install requests

使用方法:
    1. 设置环境变量 OPENAI_API_KEY，或在 generate_video() 传入 api_key
    2. 调用 generate_video() 创建任务，脚本会自动轮询直至完成
    3. 拿到视频直链后可调用 download_video() 下载

注意:
    - 本脚本严格对齐本项目的 OpenAI 兼容接口（非第三方代理 SDK 调用）。
    - 模型名建议使用：sora-2 或 sora-2-pro（以项目文档/后端限制为准）。
"""

import os
import time
import json
from typing import Optional

import requests


def _get_attr(obj, name, default=None):
    """安全读取属性/键，兼容 SDK 对象或字典。"""
    if obj is None:
        return default
    # 优先属性访问
    if hasattr(obj, name):
        try:
            return getattr(obj, name)
        except Exception:
            pass
    # 其次字典访问
    if isinstance(obj, dict):
        return obj.get(name, default)
    return default


def generate_video(
    prompt: str,
    api_key: Optional[str] = None,
    base_url: Optional[str] = None,
    model: str = "sora-2",
    seconds: int = 4,
    size: str = "720x1280",
    poll_interval_sec: float = 2.0,
) -> Optional[str]:
    """
    使用本项目标准接口创建 Sora 2 视频任务并轮询直至完成。

    参数:
        prompt: 文本提示词
        api_key: API Key（默认从 OPENAI_API_KEY 环境变量读取）
        model: 默认 'sora-2'
        seconds: 视频时长（秒），默认 4
        size: 分辨率字符串，如 '720x1280'
        poll_interval_sec: 轮询间隔秒数

    返回:
        成功: 视频直链 URL
        失败: None
    """
    key = api_key or os.getenv("OPENAI_API_KEY") or "sk-kQKMTKyEQA7X6eZ_X4xENvwOs5SmZiw2XT2sHeHMhOkz-NwEE4uIR0vdrMM"
    if not key:
        raise ValueError("请设置 OPENAI_API_KEY 环境变量或传入 api_key 参数")
    base = (
        base_url
        or os.getenv("OPENAI_BASE_URL")
        or os.getenv("BASE_URL")
        or "https://models.kapon.cloud/v1"
    )

    print("🎬 开始提交生成任务...")
    print(f"📝 提示词: {prompt}")
    print(f"🎯 模型: {model}")
    print(f"🌐 端点: {base}")
    print(f"⏱️  时长: {seconds}s")
    print(f"📐 分辨率: {size}")
    print()

    # 1) 创建视频生成任务（multipart/form-data）
    url_create = base.rstrip("/") + "/videos"
    headers = {"Authorization": f"Bearer {key}"}
    files = {
        "model": (None, model),
        "prompt": (None, prompt),
        "seconds": (None, str(int(seconds))),
        "size": (None, size),
    }
    try:
        resp = requests.post(url_create, headers=headers, files=files, timeout=60)
        if resp.status_code >= 400:
            print(f"❌ 创建任务失败: {resp.status_code} {resp.reason} -> {resp.url}")
            # 优先尝试解析 JSON 以便打印更清晰的错误
            try:
                err_json = resp.json()
                print("— 错误响应(JSON):")
                print(json.dumps(err_json, ensure_ascii=False, indent=2))
            except Exception:
                if resp.text:
                    print("— 错误响应(文本):")
                    print(resp.text)
            return None
        job = resp.json()
    except requests.exceptions.HTTPError as e:
        print(f"❌ 创建任务失败(HTTPError): {e}")
        if e.response is not None:
            print(f"— 状态: {e.response.status_code} {e.response.reason} -> {e.response.url}")
            try:
                print(json.dumps(e.response.json(), ensure_ascii=False, indent=2))
            except Exception:
                print(e.response.text)
        return None
    except Exception as e:
        print(f"❌ 创建任务失败(异常): {e}")
        return None

    job_id = _get_attr(job, "id")
    status = _get_attr(job, "status")
    if not job_id:
        print(f"❌ 创建响应缺少任务ID: {job}")
        return None
    print(f"📦 任务已创建: {job_id} (状态: {status})")

    # 2) 轮询任务状态直至完成/失败
    url_retrieve = base.rstrip("/") + f"/videos/{job_id}"
    last_status = None
    while True:
        try:
            r = requests.get(url_retrieve, headers=headers, timeout=60)
            if r.status_code >= 400:
                print(f"⚠️  轮询失败: {r.status_code} {r.reason} -> {r.url}")
                try:
                    print(json.dumps(r.json(), ensure_ascii=False, indent=2))
                except Exception:
                    if r.text:
                        print(r.text)
                time.sleep(max(1.0, poll_interval_sec))
                continue
            cur = r.json()
        except requests.exceptions.HTTPError as e:
            print(f"⚠️  轮询失败(HTTPError): {e}")
            if e.response is not None:
                print(f"— 状态: {e.response.status_code} {e.response.reason} -> {e.response.url}")
                try:
                    print(json.dumps(e.response.json(), ensure_ascii=False, indent=2))
                except Exception:
                    print(e.response.text)
            time.sleep(max(1.0, poll_interval_sec))
            continue
        except Exception as e:
            print(f"⚠️  轮询失败(异常)，稍后重试: {e}")
            time.sleep(max(1.0, poll_interval_sec))
            continue

        status = (_get_attr(cur, "status") or "").lower()
        progress = _get_attr(cur, "progress")
        if status != last_status or progress is not None:
            msg = f"⏳ 状态: {status or 'unknown'}"
            if progress is not None:
                try:
                    msg += f" | 进度: {int(progress)}%"
                except Exception:
                    msg += f" | 进度: {progress}"
            print(msg)
            last_status = status

        if status in ("completed",):
            video_url = _get_attr(cur, "video_url")
            if video_url:
                print("\n✅ 视频生成完成！")
                print(f"🔗 视频链接: {video_url}")
                return video_url
            print("\n⚠️ 完成但未获取到视频链接。")
            return None

        if status in ("failed", "error"):
            err = _get_attr(cur, "error")
            print(f"\n❌ 任务失败: {err}")
            return None

        time.sleep(max(1.0, poll_interval_sec))


def download_video(video_url: str, save_path: str = "./sora_video.mp4") -> bool:
    """
    下载视频到本地

    参数:
        video_url: 视频 URL 直链
        save_path: 保存路径

    返回:
        True/False
    """
    try:
        print("\n📥 开始下载视频...")
        with requests.get(video_url, stream=True, timeout=300) as resp:
            resp.raise_for_status()
            with open(save_path, "wb") as f:
                for chunk in resp.iter_content(chunk_size=8192):
                    if chunk:
                        f.write(chunk)
        print(f"✅ 视频已保存到: {save_path}")
        return True
    except Exception as e:
        print(f"❌ 下载失败: {e}")
        return False


# ==================== 使用示例 ====================

if __name__ == "__main__":
    # 默认指定端点与密钥（可通过环境变量覆盖）
    API_KEY = os.getenv("OPENAI_API_KEY", "sk-kQKMTKyEQA7X6eZ_X4xENvwOs5SmZiw2XT2sHeHMhOkz-NwEE4uIR0vdrMM")
    BASE_URL = os.getenv("OPENAI_BASE_URL", os.getenv("BASE_URL", "https://models.kapon.cloud/v1"))

    print("=" * 60)
    print("OpenAI Sora 2 文生视频示例")
    print("=" * 60)

    # 示例1: 基础使用（严格对齐本项目字段）
    prompt = "百事可乐宣传片"
    url = generate_video(
        prompt,
        api_key=API_KEY or None,
        base_url=BASE_URL,
        model="sora-2",
        seconds=10,
        size="720x1280",
    )

    if url:
        download_video(url, save_path="./cat_playing.mp4")

    # 其他示例可按需开启
    # prompt2 = "A futuristic city at sunset with flying cars"
    # url2 = generate_video(prompt2, api_key=API_KEY or None, base_url=BASE_URL, model="sora-2", seconds=8, size="1280x720")
    # if url2:
    #     download_video(url2, save_path="./futuristic_city.mp4")
