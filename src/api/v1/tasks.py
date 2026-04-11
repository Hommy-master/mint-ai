from __future__ import annotations

from fastapi import APIRouter, HTTPException, Query
from fastapi.responses import JSONResponse

from schemas.common import ApiResponse, ErrorResponse
from schemas.tasks import TaskCreateData, TaskCreateRequest, TaskQueryData
from services.task_service import task_service

router = APIRouter()


@router.post(
    "",
    response_model=ApiResponse[TaskCreateData],
    summary="提交成片任务",
)
async def submit_task(body: TaskCreateRequest) -> ApiResponse[TaskCreateData]:
    tid = await task_service.create_task(body.video_url)
    return ApiResponse(data=TaskCreateData(task_id=tid))


@router.get(
    "",
    summary="查询任务状态",
    response_model=None,
    responses={
        200: {"model": ApiResponse[TaskQueryData]},
        404: {"model": ErrorResponse, "description": "task_id 不存在"},
    },
)
async def query_task(
    task_id: str = Query(..., min_length=1, description="任务 ID"),
):
    rec = await task_service.get_record(task_id)
    if rec is None:
        return JSONResponse(
            status_code=404,
            content=ErrorResponse(
                code=40401,
                message="task not found",
                data=None,
            ).model_dump(),
        )
    # 设计文档 §5.2 示例允许 data 为空对象；此处返回完整状态便于轮询
    return ApiResponse(data=task_service.to_query_data(rec))
