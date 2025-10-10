import { Routes, Route } from "react-router-dom";
import Login from "@/app/features/auth/pages/login-page";
import { Home } from "@/app/features/home/pages/home-page";
import Documents from "@/app/features/documents/pages/documents-page.tsx";
import Photos from "@/app/features/photos/pages/photos-page";

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
            <Route path="/home" element={<Home />} />
            <Route path="/documents" element={<Documents />} />
            <Route path="/photos" element={<Photos />} />
        </Routes>
    );
};

export default AppRoutes;