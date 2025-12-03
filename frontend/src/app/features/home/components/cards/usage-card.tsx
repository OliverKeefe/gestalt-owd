import {Card, CardContent, CardDescription, CardFooter, CardTitle} from '@/components/ui/card';

interface UsageData {
    availableStorage: number;
    usedStorage: number;
    totalFiles: number;
    totalDocuments: number;
    totalPhotos: number;
    totalVideos: number;
    totalAudio: number;
    totalBooks: number;
    serviceHealth: string;
    pins: number;
    blocks: number;
}

export function UsageCard() {
    const usedStorage = 200;
    const totalFiles = 354;
    const totalDocuments = 243;
    return (
        <Card>
            <CardContent className={"w-full flex flex-col"}>
                <CardTitle>Status</CardTitle>
                <CardDescription>Usage information and service health.</CardDescription>
                <span className={"pt-5 w-full"}>
                    <ul className={"flex w-full items-center"}>
                        <li className={"flex-1 text-center"} >Storage Used: {usedStorage} GB</li>
                        <li className={"flex-1 text-center"}>Files: {totalFiles}</li>
                        <li className={"flex-1 text-center"}>Documents: {totalDocuments}</li>
                        <li className={"flex-1 text-center"}>Documents: {totalDocuments}</li>
                    </ul>
                </span>
            </CardContent>
            <CardFooter>Yo mama</CardFooter>
        </Card>
    );
}

function StorageUse() {
    return (
        <div>
            <h3>Storage</h3>
            <p></p>
        </div>
    );
}