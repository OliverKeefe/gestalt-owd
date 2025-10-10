import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import { AppSidebar } from "@/app/features/shared/components/navigation/sidebars/app-sidebar"
import { Outlet } from "react-router-dom"
import { AppTopbar } from "../navigation/topbars/app-topbar"

interface LayoutProps {
    children?: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
    return (
        <SidebarProvider>
            <div className="flex min-w-screen">
                <AppSidebar />
                <div className="flex flex-1 flex-col">
                    <AppTopbar>
                        <SidebarTrigger />
                    </AppTopbar>
                    <main className="flex-1">
                        {children}
                        <Outlet />
                    </main>
                </div>
            </div>
        </SidebarProvider>
    )
}