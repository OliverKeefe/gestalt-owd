import './App.css'
import { ThemeProvider } from "@/components/theme-provider"
import { RouterProvider } from "@tanstack/react-router";
import { router } from "@/routes/app-routes";
import {TooltipProvider} from "@/components/ui/tooltip.tsx";

function App() {
    return (
        <ThemeProvider defaultTheme="dark" storageKey="vite-ui-theme">
            <TooltipProvider>
                <RouterProvider router={router} />
            </TooltipProvider>
        </ThemeProvider>
    );
}

export default App;
