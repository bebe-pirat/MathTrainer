import { useEffect, useState } from "react";
import LogoutButton from "./../../components/LogoutButton";
import { BASE_URL } from "./../../constants";

function LevelMapPage() {
    const [map, setMap] = useState(null);

    useEffect(() => {
        fetch(BASE_URL + "/student/level-map", {
            credentials: "include",
        })
            .then((res) => res.json())
            .then(setMap)
            .catch(console.error);
    }, []);

    if (!map) return <div>Loading...</div>;

    console.log(map)

    return (
        <div style={styles.page}>
            {/* верхняя панель */}
            <div style={styles.header}>
                <button>Статистика</button>
                <button>Профиль</button>
                <LogoutButton />
            </div>

            {/* карта */}
            <div style={styles.map}>
                {map.sections.map((section, sIndex) => (
                    <div key={section.id} style={styles.section}>
                        <h3>{section.name}</h3>

                        <div style={styles.levelRow}>
                            {[...Array(section.levels_count)].map((_, i) => {
                                const levelNumber = i + 1;

                                const isCompleted =
                                    map.student_position.section_id > section.id ||
                                    (map.student_position.section_id === section.id &&
                                        map.student_position.level_order > levelNumber);

                                const isCurrent =
                                    map.student_position.section_id === section.id &&
                                    map.student_position.level_order === levelNumber;

                                const isBlocked =
                                    section.is_blocked ||
                                    (map.student_position.section_id < section.id);

                                return (
                                    <div
                                        key={i}
                                        style={{
                                            ...styles.level,
                                            backgroundColor: isBlocked
                                                ? "#ccc"
                                                : isCompleted
                                                ? "#4da6ff"
                                                : isCurrent
                                                ? "#66b3ff"
                                                : "#e6f2ff",
                                        }}
                                    >
                                        {levelNumber}
                                    </div>
                                );
                            })}
                        </div>
                    </div>
                ))}
            </div>
        </div>
    );
}

export default LevelMapPage;

const styles = {
    page: {
        backgroundColor: "white",
        minHeight: "100vh",
        padding: "20px",
    },

    header: {
        display: "flex",
        gap: "10px",
        marginBottom: "20px",
    },

    map: {
        display: "flex",
        flexDirection: "column",
        gap: "40px",
        alignItems: "center",
    },

    section: {
        textAlign: "center",
    },

    levelRow: {
        display: "flex",
        gap: "15px",
        flexWrap: "wrap",
        justifyContent: "center",
    },

    level: {
        width: "50px",
        height: "50px",
        borderRadius: "50%",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        cursor: "pointer",
        fontWeight: "bold",
    },
};
