import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "./../AuthContext";
import { BASE_URL } from "../constants";
import { ROLES } from "../constants";

function Login() {
    const [login, setLogin] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const { login: authLogin, loading } = useAuth();
    const navigate = useNavigate();

    const handleSubmit = async (e) => {
        e.preventDefault();
        setError("");

        const result = await authLogin(login, password);
        
        if (result.success) {
            redirectByRole(result.role, navigate);
        } else {
            setError(result.error || "Неверный логин или пароль");
        }
    };

    return (
        <form onSubmit={handleSubmit}>
            <input
                value={loginInput}
                onChange={(e) => setLoginInput(e.target.value)}
                placeholder="логин"
            />
            <input
                type="password"
                value={password}
                onChange={(e) => setPassword(e.target.value)}
                placeholder="пароль"
            />
            <button type="submit">войти</button>
        </form>
    );
}

function redirectByRole(role_id, navigate) {
    switch (role_id) {
        case ROLES.ADMIN:
            navigate("/admin/dashboard");
            break;
        case ROLES.STUDENT:
            navigate("/student/dashboard");
            break;
        case ROLES.TEACHER:
            navigate("/teacher/dashboard");
            break;
        case ROLES.HEAD:
            navigate("/director/dashboard");
            break;
        default:
            navigate("/");
    }
}

export default Login;
