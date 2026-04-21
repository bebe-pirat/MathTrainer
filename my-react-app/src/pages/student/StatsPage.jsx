import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import { BASE_URL } from "../../constants";

function StatsPage() {
    const [stats, setStats] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        fetch(BASE_URL + "/student/stats", {
            credentials: "include",
        })
            .then(res => res.json())
            .then(setStats)
            .catch(console.error);
    }, []);

    if (!stats) return <div>Loading...</div>;

    return (
        <div style={styles.page}>
            {/* header */}
            <div style={styles.header}>
                <button onClick={() => navigate("/student/dashboard")}>
                    Назад
                </button>
                <LogoutButton />
            </div>

            <h2>Статистика</h2>

            {/* основные показатели */}
            <div style={styles.cards}>
                <StatCard label="XP" value={stats.xp} />
                <StatCard label="Уровни" value={stats.levels_completed} />
                <StatCard label="Звезды" value={stats.stars_total} />
            </div>

            {/* accuracy */}
            <div style={styles.block}>
                <h3>Точность</h3>
                <div style={styles.bar}>
                    <div
                        style={{
                            ...styles.barFill,
                            width: `${stats.accuracy_percent}%`,
                        }}
                    />
                </div>
                <div>{stats.accuracy_percent}%</div>
            </div>

            {/* ответы */}
            <div style={styles.cards}>
                <StatCard label="Попытки" value={stats.total_attempts} />
                <StatCard label="Правильные" value={stats.correct_answers} />
                <StatCard label="Ошибки" value={stats.wrong_answers} />
            </div>

            {/* типы уравнений */}
            <div style={styles.block}>
                <h3>По типам уравнений</h3>

                {stats.equation_type_stats.map((t, i) => (
                    <div key={i} style={styles.typeRow}>
                        <div>{t.type}</div>

                        <div style={styles.bar}>
                            <div
                                style={{
                                    ...styles.barFill,
                                    width: `${t.accuracy}%`,
                                }}
                            />
                        </div>

                        <div>{t.accuracy}%</div>
                    </div>
                ))}
            </div>

            {/* слабые темы */}
            <div style={styles.block}>
                <h3>Слабые темы</h3>

                {stats.weak_types.length === 0 ? (
                    <div>Нет слабых мест 🎉</div>
                ) : (
                    stats.weak_types.map((w, i) => (
                        <div key={i} style={styles.weak}>
                            {w}
                        </div>
                    ))
                )}
            </div>

            {/* достижения */}
            <div style={styles.block}>
                <h3>Достижения</h3>

                <div style={styles.achievements}>
                    {stats.achievements.map((a) => (
                        <div key={a.id} style={styles.achievement}>
                            <div>🏆</div>
                            <div>{a.name}</div>
                        </div>
                    ))}
                </div>
            </div>
        </div>
    );
}

function StatCard({ label, value }) {
    return (
        <div style={styles.card}>
            <div>{label}</div>
            <div style={styles.cardValue}>{value}</div>
        </div>
    );
}

export default StatsPage;

const styles = {
    page: {
        backgroundColor: "white",
        minHeight: "100vh",
        padding: "20px",
        maxWidth: "800px",
        margin: "0 auto",
    },

    header: {
        display: "flex",
        justifyContent: "space-between",
        marginBottom: "20px",
    },

    cards: {
        display: "flex",
        gap: "10px",
        marginBottom: "20px",
    },

    card: {
        flex: 1,
        border: "1px solid #ddd",
        padding: "10px",
        textAlign: "center",
    },

    cardValue: {
        fontSize: "20px",
        fontWeight: "bold",
        color: "#4da6ff",
    },

    block: {
        marginBottom: "20px",
    },

    bar: {
        width: "100%",
        height: "10px",
        backgroundColor: "#eee",
        marginTop: "5px",
    },

    barFill: {
        height: "100%",
        backgroundColor: "#4da6ff",
    },

    typeRow: {
        marginBottom: "10px",
    },

    weak: {
        color: "red",
        marginTop: "5px",
    },

    achievements: {
        display: "flex",
        gap: "10px",
        flexWrap: "wrap",
    },

    achievement: {
        border: "1px solid #ddd",
        padding: "10px",
        textAlign: "center",
        width: "100px",
    },
};
