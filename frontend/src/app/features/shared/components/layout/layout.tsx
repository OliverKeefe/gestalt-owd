import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar.tsx"
import { AppSidebar } from "@/app/features/shared/components/navigation/sidebars/app-sidebar.tsx";
//import AppTopbar from "@/components/shared/navigation/sidebars/app-topbar.tsx";

export default function Layout({ children }: { children: React.ReactNode }) {
    return (
        <SidebarProvider>
            <AppSidebar />
            <main className={"flex-1"}>
                <SidebarTrigger />
                {children}

            </main>
        </SidebarProvider>
    )
}