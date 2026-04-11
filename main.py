"""仓库根入口：`uv run main.py` 启动服务。"""

import uvicorn

if __name__ == "__main__":
    uvicorn.run(
        "app:app",
        host="0.0.0.0",
        port=30000,
        app_dir="src",
        reload=False,
    )
