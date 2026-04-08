import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "./../AuthContext";
import { BASE_URL } from "../constants";
import { ROLES } from "../constants";

function Login() {
    const [loginInput, setLoginInput] = useState("");
    const [password, setPassword] = useState("");
    const navigate = useNavigate();
    const { login } = useAuth();

    const handleSubmit = async (e) => {
        e.preventDefault();

        const response = await fetch(BASE_URL + "/auth/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                login: loginInput,
                password: password,
            }),
        });

        if (!response.ok) {
            alert("Ошибка входа");
            return;
        }

        const data = await response.json();

        login(data); 

        redirectByRole(data.role, navigate);
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
