import { EllipsisVertical, Info, Settings, Trash2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { useAuthStore } from "@/security/auth/authstore/auth-store";
import { RestHandler } from "@/app/features/shared/api/rest/rest-handler";

interface FileDropdownProps {
    fileId: string;
    onDeleted: () => void;
}

export function FileDropdown({ fileId, onDeleted }: FileDropdownProps) {
    const userId = useAuthStore((s) => s.userId);

    async function handleDelete(e: React.MouseEvent) {
        e.stopPropagation();

        if (!userId) {
            console.error("No user ID found");
            return;
        }

        try {
            const api = new RestHandler("http://localhost:8081");

            await api.handlePost<
                { id: string; },
                void
            >("api/files/delete", {
                id: fileId,
            });

            onDeleted();

        } catch (error) {
            console.error("Failed to delete file on server:", error);
        }
    }

    return (
        <DropdownMenu>
            <DropdownMenuTrigger asChild>
                <Button
                    variant="ghost"
                    size="icon"
                    className="ms-1 cursor-pointer"
                >
                    <EllipsisVertical className="h-4 w-4" />
                </Button>
            </DropdownMenuTrigger>

            <DropdownMenuContent align="end">
                <DropdownMenuItem
                    className="cursor-pointer text-destructive focus:text-destructive"
                    onClick={handleDelete}
                >
                    <Trash2 className="mr-2 h-4 w-4" />
                    <span>Delete</span>
                </DropdownMenuItem>

                <DropdownMenuSeparator />

                <DropdownMenuItem className="cursor-pointer">
                    <Info className="mr-2 h-4 w-4" />
                    <span>File Info</span>
                </DropdownMenuItem>

                <DropdownMenuItem className="cursor-pointer">
                    <Settings className="mr-2 h-4 w-4" />
                    <span>File Settings</span>
                </DropdownMenuItem>
            </DropdownMenuContent>
        </DropdownMenu>
    );
}
