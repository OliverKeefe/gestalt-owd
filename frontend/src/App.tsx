import './App.css'
import {LoginForm} from "@/app/features/auth/components/login-form.tsx";
import { ThemeProvider } from "@/components/theme-provider"
import Layout from '@/app/features/shared/components/layout/layout';
import { BrowserRouter } from 'react-router-dom';
import AppRoutes from '@/routes/app-routes';

function App() {
  return (
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
            <BrowserRouter>
                <Layout>
                    <AppRoutes />
                </Layout>
            </BrowserRouter>
        </ThemeProvider>
  )
}

export default App
