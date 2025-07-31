import './App.css'
import { useEffect, useState } from "react";
import { ThemeProvider } from "@/components/theme-provider"
import Layout from '@/app/features/shared/components/layout/layout';
import { BrowserRouter } from 'react-router-dom';
import AppRoutes from '@/routes/app-routes';
import keycloak, { initKeycloak } from "@/security/auth/keycloak/keycloak";

function App() {
    const [isKeycloakReady, setIsKeycloakReady] = useState(false);
    const [isAuthenticated, setIsAuthenticated] = useState(false);

    useEffect(() => {
        initKeycloak().then((kc) => {
            if (kc) {
                console.log('Token:', kc.token);
                console.log('User:', kc.tokenParsed?.preferred_username);
                setIsAuthenticated(true);
            }
            setIsKeycloakReady(true);
        });
    }, []);

    if (!isKeycloakReady) {
        return <div className="p-8 text-center text-lg">🔐 Initializing authentication...</div>
    }

    return (
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
            <BrowserRouter>
                <Layout>
                    <AppRoutes />
                </Layout>
            </BrowserRouter>
        </ThemeProvider>
    );
}

export default App;
