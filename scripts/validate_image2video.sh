#!/usr/bin/env bash
set -euo pipefail

# 验证脚本：带参考图生视频（/v1/videos -> input_reference）
# 依赖：curl、python3（用于解析 JSON）。
# 用法：
#   ONEHUB_API_KEY="sk-xxxx" ./scripts/validate_image2video.sh \
#     --prompt "百事可乐宣传片" \
#     --image "/path/to/ref.jpg" \
#     --seconds 4 \
#     --size 720x1280 \
#     --base "http://localhost:3000"
#
# 说明：
# - ezlinkai 严格要求 seconds ∈ {4,8,12} 且 multipart 必须包含该字段。
# - Authorization 头按 One Hub 习惯，直接使用 "sk-xxxx"（无需 Bearer 前缀）。

PROMPT="百事可乐宣传片"
IMAGE_PATH=""
# 注意：Bash/Zsh 中 SECONDS 是保留变量（表示脚本运行秒数），不能用作自定义参数名
# 这里使用 CLIP_SECONDS 作为视频时长参数名
CLIP_SECONDS=4
SIZE="720x1280"
BASE_URL="http://localhost:3000"
POLL_MAX=120   # 最多轮询次数（120*5s=10分钟）
POLL_SLEEP=5

while [[ $# -gt 0 ]]; do
  case "$1" in
    --prompt)
      PROMPT="$2"; shift 2;;
    --image)
      IMAGE_PATH="$2"; shift 2;;
    --seconds)
      CLIP_SECONDS="$2"; shift 2;;
    --size)
      SIZE="$2"; shift 2;;
    --base)
      BASE_URL="$2"; shift 2;;
    --poll-max)
      POLL_MAX="$2"; shift 2;;
    --poll-sleep)
      POLL_SLEEP="$2"; shift 2;;
    *)
      echo "未知参数: $1" >&2; exit 2;;
  esac
done

API_KEY="${ONEHUB_API_KEY:-${OPENAI_API_KEY:-}}"
if [[ -z "${API_KEY}" ]]; then
  echo "[ERR] 未设置 ONEHUB_API_KEY 或 OPENAI_API_KEY 环境变量（需形如 sk-xxxx）。" >&2
  exit 1
fi

# 校验 seconds（ezlinkai 仅允许 4/8/12）
case "${CLIP_SECONDS}" in
  4|8|12) ;;
  *) echo "[ERR] seconds=${CLIP_SECONDS} 不符合 ezlinkai 要求（仅允许 4/8/12）。" >&2; exit 1;;
esac

TMP_DIR="$(mktemp -d 2>/dev/null || mktemp -d -t onehub-validate)"
cleanup() { rm -rf "${TMP_DIR}" >/dev/null 2>&1 || true; }
trap cleanup EXIT

# 若未提供参考图，自动下载一张示例图（竖图）
if [[ -z "${IMAGE_PATH}" ]]; then
  IMAGE_PATH="${TMP_DIR}/ref.jpg"
  echo "[INFO] 未指定 --image，下载示例图..."
  curl -fsSL -o "${IMAGE_PATH}" "https://picsum.photos/seed/pepsi/720/1280" || {
    echo "[ERR] 下载示例参考图失败。" >&2; exit 1;
  }
fi

if [[ ! -f "${IMAGE_PATH}" ]]; then
  echo "[ERR] 参考图不存在: ${IMAGE_PATH}" >&2
  exit 1
fi

CREATE_OUT="${TMP_DIR}/create.json"
RETRIEVE_OUT="${TMP_DIR}/retrieve.json"
VIDEO_FILE="${TMP_DIR}/output.mp4"

echo "[STEP] 创建带参考图的视频任务..."
HTTP_CODE=$(curl -sS -X POST "${BASE_URL%/}/v1/videos" \
  -H "Authorization: ${API_KEY}" \
  -F "model=sora-2" \
  -F "prompt=${PROMPT}" \
  -F "seconds=${CLIP_SECONDS}" \
  -F "size=${SIZE}" \
  -F "input_reference=@${IMAGE_PATH}" \
  -o "${CREATE_OUT}" -w "%{http_code}")

if [[ -z "${HTTP_CODE}" ]]; then
  echo "[ERR] 创建请求失败（无 HTTP_CODE）。" >&2
  exit 1
fi

if (( HTTP_CODE >= 400 )); then
  echo "[ERR] 创建接口返回 ${HTTP_CODE}，响应体如下：" >&2
  cat "${CREATE_OUT}" 2>/dev/null || true
  exit 1
fi

echo "[INFO] 创建响应:"
cat "${CREATE_OUT}"
echo

VIDEO_ID="$(python3 -c "import json,sys;print(json.load(open(sys.argv[1])).get('id',''))" "${CREATE_OUT}")"

if [[ -z "${VIDEO_ID}" ]]; then
  echo "[ERR] 无法解析 video id。" >&2
  exit 1
fi

echo "[STEP] 轮询任务进度 video_id=${VIDEO_ID} ..."
ATTEMPT=0
STATUS=""
PROGRESS=""
VIDEO_URL=""
while (( ATTEMPT < POLL_MAX )); do
  ATTEMPT=$((ATTEMPT+1))
  set +e
  curl -fsS "${BASE_URL%/}/v1/videos/${VIDEO_ID}" -H "Authorization: ${API_KEY}" -o "${RETRIEVE_OUT}"
  RET_RES=$?
  set -e
  if [[ ${RET_RES} -ne 0 ]]; then
    echo "[WARN] 第${ATTEMPT}次查询失败，${POLL_SLEEP}s 后重试..."
    sleep "${POLL_SLEEP}"; continue
  fi

  STATUS="$(python3 -c "import json,sys;print(json.load(open(sys.argv[1])).get('status',''))" "${RETRIEVE_OUT}")"
  PROGRESS="$(python3 -c "import json,sys;print(json.load(open(sys.argv[1])).get('progress',''))" "${RETRIEVE_OUT}")"
  VIDEO_URL="$(python3 -c "import json,sys;print(json.load(open(sys.argv[1])).get('video_url',''))" "${RETRIEVE_OUT}")"

  echo "[INFO] 状态=${STATUS} 进度=${PROGRESS}%" 
  if [[ "${STATUS}" == "completed" ]]; then
    echo "[INFO] 生成完成。"
    break
  elif [[ "${STATUS}" == "failed" ]]; then
    echo "[ERR] 任务失败："
    cat "${RETRIEVE_OUT}" || true
    exit 2
  fi
  sleep "${POLL_SLEEP}"
done

if [[ "${STATUS}" != "completed" ]]; then
  echo "[ERR] 轮询超时或未完成。最新响应："
  cat "${RETRIEVE_OUT}" || true
  exit 3
fi

echo "[STEP] 下载视频内容..."
curl -fsSL "${BASE_URL%/}/v1/videos/${VIDEO_ID}/content" -H "Authorization: ${API_KEY}" -o "${VIDEO_FILE}"

echo "[DONE] 校验完成："
echo "- video_id: ${VIDEO_ID}"
echo "- status: ${STATUS}"
echo "- progress: ${PROGRESS}%"
if [[ -n "${VIDEO_URL}" ]]; then
  echo "- video_url: ${VIDEO_URL}"
fi
echo "- file: ${VIDEO_FILE} (大小: $(stat -f%z "${VIDEO_FILE}" 2>/dev/null || stat -c%s "${VIDEO_FILE}"))"
