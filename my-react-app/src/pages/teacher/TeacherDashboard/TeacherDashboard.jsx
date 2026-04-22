import { useNavigate } from "react-router-dom";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./TeacherDashboard.module.css";

function TeacherDashboard() {
  const navigate = useNavigate();

  return (
    <div className={styles.page}>
      <div className={styles.container}>
        <div className={styles.header}>
          <LogoutButton />
        </div>

        <h1 className={styles.title}>Панель преподавателя</h1>
        <div className={styles.subtitle}>
          Управляйте классом, смотрите статистику и успехи учеников
        </div>

        <div className={styles.cardsGrid}>
          <div
            className={styles.actionCard}
            onClick={() => navigate("/teacher/class-statistics")}
          >
            <div className={styles.cardIcon}></div>
            <div className={styles.cardTitle}>Статистика класса</div>
            <div className={styles.cardDesc}>
              Общая успеваемость, прогресс по темам
            </div>
          </div>

          <div
            className={styles.actionCard}
            onClick={() => navigate("/teacher/students")}
          >
            <div className={styles.cardIcon}></div>
            <div className={styles.cardTitle}>Список учеников</div>
            <div className={styles.cardDesc}>
              Просмотр и управление учениками
            </div>
          </div>
        </div>

        <div className={styles.footer}>
          {/* Можно добавить дополнительную информацию или кнопку "Назад" */}
        </div>
      </div>
    </div>
  );
}

export default TeacherDashboard;
