"""第三步：capcut-mate 剪辑 + gen_video。占位实现。"""

from __future__ import annotations

import asyncio
from typing import TYPE_CHECKING

from schemas.tasks import TaskStatus
from services.capcut_mate_client import CapcutMateClient

if TYPE_CHECKING:
    from services.task_service import TaskService


async def run(task_service: TaskService, task_id: str) -> None:
    await task_service.patch(
        task_id,
        step=3,
        step_message="capcut-mate editing and gen_video",
        progress=0.75,
    )
    client = CapcutMateClient()
    await asyncio.sleep(0.1)
    result_url = await client.gen_video({"task_id": task_id})
    await task_service.patch(
        task_id,
        status=TaskStatus.succeeded,
        step=3,
        step_message="done",
        progress=1.0,
        result_url=result_url,
    )
