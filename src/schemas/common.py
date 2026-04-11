from typing import Any, Generic, TypeVar

from pydantic import BaseModel, Field

T = TypeVar("T")


class ApiResponse(BaseModel, Generic[T]):
    """统一业务响应包（成功时 code=0）。"""

    code: int = Field(0, description="0 表示成功")
    message: str = Field("success", description="说明")
    data: T | None = None


class ErrorResponse(BaseModel):
    code: int
    message: str
    data: Any | None = None
