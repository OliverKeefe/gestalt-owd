import * as React from "react"
import { ModeToggle } from "@/components/mode-toggle"

interface AppTopbarProps extends React.HTMLAttributes<HTMLDivElement> {
    children?: React.ReactNode;

}

export function AppTopbar({ children, className, ...props }: AppTopbarProps) {
    return (
        <header
            className={`flex w-full items-center px-4 py-2 ${className}`}
            {...props}
        >
            <div className="flex items-center gap-2">
                {children}
            </div>

            <div className="ml-auto flex items-center gap-2">
                <ModeToggle />
                {/*<SettingsToggle />*/}
            </div>
        </header>
    )
}

