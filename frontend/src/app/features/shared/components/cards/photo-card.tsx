import {Card, CardContent} from "@/components/ui/card";
import { React } from 'react';

export function PhotoCard() {

    return (
        <Card>
            <CardContent>
                <img src={"https://placehold.co/600x400/000000/FFFFFF/png"} />
            </CardContent>
        </Card>
    );
}