import {Card, CardContent} from "@/components/ui/card";
import { React } from 'react';

export function PhotoCard() {
    const placeholder: string = "https://placehold.co/600x400/000000/FFFFFF/png";

    return (
        <Card>
            <CardContent>
                <img src={"https://placehold.co/600x400/000000/FFFFFF/png" || placeholder} />
                <p>Photo title</p>
            </CardContent>
        </Card>
    );
}