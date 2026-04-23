import { useNavigate } from "react-router-dom";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../../styles/shared.module.css";

function TeacherDashboard() {
  const navigate = useNavigate();

  return (
    <div className={sharedStyles.dashboardPage}>
      <div className={sharedStyles.dashboardContainer}>
        <div className={sharedStyles.dashboardHeader}>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.dashboardTitle}>Панель преподавателя</h1>
        <div className={sharedStyles.dashboardSubtitle}>
          Управляйте классом, смотрите статистику и успехи учеников
        </div>

        <div className={sharedStyles.dashboardCardsGrid}>
          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/teacher/class-statistics")}
          >
            <div className={sharedStyles.dashboardCardIcon}>📊</div>
            <div className={sharedStyles.dashboardCardTitle}>Статистика класса</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Общая успеваемость, прогресс по темам
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/teacher/students")}
          >
            <div className={sharedStyles.dashboardCardIcon}>👥</div>
            <div className={sharedStyles.dashboardCardTitle}>Список учеников</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Просмотр и управление учениками
            </div>
          </div>
        </div>

        <div className={sharedStyles.dashboardFooter} />
      </div>
    </div>
  );
}

export default TeacherDashboard;