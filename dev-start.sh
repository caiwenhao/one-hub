#!/bin/bash

# One Hub 开发环境启动脚本
# 使用方法: ./dev-start.sh [start|stop|restart|status]

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# 项目根目录
PROJECT_ROOT=$(cd "$(dirname "$0")" && pwd)
WEB_DIR="$PROJECT_ROOT/web"

# PID 文件
BACKEND_PID_FILE="$PROJECT_ROOT/tmp/backend.pid"
FRONTEND_PID_FILE="$PROJECT_ROOT/tmp/frontend.pid"

# 创建 tmp 目录
mkdir -p "$PROJECT_ROOT/tmp"
mkdir -p "$PROJECT_ROOT/logs"
mkdir -p "$PROJECT_ROOT/cache/tiktoken"
mkdir -p "$PROJECT_ROOT/cache/data_gym"

# 日志函数
log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

log_step() {
    echo -e "${BLUE}[STEP]${NC} $1"
}

# 检查依赖
check_dependencies() {
    log_step "检查依赖..."
    
    # 检查 Docker 和 Docker Compose
    if ! command -v docker &> /dev/null; then
        log_error "Docker 未安装，请先安装 Docker"
        exit 1
    fi
    
    if ! command -v docker-compose &> /dev/null && ! docker compose version &> /dev/null; then
        log_error "Docker Compose 未安装，请先安装 Docker Compose"
        exit 1
    fi
    
    # 检查 Go
    if ! command -v go &> /dev/null; then
        log_error "Go 未安装，请先安装 Go"
        exit 1
    fi
    
    # 检查 Node.js 和 Yarn
    if ! command -v node &> /dev/null; then
        log_error "Node.js 未安装，请先安装 Node.js"
        exit 1
    fi
    
    if ! command -v yarn &> /dev/null; then
        log_error "Yarn 未安装，请先安装 Yarn"
        exit 1
    fi
    
    log_info "依赖检查完成"
}

# 启动数据库服务
start_database() {
    log_step "启动数据库服务 (MySQL + Redis)..."
    
    cd "$PROJECT_ROOT"
    
    # 使用 docker compose 或 docker-compose
    if docker compose version &> /dev/null; then
        docker compose -f docker-compose-dev.yml up -d
    else
        docker-compose -f docker-compose-dev.yml up -d
    fi
    
    log_info "等待数据库服务启动..."
    sleep 5
    
    # 检查 MySQL 连接
    local retry_count=0
    local max_retries=30
    
    while [ $retry_count -lt $max_retries ]; do
        if docker exec mysql-dev mysqladmin ping -h localhost --silent; then
            log_info "MySQL 服务已就绪"
            break
        fi
        
        retry_count=$((retry_count + 1))
        log_warn "等待 MySQL 启动... ($retry_count/$max_retries)"
        sleep 2
    done
    
    if [ $retry_count -eq $max_retries ]; then
        log_error "MySQL 启动超时"
        exit 1
    fi
    
    log_info "数据库服务启动完成"
}

# 停止数据库服务
stop_database() {
    log_step "停止数据库服务..."
    
    cd "$PROJECT_ROOT"
    
    if docker compose version &> /dev/null; then
        docker compose -f docker-compose-dev.yml down
    else
        docker-compose -f docker-compose-dev.yml down
    fi
    
    log_info "数据库服务已停止"
}

# 启动后端服务
start_backend() {
    log_step "启动后端服务..."

    cd "$PROJECT_ROOT"

    # 检查配置文件
    if [ ! -f "config-dev.yaml" ]; then
        log_error "开发配置文件 config-dev.yaml 不存在"
        exit 1
    fi

    # 生成 Swagger 文档（docs/swagger）
    log_info "生成 Swagger 文档..."
    if ! go generate ./...; then
        log_warn "go generate 执行失败，请检查生成工具或网络后重试"
    fi

    # 确保嵌入资源存在（web/build）以满足 go:embed
    if [ ! -d "web/build" ]; then
        log_info "未检测到 web/build，开始构建前端静态资源用于嵌入..."
        pushd "$WEB_DIR" >/dev/null
        if [ ! -d "node_modules" ]; then
            log_info "安装前端依赖..."
            yarn install
        fi
        if ! yarn build; then
            log_warn "前端构建失败，尝试创建最小占位文件以通过编译..."
            mkdir -p build
            # 创建最小 index.html 与 favicon 以满足 go:embed
            [ -f index.html ] && cp index.html build/index.html || echo "<html><head><meta charset=\"utf-8\"></head><body><div id=\"root\"></div></body></html>" > build/index.html
            [ -f public/favicon.ico ] && cp public/favicon.ico build/favicon.ico || true
        fi
        popd >/dev/null
    fi

    # 构建并启动后端
    log_info "构建后端服务..."
    go build -o tmp/one-hub-dev main.go
    
    log_info "启动后端服务 (端口: 3000)..."
    nohup ./tmp/one-hub-dev --config config-dev.yaml > logs/backend-dev.log 2>&1 &
    echo $! > "$BACKEND_PID_FILE"
    
    # 等待后端启动
    sleep 3
    
    if ps -p $(cat "$BACKEND_PID_FILE") > /dev/null; then
        log_info "后端服务启动成功 (PID: $(cat "$BACKEND_PID_FILE"))"
    else
        log_error "后端服务启动失败，请检查日志: logs/backend-dev.log"
        exit 1
    fi
}

# 启动前端服务
start_frontend() {
    log_step "启动前端服务..."
    
    cd "$WEB_DIR"
    
    # 检查 node_modules
    if [ ! -d "node_modules" ]; then
        log_info "安装前端依赖..."
        yarn install
    fi
    
    log_info "启动前端开发服务器 (端口: 3010)..."
    nohup yarn dev > ../logs/frontend-dev.log 2>&1 &
    echo $! > "$FRONTEND_PID_FILE"
    
    # 等待前端启动
    sleep 5
    
    if ps -p $(cat "$FRONTEND_PID_FILE") > /dev/null; then
        log_info "前端服务启动成功 (PID: $(cat "$FRONTEND_PID_FILE"))"
    else
        log_error "前端服务启动失败，请检查日志: logs/frontend-dev.log"
        exit 1
    fi
}

# 停止后端服务
stop_backend() {
    if [ -f "$BACKEND_PID_FILE" ]; then
        local pid=$(cat "$BACKEND_PID_FILE")
        if ps -p $pid > /dev/null; then
            log_step "停止后端服务 (PID: $pid)..."
            kill $pid
            rm -f "$BACKEND_PID_FILE"
            log_info "后端服务已停止"
        else
            log_warn "后端服务进程不存在"
            rm -f "$BACKEND_PID_FILE"
        fi
    else
        log_warn "后端服务 PID 文件不存在"
    fi
}

# 停止前端服务
stop_frontend() {
    if [ -f "$FRONTEND_PID_FILE" ]; then
        local pid=$(cat "$FRONTEND_PID_FILE")
        if ps -p $pid > /dev/null; then
            log_step "停止前端服务 (PID: $pid)..."
            kill $pid
            rm -f "$FRONTEND_PID_FILE"
            log_info "前端服务已停止"
        else
            log_warn "前端服务进程不存在"
            rm -f "$FRONTEND_PID_FILE"
        fi
    else
        log_warn "前端服务 PID 文件不存在"
    fi
}

# 显示服务状态
show_status() {
    log_step "检查服务状态..."
    
    echo ""
    echo "=== 数据库服务状态 ==="
    if docker ps | grep -q "mysql-dev"; then
        echo -e "${GREEN}✓${NC} MySQL: 运行中 (端口: 3306)"
    else
        echo -e "${RED}✗${NC} MySQL: 未运行"
    fi
    
    if docker ps | grep -q "redis-dev"; then
        echo -e "${GREEN}✓${NC} Redis: 运行中 (端口: 6379)"
    else
        echo -e "${RED}✗${NC} Redis: 未运行"
    fi
    
    echo ""
    echo "=== 应用服务状态 ==="
    
    if [ -f "$BACKEND_PID_FILE" ] && ps -p $(cat "$BACKEND_PID_FILE") > /dev/null; then
        echo -e "${GREEN}✓${NC} 后端服务: 运行中 (PID: $(cat "$BACKEND_PID_FILE"), 端口: 3000)"
    else
        echo -e "${RED}✗${NC} 后端服务: 未运行"
    fi
    
    if [ -f "$FRONTEND_PID_FILE" ] && ps -p $(cat "$FRONTEND_PID_FILE") > /dev/null; then
        echo -e "${GREEN}✓${NC} 前端服务: 运行中 (PID: $(cat "$FRONTEND_PID_FILE"), 端口: 3010)"
    else
        echo -e "${RED}✗${NC} 前端服务: 未运行"
    fi
    
    echo ""
    echo "=== 访问地址 ==="
    echo "🌐 开发环境: http://localhost:3010"
    echo "🔧 后端 API: http://localhost:3000"
    echo ""
}

# 主函数
main() {
    case "${1:-start}" in
        "start")
            log_info "启动 One Hub 开发环境..."
            check_dependencies
            start_database
            start_backend
            start_frontend
            echo ""
            log_info "🎉 开发环境启动完成!"
            show_status
            ;;
        "stop")
            log_info "停止 One Hub 开发环境..."
            stop_frontend
            stop_backend
            stop_database
            log_info "开发环境已停止"
            ;;
        "restart")
            log_info "重启 One Hub 开发环境..."
            $0 stop
            sleep 2
            $0 start
            ;;
        "status")
            show_status
            ;;
        *)
            echo "使用方法: $0 [start|stop|restart|status]"
            echo ""
            echo "命令说明:"
            echo "  start   - 启动完整开发环境"
            echo "  stop    - 停止所有服务"
            echo "  restart - 重启所有服务"
            echo "  status  - 查看服务状态"
            exit 1
            ;;
    esac
}

# 执行主函数
main "$@"
