import {Dialog, DialogContent, DialogTitle, DialogTrigger} from "@/components/ui/dialog.tsx";
import type {Metadata} from "@/app/features/files/hooks/types.ts";
import {useState} from "react";
import {Button} from "@/components/ui/button.tsx";
import {EllipsisVertical} from "lucide-react";

interface FileDialogProps{
    open: boolean,
    onOpenChange: (open: boolean) => void,
    metadata: Metadata,
    ipfsLink: string,
    spaceName: string,
    spaceDid: string
}

export function FileDialog({
                               open,
                               onOpenChange,
                               metadata,
                           }: FileDialogProps) {
    return (
        <Dialog open={open} onOpenChange={onOpenChange}>
            <DialogTitle title={metadata.file_name}></DialogTitle>
            <DialogContent>
                <h1>Last modified at: {metadata.modified_at}</h1>
                <h1>Last modified at: {metadata.modified_at}</h1>
            </DialogContent>
        </Dialog>
    )
}