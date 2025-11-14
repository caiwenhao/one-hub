#!/usr/bin/env python3
# -*- coding: utf-8 -*-

"""
平台任务ID（task_<base36>）统一与日志导出增强 回归脚本

使用方式：
  环境变量：
    BASE_URL    服务地址，默认 http://localhost:3000
    USER_TOKEN  普通用户 Token（用于 /v1/videos 测试）
    ADMIN_TOKEN 管理员 Token（可选，用于导出 /api/log/export）

  运行：
    python3 scripts/test_platform_task_id.py

脚本做的事情：
  1) /v1/videos 创建 → 断言返回 id 以 task_ 开头
  2) /v1/videos/{platform_id} 查询 → 断言返回 id 等于 platform_id
  3) /api/log/self/export（用户）CSV → 断言包含“平台任务ID”列
  4) /api/log/export（管理员，若提供 ADMIN_TOKEN）CSV → 断言包含“平台任务ID/上游任务ID”列

备注：实际调通取决于本地已配置的视频上游渠道（OpenAI Sora/Veo 或 NewAPI）。若创建失败，将自动跳过 1/2 相关断言，仅验证导出表头。
"""

import csv
import io
import json
import os
import sys
import time
from typing import Optional

import requests


def norm_base_url() -> str:
    # 兼容 OPENAI_BASE_URL；若以 /v1 结尾则裁剪掉，保留服务根地址
    raw = os.getenv("BASE_URL") or os.getenv("OPENAI_BASE_URL") or "http://localhost:3000"
    raw = raw.strip().rstrip("/")
    if raw.lower().endswith("/v1"):
        raw = raw[:-3]
    return raw

BASE_URL = norm_base_url()
USER_TOKEN = (os.getenv("USER_TOKEN") or os.getenv("OPENAI_API_KEY") or "").strip()
ADMIN_TOKEN = os.getenv("ADMIN_TOKEN", "").strip()


def auth_headers(token: str) -> dict:
    return {"Authorization": f"Bearer {token}"} if token else {}


def jprint(title: str, data):
    print(f"\n== {title} ==")
    try:
        print(json.dumps(data, ensure_ascii=False, indent=2))
    except Exception:
        print(data)


def assert_true(cond: bool, msg: str):
    if cond:
        print(f"[PASS] {msg}")
    else:
        print(f"[FAIL] {msg}")
        raise AssertionError(msg)


def create_video() -> Optional[str]:
    if not USER_TOKEN:
        print("[SKIP] 未提供 USER_TOKEN，跳过 /v1/videos 创建用例")
        return None
    url = f"{BASE_URL}/v1/videos"
    body = {
        "model": "sora-2",  # 需本地有可用渠道
        "prompt": "codex-e2e smoke",
        "seconds": 4,
        "size": "1280x720",
    }
    try:
        resp = requests.post(url, headers={**auth_headers(USER_TOKEN), "Content-Type": "application/json"}, json=body, timeout=30)
    except Exception as e:
        print(f"[SKIP] 调用失败: {e}")
        return None
    if resp.status_code != 200:
        print(f"[SKIP] /v1/videos 返回 {resp.status_code}: {resp.text[:200]}")
        return None
    data = resp.json()
    jprint("/v1/videos 响应", data)
    vid = data.get("id") or data.get("video_id")
    assert_true(isinstance(vid, str) and vid.startswith("task_"), "创建返回的平台任务ID 以 task_ 开头")
    return vid


def retrieve_video(vid: str):
    url = f"{BASE_URL}/v1/videos/{vid}"
    resp = requests.get(url, headers=auth_headers(USER_TOKEN), timeout=30)
    if resp.status_code != 200:
        print(f"[WARN] /v1/videos/{{id}} 返回 {resp.status_code}: {resp.text[:200]}")
        return
    data = resp.json()
    jprint("/v1/videos/{id} 响应", data)
    assert_true(data.get("id") == vid, "查询返回的平台任务ID 与入参一致")


def export_user_logs_check_header():
    if not USER_TOKEN:
        print("[SKIP] 未提供 USER_TOKEN，跳过用户日志导出")
        return
    url = f"{BASE_URL}/api/log/self/export"
    resp = requests.get(url, headers=auth_headers(USER_TOKEN), timeout=60)
    assert_true(resp.status_code == 200 and resp.headers.get("Content-Type", "").startswith("text/csv"), "用户日志导出 CSV 可用")
    content = resp.content.decode("utf-8-sig", errors="ignore")
    reader = csv.reader(io.StringIO(content))
    header = next(reader, [])
    jprint("用户导出 CSV 表头", header)
    assert_true("平台任务ID" in header, "用户导出包含‘平台任务ID’列")


def export_admin_logs_check_header():
    if not ADMIN_TOKEN:
        print("[SKIP] 未提供 ADMIN_TOKEN，跳过管理员日志导出")
        return
    url = f"{BASE_URL}/api/log/export"
    resp = requests.get(url, headers=auth_headers(ADMIN_TOKEN), timeout=60)
    assert_true(resp.status_code == 200 and resp.headers.get("Content-Type", "").startswith("text/csv"), "管理员日志导出 CSV 可用")
    content = resp.content.decode("utf-8-sig", errors="ignore")
    reader = csv.reader(io.StringIO(content))
    header = next(reader, [])
    jprint("管理员导出 CSV 表头", header)
    assert_true("平台任务ID" in header and "上游任务ID" in header, "管理员导出包含‘平台任务ID/上游任务ID’列")


def main():
    print(f"BASE_URL={BASE_URL}")
    print(f"USER_TOKEN={'SET' if bool(USER_TOKEN) else 'EMPTY'}  ADMIN_TOKEN={'SET' if bool(ADMIN_TOKEN) else 'EMPTY'}")
    platform_id = create_video() or ''
    if platform_id:
        # 部分上游会有创建-可读的短暂延时
        time.sleep(1)
        retrieve_video(platform_id)
    export_user_logs_check_header()
    export_admin_logs_check_header()
    print("\nAll checks finished.")


if __name__ == "__main__":
    try:
        main()
    except AssertionError:
        sys.exit(2)
    except KeyboardInterrupt:
        sys.exit(130)
