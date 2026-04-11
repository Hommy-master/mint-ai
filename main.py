"""仓库根入口：`uv run main.py` 启动服务。"""

from __future__ import annotations

import sys
from pathlib import Path

import uvicorn
from fastapi.routing import APIRoute
from starlette.routing import Mount

_src = Path(__file__).resolve().parent / "src"
if str(_src) not in sys.path:
    sys.path.insert(0, str(_src))


def print_routes() -> None:
    from app import app as fastapi_app

    print("Registered routes:", flush=True)
    for route in fastapi_app.routes:
        if isinstance(route, APIRoute):
            methods = sorted(m for m in route.methods if m != "HEAD")
            line = f"  {'|'.join(methods):12} {route.path}"
            if route.name:
                line += f"  [{route.name}]"
            print(line, flush=True)
        elif isinstance(route, Mount):
            print(f"  {'MOUNT':12} {route.path}", flush=True)
        else:
            path = getattr(route, "path", "")
            print(f"  {type(route).__name__:12} {path}", flush=True)
    print(flush=True)


if __name__ == "__main__":
    print_routes()
    uvicorn.run(
        "app:app",
        host="0.0.0.0",
        port=30000,
        app_dir="src",
        reload=False,
    )
