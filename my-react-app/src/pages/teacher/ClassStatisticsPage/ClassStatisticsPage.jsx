import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../../constants";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./ClassStatisticsPage.module.css";

function ClassStatistics() {
  const navigate = useNavigate();
  const [stats, setStats] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchClassStats();
  }, []);

  const fetchClassStats = async () => {
    try {
      const response = await fetch(BASE_URL + "/teacher/class/stats", {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка загрузки статистики");
      const data = await response.json();
      setStats(data);
    } catch (err) {
      setError(err.message);
      console.error(err);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return <div className={styles.loader}>Загрузка статистики...</div>;
  }

  if (error) {
    return (
      <div className={styles.errorBox}>
        <p>Ошибка: {error}</p>
        <button
          className={sharedStyles.headerButton}
          onClick={() => navigate("/teacher/dashboard")}
        >
          Назад
        </button>
      </div>
    );
  }

  if (!stats) {
    return <div className={styles.noData}>Нет данных</div>;
  }

  return (
    <div className={styles.page}>
      <div className={styles.container}>
        <div className={styles.header}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/teacher/dashboard")}
          >
            Назад
          </button>
        </div>

        <h1 className={styles.title}>Статистика класса</h1>

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
            <h3>Точность по типам уравнений</h3>
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

export default ClassStatistics;