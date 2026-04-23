import { useNavigate } from "react-router-dom";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../styles/shared.module.css";

function AdminDashboard() {
  const navigate = useNavigate();

  return (
    <div className={sharedStyles.dashboardPage}>
      <div className={sharedStyles.dashboardContainer}>
        <div className={sharedStyles.dashboardHeader}>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.dashboardTitle}>Административная панель</h1>
        <div className={sharedStyles.dashboardSubtitle}>
          Управление школами, учителями, классами и пользователями
        </div>

        <div className={sharedStyles.dashboardCardsGrid}>
          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/schools")}
          >
            <div className={sharedStyles.dashboardCardIcon}>🏫</div>
            <div className={sharedStyles.dashboardCardTitle}>Школы</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Добавление, редактирование, удаление школ
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/teachers")}
          >
            <div className={sharedStyles.dashboardCardIcon}>👩‍🏫</div>
            <div className={sharedStyles.dashboardCardTitle}>Учителя</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Управление учителями и их данными
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/classes")}
          >
            <div className={sharedStyles.dashboardCardIcon}>📚</div>
            <div className={sharedStyles.dashboardCardTitle}>Классы</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Создание и настройка классов
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/users")}
          >
            <div className={sharedStyles.dashboardCardIcon}>👤</div>
            <div className={sharedStyles.dashboardCardTitle}>Пользователи</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Управление пользователями системы
            </div>
          </div>
        </div>

        <div className={sharedStyles.dashboardFooter} />
      </div>
    </div>
  );
}

export default AdminDashboard;