import React, { useState } from 'react';
import {
    Dialog,
    DialogContent,
    DialogTitle,
    DialogTrigger,
    DialogHeader,
    DialogDescription
} from "@/components/ui/dialog";
import { Button } from "@/components/ui/button.tsx";
import { Dropzone, DropzoneContent, DropzoneEmptyState } from "@/components/ui/shadcn-io/dropzone";
import { Upload } from "lucide-react";
import { RestHandler } from "@/app/features/shared/api/rest/rest-handler.ts";
import {UploadForm} from "@/app/features/shared/files/upload.ts";

type UploadDialogProps = {
    onUploaded?: (file: any) => void;
};

const api = new RestHandler(`http://localhost:8081`);


export function UploadDialog() {
    const [open, setDialogOpen] = useState(false);
    const [files, setFiles] = useState<File[] | null>(null);

    const handleDrop = (files: File[]) => setFiles(files);

    async function handleDialogUpload(): Promise<void> {
        if (!files || files.length === 0) {
            alert("Can't upload an empty file.");
            return;
        }

        try {
            const upload = new UploadForm(files);
            await upload.prepare();

            console.log("Sending upload form data...")
            const response = await upload.send();
            console.log(response)

            setDialogOpen(false);
            setFiles(null);
        } catch (err) {
            console.log(err);
        }
    }

    return (
        <Dialog open={open}
                onOpenChange={(isOpen) => {
                    setDialogOpen(isOpen)
                    if (!isOpen) setFiles(null)
                }}>
            <DialogTrigger asChild>
                <Button variant="default">
                    <Upload /> Upload
                </Button>
            </DialogTrigger>

            <DialogContent className="max-w-3xl">
                <DialogHeader>
                    <DialogTitle>Upload File</DialogTitle>
                    <DialogDescription>Upload a file or folder.</DialogDescription>
                </DialogHeader>

                <Dropzone
                    accept={{ "*/*": [] }}
                    maxFiles={10}
                    maxSize={1024 * 1024 * 1024 * 15}
                    minSize={1024}
                    onDrop={handleDrop}
                    onError={console.error}
                    src={files}
                >
                    <DropzoneEmptyState />
                    <DropzoneContent />
                </Dropzone>

                <div className="flex flex-col space-y-2 mt-3">
                    <Button className="cursor-pointer" onClick={handleDialogUpload}>
                        Upload
                    </Button>
                </div>
            </DialogContent>
        </Dialog>
    );
}
