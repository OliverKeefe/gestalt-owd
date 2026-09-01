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

        /** Read identity claims directly from the locally-decoded access token.
         * This avoids the Keycloak account/userinfo endpoints entirely, which
         * are cross-origin here and can fail with CORS/401 during bootstrap.
         *
         * * **/
        const parsed = kc.tokenParsed as {
            sub?: string;
            preferred_username?: string;
            email?: string;
        } | null;

        useAuthStore.getState().setToken(kc.token ?? null);
        useAuthStore.getState().setUserId(parsed?.sub ?? null);
        useAuthStore.getState().setUsername(parsed?.preferred_username ?? parsed?.email ?? null);

        root.render(
            <StrictMode>
                <App />
            </StrictMode>
        );
    }

    bootstrap();