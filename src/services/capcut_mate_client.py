"""capcut-mate 客户端占位；对接真实服务时在此实现 HTTP 调用与 gen_video。"""

from core.config import get_settings


class CapcutMateClient:
    def __init__(self) -> None:
        self._base = get_settings().capcut_mate_base_url.rstrip("/")

    async def gen_video(self, project_payload: dict) -> str:
        """调用 capcut-mate 生成成片，返回可访问 URL（占位）。"""
        _ = project_payload
        _ = self._base
        # 实现阶段：POST self._base + "/gen_video" 等
        return "https://example.com/placeholder-output.mp4"
