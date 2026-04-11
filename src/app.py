from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware

from api.v1.router import router as v1_router
from core.logger import get_logger, setup_logging

setup_logging()
logger = get_logger(__name__)

app = FastAPI(
    title="Mint-AI API",
    description="输入视频一键成片 — 业务前缀 `/openapi/mint-ai/v1`，与 `/openapi.json` 规范路径不同。",
    version="0.1.0",
)

app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(v1_router, prefix="/openapi/mint-ai/v1")

logger.info("Mint-AI application loaded")


@app.get("/openapi/health", tags=["health"])
async def health() -> dict[str, str]:
    return {"status": "ok"}
