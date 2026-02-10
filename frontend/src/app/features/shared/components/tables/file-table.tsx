import {
    Table,
    TableBody,
    TableCell,
    TableHead,
    TableHeader,
    TableRow,
} from "@/components/ui/table"
import { useEffect, useMemo, useState, useOptimistic } from "react"
import { Checkbox } from "@/components/ui/checkbox"
import { UploadDialog } from "@/app/features/shared/components/dialog/upload-dialog"
import { Button } from "@/components/ui/button"
import { Clock, FolderPlus, Star, EllipsisVertical } from "lucide-react"
import {
    type CursorReq,
    getAllMetadata,
    type GetAllMetadataReq,
} from "@/app/features/files/hooks/handler"
import type { Metadata } from "@/app/features/files/hooks/types"
import { useAuthStore } from "@/security/auth/authstore/auth-store"
import { getIconForFile } from "@react-symbols/icons/utils"
import {DialogTrigger} from "@/components/ui/dialog.tsx";
import {FileDialog} from "@/app/features/files/components/dialogs/file-dialog.tsx";
import {FileDropdown} from "@/app/features/files/components/dropdowns/file-dropdown.tsx";



export function FileTable() {
    const [dialogOpen, setDialogOpen] = useState(false)
    const [activeFile, setActiveFile] = useState<Metadata | null>(null)

    function openDialog(file: Metadata) {
        setActiveFile(file)
        setDialogOpen(true)
    }

    const userId = useAuthStore((s) => s.userId)

    const cursor = useMemo<CursorReq>(
        () => ({ modified_at: null, id: null }),
        []
    )

    const req = useMemo<GetAllMetadataReq>(
        () => ({
            user_id: userId,
            cursor,
            limit: 20,
        }),
        [userId, cursor]
    )

    const [files, setFiles] = useState<Metadata[]>([])
    const [optimisticFiles, addOptimisticFiles] = useOptimistic<
        Metadata[],
        Metadata[]
    >(files, (state, action) => [...action, ...state])

    const [selected, setSelected] = useState<string[]>([])

    useEffect(() => {
        if (!userId) return
        getAllMetadata(req).then((resp) => setFiles(resp.metadata))
    }, [req, userId])

    const toggleSelect = (id: string) => {
        setSelected((prev) =>
            prev.includes(id) ? prev.filter((x) => x !== id) : [...prev, id]
        )
    }

    const selectAll = (checked: boolean) => {
        setSelected(checked ? optimisticFiles.map((f) => f.uuid) : [])
    }

    return (
        <div>
            <h1 className="text-2xl font-semibold pb-4 pt-4 m-1">All files</h1>

            <nav className="w-full flex gap-3">
                <Button variant="outline">
                    <Clock /> Recents
                </Button>

                <UploadDialog
                    onUploaded={(newFiles) => {
                        if (!newFiles.length) return
                        addOptimisticFiles(newFiles)
                        setFiles((prev) => {
                            return [...newFiles, ...prev];
                        })
                    }}
                />

                <Button variant="outline">
                    <FolderPlus /> New Folder
                </Button>

                <Button variant="outline">
                    <Star /> Favorites
                </Button>
            </nav>

            <Table className="mt-2 w-full table-fixed">
                <TableHeader>
                    <TableRow>
                        <TableHead className="w-[30px]">
                            <Checkbox
                                checked={
                                    optimisticFiles.length > 0 &&
                                    selected.length === optimisticFiles.length
                                }
                                onCheckedChange={(v) => selectAll(v === true)}
                            />
                        </TableHead>
                        <TableHead className="w-[30px]" />
                        <TableHead className={"full"}>Name</TableHead>
                        <TableHead className={"w-[150px]"}>Last Modified</TableHead>
                        <TableHead className="w-[50px]" />
                    </TableRow>
                </TableHeader>

                <TableBody>
                    {optimisticFiles.map((file) => (
                        <TableRow key={file.uuid}>
                            <TableCell>
                                <Checkbox
                                    checked={selected.includes(file.uuid)}
                                    onCheckedChange={() => toggleSelect(file.uuid)}
                                />
                            </TableCell>

                            <TableCell onClick={() => openDialog(file)}>
                                <div className="w-4">
                                    {getIconForFile({ fileName: file.file_name })}
                                </div>
                            </TableCell>

                            <TableCell onClick={() => openDialog(file)}>
                                <p className="truncate whitespace-nowrap overflow-hidden">
                                    {file.file_name}
                                </p>
                            </TableCell>

                            <TableCell onClick={() => openDialog(file)}>
                                <p className="truncate whitespace-nowrap overflow-hidden">
                                    {formatDate(file.modified_at)}
                                </p>
                            </TableCell>

                            <TableCell className={"w-[44px]"}>
                                    <FileDropdown />
                            </TableCell>
                        </TableRow>
                    ))}
                </TableBody>
            </Table>
            {activeFile && (
                <FileDialog
                    open={dialogOpen}
                    onOpenChange={setDialogOpen}
                    metadata={activeFile}
                    ipfsLink=""
                    spaceName=""
                    spaceDid=""
                />
            )}
        </div>
    )
}

function formatDate(date: string): string {
    return new Date(date).toLocaleString()
}
