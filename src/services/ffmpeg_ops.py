"""FFmpeg 封装占位；真实环境需调用本地 ffmpeg 可执行文件。"""

from core.config import get_settings


def ffmpeg_binary() -> str:
    return get_settings().ffmpeg_path
