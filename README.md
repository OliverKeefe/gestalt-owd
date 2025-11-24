# Gestalto
![Backend](https://img.shields.io/github/actions/workflow/status/OliverKeefe/gestalto/go-tests.yml?branch=main)


### A Client-Side Encrypted, Privacy Focused Cloud Storage Platform Leveraging IPFS

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

#### What is Gestalt
Gestalt is a decentralized cloud storage platform that combines client-side encryption with IPFS-based distributed storage. Users maintain full control over their encryption keys, ensuring that files are encrypted locally before upload and decrypted only on their devices. Gestalt eliminates single points of failure and provides privacy and resilience without relying on centralized cloud providers.

#### Why?
- Protect User Privacy: Centralized cloud providers are vulnerable to government legal requests like the UK's TCN, data breaches, and insider threats. Gestalt keeps user data private by design.

- Decentralization: IPFS distributes data across multiple nodes, eliminating reliance on any single server or provider.

- Security: End-to-end client-side encryption ensures that only users can access their files, even if storage nodes are compromised.

- Disaster Resilience: Distributed storage increases availability and fault tolerance.

#### High Level Aims 

**December 2025: Minimum Viable Product**
- User authentication and account management
    - Implement secure login and registration using Keycloak.

    - Store only metadata in the backend; no plaintext passwords or keys leave the client.

    - Provide a web UI displaying account info, files, file information, and storage usage.

- File upload, download, and deletion via IPFS

    - Integrate IPFS client library for peer-to-peer file storage.

    - Ensure files are pinned to a test network for redundancy during MVP.

    - Allow users to delete files from their dashboard with instant confirmation.

- Client-side encryption/decryption
    
    - Use browser-native WebCrypto API to encrypt files locally before upload.

    - Keys are generated per file/session and never sent to the server.

    - Decryption occurs only in the browser on download, ensuring end-to-end privacy.

- Core UI and functional SPA Frontend WebApp.
    
    - Build a responsive Single-Page Application (SPA) using React.js, TypeScript, Vite, SWC and Tailwind CSS.

    - Provide intuitive interfaces for app functions.

    - Ensure responsive UI design and basic UX consistency across devices.

**April 2026: Final Product Release**
- File sharing and permission management
    
    - Implement p2p file sharing using encrypted links.

    - Allow time-limited or revocable access to shared files.

    - Track and modify file access permissions.

- Version control and collaboration systems
    
    - Maintain file version history.

    - Enable collaborative editing or upload of new versions while preserving previous file states.

    - Provide UI for comparing versions and restoring previous iterations.

- Fully implemented search feature

    - Index metadata (filename, tags, upload date) and query via backend API.

    - Support keyword, date, and tag-based searches while maintaining encrypted file privacy.

- Fine-grained RBAC and ABAC controls, support for identity federation.

    - Implement Role-Based Access Control (RBAC) and Attribute-Based Access Control (ABAC) for enterprise users.

    - Integrate federated identity providers (e.g., OAuth2, SAML).

- Build tools for easy deployments for home and business users
    
    - Provide scripts and containerized deployments for self-hosted environments.

    - Include documentation and setup automation for ease-of-use.