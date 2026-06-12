import { useNavigate } from "react-router-dom";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../styles/shared.module.css";
import { ReactComponent as SchoolIcon } from "../../assets/building.svg"
import { ReactComponent as GradeIcon } from "../../assets/building-user.svg"
import { ReactComponent as ClassIcon } from "../../assets/user-list.svg"

function DirectorDashboard() {
  const navigate = useNavigate();

  return (
    <div className={sharedStyles.dashboardPage}>
      <div className={sharedStyles.dashboardContainer}>
        <div className={sharedStyles.dashboardHeader}>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.dashboardTitle}>Панель завуча</h1>
        <div className={sharedStyles.dashboardSubtitle}>
          Просмотр статистики по разным уровням: школа, параллель, класс и ученик 
        </div>

        <div className={sharedStyles.dashboardCardsGrid}>
            { /* Школа */}
          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/director/school-stats")}
          >
            <SchoolIcon className={sharedStyles.iconBlue}/>
            <div className={sharedStyles.dashboardCardTitle}>Статистика школы</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Просмотр общей статистики по школе
            </div>
          </div>

            { /* Класс */}
           <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/director/classes")}
          >
            <ClassIcon className={sharedStyles.iconBlue}/>
            <div className={sharedStyles.dashboardCardTitle}>Статистика классов и учеников</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Просмотр статистики по классам и ученикам
            </div>
          </div>

        </div>
        <div className={sharedStyles.dashboardFooter} />
      </div>
    </div>
  );
}

export default DirectorDashboard;