from fastapi import APIRouter

from api.v1.tasks import router as tasks_router

router = APIRouter()
# 完整业务前缀由 app.py 挂载：/openapi/mint-ai/v1
router.include_router(tasks_router, prefix="/tasks", tags=["tasks"])
