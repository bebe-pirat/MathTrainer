import { createContext, useContext, useState, useEffect } from "react";
import { BASE_URL } from "./constants";

const AuthContext = createContext();

export function AuthProvider({ children }) {
    const [user, setUser] = useState(null); 

    useEffect(() => {
        const storedUser = localStorage.getItem("user");
        if (storedUser ) {
            setUser(JSON.parse(storedUser));
        }
    }, []);

    const login = (data) => {
        const userData = {
            user_id: data.user_id,
            role_id: data.role, 
        };

        setUser(userData);
        localStorage.setItem("user", JSON.stringify(userData));
    };

    const logout = async () => {
        try {
            await fetch(BASE_URL + "/auth/logout", {
                method: "POST", 
                credentials: "include",
            })
        }
        catch (e) {
            console.error(e);
        }

        setUser(null);
        localStorage.removeItem("user");
    };

    return (
        <AuthContext.Provider value={{ user, login, logout }}>
            {children}
        </AuthContext.Provider>
    );
}

export function useAuth() {
    return useContext(AuthContext)
}