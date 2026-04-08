import { createContext, useContext, useState, useEffect } from "react";
import { BASE_URL } from "./constants";

const AuthContext = createContext();

export function AuthProvider({ children }) {
    const [user, setUser] = useState(null);
    const [loading, setLoading] = useState(true); 

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
                    role_id: Number(data.role),
                });
            } catch (err) {
                console.error(err);
                setUser(null);
            }

            setLoading(false);
        };

        checkSession();
    }, []);

    const login = (data) => {
        setUser({
            user_id: Number(data.user_id),
            role_id: Number(data.role),
        });
    };

    const logout = async () => {
        try {
            await fetch(BASE_URL + "/auth/logout", {
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
}

export function useAuth() {
    return useContext(AuthContext);
}
