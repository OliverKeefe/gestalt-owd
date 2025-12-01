import path from "path"
import tailwindcss from '@tailwindcss/vite'
import react from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import { keycloakify } from "keycloakify/vite-plugin";

export default defineConfig({
    plugins: [
        react(),
        tailwindcss(),
        keycloakify({
            accountThemeImplementation: "none",
        })
    ],
    resolve: {
        alias: {
            "@": path.resolve(__dirname, "./src"),
        },
    },
});