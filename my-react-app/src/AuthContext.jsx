import { createContext, useContext, useState, useEffect } from "react";
import { BASE_URL } from "./constants";

const AuthContext = createContext(null);

export const useAuth = () => {
    const context = useContext(AuthContext);
    if (!context) {
        throw new Error("useAuth must be used within AuthProvider");
    }
    return context;
};

export const AuthProvider = ({ children }) => {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const checkSession = async () => {
            try {
                const response = await fetch(BASE_URL + "/auth/session", {
                    method: "GET",
                    credentials: "include",
                });

                if (!response.ok) {
                    setUser(null);
                    setLoading(false);
                    return;
                }

                const data = await response.json();

                setUser({
                    user_id: Number(data.user_id),
                    role_id: Number(data.role_id),
                });
            } catch (err) {
                console.error(err);
                setUser(null);
            }

            setLoading(false);
        };

        checkSession();
    }, []);

    // Функция входа
    const login = async (login, password) => {
        setLoading(true);
        setError(null);
        
        try {
            const response = await fetch(`${BASE_URL}/auth/login`, {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({ login, password }),
            });

            if (!response.ok) {
                const errorData = await response.json();
                throw new Error(errorData.message || "Ошибка входа");
            }

            const data = await response.json();
            
            await checkSession();
            
            return { success: true, role: data.role };
        } catch (err) {
            setError(err.message);
            return { success: false, error: err.message };
        } finally {
            setLoading(false);
        }
    };

    // Функция выхода
     const logout = async () => {
        try {
            await fetch("/auth/logout", {
                method: "POST",
                credentials: "include",
            });
        } catch (e) {
            console.error(e);
        }

        setUser(null);
    };

    return (
        <AuthContext.Provider value={{ user, login, logout, loading }}>
            {children}
        </AuthContext.Provider>
    );
};

export function useAuth() {
    return useContext(AuthContext);
}
