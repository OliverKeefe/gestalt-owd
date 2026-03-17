import { FileTable } from "@/app/features/files/components/file-table";
import { Container } from "../../shared/components/layout/container";


export function Files() {
    return (
        <div className="flex flex-1 flex-col">
                    <Container>
                        <FileTable />
                    </Container>
            </div>
    );
}