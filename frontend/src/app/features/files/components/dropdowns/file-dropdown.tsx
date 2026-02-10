import {EllipsisVertical, Info, Moon, Settings, Sun, Trash2} from "lucide-react"

import { Button } from "@/components/ui/button"
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem, DropdownMenuPortal, DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu"
import { useTheme } from "@/components/theme-provider"

export function FileDropdown() {
    const { setTheme } = useTheme()

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button variant="ghost" size="icon" className="ms-1 cursor-pointer">
                    <EllipsisVertical />
                </Button>
            </DropdownMenuTrigger>
            <DropdownMenuContent align="end">
                <DropdownMenuItem className={"cursor-pointer"} onClick={() => deleteFile(id)}>
                        {<Trash2 className={"align-left"} />}{<p>Delete</p>}
                </DropdownMenuItem>
                <DropdownMenuSeparator />
                <DropdownMenuItem>
                    {<Info />}{<p>File Info</p>}
                </DropdownMenuItem>
                <DropdownMenuPortal></DropdownMenuPortal>
                <DropdownMenuItem>
                    {<Settings />}{<p>File Settings</p>}
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    )
}