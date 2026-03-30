import { useAuth } from "../AuthContext";
import { useNavigate } from "react-router-dom";

function LogoutButton() {
    const {logout} = useAuth();
    const navigate = useNavigate();

    const handleLogout = async () => {
        await logout();
        navigate("/login");
    };

    return <button onClick={handleLogout}>выйти</button>;
}

export default LogoutButton