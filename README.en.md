<h1 align="center">Intelligent Video Slicing System</h1>

<p align="center">
  <a href="README.md">简体中文</a> | <b>English</b>
</p>

<p align="center"><strong>Highlights, CapCut-ready in seconds</strong></p>

**mint-ai** turns long-form and multi-source footage into editable shorts and exports: **ingestion → intelligent slicing → CapCut draft packaging → optional rendered video**, with less manual rough-cut work.

### Deployment

#### Prerequisites

- **Docker** and **Docker Compose** (recommended)  
- Credentials for **LLM**, **ASR** (e.g. Alibaba DashScope), and **object storage** (e.g. Tencent COS)—see `docker/.env.example` for Worker-related keys  
- Local development and API details: **`backend/worker/README.md`**

#### Docker Compose (recommended)

The repo’s `docker/docker-compose.yml` may define several services. To run **only** the slicing / draft / render pipeline, start Worker, capcut-mate, and Nginx:

```bash
cd docker

cp .env.example .env   # fill in LLM_*, ASR_*, COS_*, etc.

# Adjust WEB_ROOT_URL, DOWNLOAD_URL, DRAFT_URL, CAPCUT_MATE_URL, etc. to match your host or domain

docker compose up -d mint-worker mint-capcut-mate mint-nginx
```

Notes:

- **mint-worker** is usually published on host port **32001→30000** inside the container (confirm in your compose file).
- **mint-capcut-mate** must match **`CAPCUT_MATE_URL`** and any URLs used in Nginx/capcut env.
- **mint-nginx** may proxy `/openapi/mint-worker/` and `/openapi/capcut-mate/` depending on the image template.
- Ensure **`WEB_ROOT_DIR` / `WEB_ROOT_URL`** match how clients reach generated draft/video URLs.

Smoke test:

```bash
curl http://localhost:32001/openapi/health
```

OpenAPI UI: `http://localhost:32001/docs`.
