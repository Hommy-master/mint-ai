"""编排 §2 三步流水线。"""

from __future__ import annotations

from typing import TYPE_CHECKING

from services.pipeline import step1_initial_slice, step2_ai_highlights, step3_capcut_mate

if TYPE_CHECKING:
    from services.task_service import TaskService


async def run_pipeline(task_service: TaskService, task_id: str) -> None:
    await step1_initial_slice.run(task_service, task_id)
    await step2_ai_highlights.run(task_service, task_id)
    await step3_capcut_mate.run(task_service, task_id)
