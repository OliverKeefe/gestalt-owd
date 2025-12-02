import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { initKeycloak } from '@/security/auth/keycloak/keycloak.ts';
import { useAuthStore } from "@/security/auth/authstore/auth-store.ts";

const rootElement = document.getElementById("root")!;
rootElement.innerHTML = `
    <div class="w-18rem h-screen flex items-center justify-center">
        <img 
         src="/media/spinner.svg"
         alt="loading..." 
         class="w-[100px] h-[100px] invert-0 dark:invert"
         />
    </div>
`;

async function auth() {
    const kc = await initKeycloak();

    if (kc) {
        console.log('Token:', kc.token);
        console.log('User:', kc.tokenParsed?.preferred_username);
    }

    if (kc?.tokenParsed) {
        const userId = kc.tokenParsed.sub;
        useAuthStore.getState().setUserId(userId);
    }

    const isAuthenticated = !!kc;

    createRoot(document.getElementById('root')!).render(
        <StrictMode>
            <App isAuthenticated={isAuthenticated}/>
        </StrictMode>,
    );

}

auth();