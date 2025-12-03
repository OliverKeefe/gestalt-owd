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

export function MeterGauge({ segmentData, total, children }: MeterGaugeProps) {
    const processedSegments = segmentData.map(seg => ({
        ...seg,
        percentage: (seg.value / total) * 100,
    }));

    return (
        <Background>
            {processedSegments.map((seg, i) => (
                <Segment
                    key={i}
                    label={seg.label}
                    value={seg.value}
                    color={seg.color}
                />
            ))}
        </Background>
    );
}

function Segment( { label, value, color}: MeterGaugeSegment ) {
    return (
        <div
            className={`rounded-b-full h[10px] ${color}`}>
            title={label}
            value={value}
        </div>
    );
}

type BackgroundProps = React.PropsWithChildren<{
    children: React.ReactNode | undefined;
}>;

function Background({ children }: BackgroundProps) {
    return (
        <div className={"rounded-b-full h[10px] w-full"}>{children}</div>
    );
}

