@echo off
chcp 65001 >nul
setlocal enabledelayedexpansion

REM One Hub 开发环境启动脚本 (Windows)
REM 使用方法: dev-start.bat [start|stop|restart|status]

set "PROJECT_ROOT=%~dp0"
set "WEB_DIR=%PROJECT_ROOT%web"
set "BACKEND_PID_FILE=%PROJECT_ROOT%tmp\backend.pid"
set "FRONTEND_PID_FILE=%PROJECT_ROOT%tmp\frontend.pid"

REM 创建必要目录
if not exist "%PROJECT_ROOT%tmp" mkdir "%PROJECT_ROOT%tmp"
if not exist "%PROJECT_ROOT%logs" mkdir "%PROJECT_ROOT%logs"
if not exist "%PROJECT_ROOT%cache\tiktoken" mkdir "%PROJECT_ROOT%cache\tiktoken"
if not exist "%PROJECT_ROOT%cache\data_gym" mkdir "%PROJECT_ROOT%cache\data_gym"

REM 颜色定义 (Windows 10+)
set "GREEN=[92m"
set "RED=[91m"
set "YELLOW=[93m"
set "BLUE=[94m"
set "NC=[0m"

goto main

:log_info
echo %GREEN%[INFO]%NC% %~1
goto :eof

:log_warn
echo %YELLOW%[WARN]%NC% %~1
goto :eof

:log_error
echo %RED%[ERROR]%NC% %~1
goto :eof

:log_step
echo %BLUE%[STEP]%NC% %~1
goto :eof

:check_dependencies
call :log_step "检查依赖..."

REM 检查 Docker
docker --version >nul 2>&1
if errorlevel 1 (
    call :log_error "Docker 未安装，请先安装 Docker Desktop"
    exit /b 1
)

REM 检查 Docker Compose
docker compose version >nul 2>&1
if errorlevel 1 (
    docker-compose --version >nul 2>&1
    if errorlevel 1 (
        call :log_error "Docker Compose 未安装"
        exit /b 1
    )
)

REM 检查 Go
go version >nul 2>&1
if errorlevel 1 (
    call :log_error "Go 未安装，请先安装 Go"
    exit /b 1
)

REM 检查 Node.js
node --version >nul 2>&1
if errorlevel 1 (
    call :log_error "Node.js 未安装，请先安装 Node.js"
    exit /b 1
)

REM 检查 Yarn
yarn --version >nul 2>&1
if errorlevel 1 (
    call :log_error "Yarn 未安装，请先安装 Yarn"
    exit /b 1
)

call :log_info "依赖检查完成"
goto :eof

:start_database
call :log_step "启动数据库服务 (MySQL + Redis)..."

cd /d "%PROJECT_ROOT%"

REM 使用 docker compose 或 docker-compose
docker compose -f docker-compose-dev.yml up -d >nul 2>&1
if errorlevel 1 (
    docker-compose -f docker-compose-dev.yml up -d >nul 2>&1
    if errorlevel 1 (
        call :log_error "启动数据库服务失败"
        exit /b 1
    )
)

call :log_info "等待数据库服务启动..."
timeout /t 5 /nobreak >nul

REM 检查 MySQL 连接
set /a retry_count=0
set /a max_retries=30

:mysql_check_loop
docker exec mysql-dev mysqladmin ping -h localhost --silent >nul 2>&1
if not errorlevel 1 (
    call :log_info "MySQL 服务已就绪"
    goto mysql_ready
)

set /a retry_count+=1
if !retry_count! geq !max_retries! (
    call :log_error "MySQL 启动超时"
    exit /b 1
)

call :log_warn "等待 MySQL 启动... (!retry_count!/!max_retries!)"
timeout /t 2 /nobreak >nul
goto mysql_check_loop

:mysql_ready
call :log_info "数据库服务启动完成"
goto :eof

:stop_database
call :log_step "停止数据库服务..."

cd /d "%PROJECT_ROOT%"

docker compose -f docker-compose-dev.yml down >nul 2>&1
if errorlevel 1 (
    docker-compose -f docker-compose-dev.yml down >nul 2>&1
)

call :log_info "数据库服务已停止"
goto :eof

:start_backend
call :log_step "启动后端服务..."

cd /d "%PROJECT_ROOT%"

if not exist "config-dev.yaml" (
    call :log_error "开发配置文件 config-dev.yaml 不存在"
    exit /b 1
)

call :log_info "构建后端服务..."
go build -o tmp\one-hub-dev.exe main.go
if errorlevel 1 (
    call :log_error "后端构建失败"
    exit /b 1
)

call :log_info "启动后端服务 (端口: 3000)..."
start /b "" tmp\one-hub-dev.exe --config config-dev.yaml > logs\backend-dev.log 2>&1

REM 获取进程 ID (Windows 方式)
timeout /t 2 /nobreak >nul
for /f "tokens=2" %%i in ('tasklist /fi "imagename eq one-hub-dev.exe" /fo csv ^| find "one-hub-dev.exe"') do (
    echo %%~i > "%BACKEND_PID_FILE%"
)

if exist "%BACKEND_PID_FILE%" (
    set /p backend_pid=<"%BACKEND_PID_FILE%"
    call :log_info "后端服务启动成功 (PID: !backend_pid!)"
) else (
    call :log_error "后端服务启动失败，请检查日志: logs\backend-dev.log"
    exit /b 1
)
goto :eof

:start_frontend
call :log_step "启动前端服务..."

cd /d "%WEB_DIR%"

if not exist "node_modules" (
    call :log_info "安装前端依赖..."
    yarn install
    if errorlevel 1 (
        call :log_error "前端依赖安装失败"
        exit /b 1
    )
)

call :log_info "启动前端开发服务器 (端口: 3010)..."
start /b "" yarn dev > ..\logs\frontend-dev.log 2>&1

timeout /t 3 /nobreak >nul

REM 检查前端进程
for /f "tokens=2" %%i in ('tasklist /fi "imagename eq node.exe" /fo csv ^| find "node.exe"') do (
    echo %%~i > "%FRONTEND_PID_FILE%"
    goto frontend_started
)

:frontend_started
if exist "%FRONTEND_PID_FILE%" (
    set /p frontend_pid=<"%FRONTEND_PID_FILE%"
    call :log_info "前端服务启动成功"
) else (
    call :log_error "前端服务启动失败，请检查日志: logs\frontend-dev.log"
    exit /b 1
)
goto :eof

:stop_backend
if exist "%BACKEND_PID_FILE%" (
    set /p pid=<"%BACKEND_PID_FILE%"
    call :log_step "停止后端服务 (PID: !pid!)..."
    taskkill /f /im one-hub-dev.exe >nul 2>&1
    del "%BACKEND_PID_FILE%" >nul 2>&1
    call :log_info "后端服务已停止"
) else (
    call :log_warn "后端服务 PID 文件不存在"
)
goto :eof

:stop_frontend
call :log_step "停止前端服务..."
taskkill /f /fi "windowtitle eq*yarn*" >nul 2>&1
taskkill /f /fi "windowtitle eq*vite*" >nul 2>&1
if exist "%FRONTEND_PID_FILE%" del "%FRONTEND_PID_FILE%" >nul 2>&1
call :log_info "前端服务已停止"
goto :eof

:show_status
call :log_step "检查服务状态..."

echo.
echo === 数据库服务状态 ===
docker ps | find "mysql-dev" >nul 2>&1
if not errorlevel 1 (
    echo %GREEN%✓%NC% MySQL: 运行中 (端口: 3306)
) else (
    echo %RED%✗%NC% MySQL: 未运行
)

docker ps | find "redis-dev" >nul 2>&1
if not errorlevel 1 (
    echo %GREEN%✓%NC% Redis: 运行中 (端口: 6379)
) else (
    echo %RED%✗%NC% Redis: 未运行
)

echo.
echo === 应用服务状态 ===

tasklist /fi "imagename eq one-hub-dev.exe" | find "one-hub-dev.exe" >nul 2>&1
if not errorlevel 1 (
    echo %GREEN%✓%NC% 后端服务: 运行中 (端口: 3000)
) else (
    echo %RED%✗%NC% 后端服务: 未运行
)

tasklist /fi "imagename eq node.exe" | find "node.exe" >nul 2>&1
if not errorlevel 1 (
    echo %GREEN%✓%NC% 前端服务: 运行中 (端口: 3010)
) else (
    echo %RED%✗%NC% 前端服务: 未运行
)

echo.
echo === 访问地址 ===
echo 🌐 开发环境: http://localhost:3010
echo 🔧 后端 API: http://localhost:3000
echo.
goto :eof

:main
set "action=%~1"
if "%action%"=="" set "action=start"

if "%action%"=="start" (
    call :log_info "启动 One Hub 开发环境..."
    call :check_dependencies
    if errorlevel 1 exit /b 1
    call :start_database
    if errorlevel 1 exit /b 1
    call :start_backend
    if errorlevel 1 exit /b 1
    call :start_frontend
    if errorlevel 1 exit /b 1
    echo.
    call :log_info "🎉 开发环境启动完成!"
    call :show_status
) else if "%action%"=="stop" (
    call :log_info "停止 One Hub 开发环境..."
    call :stop_frontend
    call :stop_backend
    call :stop_database
    call :log_info "开发环境已停止"
) else if "%action%"=="restart" (
    call :log_info "重启 One Hub 开发环境..."
    call "%~f0" stop
    timeout /t 2 /nobreak >nul
    call "%~f0" start
) else if "%action%"=="status" (
    call :show_status
) else (
    echo 使用方法: %~nx0 [start^|stop^|restart^|status]
    echo.
    echo 命令说明:
    echo   start   - 启动完整开发环境
    echo   stop    - 停止所有服务
    echo   restart - 重启所有服务
    echo   status  - 查看服务状态
    exit /b 1
)

goto :eof
