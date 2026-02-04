import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import {useEffect, useMemo, useState} from "react";
import { Checkbox } from "@/components/ui/checkbox";
import { UploadDialog } from "@/app/features/shared/components/dialog/upload-dialog.tsx";
import { Button } from "@/components/ui/button.tsx";
import { Clock, FolderPlus, Star } from "lucide-react";
import {type CursorReq, getAllMetadata, type GetAllMetadataReq} from "@/app/features/home/hooks/handler.ts";
import type {Metadata} from "@/app/features/home/hooks/data.ts";
import {useAuthStore} from "@/security/auth/authstore/auth-store.ts";


export function FileTable() {
    const cur = useMemo<CursorReq>(() => ({
        modified_at: null, //"2025-02-13T11:21:04.791Z",
        id: null, //"c7a1735e-504e-47d9-a8c0-a0e37f7df8b3",
    }), []);

    const userId = useAuthStore((s) => s.userId);

    const req = useMemo<GetAllMetadataReq>(() => ({
        user_id: userId,
        cursor: cur,
        limit: 20,
    }), [cur, userId]);

    const [files, setFiles] = useState<Metadata[]>([]);

    useEffect(() => {
        if (!userId) return;

        getAllMetadata(req).then((resp) => {
            setFiles(resp.metadata);
        });
    }, [req, userId]);


    const [selected, setSelected] = useState<string[]>([]);

    const toggleSelect = (id: string) => {
        setSelected((prev) =>
            prev.includes(id)
                ? prev.filter((item) => item !== id)
                : [...prev, id]
        );
    };

    return (
        <div>
            <h1 className="text-2xl font-semibold pb-4 pt-4 m-1">All files</h1>

            <nav className="w-full flex gap-3">
                <Button variant="default">
                    <Clock /> Recents
                </Button>

                <UploadDialog
                    onUploaded={(newFile) => setFiles((prev) => [...prev, newFile])}
                />

                <Button variant="default">
                    <FolderPlus /> New Folder
                </Button>

                <Button variant="default">
                    <Star /> Favorites
                </Button>
            </nav>

            <Table className="mt-2">
                <TableHeader>
                    <TableRow>
                        <TableHead className="w-[30px]">
                            <Checkbox
                                checked={selected.length === files.length}
                                onCheckedChange={(checked) => {
                                    if (checked) {
                                        setSelected(files.map((file) => file.uuid));
                                    } else {
                                        setSelected([]);
                                    }
                                }}
                            />
                        </TableHead>
                        <TableHead className="w-[100px]">Last Modified</TableHead>
                        <TableHead>Name</TableHead>
                        <TableHead>Path</TableHead>
                        <TableHead className="text-right">Type</TableHead>
                    </TableRow>
                </TableHeader>

                <TableBody>
                    {files.map((file) => (
                        <TableRow key={file.uuid}>
                            <TableCell>
                                <Checkbox
                                    checked={selected.includes(file.uuid)}
                                    onCheckedChange={() => toggleSelect(file.uuid)}
                                />
                            </TableCell>
                            <TableCell>{file.modified_at}</TableCell>
                            <TableCell>{file.file_name}</TableCell>
                            <TableCell>{file.path}</TableCell>
                            <TableCell className="text-right">{file.file_type}</TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
        </div>
    );
}
