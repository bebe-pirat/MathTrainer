import { useNavigate } from "react-router-dom";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../../styles/shared.module.css";
import { ReactComponent as ChartIcon } from "../../../assets/chart-column.svg"
import { ReactComponent as ClassIcon } from "../../../assets/user-list.svg"

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
            <ChartIcon className={sharedStyles.iconBlue}/>
            <div className={sharedStyles.dashboardCardTitle}>Статистика класса</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Общая успеваемость, прогресс по темам
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/teacher/students")}
          >
            <ClassIcon className={sharedStyles.iconBlue}/>
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