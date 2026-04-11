from enum import StrEnum

from pydantic import BaseModel, Field, field_validator


class TaskStatus(StrEnum):
    pending = "pending"
    running = "running"
    succeeded = "succeeded"
    failed = "failed"


class TaskCreateRequest(BaseModel):
    video_url: str = Field(
        ...,
        min_length=1,
        max_length=2048,
        description="原始视频可访问地址（仅 http/https）",
    )

    @field_validator("video_url")
    @classmethod
    def http_https_only(cls, v: str) -> str:
        if not v.startswith(("http://", "https://")):
            raise ValueError("video_url must start with http:// or https://")
        return v


class TaskCreateData(BaseModel):
    task_id: str


class TaskQueryData(BaseModel):
    status: TaskStatus
    step: int | None = Field(None, description="当前主阶段 1～3，未开始时为 null")
    step_message: str | None = None
    progress: float | None = Field(None, ge=0.0, le=1.0, description="0～1")
    result_url: str | None = None
    error_message: str | None = None
