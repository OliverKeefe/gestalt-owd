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
    const totalPhotos = 1203;
    const availableStorage = 500;
    const segdat: MeterGaugeSegment[] = [
        {
            label: "Photos",
            value: 200,
            color: "bg-yellow-500",
            percentage: 0.25,
        },
        {
            label: "Documents",
            value: 15.6,
            color: "bg-blue-500",
            percentage: 0.156,
        },
        {
            label: "Code Repos",
            value: 3.9,
            color: "bg-green-500",
            percentage: 0.039,
        }
    ];
    return (
        <Card>
            <CardContent className={"w-full flex flex-col"}>
                <CardTitle>Status</CardTitle>
                <CardDescription>Usage information and service health.</CardDescription>
                <span className={"pt-5 w-full"}>
                    <ul className={"flex w-full items-center"}>
                        <li className={"flex-1 text-center items-center"} >

                            <div className="flex items-center w-[100%]">
                                 <MeterGauge
                                     segmentData={segdat}
                                     total={200}
                                     children={undefined}>
                                 </MeterGauge>
                            </div>
                            <CardTitle className={"mt-2"}>Storage {usedStorage} GB</CardTitle>
                            <CardTitle>{usedStorage} used • {availableStorage} available</CardTitle>
                        </li>
                        <li className={"flex-1 text-center"}>
                           <CardTitle>{totalFiles} Files</CardTitle>
                        </li>
                        <li className={"flex-1 text-center"}>
                            <CardTitle>{totalDocuments} Documents</CardTitle>
                        </li>
                        <li className={"flex-1 text-center"}>
                            <CardTitle>{totalPhotos} Photos</CardTitle>
                        </li>
                    </ul>
                </span>
            </CardContent>
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