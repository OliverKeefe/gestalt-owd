import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar"
import { AppSidebar } from "@/app/features/shared/components/navigation/sidebars/app-sidebar"
import { Outlet } from "react-router-dom"
import { AppTopbar } from "../navigation/topbars/app-topbar"
import {SecondarySidebar} from "@/app/features/shared/components/navigation/sidebars/secondary-sidebar.tsx";

interface LayoutProps {
    children?: React.ReactNode;
}

export default function Layout({ children }: LayoutProps) {
    return (
        <SidebarProvider>
            <div className="flex w-screen overflow-x-hidden ">
                <AppSidebar />
                <SecondarySidebar />
                <div className="pl-16 flex flex-1 min-w-0 flex-col">
                    <AppTopbar>
                        <SidebarTrigger />
                    </AppTopbar>
                    <main className="flex-1 min-w-0 overflow-x-hidden theme-scroll">
                        {children}
                        <Outlet />
                    </main>

                </div>
            </div>
        </SidebarProvider>
    )
}
