import React, {useState} from 'react'
import {
    Dialog,
    DialogContent,
    DialogTitle,
    DialogTrigger,
    DialogHeader,
    DialogDescription
} from "@/components/ui/dialog"
import {Button} from "@/components/ui/button.tsx";
import {Dropzone, DropzoneContent, DropzoneEmptyState} from "@/components/ui/shadcn-io/dropzone";
import {Upload} from "lucide-react";

type UploadDialogProps = {
    children?: React.ReactNode
}


export function UploadDialog({children}: UploadDialogProps) {
    const [open, setDialogOpen] = useState(false);
    const [files, setFiles] = useState<File[] | null>(null);
    const handleDrop = (files: File[]) => {
        console.log(files);
        setFiles(files);
    };

    const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files?.length) {
            setFiles(Array.from(e.target.files));
        }
    }

    async function handleDialogUpload(): Promise<void> {
        if (!files || files.length === 0) {
            alert("Can't upload an empty file.");
            return;
        }

        const formData = new FormData();
        files.forEach((file) => {
            formData.append("file", file);
        })


        try {
            const response = await fetch("http://127.0.0.1:8081/files/upload", {
                method: "PUT",
                body: formData,
            });

            if (!response.ok) {
                throw new Error(`Upload failed with status ${response.status}`);
            }

            setDialogOpen(false);
            setFiles(null);
        } catch (err) {
            throw new Error("failed to upload file.", err);
        }
    };

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
