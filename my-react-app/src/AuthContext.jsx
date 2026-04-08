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

    const checkSession = async () => {
        try {
            const response = await fetch(`${BASE_URL}/auth/session`, {
                method: "GET",
                credentials: "include",
            });

            if (!response.ok) {
                setUser(null);
                return false;
            }

            const sessionData = await response.json();
            
            const userData = {
                id: sessionData.user_id,
                role: sessionData.role,
                sessionId: sessionData.session_id
            };
            
            setUser(userData);
            return true;
        } catch (err) {
            console.error("Session check failed:", err);
            setUser(null);
            return false;
        } finally {
            setLoading(false);
        }
    };

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
            
            // После успешного входа проверяем сессию
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
            const response = await fetch(`${BASE_URL}/auth/logout`, {
                method: "POST",
                credentials: "include",
            });

            if (!response.ok) {
                console.error("Logout failed");
            }
        } catch (err) {
            console.error("Logout error:", err);
        } finally {
            setUser(null);
            window.location.href = "/login";
        }
    };

    useEffect(() => {
        checkSession();
    }, []);

    const value = {
        user,
        loading,
        error,
        login,
        logout,
        checkSession,
        isAuthenticated: !!user,
    };

    return (
        <AuthContext.Provider value={value}>
            {children}
        </AuthContext.Provider>
    );
};