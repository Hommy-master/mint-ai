"""第二步：大模型分析 + 清洗 + FFmpeg 高光切片。占位实现。"""

from __future__ import annotations

import asyncio
from typing import TYPE_CHECKING

from schemas.tasks import TaskStatus

if TYPE_CHECKING:
    from services.task_service import TaskService


async def run(task_service: TaskService, task_id: str) -> None:
    await task_service.patch(
        task_id,
        step=2,
        step_message="AI analysis and highlight extraction",
        progress=0.45,
    )
    await asyncio.sleep(0.2)
