import {PhotoCard} from "@/app/features/shared/components/cards/photo-card.tsx";

export default function Photos() {
    return (
        <div className="flex min-h-svh w-full items-center justify-center p-6 md:p-10">
            <div className="w-full max-w-sm">
                <p>Photos Page Route Success</p>
                <PhotoCard />
            </div>
        </div>
    )
}