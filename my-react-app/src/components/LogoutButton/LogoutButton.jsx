import { useAuth } from "../../AuthContext";
import { useNavigate } from "react-router-dom";
import styles from './LogoutButton.module.css';

function LogoutButton() {
    const {logout} = useAuth();
    const navigate = useNavigate();

    const handleLogout = async () => {
        await logout();
        navigate("/login");
    };

    return <button onClick={handleLogout} className={styles.LogoutButton} >выйти</button>;
}

export default LogoutButton