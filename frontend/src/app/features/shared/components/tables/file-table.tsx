import { Card } from "@/components/ui/card";
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";
import { useState } from "react";
import { Checkbox } from "@/components/ui/checkbox";

const data = {
    files: [
        {
            id: 1,
            name: "Santander Bank Statement 2025",
            lastModified: "10-01-25",
            icon: "🀄",
            type: ".PDF",
            size: "1.52 MB",
            owner: "Steve Smith",
            access: "Everyone",
        },
        {
            id: 2,
            name: "CV",
            lastModified: "10-01-25",
            icon: "📄",
            type: ".DOCX",
            size: "7.9 MB",
            owner: "Steve Smith",
            access: "Only You",
        },
    ]
}

export function FileTable() {
    const [selected, setSelected] = useState<number[]>([]);

    const toggleSelect = (id: number) => {
        setSelected((prev) =>
            prev.includes(id) ? prev.filter((item) => item !== id) : [...prev, id]
        );
    };

    return (
            <Card className="pt-0">
                <Table>
                    <TableHeader>
                        <TableRow>
                            <TableHead className="w-[30px]">
                                <Checkbox
                                    checked={selected.length === data.files.length}
                                    onCheckedChange={(checked) => {
                                        if (checked) {
                                            setSelected(data.files.map((file) => file.id));
                                        } else {
                                            setSelected([]);
                                        }
                                    }}
                                />
                            </TableHead>
                            <TableHead className="w-[100px]">Last Modified</TableHead>
                            <TableHead>Access</TableHead>
                            <TableHead>Name</TableHead>
                            <TableHead className="text-right">Type</TableHead>
                        </TableRow>
                    </TableHeader>
                    <TableBody>
                        {data.files.map((file) => (
                            <TableRow key={file.id}>
                                <TableCell>
                                    <Checkbox
                                        checked={selected.includes(file.id)}
                                        onCheckedChange={() => toggleSelect(file.id)}
                                    />
                                </TableCell>
                                <TableCell className="font-medium">{file.lastModified}</TableCell>
                                <TableCell>{file.access}</TableCell>
                                <TableCell>{file.icon} {file.name}</TableCell>
                                <TableCell className="text-right">{file.type}</TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            </Card>
    );
}