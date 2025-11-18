import {PhotoCard} from "@/app/features/shared/components/cards/photo-card.tsx";
import {MasonryGridGallery} from "@/app/features/shared/components/galleries/masonry-gallery.tsx";

export default function Photos() {
    return (
        <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
            <MasonryGridGallery />
        </div>
    )
}