import {SectionCards} from "@/app/features/shared/components/cards/section-cards.tsx";
import {Table} from "@/components/ui/table.tsx";
import {FileTable} from "@/app/features/shared/components/tables/file-table.tsx";
import { Container } from "../../shared/components/layout/container";
import { FileCardCarousel } from "../../shared/components/carousel/file-card-carousel";


export function Home() {
    return (
        <div className="flex flex-1 flex-col">
                    <Container>
                        <h1 className="text-2xl font-semibold p-6 m-1">Folders</h1>
                        <FileCardCarousel />
                    </Container>
                    <Container>
                        <h1 className="text-2xl font-semibold p-6 m-1">All files</h1>
                        <FileTable />
                    </Container>
            </div>
    );
}