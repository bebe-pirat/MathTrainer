import { useNavigate } from "react-router-dom";
import LogoutButton from "../../components/LogoutButton/LogoutButton";

function TeacherDashboard() {
    const navigate = useNavigate();

    return (
        <div>
            <h1>Панель преподавателя</h1>

            <div>
                <button onClick={() => navigate("/teacher/class-statistics")}>
                    Статистика класса
                </button>
            </div>

            <div>
                <button onClick={() => navigate("/teacher/students")}>
                    Список учеников
                </button>
            </div>

            <LogoutButton />
        </div>
    );
}

export default TeacherDashboard;