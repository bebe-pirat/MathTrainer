import { useEffect, useState, useRef, useCallback } from "react";
import { useNavigate } from "react-router-dom";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import { BASE_URL } from "../../../constants";
import styles from "./LevelsMapPage.module.css";

function LevelMapPage() {
  const navigate = useNavigate();
  const [map, setMap] = useState(null);

  useEffect(() => {
    fetch(BASE_URL + "/student/level-map", { credentials: "include" })
      .then((res) => res.json())
      .then(setMap)
      .catch(console.error);
  }, []);

  if (!map) return <div style={{ textAlign: "center", padding: "50px" }}>🌀 Загрузка карты...</div>;

  return (
    <div className={styles.page}>
      <div className={styles.header}>
        <button className={styles.headerButton} onClick={() => navigate("/student/stats")}>
          Статистика
        </button>
        <button className={styles.headerButton} onClick={() => navigate("/student/profile")}>
          Профиль
        </button>
        <LogoutButton />
      </div>

      <div className={styles.map}>
        {map.sections.map((section) => (
          <SectionSnake
            key={section.id}
            section={section}
            position={map.student_position}
          />
        ))}
      </div>
    </div>
  );
}

function SectionSnake({ section, position }) {
  const navigate = useNavigate();
  const containerRef = useRef(null);
  const [levelsWithPos, setLevelsWithPos] = useState([]);
  const [containerWidth, setContainerWidth] = useState(0);

  const RADIUS = 100;           // половина ширины кружка (140/2 + 30)
  const VERTICAL_STEP = 160;   // расстояние между центрами по вертикали

  const computePositions = useCallback(() => {
    const total = section.levels_count;
    if (total <= 1 || containerWidth === 0) return [];

    const amplitude = Math.min(containerWidth * 0.15, 150);
    const startX = containerWidth / 2;
    const cycles = 1.5;

    const positions = [];
    for (let i = 0; i < total; i++) {
      const t = i / (total - 1);
      const angle = t * 2 * Math.PI * cycles;
      const x = startX + amplitude * Math.sin(angle);
      const y = i * VERTICAL_STEP + RADIUS; // сдвиг вниз на радиус
      positions.push({
        left: `${x}px`,
        top: `${y}px`,
      });
    }
    return positions;
  }, [section.levels_count, containerWidth]);

  // Пересчёт при изменении ширины или количества уровней
  useEffect(() => {
    const currentRef = containerRef.current;
    if (!currentRef) return;

    const updateWidth = () => {
      if (containerRef.current) {
        setContainerWidth(containerRef.current.clientWidth);
      }
    };

    updateWidth();
    const resizeObserver = new ResizeObserver(updateWidth);
    resizeObserver.observe(currentRef);

    return () => {
      resizeObserver.disconnect();
    };
  }, []); // пустой массив зависимостей – только при монтировании

  useEffect(() => {
    setLevelsWithPos(computePositions());
  }, [computePositions]);

  const getLevelClass = (levelOrder) => {
    const isCompleted =
      position.section_id > section.id ||
      (position.section_id === section.id && position.level_order > levelOrder);
    const isCurrent =
      position.section_id === section.id && position.level_order === levelOrder;
    const isBlocked = section.is_blocked || position.section_id < section.id;

    if (isBlocked) return styles.blocked;
    if (isCompleted) return styles.completed;
    if (isCurrent) return styles.current;
    return styles.available;
  };

  const handleLevelClick = (levelOrder, isBlocked) => {
    if (isBlocked) return;
    navigate("/game", {
      state: {
        sectionId: section.id,
        levelOrder: levelOrder,
      },
    });
  };

   const total = section.levels_count;
  const containerHeight = (total - 1) * VERTICAL_STEP + 2 * RADIUS; // высота с учётом кружков

  return (
    <div className={styles.section}>
      <h2>{section.name}</h2>

      <div
        className={styles.snakeLine}
        ref={containerRef}
        style={{ position: "relative", width: "100%", height: `${containerHeight}px` }}
      >
        {levelsWithPos.map((pos, idx) => {
          const levelOrder = idx + 1;
          const isBlocked = section.is_blocked || position.section_id < section.id;
          return (
            <div
              key={idx}
              className={`${styles.level} ${getLevelClass(levelOrder)}`}
              style={{
                left: pos.left,
                top: pos.top,
                transform: "translate(-50%, -50%)",
              }}
              onClick={() => handleLevelClick(levelOrder, isBlocked)}
            >
              {levelOrder}
            </div>
          );
        })}
      </div>
    </div>
  );
}

export default LevelMapPage;