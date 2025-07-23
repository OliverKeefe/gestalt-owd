import * as React from "react"
import { Button } from "@/components/ui/button"
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar"
import {
    DropdownMenu,
    DropdownMenuTrigger,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
} from "@/components/ui/dropdown-menu"
import { ModeToggle } from "@/components/mode-toggle"
import { SettingsToggle } from "../../buttons/settings-toggle"

interface AppTopbarProps extends React.HTMLAttributes<HTMLDivElement> {}

export function AppTopbar({ className, ...props }: AppTopbarProps) {
    return (
        <header
            className={`flex items-center justify-between px-4 py-2 ${className}`}
            {...props}
        >
            <div className="text-lg font-semibold"></div>
            <div className="flex items-center gap-1">
                <ModeToggle />
                <SettingsToggle />
            </div>
        </header>
    )
}
