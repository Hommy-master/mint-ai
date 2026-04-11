from functools import lru_cache

from pydantic_settings import BaseSettings, SettingsConfigDict


class Settings(BaseSettings):
    model_config = SettingsConfigDict(env_prefix="MINT_", env_file=".env", extra="ignore")

    capcut_mate_base_url: str = "http://127.0.0.1:8080"
    ffmpeg_path: str = "ffmpeg"
    llm_api_base: str = ""
    llm_api_key: str = ""


@lru_cache
def get_settings() -> Settings:
    return Settings()
