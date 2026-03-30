import { useNavigate } from "react-router-dom";
import LogoutButton from "./../components/LogoutButton";

function AdminDashboard() {
    const navigate = useNavigate();

    return (
        <div>
            <h1>Админ панель</h1>

            <div>
                <button onClick={() => navigate("/admin/schools")}>
                    Управление школами
                </button>
            </div>

            <div>
                <button onClick={() => navigate("/admin/teachers")}>
                    Управление учителями
                </button>
            </div>

            <div>
                <button onClick={() => navigate("/admin/classes")}>
                    Управление классами
                </button>
            </div>

            <div>
                <button onClick={() => navigate("/admin/users")}>
                    Управление пользователями
                </button>
            </div>

            <LogoutButton />
        </div>
    );
}

export default AdminDashboard;
