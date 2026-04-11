# ClipFlow AI Docker 镜像
# 基于 Python 3.13 + uv

FROM python:3.13-slim

# 安装系统依赖（健康检查需 curl；FFmpeg 若流水线需要可在此追加）
RUN apt-get update \
    && apt-get install -y --no-install-recommends curl \
    && rm -rf /var/lib/apt/lists/*

# 安装 uv 包管理器
RUN pip install uv

# 设置工作目录
WORKDIR /app

# 复制依赖文件（利用 Docker 缓存层）
COPY pyproject.toml uv.lock README.md ./

# 使用 uv 安装依赖（生产环境，不包含 dev 依赖）
RUN uv sync --frozen --no-dev

# 复制应用代码
COPY main.py ./
COPY src/ ./src/

# 仓库内 app/bin（如 ffprobe、ffmpeg 静态/自带二进制）加入 PATH，供 ffmpeg 模块子进程调用
RUN cp tools/* /usr/local/bin/ && chmod +x /usr/local/bin/ffmpeg && chmod +x /usr/local/bin/ffprobe

# 创建数据目录
RUN mkdir -p /app/data /app/temp

# 暴露服务端口
EXPOSE 30000

# 健康检查
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=3 \
    CMD curl -f http://localhost:30000/openapi/health || exit 1

# 启动命令
CMD ["uv", "run", "main.py", "--workers", "4"]