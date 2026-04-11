"""应用日志：默认同时输出到控制台与滚动日志文件。"""

from __future__ import annotations

import logging
import sys
from logging.handlers import RotatingFileHandler
from pathlib import Path

from core.config import get_settings


def setup_logging() -> None:
    """初始化 `mint_ai` 命名空间下的 logger（控制台 + 文件）。可重复调用，不会重复挂载 handler。"""
    log = logging.getLogger("mint_ai")
    if log.handlers:
        return

    s = get_settings()
    level_name = s.log_level.upper()
    level = getattr(logging, level_name, logging.INFO)

    # 使用调用处的源文件名与行号，便于定位；不再输出 logger 名（如 mint_ai.app）
    fmt = logging.Formatter(
        fmt="%(asctime)s | %(levelname)-8s | %(filename)s:%(lineno)d | %(message)s",
        datefmt="%Y-%m-%d %H:%M:%S",
    )

    log.setLevel(level)
    log.propagate = False

    console = logging.StreamHandler(sys.stderr)
    console.setLevel(level)
    console.setFormatter(fmt)

    log_dir = Path(s.log_dir)
    log_dir.mkdir(parents=True, exist_ok=True)
    log_path = log_dir / s.log_file

    file_handler = RotatingFileHandler(
        log_path,
        maxBytes=s.log_max_bytes,
        backupCount=s.log_backup_count,
        encoding="utf-8",
    )
    file_handler.setLevel(level)
    file_handler.setFormatter(fmt)

    log.addHandler(console)
    log.addHandler(file_handler)


def get_logger(name: str) -> logging.Logger:
    """
    获取业务 logger。`name` 一般传 `__name__`，实际 logger 名为 `mint_ai.<name>`。
    日志格式中展示的是调用处的「文件名:行号」，而非 logger 名。
    """
    setup_logging()
    if name.startswith("mint_ai"):
        return logging.getLogger(name)
    return logging.getLogger(f"mint_ai.{name}")
