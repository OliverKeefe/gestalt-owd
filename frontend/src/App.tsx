import './App.css'
import {LoginForm} from "@/app/features/auth/components/login-form.tsx";
import { ThemeProvider } from "@/components/theme-provider"
import Layout from '@/app/features/shared/components/layout/layout';

function App() {
  return (
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">

            <Layout>
                <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
                    <div className="w-full max-w-sm">
                        <LoginForm />
                    </div>
                </div>
            </Layout>
        </ThemeProvider>
  )
}

export default App
