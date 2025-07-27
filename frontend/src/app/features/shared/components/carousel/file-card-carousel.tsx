import { FileCard } from "../cards/file-card";
import {Carousel, CarouselContent, CarouselItem} from "@/components/ui/carousel.tsx";

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
            screenshot: "https://placehold.co/600x600",
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
            screenshot: "https://placehold.co/600x600",
        },
    ]
}

export function FileCardCarousel() {
    return (
        <Carousel>
            <CarouselContent className="gap-1" >
                {data.files.map(file => (
                    <CarouselItem
                        key={file.id}
                        className="basis-full sm:basis-1/2 lg:basis-1/4 xl:basis-1/5"
                    >
                        <FileCard {...file} />
                    </CarouselItem>
                ))}
            </CarouselContent>
        </Carousel>
    );
}

