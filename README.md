<h1 align="left">Open Web Drive</h1>

<p align="left"> 

</p>

## 🛠️ About This Project

<p align="left">
  <img alt="Go" src="https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go">
  <img alt="TypeScript" src="https://img.shields.io/badge/TypeScript-5.8-3178C6?logo=typescript">
  <img alt="backend" src="https://img.shields.io/github/actions/workflow/status/OliverKeefe/open-web-drive/backend-ci.yml?branch=main&label=backend">
  <img alt="frontend" src="https://img.shields.io/github/actions/workflow/status/OliverKeefe/open-web-drive/frontend-ci.yml?branch=main&label=frontend">
  <img alt="License" src="https://img.shields.io/badge/license-MIT-white.svg">
</p>

**Open Web Drive** is a Dropbox / Google Drive clone I built in TypeScript (React.js), Go and Postgres, in an effort to learn Go. Future
features will also include a document editor and IPFS gateway service.


<h2 align="center">UI Screenshots</h2>
<table>
  <tr>
    <td width="50%"><img src="docs/media/readme/Screenshot from 2026-08-05 17-01-55.png" width="100%" alt="Image 1"></td>
    <td width="50%"><img src="docs/media/readme/Screenshot from 2026-08-05 17-02-57.png" width="100%" alt="Image 2"></td>
  </tr>
  <tr>
    <td><img src="docs/media/readme/Screenshot from 2026-08-05 22-06-36.png" width="100%" alt="Image 3"></td>
    <td><img src="docs/media/readme/Screenshot from 2026-08-05 17-02-40.png" width="100%" alt="Image 4"></td>
  </tr>
</table>

> [!NOTE]
> 🏗️ Auth / Client-Side Encryption are currently undergoing a full re-write.

## 🚀 Quick Start
**📋 Prerequisites**

**[Terraform](https://developer.hashicorp.com/terraform/install)**

**[Minikube](https://minikube.sigs.k8s.io/docs/start/?arch=%2Flinux%2Fx86-64%2Fstable%2Fbinary+download)**

**[Docker](https://docs.docker.com/engine/install/)**

**[Go](https://go.dev/dl/)**

**[Node.js](https://nodejs.org/en/download)**

**[AWS CLI](https://docs.aws.amazon.com/cli/latest/userguide/getting-started-install.html)**

**[LocalStack](https://docs.localstack.cloud/aws/getting-started/installation/)**

**Clone the repository**
```shell
git clone https://github.com/OliverKeefe/open-web-drive.git 
```

Once you've cloned the repository, change your working directory to ~/open-web-drive and
execute the following setup and dev build scripts.
**Setup Dev Environment**
```shell
chmod +x /scripts/setup.sh && ./scripts/setup.sh
```

**Run in Dev Mode**
```shell
chmod +x ./scripts/build.sh &&./scripts/build.sh
```

**Deploy with Docker Compose** (development / homelab / small orgs)
```shell
cd deployments/compose

cat > .env <<'EOF'
# Copy to .env in this directory and adjust before `docker compose up`.
# `.env` is gitignored; `.env.example` is committed.

# PostgreSQL (Open Web Drive metadata database)
POSTGRES_USER=admin
POSTGRES_PASSWORD=passwd
POSTGRES_DB=metadatadb

# Keycloak bootstrap admin (applies on first boot only)
KEYCLOAK_ADMIN_USERNAME=admin
KEYCLOAK_ADMIN_PASSWORD=change-me

# LocalStack S3 static credentials (used by both localstack and backend)
AWS_ACCESS_KEY_ID=test
AWS_SECRET_ACCESS_KEY=test
EOF

docker compose up -d --build
```

For dev with hot reload: `docker compose -f docker-compose.yml -f docker-compose.dev.yml up`.
Kubernetes deploys for larger/distributed setups are in progress under `deployments/kubernetes/`.
The compose stack runs Keycloak, LocalStack S3, and Open Web Drive end to end —
see `deployments/README.md` for the full breakdown (URLs, pinned images).

Kubernetes deployments are still very much a work in progress.

📄 License
Distributed under the MIT License. See `LICENSE.md` for more information.
