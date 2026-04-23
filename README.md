<h1 align="center">智能视频切片系统</h1>

<p align="center">
  <b>简体中文</b> | <a href="README.en.md">English</a>
</p>

<p align="center"><strong>高光不遗漏，智能切片秒出稿</strong></p>

**mint-ai** 面向长视频与多源素材，提供从「素材接入 → 智能理解与选片 → 剪映草稿 → 成片导出」的一体化能力，降低人工粗剪与封装成本。

### 部署方法

#### 环境要求

- **Docker** 与 **Docker Compose**（推荐用于生产或一键体验）
- 需自备：**LLM**、**ASR**（如阿里云 DashScope）及 **对象存储**（如腾讯云 COS）等密钥与参数，参见 `docker/.env.example` 中与 Worker 相关的变量
- 本地开发、接口细节见 **`backend/worker/README.md`**

#### Docker Compose（推荐）

仓库中 `docker/docker-compose.yml` 可能包含多个服务。若你只需要 **切片 / 草稿 / 成片流水线**，可只启动 Worker、capcut-mate 与 Nginx（不启动与本流水线无关的其他容器）：

```bash
cd docker

# 复制并按说明填写密钥（至少配置 LLM_*、ASR_*、COS_* 等与 Worker 相关的项）
copy .env.example .env    # Windows
# cp .env.example .env    # Linux / macOS

# 根据实际主机 IP / 域名修改 compose 中与 WEB_ROOT_URL、DOWNLOAD_URL、DRAFT_URL、CAPCUT_MATE_URL 等有关的值

docker compose up -d mint-worker mint-capcut-mate mint-nginx
```

要点说明：

- **mint-worker**：对外默认映射宿主机端口 **32001→容器 30000**（以你本地 `docker-compose.yml` 为准）。
- **mint-capcut-mate**：草稿生成与渲染相关服务，需与 Worker 提供的 `CAPCUT_MATE_URL`、以及 Nginx/capcut 的环境变量指向一致。
- **mint-nginx**：可为静态目录与 `/openapi/mint-worker/`、`/openapi/capcut-mate/` 等路径提供反向代理（具体以镜像内模板为准）。
- Worker 会持续写入 **`docker/html`**（或你在 compose 中挂载的路径），用于草稿与成片 URL 的可访问性；请确保 **`WEB_ROOT_DIR` / `WEB_ROOT_URL`** 与实际访问方式一致。

验证 Worker 是否正常：

```bash
curl http://localhost:32001/openapi/health
```

浏览器打开：`http://localhost:32001/docs` 查看 OpenAPI 文档。
