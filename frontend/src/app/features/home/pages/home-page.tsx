import { FileTable } from "@/app/features/shared/components/tables/file-table.tsx";
import { Container } from "../../shared/components/layout/container";


export function Home() {
    return (
        <div className="flex flex-1 flex-col">
                    <Container>
                        <FileTable />
                    </Container>
            </div>
    );
}