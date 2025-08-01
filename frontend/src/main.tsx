import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { initKeycloak } from '@/security/auth/keycloak/keycloak.ts';

async function auth() {
    const kc = await initKeycloak();

    if (kc) {
        console.log('Token:', kc.token);
        console.log('User:', kc.tokenParsed?.preferred_username);
    }

    const isAuthenticated = !!kc;

    createRoot(document.getElementById('root')!).render(
        <StrictMode>
            <App isAuthenticated={isAuthenticated}/>
        </StrictMode>,
    );
}

document.getElementById('root')!.innerHTML = `<div style="padding:2rem;text-align:center;">🔐 Initializing authentication...</div>`;

auth();