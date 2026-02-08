import { Routes, Route } from "react-router-dom";
import Login from "@/app/features/auth/pages/login-page";
import { Files } from "@/app/features/files/pages/files-page";
import Documents from "@/app/features/documents/pages/documents-page.tsx";
import Photos from "@/app/features/photos/pages/photos-page";
//import Settings from "@/app/features/settings/pages/settings-page";

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<Files />} />
            <Route path="/login" element={<Login />} />
            <Route path="/files" element={<Files />} />
            <Route path="/documents" element={<Documents />} />
            <Route path="/photos" element={<Photos />} />
            {/*<Route path="/settings" element={<Settings />} />*/}
        </Routes>
    );
};

export default AppRoutes;