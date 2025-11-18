import {
    Card, CardContent,
    CardHeader,
    CardTitle,
} from "@/components/ui/card"
import { FilePlus, FolderPlus, Share, Upload } from "lucide-react"

interface FileCardProps {
    id: number;
    name: string;
    lastModified: string;
    icon: string;
    type: string;
    size: string;
    owner: string;
    access: string;
    screenshot: string;
}

export function FileCard({ id, name, lastModified, icon, type, size, owner, access, screenshot }: FileCardProps) {
    const placeholder: string = "https://placehold.co/600x400/000000/FFFFFF/png"


    return (
        <Card className={"min-w-[250px] min-h-[400px] max-h-[400px]"}>
            <CardHeader>{icon} {name}{type}</CardHeader>
            <img
                src={screenshot || placeholder}
                alt={`${name} preview`}
                onError={(err) => {
                    err.currentTarget.src = placeholder;
                }}
                className={"w-full h-[200px] object-cover rounded-md"}
            />
            <CardContent>
                <div>
                    <p>{lastModified}</p>
                    <p>{size}</p>
                    <p>{owner}</p>
                    <p>{access}</p>
                </div>
            </CardContent>
        </Card>
    );
}