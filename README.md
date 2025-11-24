# Gestalto
Client-Side Encrypted, Decentralized File Storage & Compute Platform Built on IPFS (Interplanetary File System).
<p align="center"><img alt="License" src="https://img.shields.io/badge/license-MIT-blue.svg" /> <img alt="Status" src="https://img.shields.io/badge/status-in%20development-orange" /> 
<img alt="backend" src="https://img.shields.io/github/actions/workflow/status/OliverKeefe/gestalto/go-tests.yml?branch=main">
<img alt="frontend" src="https://img.shields.io/github/actions/workflow/status/OliverKeefe/gestalto/ts-tests.yml?branch=main">
</p>

## About
Gestalto is a decentralized cloud storage and compute platform combining client-side encryption, zero-trust principles, and IPFS-based storage.
Data is always encrypted before leaving the client, and only decrypted locally by the user. No plaintext ever touches a backend.

Gestalto removes single points of failure, central surveillance vectors, and reliance on traditional cloud vendors.

#### Why?
- Protect User Privacy: Centralized cloud providers are vulnerable to government legal requests like the UK's TCN, data breaches, and insider threats. Gestalt keeps user data private by design.

- Decentralization: IPFS distributes data across multiple nodes, eliminating reliance on any single data silo, datacenter or cloud provider.

- Security: End-to-end client-side encryption ensures that only users can access their files, even if storage nodes are compromised.

- Disaster Resilience: Distributed storage increases availability and fault tolerance.

## Project Roadmap
**Minimum Viable Product – December 2025**
- User Authentication
- Secure login/registration via Keycloak
- Only encrypted metadata stored server-side
- Web UI for accounts, files, metadata, usage
- IPFS integration 
- Fast access namespace / cgroup object storage Upload, download, delete files
- Immediate deletion from dashboard
- Client-Side Encryption
    - WebCrypto API
- Per-file/session keys
- Core UX and functionality implemented

**Final Product Release – April 2026**
- File Sharing & Permissions
- RBAC / ABAC encrypted file sharing link generation
- Time-limited/revocable access
- Permission editing & tracking
- Versioning & Collaboration
- Version history
- Restore previous versions
- UI for comparing file versions
- Search Feature
- Metadata indexing
- Keyword, date, and tag search
- Privacy preserved (no plaintext indexing)
- Advanced Access Control
- RBAC / ABAC for teams / orgs
- Federated identity support (OAuth2/SAML)
- Deployment Tools

#### Quickstart: Run in Dev Environment

**Run frontend in dev mode**
```shell
cd frontend \
npm run dev
```

**Run test Auth Docker Container**
```shell
docker run -p 127.0.0.1:8080:8080 \
  -e KC_BOOTSTRAP_ADMIN_USERNAME=$USERNAME \
  -e KC_BOOTSTRAP_ADMIN_PASSWORD=$PASS \
  quay.io/keycloak/keycloak:26.4.0 start-dev
```
