import './App.css'
import { useEffect, useState } from "react";
import { ThemeProvider } from "@/components/theme-provider"
import Layout from '@/app/features/shared/components/layout/layout';
import { BrowserRouter } from 'react-router-dom';
import AppRoutes from '@/routes/app-routes';

interface AppProps {
    isAuthenticated: boolean;
}

function App({ isAuthenticated }: AppProps) {
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
