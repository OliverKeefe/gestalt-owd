import { Table, TableBody, TableCaption, TableCell, TableHead, TableHeader, TableRow } from "@/components/ui/table";

export function FileTable() {
    return (
        <Table>
            <TableCaption>...</TableCaption>
            <TableHeader>
                <TableRow>
                    <TableHead className="w-[100px]">Last Modified</TableHead>
                    <TableHead>Access</TableHead>
                    <TableHead>Name</TableHead>
                    <TableHead className="text-right">Type</TableHead>
                </TableRow>
            </TableHeader>
            <TableBody>
                <TableRow>
                    <TableCell className="font-medium">10-01-25</TableCell>
                    <TableCell>Only You</TableCell>
                    <TableCell>📄 Santander Bank Statement Oct 24.pdf</TableCell>
                    <TableCell className="text-right">.PDF</TableCell>
                </TableRow>
                <TableRow>
                    <TableCell className="font-medium">11-01-25</TableCell>
                    <TableCell>Only You</TableCell>
                    <TableCell>📝 CV.docx</TableCell>
                    <TableCell className="text-right">.DOCX</TableCell>
                </TableRow>
                <TableRow>
                    <TableCell className="font-medium">10-01-25</TableCell>
                    <TableCell>Only You</TableCell>
                    <TableCell>📝 Untitled1.pdf</TableCell>
                    <TableCell className="text-right">.DOCX</TableCell>
                </TableRow>
                <TableRow>
                    <TableCell className="font-medium">10-01-25</TableCell>
                    <TableCell>Everyone</TableCell>
                    <TableCell>🖼️ picture_of_cat.jpeg</TableCell>
                    <TableCell className="text-right">.JPEG</TableCell>
                </TableRow>
            </TableBody>
        </Table>
    );
}