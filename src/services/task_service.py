from __future__ import annotations

import asyncio
import uuid
from dataclasses import dataclass

from schemas.tasks import TaskQueryData, TaskStatus
from services.pipeline.runner import run_pipeline


@dataclass
class TaskRecord:
    task_id: str
    video_url: str
    status: TaskStatus = TaskStatus.pending
    step: int | None = None
    step_message: str | None = None
    progress: float | None = None
    result_url: str | None = None
    error_message: str | None = None


class TaskService:
    def __init__(self) -> None:
        self._tasks: dict[str, TaskRecord] = {}
        self._lock = asyncio.Lock()

    async def create_task(self, video_url: str) -> str:
        task_id = str(uuid.uuid4())
        async with self._lock:
            self._tasks[task_id] = TaskRecord(task_id=task_id, video_url=video_url)
        asyncio.create_task(self._run_pipeline(task_id))
        return task_id

    async def get_record(self, task_id: str) -> TaskRecord | None:
        async with self._lock:
            return self._tasks.get(task_id)

    def to_query_data(self, rec: TaskRecord) -> TaskQueryData:
        return TaskQueryData(
            status=rec.status,
            step=rec.step,
            step_message=rec.step_message,
            progress=rec.progress,
            result_url=rec.result_url,
            error_message=rec.error_message,
        )

    async def patch(self, task_id: str, **kwargs: object) -> None:
        async with self._lock:
            rec = self._tasks.get(task_id)
            if rec is None:
                return
            for key, val in kwargs.items():
                setattr(rec, key, val)

    async def _run_pipeline(self, task_id: str) -> None:
        try:
            await run_pipeline(self, task_id)
        except Exception as e:
            await self.patch(
                task_id,
                status=TaskStatus.failed,
                error_message=str(e),
            )


task_service = TaskService()
