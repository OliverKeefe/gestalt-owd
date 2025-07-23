import {
    Card,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { FilePlus, FolderPlus, Share, Upload } from "lucide-react"

export function SectionCards() {
    return (
        <div className=" grid grid-cols-1 gap-4 px-4 *:data-[slot=card]:bg-gradient-to-t *:data-[slot=card]:shadow-xs lg:px-6 @xl/main:grid-cols-2 @5xl/main:grid-cols-6">
            <Card className="@container/card max-h-34">
                <CardHeader>
                    <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
                        <Upload className="w-5 h-5" />
                        <p>Upload files</p>
                    </CardTitle>
                </CardHeader>
            </Card>
            <Card className="@container/card max-h-34">
                <CardHeader>
                    <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
                        <FilePlus className="w-5 h-5" />
                        <p>New File</p>
                    </CardTitle>
                </CardHeader>
            </Card>
            <Card className="@container/card max-h-34">
                <CardHeader>
                    <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
                        <FolderPlus className="w-5 h-5" />
                        <p>Create folder</p>
                    </CardTitle>
                </CardHeader>
            </Card>
            <Card className="@container/card max-h-34">
                <CardHeader>
                    <CardTitle className="text-2xl font-semibold tabular-nums @[250px]/card:text-3xl">
                        <Share className="w-5 h-5" />
                        <p>Share</p>
                    </CardTitle>
                </CardHeader>
            </Card>
        </div>
    )
}
