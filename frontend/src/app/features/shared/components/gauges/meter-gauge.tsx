export type MeterGaugeSegment = {
    label: string;
    value: number;
    color: string;
    percentage?: number;
}

type MeterGaugeProps = React.PropsWithChildren<{
    segmentData: MeterGaugeSegment[];
    total: number;
    children: React.ReactNode | undefined;
}>;


type BackgroundProps = React.PropsWithChildren<{
    children: React.ReactNode | undefined;
}>;

function Background({ children }: BackgroundProps) {
    return (
        <div className={"rounded-b-full h[10px] w-full"}>{children}</div>
    );
}

