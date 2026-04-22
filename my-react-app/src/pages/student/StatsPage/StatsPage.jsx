import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import { BASE_URL } from "../../../constants";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./StatsPage.module.css";

function StatsPage() {
  const [stats, setStats] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetch(BASE_URL + "/student/stats", { credentials: "include" })
      .then((res) => res.json())
      .then(setStats)
      .catch(console.error);
  }, []);

  if (!stats) return <div className={sharedStyles.loader}>Загрузка статистики...</div>;

  return (
    <div className={styles.page}>
      <div className={styles.container}>
        <div className={styles.header}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/student/dashboard")}
          >
            Назад
          </button>
          <LogoutButton />
        </div>

        <h2 className={styles.title}>Статистика</h2>

        {/* основные показатели */}
        <div className={styles.cardsGrid}>
          <StatCard label="XP" value={stats.xp} />
          <StatCard label="Уровни" value={stats.levels_completed} />
          <StatCard label="Звёзды" value={stats.stars_total} />
        </div>

        {/* точность */}
        <div className={styles.block}>
          <h3>Точность</h3>
          <div className={styles.barContainer}>
            <div className={styles.bar}>
              <div
                className={styles.barFill}
                style={{ width: `${stats.accuracy_percent}%` }}
              />
            </div>
            <div className={styles.percentText}>{stats.accuracy_percent}%</div>
          </div>
        </div>

        {/* попытки, правильные, ошибки */}
        <div className={styles.cardsGrid}>
          <StatCard label="Попытки" value={stats.total_attempts} />
          <StatCard label="Правильные" value={stats.correct_answers} />
          <StatCard label="Ошибки" value={stats.wrong_answers} />
        </div>

        {/* типы уравнений */}
        <div className={styles.block}>
          <h3>По типам уравнений</h3>
          {stats.equation_type_stats.map((t, i) => (
            <div key={i} className={styles.typeRow}>
              <div className={styles.typeName}>{t.type}</div>
              <div className={styles.typeBar}>
                <div className={styles.bar}>
                  <div
                    className={styles.barFill}
                    style={{ width: `${t.accuracy_percent}%` }}
                  />
                </div>
              </div>
              <div className={styles.percentText}>{t.accuracy_percent}%</div>
            </div>
          ))}
        </div>

        {/* слабые темы */}
        <div className={styles.block}>
          <h3>Слабые темы</h3>
          {stats.weak_types.length === 0 ? (
            <div className={styles.weakItem} style={{ background: "#d5f5e3", color: "#27ae60" }}>
              Нет слабых мест, отлично!
            </div>
          ) : (
            <div className={styles.weakList}>
              {stats.weak_types.map((w, i) => (
                <div key={i} className={styles.weakItem}>
                  {w}
                </div>
              ))}
            </div>
          )}
        </div>

        {/* достижения */}
        <div className={styles.block}>
          <h3>Достижения</h3>
          <div className={styles.achievementsGrid}>
            {stats.achievements.map((a) => (
              <div key={a.id} className={styles.achievementCard}>
                <div className={styles.achievementIcon}>🏆</div>
                <div className={styles.achievementName}>{a.name}</div>
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
    <div className={styles.statCard}>
      <div className={styles.statCardLabel}>{label}</div>
      <div className={styles.statCardValue}>{value}</div>
    </div>
  );
}

export default StatsPage;