import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { BASE_URL } from "../../../constants";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "../../teacher/ClassStatisticsPage/ClassStatisticsPage.module.css";

function DirectorClassStats() {
  const navigate = useNavigate();
  const { classId } = useParams();
  const [stats, setStats] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    const fetchClassStats = async () => {
      try {
        const res = await fetch(`${BASE_URL}/director/class-stats/${classId}`, {
          credentials: "include",
        });
        if (!res.ok) throw new Error("Ошибка загрузки статистики класса");
        const data = await res.json();
        setStats(data);
      } catch (err) {
        console.error(err);
        setError(err.message);
      } finally {
        setLoading(false);
      }
    };
    fetchClassStats();
  }, [classId]);

  if (loading) return <div className={sharedStyles.loader}>Загрузка...</div>;
  if (error)
    return (
      <div className={sharedStyles.errorBox}>
        <p>Ошибка: {error}</p>
        <button className={sharedStyles.formButton} onClick={() => navigate(-1)}>
          Назад
        </button>
      </div>
    );
  if (!stats) return <div className={sharedStyles.emptyMessage}>Нет данных</div>;

  return (
    <div className={sharedStyles.dashboardPage}>
      <div className={sharedStyles.dashboardContainer}>
        <div className={sharedStyles.dashboardHeader}>
          <button className={sharedStyles.headerButton} onClick={() => navigate(-1)}>
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.dashboardTitle}>Статистика класса</h1>
        {/* Общая статистика */}
        <div className={styles.statsCard}>
          <h3>Общие показатели</h3>
          <div className={styles.statsGrid}>
            <div className={styles.statItem}>
              <div className={styles.statLabel}>Учеников</div>
              <div className={styles.statValue}>{stats.students_count}</div>
            </div>
            <div className={styles.statItem}>
              <div className={styles.statLabel}>Всего попыток</div>
              <div className={styles.statValue}>{stats.total_attempts}</div>
            </div>
            <div className={styles.statItem}>
              <div className={styles.statLabel}>Правильных ответов</div>
              <div className={styles.statValue}>{stats.correct_answers}</div>
            </div>
            <div className={styles.statItem}>
              <div className={styles.statLabel}>Неправильных ответов</div>
              <div className={styles.statValue}>{stats.wrong_answers}</div>
            </div>
            <div className={styles.statItem}>
              <div className={styles.statLabel}>Общая точность</div>
              <div className={styles.statValue}>{stats.accuracy_percent.toFixed(1)}%</div>
            </div>
          </div>
        </div>

        {/* Точность по типам уравнений */}
        {stats.equation_types_stats && stats.equation_types_stats.length > 0 && (
          <div className={styles.statsCard}>
            <h3>Точность по типам примеров</h3>
            <div className={styles.typeStats}>
              {stats.equation_types_stats.map((type, idx) => (
                <div key={idx} className={styles.typeRow}>
                  <div className={styles.typeName}>{type.type}</div>
                  <div className={styles.progressBarWrapper}>
                    <div className={styles.progressBar}>
                      <div
                        className={styles.progressFill}
                        style={{ width: `${type.accuracy_percent}%` }}
                      />
                    </div>
                    <div className={styles.progressValue}>{type.accuracy_percent}%</div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}

        {/* Таблица учеников */}
        {stats.students && stats.students.length > 0 && (
          <div className={styles.statsCard}>
            <h3>Успеваемость учеников</h3>
            <div className={styles.tableWrapper}>
              <table className={styles.statsTable}>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Имя</th>
                    <th>Точность</th>
                    <th>Уровней пройдено</th>
                    <th>Действия</th>
                  </tr>
                </thead>
                <tbody>
                  {stats.students.map((student) => (
                    <tr key={student.student_id}>
                      <td>{student.student_id}</td>
                      <td>{student.name}</td>
                      <td>
                        <div className={styles.tableProgress}>
                          <div
                            className={styles.tableProgressFill}
                            style={{ width: `${student.accuracy}%` }}
                          />
                        </div>
                        {student.accuracy}%
                      </td>
                      <td>{student.levels_complited}</td>
                      <td>
                        <button
                          className={sharedStyles.smallButton}
                          onClick={() => navigate(`/teacher/student-attempts/${student.student_id}`)}
                        >
                          Попытки
                        </button>
                        <button
                          className={sharedStyles.smallButton}
                          onClick={() => navigate(`/director/student-stats/${student.student_id}`)}
                        >
                          Посмотреть статистику
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          </div>
        )}

        {/* Сравнение точности учеников */}
        {stats.students && stats.students.length > 0 && (
          <div className={styles.statsCard}>
            <h3>Сравнение точности учеников</h3>
            <div className={styles.comparisonBlock}>
              {stats.students.map((student) => (
                <div key={student.student_id} className={styles.comparisonItem}>
                  <div className={styles.comparisonLabel}>{student.name}</div>
                  <div className={styles.progressBarWrapper}>
                    <div className={styles.progressBar}>
                      <div
                        className={styles.progressFill}
                        style={{ width: `${student.accuracy}%` }}
                      />
                    </div>
                    <div className={styles.progressValue}>{student.accuracy}%</div>
                  </div>
                </div>
              ))}
            </div>
          </div>
        )}

        {/* График пройденных уровней */}
        {stats.students && stats.students.length > 0 && (
          <div className={styles.statsCard}>
            <h3>Пройденные уровни</h3>
            <div className={styles.comparisonBlock}>
              {stats.students.map((student) => {
                const maxLevels = Math.max(...stats.students.map(s => s.levels_complited));
                const percent = maxLevels ? (student.levels_complited / maxLevels) * 100 : 0;
                return (
                  <div key={student.student_id} className={styles.comparisonItem}>
                    <div className={styles.comparisonLabel}>{student.name}</div>
                    <div className={styles.progressBarWrapper}>
                      <div className={styles.progressBar}>
                        <div
                          className={styles.progressFill}
                          style={{ width: `${percent}%` }}
                        />
                      </div>
                      <div className={styles.progressValue}>{student.levels_complited}</div>
                    </div>
                  </div>
                );
              })}
            </div>
          </div>
        )}
      </div>
    </div>
  );
}

export default DirectorClassStats;