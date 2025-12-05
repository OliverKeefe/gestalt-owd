import type { TemplateProps } from "keycloakify/login/TemplateProps";

export default function Template(props: TemplateProps<any, any>) {

    const { children } = props;

    return (
        <div
            className="min-h-screen w-full bg-cover bg-center text-foreground"
            style={{ backgroundImage: "url('/login-bg.jpg')" }}
        >
            <main className="flex items-center justify-center p-6">
                {children}
            </main>
        </div>
    );
}
