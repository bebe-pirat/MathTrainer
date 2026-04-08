import { useNavigate } from "react-router-dom";
import { useAuth } from "./../AuthContext";
import { useEffect } from "react";
import { ROLES } from "../constants";

function HomePage() {
    const { user, loading } = useAuth();
    const navigate = useNavigate();

    useEffect(() => {
        if (loading) return;
        
        if (!user) {
            return;
        }

        console.log("HomePage - user role:", user.role_id);
        
        switch (user.role_id) {  
            case ROLES.ADMIN: 
                navigate("/admin/dashboard", { replace: true });
                break;
            case ROLES.TEACHER:
                navigate("/teacher/dashboard", { replace: true });
                break;
            case ROLES.STUDENT:
                navigate("/student/dashboard", { replace: true });
                break;
            case ROLES.HEAD:  
                navigate("/director/dashboard", { replace: true });
                break;
            default:
                console.log("Unknown role:", user.role);
        }
    }, [user, loading, navigate]);

    if (loading) {
        return <div>Loading...</div>;
    }

    if (!user) {
        return (
            <div>
                <h1>Добро пожаловать</h1>
                <button onClick={() => navigate("/login")}>
                    Войти
                </button>
            </div>
        );
    }

    return <div>Перенаправление...</div>;
}

export default HomePage;