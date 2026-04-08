import { useState, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../AuthContext";
import { ROLES } from "../constants";

function Login() {
    const [login, setLogin] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const { login: authLogin, loading, user, isAuthenticated } = useAuth();
    const navigate = useNavigate();

    // Редирект, если пользователь уже авторизован
    useEffect(() => {
        console.log("Login page - isAuthenticated:", isAuthenticated);
        console.log("Login page - user:", user);
        
        if (isAuthenticated && user) {
            console.log("User already has session, redirecting...");
            redirectByRole(user.role, navigate);
        }
    }, [isAuthenticated, user, navigate]);

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

    // Показываем форму только если нет пользователя
    if (isAuthenticated && user) {   
        return <div>Перенаправление...</div>;
    }

    return (
        <div>
            <h2>Вход в систему</h2>
            {error && <div style={{ color: "red" }}>{error}</div>}
            <form onSubmit={handleSubmit}>
                <div>
                    <input
                        type="text"
                        value={login}
                        onChange={(e) => setLogin(e.target.value)}
                        placeholder="Логин"
                        required
                    />
                </div>
                <div>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        placeholder="Пароль"
                        required
                    />
                </div>
                <button type="submit" disabled={loading}>
                    {loading ? "Вход..." : "Войти"}
                </button>
            </form>
        </div>
    );
}

function redirectByRole(role_id, navigate) {
    console.log("Redirecting with role_id:", role_id);
    console.log("ROLES.ADMIN:", ROLES.ADMIN);
    console.log("ROLES.TEACHER:", ROLES.TEACHER);
    
    switch (role_id) {
        case ROLES.ADMIN:
            navigate("/admin/dashboard");
            break;
        case ROLES.TEACHER:
            navigate("/teacher/dashboard");
            break;
        case ROLES.STUDENT:
            navigate("/student/dashboard");
            break;
        case ROLES.HEAD:
            navigate("/director/dashboard");
            break;
        default:
            console.log("Unknown role, going to login");
            navigate("/login");
    }
}

export default Login;