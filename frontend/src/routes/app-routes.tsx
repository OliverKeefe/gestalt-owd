import { Routes, Route } from "react-router-dom";
import Login from "@/app/features/auth/pages/login-page";
import { Home } from "@/app/features/home/pages/home-page";

const AppRoutes = () => {
    return (
        <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/login" element={<Login />} />
        </Routes>
    );
};

export default AppRoutes;