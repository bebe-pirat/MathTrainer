import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import { BASE_URL } from "../../constants";
import sharedStyles from "../../styles/shared.module.css";

function TeacherStudentStatsPage() {
  const [stats, setStats] = useState(null);
  const navigate = useNavigate();
  const { studentId } = useParams();

  useEffect(() => {
    fetch(BASE_URL + "/teacher/students/stats/" + studentId, { credentials: "include" })
      .then((res) => res.json())
      .then(setStats)
      .catch(console.error);
  }, []);

  if (!stats) return <div className={sharedStyles.loader}>Загрузка статистики...</div>;

  return (
    <div className={sharedStyles.page}>
      <div className={sharedStyles.container}>
        <div className={sharedStyles.header}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/director/dashboard")}
          >
            Назад
          </button>
          <LogoutButton />
        </div>

        <h2 className={sharedStyles.title}>Статистика</h2>

        {/* основные показатели */}
        <div className={sharedStyles.cardsGrid}>
          <StatCard label="XP" value={stats.xp} />
          <StatCard label="Уровни" value={stats.levels_completed} />
          <StatCard label="Звёзды" value={stats.stars_total} />
        </div>

        {/* точность */}
        <div className={sharedStyles.block}>
          <h3>Точность</h3>
          <div className={sharedStyles.barContainer}>
            <div className={sharedStyles.bar}>
              <div
                className={sharedStyles.barFill}
                style={{ width: `${stats.accuracy_percent}%` }}
              />
            </div>
            <div className={sharedStyles.percentText}>{stats.accuracy_percent}%</div>
          </div>
        </div>

        {/* попытки, правильные, ошибки */}
        <div className={sharedStyles.cardsGrid}>
          <StatCard label="Попытки" value={stats.total_attempts} />
          <StatCard label="Правильные" value={stats.correct_answers} />
          <StatCard label="Ошибки" value={stats.wrong_answers} />
        </div>

        {/* типы уравнений */}
        <div className={sharedStyles.block}>
          <h3>По типам уравнений</h3>
          {stats.equation_type_stats.map((t, i) => (
            <div key={i} className={sharedStyles.typeRow}>
              <div className={sharedStyles.typeName}>{t.type}</div>
              <div className={sharedStyles.typeBar}>
                <div className={sharedStyles.bar}>
                  <div
                    className={sharedStyles.barFill}
                    style={{ width: `${t.accuracy_percent}%` }}
                  />
                </div>
              </div>
              <div className={sharedStyles.percentText}>{t.accuracy_percent}%</div>
            </div>
          ))}
        </div>

        {/* слабые темы */}
        <div className={sharedStyles.block}>
          <h3>Слабые темы</h3>
          {stats.weak_types.length === 0 ? (
            <div className={sharedStyles.weakItem} style={{ background: "#d5f5e3", color: "#27ae60" }}>
              Нет слабых мест, отлично!
            </div>
          ) : (
            <div className={sharedStyles.weakList}>
              {stats.weak_types.map((w, i) => (
                <div key={i} className={sharedStyles.weakItem}>
                  {w}
                </div>
              ))}
            </div>
          )}
        </div>

        {/* достижения */}
        <div className={sharedStyles.block}>
          <h3>Достижения</h3>
          <div className={sharedStyles.achievementsGrid}>
            {stats.achievements.map((a) => (
              <div key={a.id} className={sharedStyles.achievementCard}>
                <div className={sharedStyles.achievementIcon}>🏆</div>
                <div className={sharedStyles.achievementName}>{a.name}</div>
              </div>
            ))}
          </div>
        </div>
      </div>
    </div>
  );
}

function StatCard({ label, value }) {
  return (
    <div className={sharedStyles.statCard}>
      <div className={sharedStyles.statCardLabel}>{label}</div>
      <div className={sharedStyles.statCardValue}>{value}</div>
    </div>
  );
}

export default TeacherStudentStatsPage;