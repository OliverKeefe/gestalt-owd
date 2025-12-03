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

