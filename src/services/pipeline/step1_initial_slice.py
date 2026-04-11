"""第一步：初步切片（≈20 分钟/段）。占位实现，不调用真实 FFmpeg。"""

from __future__ import annotations

import asyncio
from typing import TYPE_CHECKING

from schemas.tasks import TaskStatus

if TYPE_CHECKING:
    from services.task_service import TaskService


async def run(task_service: TaskService, task_id: str) -> None:
    await task_service.patch(
        task_id,
        status=TaskStatus.running,
        step=1,
        step_message="initial slice (~20 min segments)",
        progress=0.1,
    )
    await asyncio.sleep(0.15)
