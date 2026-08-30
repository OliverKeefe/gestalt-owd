import {Card, CardContent} from "@/components/ui/card";

export function PhotoCard() {
    return (
        <Card>
            <CardContent>
                <img src={"https://placehold.co/600x400/000000/FFFFFF/png"} />
                <p>Photo title</p>
            </CardContent>
        </Card>
    );
}