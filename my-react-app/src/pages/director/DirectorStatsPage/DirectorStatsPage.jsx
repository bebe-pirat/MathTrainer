import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { BASE_URL } from "../../constants";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../styles/shared.module.css";
import styles from "./DirectorSchoolStats.module.css";

function DirectorSchoolStats() {
  const navigate = useNavigate();
  const { schoolId } = useParams(); 
  const [stats, setStats] = useState(null);
  const [classes, setClasses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  const fetchSchoolStats = async () => {
    try {
      const response = await fetch(`${BASE_URL}/director/school-stats/${schoolId}`, {
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка загрузки статистики школы");
      const data = await response.json();
      setStats(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    }
  };

  const fetchClasses = async () => {
    try {
      const response = await fetch(`${BASE_URL}/director/classes?school_id=${schoolId}`, {
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка загрузки списка классов");
      const data = await response.json();
      setClasses(data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    if (!schoolId) {
      setError("Не указан идентификатор школы");
      setLoading(false);
      return;
    }
    Promise.all([fetchSchoolStats(), fetchClasses()]).finally(() => setLoading(false));
  }, [schoolId]);

  const handleClassClick = (classId) => {
    navigate(`/director/class-stats/${classId}`);
  };

  if (loading) return <div className={sharedStyles.loader}>Загрузка статистики...</div>;
  if (error)
    return (
      <div className={sharedStyles.errorBox}>
        <p>Ошибка: {error}</p>
        <button className={sharedStyles.formButton} onClick={() => navigate("/director/dashboard")}>
          Назад
        </button>
      </div>
    );
  if (!stats) return <div className={sharedStyles.emptyMessage}>Нет данных</div>;

  const classList = stats.classes && stats.classes.length ? stats.classes : classes;

  return (
    <div className={sharedStyles.dashboardPage}>
      <div className={sharedStyles.dashboardContainer}>
        <div className={sharedStyles.dashboardHeader}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/director/dashboard")}
          >
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.dashboardTitle}>Статистика школы</h1>

        <div className={sharedStyles.formCard}>
          <h3>Общие показатели</h3>
          <div className={sharedStyles.statsGrid}>
            <div className={sharedStyles.statItem}>
              <div className={sharedStyles.statLabel}>Учеников</div>
              <div className={sharedStyles.statValue}>{stats.students_count}</div>
            </div>
            <div className={sharedStyles.statItem}>
              <div className={sharedStyles.statLabel}>Классов</div>
              <div className={sharedStyles.statValue}>{stats.classes_count}</div>
            </div>
            <div className={sharedStyles.statItem}>
              <div className={sharedStyles.statLabel}>Решено уравнений</div>
              <div className={sharedStyles.statValue}>{stats.total_equation_solved}</div>
            </div>
            <div className={sharedStyles.statItem}>
              <div className={sharedStyles.statLabel}>Ошибок</div>
              <div className={sharedStyles.statValue}>{stats.wrong_answers}</div>
            </div>
            <div className={sharedStyles.statItem}>
              <div className={sharedStyles.statLabel}>Точность</div>
              <div className={sharedStyles.statValue}>{stats.accuracy_percent.toFixed(1)}%</div>
            </div>
          </div>
        </div>

        {stats.equation_types && stats.equation_types.length > 0 && (
          <div className={sharedStyles.formCard}>
            <h3>Точность по типам уравнений</h3>
            <div className={styles.typeStats}>
              {stats.equation_types.map((type, idx) => (
                <div key={idx} className={styles.typeRow}>
                  <div className={styles.typeName}>{type.name}</div>
                  <div className={sharedStyles.progressBarWrapper}>
                    <div className={sharedStyles.progressBar}>
                      <div
                        className={sharedStyles.progressFill}
                        style={{ width: `${type.accuracy_percent}%` }}
                      />
                    </div>
                    <div className={sharedStyles.progressValue}>{type.accuracy_percent}%</div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}

        {classList.length > 0 && (
          <div className={sharedStyles.formCard}>
            <h3>Классы школы</h3>
            <div className={sharedStyles.dataTableWrapper}>
              <table className={sharedStyles.dataTable}>
                <thead>
                  <tr>
                    <th>Название класса</th>
                    <th>Точность</th>
                    <th>Действия</th>
                  </tr>
                </thead>
                <tbody>
                  {classList.map((cls, idx) => (
                    <tr key={idx}>
                      <td>{cls.name}</td>
                      <td>
                        <div className={sharedStyles.tableProgress}>
                          <div
                            className={sharedStyles.tableProgressFill}
                            style={{ width: `${cls.accuracy_percent}%` }}
                          />
                        </div>
                        {cls.accuracy_percent}%
                      </td>
                      <td>
                        <button
                          className={sharedStyles.smallButton}
                          onClick={() => handleClassClick(cls.id)}
                        >
                          Детали
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default DirectorSchoolStats;