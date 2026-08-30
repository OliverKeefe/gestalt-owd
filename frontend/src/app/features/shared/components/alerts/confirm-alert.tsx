import {
    AlertDialog,
    AlertDialogAction,
    AlertDialogCancel,
    AlertDialogContent,
    AlertDialogDescription,
    AlertDialogFooter,
    AlertDialogHeader,
    AlertDialogTitle,
    AlertDialogTrigger,
} from "@/components/ui/alert-dialog"

interface ConfirmAlertDialogProps {
    onConfirm: () => void;
    message: string;
    icon: React.ComponentType<{ className?: string }>;
    children: React.ReactNode;
}

export function ConfirmAlertDialog({
    onConfirm,
    message,
    icon,
    children,
}: ConfirmAlertDialogProps) {
    const Icon = icon;

    return (
        <AlertDialog>
            <AlertDialogTrigger asChild>
                {children}
            </AlertDialogTrigger>
            <AlertDialogContent>
                <AlertDialogHeader>
                    <AlertDialogTitle>Are you sure?</AlertDialogTitle>
                    <AlertDialogDescription>
                        {message}
                    </AlertDialogDescription>
                </AlertDialogHeader>
                <AlertDialogFooter>
                    <AlertDialogCancel className={"cursor-pointer"}>Cancel</AlertDialogCancel>
                    <AlertDialogAction className={"cursor-pointer"} onClick={onConfirm}>
                        <Icon className="h-4 w-4" /> Delete
                    </AlertDialogAction>
                </AlertDialogFooter>
            </AlertDialogContent>
        </AlertDialog>
    )
}
