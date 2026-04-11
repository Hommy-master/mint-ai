import asyncio

import httpx
import pytest
from httpx import ASGITransport

from app import app


@pytest.mark.asyncio
async def test_submit_and_query_task() -> None:
    transport = ASGITransport(app=app)
    async with httpx.AsyncClient(transport=transport, base_url="http://test") as client:
        r = await client.post(
            "/openapi/v1/tasks",
            json={"video_url": "https://example.com/a.mp4"},
        )
        assert r.status_code == 200
        body = r.json()
        assert body["code"] == 0
        assert "task_id" in body["data"]
        tid = body["data"]["task_id"]

        await asyncio.sleep(0.6)

        r2 = await client.get("/openapi/v1/tasks", params={"task_id": tid})
        assert r2.status_code == 200
        data = r2.json()["data"]
        assert data["status"] == "succeeded"
        assert data["result_url"]


@pytest.mark.asyncio
async def test_query_unknown_task() -> None:
    transport = ASGITransport(app=app)
    async with httpx.AsyncClient(transport=transport, base_url="http://test") as client:
        r = await client.get(
            "/openapi/v1/tasks",
            params={"task_id": "00000000-0000-0000-0000-000000000000"},
        )
        assert r.status_code == 404
        assert r.json()["code"] == 40401


@pytest.mark.asyncio
async def test_invalid_video_url() -> None:
    transport = ASGITransport(app=app)
    async with httpx.AsyncClient(transport=transport, base_url="http://test") as client:
        r = await client.post(
            "/openapi/v1/tasks",
            json={"video_url": "ftp://example.com/a.mp4"},
        )
        assert r.status_code == 422
