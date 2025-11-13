#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
MiniMax 语音（PPInfra/官方）验证脚本

功能
- 同步 TTS：
  - OpenAI 兼容：POST /v1/audio/speech
  - minimaxi 官方：POST /minimaxi/v1/t2a_v2
- 异步 TTS：
  - 创建：POST /minimaxi/v1/t2a_async_v2
  - 查询：GET  /minimaxi/v1/query/t2a_async_query_v2?task_id=...

说明
- 用于验证 one-api 中 minimaxi 渠道，且渠道上游可选择官方或 PPInfra。
- 自动根据返回的 Content-Type 判断 JSON/音频流；若为音频则保存到文件。
- 支持“简写格式”指定音频：mp3-声道-采样率-比特率，例如 mp3-1-32000-128000。

依赖
  pip install requests
"""

import argparse
import base64
import json
import os
import sys
import time
from pathlib import Path

import requests

DEFAULT_TEXT = "你好，欢迎使用 MiniMax 语音 2.6；This is a test for MiniMax speech 2.6."
DEFAULT_VOICE = "alloy"  # 在后端会映射到官方 voice_id


def parse_audio_format(fmt: str):
    """将 mp3-1-32000-128000 拆分为 (format, channel, sample_rate, bitrate)"""
    if not fmt:
        return None, None, None, None
    parts = fmt.split("-")
    kind = parts[0]
    ch = int(parts[1]) if len(parts) > 1 and parts[1].isdigit() else None
    sr = int(parts[2]) if len(parts) > 2 and parts[2].isdigit() else None
    br = int(parts[3]) if len(parts) > 3 and parts[3].isdigit() else None
    return kind, ch, sr, br


def save_audio(resp: requests.Response, out: Path):
    out.parent.mkdir(parents=True, exist_ok=True)
    with open(out, "wb") as f:
        for chunk in resp.iter_content(chunk_size=8192):
            if chunk:
                f.write(chunk)
    print(f"[OK] 已保存音频: {out} ({resp.headers.get('Content-Type')}, {resp.headers.get('Content-Length', '?')} bytes)")
    # 补充打印 minimaxi 扩展头
    for k in ["X-Minimax-Subtitle-URL", "X-Minimax-Audio-Channel"]:
        if k in resp.headers:
            print(f"  {k}: {resp.headers[k]}")


def do_sync_openai(base_url: str, api_key: str, model: str, text: str, voice: str, 
                    fmt: str, out: Path):
    url = base_url.rstrip("/") + "/v1/audio/speech"
    headers = {"Authorization": f"Bearer {api_key}", "Content-Type": "application/json"}
    payload = {
        "model": model,
        "input": text,
        "voice": voice,
    }
    if fmt:
        payload["response_format"] = fmt
    print("[SYNC][OpenAI] POST", url)
    r = requests.post(url, headers=headers, json=payload, stream=True, timeout=300)
    if r.status_code != 200:
        print("[ERR]", r.status_code, r.text)
        r.raise_for_status()
    ctype = r.headers.get("Content-Type", "")
    if ctype.startswith("audio/") or ctype == "application/octet-stream" or not ctype:
        save_audio(r, out)
    else:
        # 兼容返回 JSON（官方结构）
        try:
            data = r.json()
        except Exception:
            out.write_bytes(r.content)
            print(f"[WARN] Content-Type={ctype} 非 JSON，已按二进制保存: {out}")
            return
        # 官方结构一般为 data.audio(hex/base64)；我们在后端多为直接音频流，这里兜底解析
        if isinstance(data, dict) and "data" in data and isinstance(data["data"], dict):
            audio = data["data"].get("audio")
            if audio:
                try:
                    raw = bytes.fromhex(audio)
                except ValueError:
                    raw = base64.b64decode(audio)
                out.write_bytes(raw)
                print(f"[OK] 已从 JSON 中解析音频并保存: {out}")
                return
        print(json.dumps(data, ensure_ascii=False, indent=2))


def do_sync_official(base_url: str, api_key: str, model: str, text: str, voice: str,
                      fmt: str, out: Path, stream: bool):
    url = base_url.rstrip("/") + "/minimaxi/v1/t2a_v2"
    headers = {"Authorization": f"Bearer {api_key}", "Content-Type": "application/json"}
    fmt_kind, ch, sr, br = parse_audio_format(fmt) if fmt else (None, None, None, None)
    body = {
        "model": model,
        "text": text,
        "voice_setting": {"voice_id": voice},
    }
    if fmt_kind:
        body["audio_setting"] = {"format": fmt_kind}
        if ch: body["audio_setting"]["channel"] = ch
        if sr: body["audio_setting"]["sample_rate"] = sr
        if br: body["audio_setting"]["bitrate"] = br
    if stream:
        body["stream"] = True
    print("[SYNC][Official] POST", url)
    r = requests.post(url, headers=headers, json=body, stream=True, timeout=300)
    if r.status_code != 200:
        print("[ERR]", r.status_code, r.text)
        r.raise_for_status()
    ctype = r.headers.get("Content-Type", "")
    if ctype.startswith("audio/") or ctype == "application/octet-stream" or not ctype:
        # 明确的音频或未知类型：按音频保存
        save_audio(r, out)
    else:
        try:
            data = r.json()
        except Exception:
            # 声明非音频但也不是 JSON，直接落为二进制文件
            out.write_bytes(r.content)
            print(f"[WARN] Content-Type={ctype} 非 JSON，已按二进制保存: {out}")
            return
        # 兼容：官方 JSON（通常 data.audio 为十六进制/BASE64 字符串）
        audio = None
        if isinstance(data, dict) and "data" in data and isinstance(data["data"], dict):
            audio = data["data"].get("audio")
        if audio:
            try:
                raw = bytes.fromhex(audio)
            except ValueError:
                raw = base64.b64decode(audio)
            out.write_bytes(raw)
            print(f"[OK] 已从 JSON 中解析音频并保存: {out}")
        else:
            print(json.dumps(data, ensure_ascii=False, indent=2))


def do_async(base_url: str, api_key: str, model: str, text: str, voice: str,
             fmt: str, out: Path, poll_interval: float):
    create_url = base_url.rstrip("/") + "/minimaxi/v1/t2a_async_v2"
    query_url = base_url.rstrip("/") + "/minimaxi/v1/query/t2a_async_query_v2"
    headers = {"Authorization": f"Bearer {api_key}", "Content-Type": "application/json"}

    fmt_kind, ch, sr, br = parse_audio_format(fmt) if fmt else (None, None, None, None)
    body = {
        "model": model,
        "text": text,
        "voice_setting": {"voice_id": voice},
    }
    if fmt_kind:
        body["audio_setting"] = {"format": fmt_kind}
        if ch: body["audio_setting"]["channel"] = ch
        if sr: body["audio_setting"]["sample_rate"] = sr
        if br: body["audio_setting"]["bitrate"] = br

    print("[ASYNC][Create] POST", create_url)
    r = requests.post(create_url, headers=headers, json=body, timeout=60)
    if r.status_code != 200:
        print("[ERR]", r.status_code, r.text)
        r.raise_for_status()
    data = r.json()
    task_id = data.get("task_id") or data.get("TaskId") or data.get("taskId")
    if not task_id:
        print("[ERR] 无法获取 task_id:")
        print(json.dumps(data, ensure_ascii=False, indent=2))
        sys.exit(2)
    print("[ASYNC] task_id=", task_id)

    # 轮询查询
    while True:
        time.sleep(poll_interval)
        params = {"task_id": task_id}
        print("[ASYNC][Query] GET", query_url, params)
        q = requests.get(query_url, headers=headers, params=params, timeout=60)
        if q.status_code != 200:
            print("[ERR]", q.status_code, q.text)
            q.raise_for_status()
        qd = q.json()
        status = (
            qd.get("status") or qd.get("Status") or
            qd.get("data", {}).get("status")
        )
        if not status:
            print(json.dumps(qd, ensure_ascii=False, indent=2))
            continue
        print("[ASYNC] status=", status)
        if status in ("success", "failed"):
            # 尝试获取音频
            audio = None
            if "data" in qd and isinstance(qd["data"], dict):
                audio = qd["data"].get("audio")
            if audio:
                try:
                    raw = bytes.fromhex(audio)
                except ValueError:
                    raw = base64.b64decode(audio)
                out.write_bytes(raw)
                print(f"[OK] 已保存异步结果音频: {out}")
            else:
                print(json.dumps(qd, ensure_ascii=False, indent=2))
            break


def main():
    p = argparse.ArgumentParser(description="MiniMax TTS 验证脚本（one-api relay）")
    p.add_argument("--base-url", required=True, help="one-api 服务地址，例如 http://127.0.0.1:3000")
    p.add_argument("--api-key", required=True, help="令牌（Bearer）")
    p.add_argument("--model", default="speech-2.6-hd", help="模型：speech-2.6-hd|speech-2.6-turbo 等")
    p.add_argument("--text", default=DEFAULT_TEXT, help="合成文本，不填用默认测试文本")
    p.add_argument("--voice", default=DEFAULT_VOICE, help="音色别名或官方 voice_id，默认 alloy")
    p.add_argument("--format", dest="fmt", default="mp3-1-32000-128000", help="音频格式简写，例如 mp3-1-32000-128000；留空用上游默认")
    p.add_argument("--out", default="./tmp/minimax_tts.mp3", help="输出路径（音频或 JSON）")

    sub = p.add_subparsers(dest="cmd", required=True)

    s1 = sub.add_parser("sync", help="同步 TTS")
    s1.add_argument("--official", action="store_true", help="使用 minimaxi 官方路径 /minimaxi/v1/t2a_v2；默认 OpenAI 兼容 /v1/audio/speech")
    s1.add_argument("--stream", action="store_true", help="官方模式下请求流式（后端将透传或转换为音频流）")
    # 兼容把参数写在子命令后面的用法
    s1.add_argument("--out", help="输出路径（音频或 JSON）")
    s1.add_argument("--format", dest="fmt", help="音频格式简写，例如 mp3-1-32000-128000；留空用上游默认")

    s2 = sub.add_parser("async", help="异步 TTS（创建+轮询）")
    s2.add_argument("--poll", type=float, default=2.0, help="查询轮询间隔，秒")
    # 兼容把参数写在子命令后面的用法
    s2.add_argument("--out", help="输出路径（音频或 JSON）")
    s2.add_argument("--format", dest="fmt", help="音频格式简写，例如 mp3-1-32000-128000；留空用上游默认")

    args = p.parse_args()

    out = Path(args.out)
    out.parent.mkdir(parents=True, exist_ok=True)

    if args.cmd == "sync":
        if args.official:
            do_sync_official(args.base_url, args.api_key, args.model, args.text, args.voice, args.fmt, out, args.stream)
        else:
            do_sync_openai(args.base_url, args.api_key, args.model, args.text, args.voice, args.fmt, out)
    elif args.cmd == "async":
        do_async(args.base_url, args.api_key, args.model, args.text, args.voice, args.fmt, out, args.poll)


if __name__ == "__main__":
    main()
