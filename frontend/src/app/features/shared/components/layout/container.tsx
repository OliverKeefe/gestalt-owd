import React from "react";
import { cn } from "@/lib/utils.ts";

interface ContainerProps extends React.HTMLAttributes<HTMLDivElement> {

}

export const Container = ({ className, ...props }: ContainerProps) => {
    return (
        <div
            className={cn("mx-auto w-full max-w-[100rem] px-4 sm:px-6 lg:px-8", className)}
            {...props}
        />
    );
}