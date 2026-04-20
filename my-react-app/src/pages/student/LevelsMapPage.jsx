import { useEffect, useState } from "react";
import LogoutButton from "./../../components/LogoutButton";
import { BASE_URL } from "../../constants";
import { useNavigate } from "react-router-dom";

function LevelMapPage() {
    const navigate = useNavigate();
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

    return (
        <div style={styles.page}>
            {/* header */}
            <div style={styles.header}>
                <button onClick={() => navigate("/student/stats")}>
                    Статистика
                </button>

                <button onClick={() => navigate("/student/profile")}>
                    Профиль
                </button>
                <LogoutButton />
            </div>

            {/* карта */}
            <div style={styles.map}>
                {map.sections.map((section, sIndex) => (
                    <SectionSnake
                        key={section.id}
                        section={section}
                        index={sIndex}
                        position={map.student_position}
                    />
                ))}
            </div>
        </div>
    );
}

export default LevelMapPage;

function SectionSnake({ section, index, position }) {
    const navigate = useNavigate(); 
    const direction = index % 2 === 0 ? "right" : "left";

    return (
        <div style={styles.section}>
            <h3>{section.name}</h3>

            <div
                style={{
                    ...styles.levelRow,
                    flexDirection: direction === "right" ? "row" : "row-reverse",
                }}
            >
                {[...Array(section.levels_count)].map((_, i) => {
                    const level = i + 1;

                    const isCompleted =
                        position.section_id > section.id ||
                        (position.section_id === section.id &&
                            position.level_order > level);

                    const isCurrent =
                        position.section_id === section.id &&
                        position.level_order === level;

                    const isBlocked =
                        section.is_blocked ||
                        position.section_id < section.id;

                    return (
                        <div key={i} style={styles.nodeWrapper}>
                            <div onClick={() =>
                                    navigate("/game", {
                                        state: {
                                            sectionId: section.id,
                                            levelOrder: level,
                                        },
                                    })
                                }

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
                                {level}
                            </div>

                            {/* горизонтальная линия */}
                            {i < section.levels_count - 1 && (
                                <div style={styles.horizontalLine} />
                            )}
                        </div>
                    );
                })}
            </div>

            {/* вертикальная связь к следующей секции */}
            <div style={styles.verticalConnector} />
        </div>
    );
}

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
        alignItems: "center",
        gap: "60px",
    },

    section: {
        display: "flex",
        flexDirection: "column",
        alignItems: "center",
    },

    levelRow: {
        display: "flex",
        alignItems: "center",
        position: "relative",
    },

    nodeWrapper: {
        display: "flex",
        alignItems: "center",
    },

    level: {
        width: "50px",
        height: "50px",
        borderRadius: "50%",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        fontWeight: "bold",
        cursor: "pointer",
        zIndex: 2,
    },

    horizontalLine: {
        width: "40px",
        height: "4px",
        backgroundColor: "#4da6ff",
    },

    verticalConnector: {
        width: "4px",
        height: "40px",
        backgroundColor: "#4da6ff",
        marginTop: "10px",
    },
};
