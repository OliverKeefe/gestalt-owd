    import { StrictMode } from 'react'
    import { createRoot } from 'react-dom/client'
    import './index.css'
    import App from './App.tsx'
    import { initKeycloak } from '@/security/auth/keycloak/keycloak.ts';
    import { useAuthStore } from "@/security/auth/authstore/auth-store.ts";

    const root = createRoot(document.getElementById('root')!);

    async function bootstrap() {
        const kc = await initKeycloak();

        // If Keycloak determines user not signed in, redirect to login page.
        // The spinner in index.html stays visible until the browser actually leaves the page.
        if (!kc || !kc.authenticated) {
            kc?.login();
            return;
        }

        useAuthStore.getState().setToken(kc.token ?? null);
        useAuthStore.getState().setUserId(kc.tokenParsed?.sub ?? null);

        root.render(
            <StrictMode>
                <App isAuthenticated={true} />
            </StrictMode>
        );
    }

    bootstrap();