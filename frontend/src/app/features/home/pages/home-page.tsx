import { FileTable } from "@/app/features/shared/components/tables/file-table.tsx";
import { Container } from "../../shared/components/layout/container";


export function Home() {
    return (
        <div className="flex flex-1 flex-col">
                    <Container>
                        <h1 className="text-2xl font-semibold p-6 m-1">All files </h1>
                        <FileTable />
                    </Container>
            </div>
    );
}