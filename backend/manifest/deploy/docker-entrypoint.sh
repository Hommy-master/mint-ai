#!/bin/sh

# 标记文件路径
INIT_FLAG_FILE="/app/etc/init_flag"

echo "=== arg list ==="
echo "all args: $@"
echo "args count: $#"
echo "=== begin run ==="

# 检查是否是第一次启动
if [ ! -e "$INIT_FLAG_FILE" ]; then
    echo "first run，exec envinit..."
    /app/bin/envinit -c /app/etc/config.yaml
    # 创建标记文件
    touch "$INIT_FLAG_FILE"
else
    echo "just restart, skip envinit..."
fi

# 启动服务（此时$@包含CMD的指令）
exec "$@"
