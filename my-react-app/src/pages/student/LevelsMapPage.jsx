import { useEffect, useState, useRef } from "react";
import { useNavigate } from "react-router-dom";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import { BASE_URL } from "../../constants";

const styles = {
  page: {
    background: "linear-gradient(135deg, #f5f9ff 0%, #ffffff 100%)",
    minHeight: "100vh",
    padding: "30px 20px",
    fontFamily: "'Segoe UI', Roboto, 'Helvetica Neue', sans-serif",
  },

  header: {
    display: "flex",
    justifyContent: "flex-end",
    gap: "15px",
    marginBottom: "40px",
    maxWidth: "1200px",
    margin: "0 auto 40px auto",
    padding: "0 20px",
  },

  headerButton: {
    background: "white",
    border: "none",
    padding: "10px 24px",
    borderRadius: "40px",
    fontSize: "1rem",
    fontWeight: 600,
    color: "#1e6f9f",
    cursor: "pointer",
    boxShadow: "0 2px 8px rgba(0,0,0,0.05)",
    transition: "all 0.2s ease",
  },

  map: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    gap: "80px",
    maxWidth: "1200px",
    margin: "0 auto",
  },

  section: {
    display: "flex",
    flexDirection: "column",
    alignItems: "center",
    width: "100%",
    position: "relative",
  },

  sectionTitle: {
    fontSize: "1.8rem",
    fontWeight: "700",
    background: "linear-gradient(135deg, #1e6f9f, #4da6ff)",
    backgroundClip: "text",
    WebkitBackgroundClip: "text",
    color: "transparent",
    marginBottom: "30px",
    letterSpacing: "-0.5px",
  },

  levelRow: {
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    flexWrap: "wrap",
    position: "relative",
    padding: "20px 10px",
  },

  nodeWrapper: {
    display: "flex",
    alignItems: "center",
    position: "relative",
  },

  level: {
    width: "64px",
    height: "64px",
    borderRadius: "50%",
    display: "flex",
    alignItems: "center",
    justifyContent: "center",
    fontSize: "1.3rem",
    fontWeight: "bold",
    cursor: "pointer",
    transition: "all 0.25s cubic-bezier(0.2, 0.9, 0.4, 1.1)",
    boxShadow: "0 8px 20px rgba(0,0,0,0.1)",
    border: "2px solid rgba(255,255,255,0.6)",
    backdropFilter: "blur(2px)",
    position: "relative",
    zIndex: 2,
  },

//   horizontalLine: {
//     width: "60px",
//     height: "4px",
//     background: "linear-gradient(90deg, #4da6ff, #a0d0ff)",
//     borderRadius: "4px",
//     margin: "0 5px",
//     position: "relative",
//     transform: "skewX(-5deg)", // лёгкий изгиб
//   },

  verticalConnector: {
    width: "6px",
    height: "60px",
    background: "linear-gradient(180deg, #4da6ff, #a0d0ff, #4da6ff)",
    borderRadius: "3px",
    marginTop: "20px",
    position: "relative",
  },

  levelBlocked: {
    background: "#e0e5ec",
    color: "#8a99aa",
    cursor: "not-allowed",
    boxShadow: "none",
    border: "1px solid #ccd4e0",
  },
  levelCompleted: {
    background: "linear-gradient(135deg, #2b7fb5, #4da6ff)",
    color: "white",
    boxShadow: "0 8px 20px rgba(77,166,255,0.3)",
  },
  levelCurrent: {
    background: "radial-gradient(circle at 30% 20%, #66ccff, #2b7fb5)",
    color: "white",
    boxShadow: "0 0 0 4px rgba(77,166,255,0.5), 0 8px 20px rgba(0,0,0,0.2)",
    animation: "pulse 1.5s infinite",
  },
  levelAvailable: {
    background: "white",
    color: "#2b7fb5",
    border: "2px solid #4da6ff",
  },
    svgLayer: {
    position: "absolute",
    top: 0,
    left: 0,
    width: "100%",
    height: "100%",
    pointerEvents: "none",   // чтобы клики проходили сквозь SVG
    overflow: "visible",     // чтобы линии не обрезались
  },
};

function calculateCubicPath(start, end) {
  const dx = end.x - start.x;
  const dy = end.y - start.y;
  const offsetX = dx * 0.5;
  const offsetY = Math.min(Math.abs(dy) * 0.8, 50) * (dy > 0 ? 1 : -1);
  
  const smoothY = Math.abs(dy) < 5 ? (start.y - 15) : (start.y + offsetY);
  
  return `
    M ${start.x},${start.y}
    C ${start.x + offsetX},${smoothY}
      ${end.x - offsetX},${end.y - offsetY}
      ${end.x},${end.y}
  `;
}

const injectKeyframes = () => {
  if (document.getElementById("level-map-keyframes")) return;
  const style = document.createElement("style");
  style.id = "level-map-keyframes";
  style.textContent = `
    @keyframes pulse {
      0% { box-shadow: 0 0 0 0 rgba(77,166,255,0.7); }
      70% { box-shadow: 0 0 0 10px rgba(77,166,255,0); }
      100% { box-shadow: 0 0 0 0 rgba(77,166,255,0); }
    }
    .level-hover:hover {
      transform: scale(1.08);
      filter: brightness(1.02);
    }
  `;
  document.head.appendChild(style);
};

function LevelMapPage() {
  const navigate = useNavigate();
  const [map, setMap] = useState(null);

  useEffect(() => {
    injectKeyframes();
    fetch(BASE_URL + "/student/level-map", { credentials: "include" })
      .then((res) => res.json())
      .then(setMap)
      .catch(console.error);
  }, []);

  if (!map) return <div style={{ textAlign: "center", padding: "50px" }}>🌀 Загрузка карты...</div>;

  return (
    <div style={styles.page}>
      <div style={styles.header}>
        <button
          style={styles.headerButton}
          onClick={() => navigate("/student/stats")}
          onMouseEnter={(e) => (e.currentTarget.style.transform = "translateY(-2px)")}
          onMouseLeave={(e) => (e.currentTarget.style.transform = "translateY(0)")}
        >
          Статистика
        </button>
        <button
          style={styles.headerButton}
          onClick={() => navigate("/student/profile")}
          onMouseEnter={(e) => (e.currentTarget.style.transform = "translateY(-2px)")}
          onMouseLeave={(e) => (e.currentTarget.style.transform = "translateY(0)")}
        >
          Профиль
        </button>
        <LogoutButton />
      </div>

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
function SectionSnake({ section, index, position }) {
  const navigate = useNavigate();
  const direction = index % 2 === 0 ? "row" : "row-reverse";
  const rowRef = useRef(null);
  const [paths, setPaths] = useState([]);

  // Получаем координаты кружков после рендера
  useEffect(() => {
    if (!rowRef.current) return;
    const levelNodes = rowRef.current.querySelectorAll(".level-node");
    if (levelNodes.length < 2) return;

    const newPaths = [];
    for (let i = 0; i < levelNodes.length - 1; i++) {
      const startRect = levelNodes[i].getBoundingClientRect();
      const endRect = levelNodes[i + 1].getBoundingClientRect();
      const containerRect = rowRef.current.parentElement.getBoundingClientRect();

      // Координаты относительно контейнера с SVG
      const start = {
        x: startRect.left + startRect.width / 2 - containerRect.left,
        y: startRect.top + startRect.height / 2 - containerRect.top,
      };
      const end = {
        x: endRect.left + endRect.width / 2 - containerRect.left,
        y: endRect.top + endRect.height / 2 - containerRect.top,
      };
      newPaths.push(calculateCubicPath(start, end));
    }
    setPaths(newPaths);
  }, [section.levels_count]); // пересчёт при изменении количества уровней

  const getLevelStyle = (level) => {
    const isCompleted =
      position.section_id > section.id ||
      (position.section_id === section.id && position.level_order > level);
    const isCurrent =
      position.section_id === section.id && position.level_order === level;
    const isBlocked = section.is_blocked || position.section_id < section.id;

    if (isBlocked) return { ...styles.level, ...styles.levelBlocked };
    if (isCompleted) return { ...styles.level, ...styles.levelCompleted };
    if (isCurrent) return { ...styles.level, ...styles.levelCurrent };
    return { ...styles.level, ...styles.levelAvailable };
  };

  const handleClick = (level) => {
    const isBlocked = section.is_blocked || position.section_id < section.id;
    if (!isBlocked) {
      navigate("/game", { state: { sectionId: section.id, levelOrder: level } });
    }
  };

  return (
    <div style={styles.section}>
      <h2 style={styles.sectionTitle}>{section.name}</h2>
      <div style={{ position: "relative", width: "100%" }}>
        {/* SVG слой для кривых */}
        <svg style={styles.svgLayer}>
          <g>
            {paths.map((d, idx) => (
              <path
                key={idx}
                d={d}
                fill="none"
                stroke="#4da6ff"
                strokeWidth="4"
                strokeLinecap="round"
                strokeDasharray={section.is_blocked ? "6 4" : "none"}
              />
            ))}
          </g>
        </svg>

        {/* Контейнер с кружками */}
        <div
          ref={rowRef}
          style={{
            ...styles.levelRow,
            flexDirection: direction,
            gap: "20px",
            justifyContent: "center",
          }}
        >
          {[...Array(section.levels_count)].map((_, i) => {
            const level = i + 1;
            const levelStyle = getLevelStyle(level);
            return (
              <div key={i} style={{ position: "relative" }}>
                <div
                  className="level-node"
                  onClick={() => handleClick(level)}
                  style={levelStyle}
                >
                  {level}
                </div>
              </div>
            );
          })}
        </div>
      </div>
      {/* Вертикальный соединитель */}
      <div style={styles.verticalConnector} />
    </div>
  );
}

export default LevelMapPage;