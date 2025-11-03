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
    const [files, setFiles] = useState<File[] | undefined>();
    const handleDrop = (files: File[]) => {
        console.log(files);
        setFiles(files);
    };

    async function handleUpload(): Promise<void> {
        return
    }

    return (
        <Dialog open={open} onOpenChange={setDialogOpen}>
            <DialogTrigger asChild>
                <Button variant="ghost" className="w-full justify-start flex items-center space-x-2">
                    <Upload />
                    <p>Upload</p>
                </Button>
            </DialogTrigger>

            <DialogContent className="max-w-3xl">
                <DialogHeader>
                    <div className="flex items-center space-x-2">

                        <DialogTitle> Upload File</DialogTitle>
                    </div>
                    <DialogDescription>Upload a file or folder.</DialogDescription>
                </DialogHeader>
                <Dropzone
                    accept={{ 'file/*': [] }}
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

                    <div className="flex flex-col space-y-2" >
                        <Button
                            className="cursor-pointer"
                            onClick={handleUpload}>Upload</Button>
                    </div>
            </DialogContent>
        </Dialog>
    );
}