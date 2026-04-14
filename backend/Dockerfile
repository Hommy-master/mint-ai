# 使用官方Python镜像作为基础镜像
FROM alpine:latest

# 设置工作目录
WORKDIR /

# 将应用程序代码复制到工作目录
COPY _ROOT_/ .

# 将所有文件的换行符从Windows格式转换为Unix格式
RUN find . -type f \( -name "*.py" -o -name "*.sh" \) -exec sed -i 's/\r$//' {} +

# 添加权限
RUN chmod 0755 /app/bin/*

# 暴露容器端口
EXPOSE 30000

# 设置入口脚本
ENTRYPOINT ["/app/bin/docker-entrypoint.sh"]

# 定义容器启动时运行的命令
CMD ["/app/bin/webserver", "-c", "/app/etc/config.yaml"]
