import React, {useState} from 'react';
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogHeader,
    DialogTitle,
    DialogTrigger
} from "@/components/ui/dialog";
import {Button} from "@/components/ui/button.tsx";
import {Dropzone, DropzoneContent, DropzoneEmptyState} from "@/components/ui/shadcn-io/dropzone";
import {Upload} from "lucide-react";
import {RestHandler} from "@/app/features/shared/api/rest/rest-handler.ts";
import {UploadForm} from "@/app/features/shared/files/upload.ts";
import {uploadObject} from "@/app/features/shared/storacha/upload.ts";
import {ObjectType} from "@/app/features/shared/storacha/types.ts";
import type {Metadata} from "@/app/features/shared/files/metadata.ts";


type UploadDialogProps = {
    onUploaded?: (files: Metadata[]) => void;
};

const api = new RestHandler(`http://localhost:8081`);


export function UploadDialog({onUploaded}: UploadDialogProps) {
    const [alert, setAlertVisible] = useState(false);
    const [open, setDialogOpen] = useState(false);
    const [files, setFiles] = useState<File[] | null>(null);

    const handleDrop = (newFiles: File[]) => {
        setFiles(prev => {
            const existing = prev ?? [];
            const merged = [...existing, ...newFiles];

            return Array.from(
                new Map(merged.map(f => [`${f.name}-${f.size}-${f.lastModified}`, f])).values()
            );
        });
    };

    async function handleDialogUpload(): Promise<void> {
        if (!files || files.length === 0) {
            alert("Can't upload an empty file.");
            return;
        }

        try {
            const upload = new UploadForm(files);
            await upload.prepare();

            console.log("Sending upload form data...")
            const optimistic: Metadata[] = files.map((f) => ({
                uuid: crypto.randomUUID(), // temp ID
                file_name: f.name,
                path: "",
                size: f.size,
                file_type: f.type,
                modified_at: new Date().toISOString(),
                created_at: new Date().toISOString(),
                owner_id: "", // or userId if available
                access_to: [],
                group_id: [],
                checksum: new Uint8Array(),
                version: new Date().toISOString(),
            }));
            const response = await upload.send();

            onUploaded?.(optimistic);


            const bytes = new Uint8Array(100);
            crypto.getRandomValues(bytes);

            const testMetadata: Metadata[] = [
                {
                    id: crypto.randomUUID(),
                    ownerId: "test-user",
                    checkSum: "deadbeef",

                    path: "blegowe.bin",
                    relativePath: "blwgo.bin",
                    lastModified: Date.now(),
                    lastModifiedDate: new Date().toISOString(),
                    size: 1024,
                    fileType: "application/octet-stream",
                }
            ]

            await uploadObject([bytes], testMetadata, ObjectType.FILE);

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
                    <Upload/> Upload
                </Button>
            </DialogTrigger>

            <DialogContent className="max-w-3xl">
                <DialogHeader>
                    <DialogTitle>Upload File</DialogTitle>
                    <DialogDescription>Upload a file or folder.</DialogDescription>
                </DialogHeader>

                <Dropzone
                    accept={{"*/*": []}}
                    maxFiles={10}
                    maxSize={1024 * 1024 * 1024 * 15}
                    minSize={1}
                    onDrop={handleDrop}
                    onError={console.error}
                    src={files}
                >
                    <DropzoneEmptyState/>
                    <DropzoneContent/>
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

