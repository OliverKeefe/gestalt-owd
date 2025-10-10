import { FileCard } from "../cards/file-card";
import {Carousel, CarouselContent, CarouselItem, CarouselNext, CarouselPrevious} from "@/components/ui/carousel.tsx";
import { Container } from "../layout/container";

const data = {
    files: [
        {
            id: 1,
            name: "Santander Bank Statements 2025",
            lastModified: "",
            icon: "🀄",
            type: "",
            size: "",
            owner: "Owner: John Smith . ",
            access: "Access: Everyone",
            screenshot: "/src/assets/macos-folder-original.png",
        },
        {
            id: 2,
            name: "House Move",
            lastModified: "",
            icon: "📄",
            type: "",
            size: "",
            owner: "Owner: John Smith",
            access: "Access: Only You",
            screenshot: "/src/assets/macos-folder-original.png",
        },
    ]
}

export function FileCardCarousel() {
    return (
        <Container>
            <Carousel className="relative">
                <CarouselContent className="-ml-1">
                    {data.files.map((file) => (
                        <CarouselItem
                            key={file.id}
                            className="pl-1 basis-full sm:basis-1/2 lg:basis-1/4 xl:basis-1/5"
                        >
                            <div className="p-1">
                                <FileCard {...file} />
                            </div>
                        </CarouselItem>
                    ))}
                </CarouselContent>
                <CarouselPrevious className="-left-6" />
                <CarouselNext className="-right-6" />
            </Carousel>
        </Container>
    );
}


