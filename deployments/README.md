# Gestalt — Deployments

Every micro-frontend authenticates against a single Keycloak
realm (`gestalt`) using its own client; Open Web Drive uses client `open-web-drive`.

Two deployment types:

- **Docker Compose** (`compose/`) — development, homelab, and small orgs.
  One `docker compose up` runs the full product: identity (Keycloak), storage
  (LocalStack S3), Open Web Drive (db, backend, frontend).
- **Kubernetes** (`kubernetes/`) — larger / distributed deployments (work in progress).
  k8s manifests will live here and consume the same per-service images.

## Docker Compose

Prerequisite: Docker Compose v2.24+.

```shell
cd deployments/compose      

cat > .env <<'EOF'
# Copy to .env in this directory and adjust before `docker compose up`.
# `.env` is gitignored; `.env.example` is committed.

# PostgreSQL (OWD file metadata database)
POSTGRES_USER=admin
POSTGRES_PASSWORD=passwd
POSTGRES_DB=metadatadb

# Keycloak bootstrap admin (applies on first boot only)
KEYCLOAK_ADMIN_USERNAME=admin
KEYCLOAK_ADMIN_PASSWORD=change-me

# LocalStack S3 static credentials (used by localstack and backend containers)
AWS_ACCESS_KEY_ID=test
AWS_SECRET_ACCESS_KEY=test
EOF

docker compose up -d --build
```

```shell```

| Service | Host URL | Notes |
| --- | --- | --- |
| Frontend (OWD) | http://localhost | prod: nginx-served static build |
| Backend API (OWD) | http://localhost:8081 | |
| Keycloak admin | http://localhost:8080/admin | realm `gestalt`, client `open-web-drive` |
| LocalStack S3 | http://localhost:4567 | bucket `temp-buck` auto-created |

> Host port 4567 (not the default 4566) avoids colliding with a LocalStack that is
> already running on the host; containers reach LocalStack over the internal network,
> so the host mapping is only for manual `aws`/`awslocal` use.

The stack persists PostgreSQL, Keycloak, and LocalStack state in named volumes.

### Pinned images

- `keycloak/keycloak:26.2` is pulled from Docker Hub (quay.io's CDN can be slow or
  unavailable; quay.io is the upstream source and can be re-enabled).
- `localstack/localstack:3.8.1` — the community image. `latest` moved behind the
  LocalStack Pro license gate (`LOCALSTACK_AUTH_TOKEN`), so it is pinned to the last
  community build.

### Dev mode (hot reloading)

```shell
cd deployments/compose
docker compose -f docker-compose.yml -f docker-compose.dev.yml up
```

Frontend runs on http://localhost:5173, DB exposed on host port 5433, and source
directories are bind-mounted for live edits. Dev needs no product Dockerfiles — it
uses `golang` and `node` toolchain images directly.

### Adding another app (micro-frontend path)

Add a per-product client to the `gestalt` realm (or a second realm export), then add
product-scoped services here — e.g. `docs-backend`, `docs-frontend` — sharing the same
Keycloak and storage services. No per-product identity infrastructure.

This will become more relevant as the project is expanded to include more office applications
and an IPFS gateway.