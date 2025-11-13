#!/bin/bash

# 仅启动本地后端（不启动数据库/前端）
# 用法：./backend-local.sh [start|stop|restart|status] [可选: 配置文件路径，默认 config-dev.yaml]

set -e

# 颜色
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
BLUE='\033[0;34m'
NC='\033[0m'

log_info() { echo -e "${GREEN}[INFO]${NC} $1"; }
log_warn() { echo -e "${YELLOW}[WARN]${NC} $1"; }
log_error(){ echo -e "${RED}[ERROR]${NC} $1"; }
log_step() { echo -e "${BLUE}[STEP]${NC} $1"; }

PROJECT_ROOT=$(cd "$(dirname "$0")" && pwd)
WEB_DIR="$PROJECT_ROOT/web"
BACKEND_PID_FILE="$PROJECT_ROOT/tmp/backend.pid"
LOG_FILE="$PROJECT_ROOT/logs/backend-dev.log"
CONFIG_FILE="${2:-config-dev.yaml}"

ensure_dirs() {
  mkdir -p "$PROJECT_ROOT/tmp" "$PROJECT_ROOT/logs" \
           "$PROJECT_ROOT/cache/tiktoken" "$PROJECT_ROOT/cache/data_gym"
}

check_go() {
  if ! command -v go >/dev/null 2>&1; then
    log_error "未检测到 Go，请先安装 Go"
    exit 1
  fi
}

prepare_embed_assets() {
  # 确保 web/build 存在，优先尝试 yarn build，否则创建占位文件
  if [ ! -d "$WEB_DIR/build" ]; then
    log_step "准备前端嵌入资源 (web/build)..."
    if command -v yarn >/dev/null 2>&1; then
      pushd "$WEB_DIR" >/dev/null
      [ -d node_modules ] || { log_info "安装前端依赖..."; yarn install; }
      if ! yarn build; then
        log_warn "yarn build 失败，使用占位文件兜底"
        mkdir -p build
      fi
      popd >/dev/null
    fi
    mkdir -p "$WEB_DIR/build"
    # 兜底：最小 index.html 与 favicon
    [ -f "$WEB_DIR/build/index.html" ] || {
      if [ -f "$WEB_DIR/index.html" ]; then
        cp -f "$WEB_DIR/index.html" "$WEB_DIR/build/index.html"
      else
        cat > "$WEB_DIR/build/index.html" <<'EOF'
<!doctype html><html lang="zh-CN"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1"></head><body><div id="root"></div></body></html>
EOF
      fi
    }
    [ -f "$WEB_DIR/build/favicon.ico" ] || {
      [ -f "$WEB_DIR/public/favicon.ico" ] && cp -f "$WEB_DIR/public/favicon.ico" "$WEB_DIR/build/favicon.ico" || true
    }
  fi
}

generate_swagger() {
  log_step "生成 Swagger 文档 (docs/swagger)..."
  # 先确保 embed 目录存在，避免 go generate 遍历时告警
  prepare_embed_assets
  if ! (cd "$PROJECT_ROOT" && go generate ./...); then
    log_warn "go generate 执行失败，继续后续步骤（若缺少 docs/swagger 请手动排查）"
  fi
}

build_backend() {
  log_step "构建后端二进制..."
  (cd "$PROJECT_ROOT" && go build -o tmp/one-hub-dev main.go)
}

run_backend() {
  log_step "启动后端 (配置: $CONFIG_FILE)..."
  if [ ! -f "$PROJECT_ROOT/$CONFIG_FILE" ]; then
    log_error "配置文件不存在：$CONFIG_FILE"
    exit 1
  fi
  nohup "$PROJECT_ROOT/tmp/one-hub-dev" --config "$CONFIG_FILE" > "$LOG_FILE" 2>&1 &
  echo $! > "$BACKEND_PID_FILE"
  sleep 2
  if ps -p $(cat "$BACKEND_PID_FILE") >/dev/null 2>&1; then
    log_info "后端已启动 (PID: $(cat "$BACKEND_PID_FILE"))，日志：$LOG_FILE"
  else
    log_error "后端启动失败，请查看日志：$LOG_FILE"
    exit 1
  fi
}

start() {
  ensure_dirs
  check_go
  generate_swagger
  build_backend
  run_backend
}

stop() {
  if [ -f "$BACKEND_PID_FILE" ]; then
    local pid=$(cat "$BACKEND_PID_FILE")
    if ps -p $pid >/dev/null 2>&1; then
      log_step "停止后端 (PID: $pid)..."
      kill $pid || true
      rm -f "$BACKEND_PID_FILE"
      log_info "后端已停止"
    else
      log_warn "进程不存在，清理残留 PID 文件"
      rm -f "$BACKEND_PID_FILE"
    fi
  else
    log_warn "未发现后端 PID 文件"
  fi
}

status() {
  if [ -f "$BACKEND_PID_FILE" ] && ps -p $(cat "$BACKEND_PID_FILE") >/dev/null 2>&1; then
    echo -e "${GREEN}✓${NC} 后端运行中 (PID: $(cat "$BACKEND_PID_FILE"))"
  else
    echo -e "${RED}✗${NC} 后端未运行"
  fi
  echo "日志位置：$LOG_FILE"
}

case "${1:-start}" in
  start)
    start ;;
  stop)
    stop ;;
  restart)
    stop || true
    start ;;
  status)
    status ;;
  *)
    echo "用法：$0 [start|stop|restart|status] [配置文件，默认 config-dev.yaml]"; exit 1 ;;
esac

