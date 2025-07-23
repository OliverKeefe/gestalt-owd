import {SectionCards} from "@/app/features/shared/components/cards/section-cards.tsx";
import {Table} from "@/components/ui/table.tsx";
import {FileTable} from "@/app/features/shared/components/tables/file-table.tsx";
import { Container } from "../../shared/components/layout/container";


export function Home() {
    return (
        <div className="flex flex-1 flex-col">
            <div className="@container/main flex flex-1 flex-col gap-2">
                <div className="flex flex-col gap-4 py-4 md:gap-6 md:py-6">
                    <h1 className="text-2xl font-semibold p-6 m-1">Actions</h1>
                    <SectionCards />
                </div>
                <div>
                    <h1 className="text-2xl font-semibold p-6 m-1">All files</h1>
                    <Container>
                        <FileTable />
                    </Container>

                </div>
            </div>
        </div>
    );
}