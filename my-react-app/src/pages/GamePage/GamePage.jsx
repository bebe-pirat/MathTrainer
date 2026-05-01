import { useEffect, useState } from "react";
import { useLocation, useNavigate } from "react-router-dom";
import { BASE_URL } from "./../../constants";
import sharedStyles from "../../styles/shared.module.css";
import styles from "./GamePage.module.css";

function GamePage() {
  const location = useLocation();
  const navigate = useNavigate();
  const { sectionId, levelOrder } = location.state || {};

  const [equations, setEquations] = useState([]);
  const [answers, setAnswers] = useState({});
  const [feedback, setFeedback] = useState(null);
  const [result, setResult] = useState(null);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  // загрузка уравнений
  useEffect(() => {
    if (!sectionId) {
      setError("Отсутствует информация об уровне");
      setLoading(false);
      return;
    }

    fetch(BASE_URL + "/game/equations-set", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify(sectionId),
    })
      .then((res) => {
        if (!res.ok) throw new Error("Ошибка загрузки уравнений");
        return res.json();
      })
      .then((data) => {
        setEquations(data);
        setLoading(false);
      })
      .catch((err) => {
        console.error(err);
        setError(err.message);
        setLoading(false);
      });
  }, [sectionId]);

  const handleChange = (id, value) => {
    setAnswers((prev) => ({ ...prev, [id]: value }));
  };

  const handleCheck = async () => {
    if (Object.keys(answers).length !== equations.length) {
      alert("Ответьте на все вопросы");
      return;
    }

    const payload = equations.map((eq) => ({
      equation_id: eq.id,
      equation_text: eq.equation_text,
      correct_answer: eq.correct_answer,
      user_answer: Number(answers[eq.id]),
      equation_type_id: eq.equation_type_id,
    }));

    try {
      const res = await fetch(BASE_URL + "/game/check", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify(payload),
      });
      const data = await res.json();
      setFeedback(data);
    } catch (err) {
      console.error(err);
      alert("Ошибка при проверке");
    }
  };

  const handleFinish = async () => {
    try {
      const res = await fetch(BASE_URL + "/game/finish-level", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          feedback: feedback,
          section_id: sectionId,
          level_order: levelOrder,
        }),
      });
      const data = await res.json();
      setResult(data);
    } catch (err) {
      console.error(err);
      alert("Ошибка завершения уровня");
    }
  };

  if (loading) return <div className={styles.loader}>Загрузка уравнений...</div>;
  if (error) return <div className={styles.loader}>Ошибка: {error}</div>;
  if (!equations.length) return <div className={styles.loader}>Нет уравнений для этого уровня</div>;

  return (
    <div className={styles.page}>
      <div className={styles.container}>
        <div className={styles.header}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/student/Dashboard")}
          >
            Назад
          </button>
        </div>

        <h1 className={styles.title}>Решение уравнений</h1>
        <div className={styles.levelInfo}>
          Уровень {levelOrder} • Секция {sectionId}
        </div>

        {equations.map((eq) => {
          const fb = feedback?.find((f) => f.equation_id === eq.id);
          return (
            <div key={eq.id} className={styles.equationCard}>
              <div className={styles.equationText}>{eq.equation_text}</div>
              <input
                type="number"
                className={styles.inputField}
                placeholder="Введите ответ"
                value={answers[eq.id] || ""}
                onChange={(e) => handleChange(eq.id, e.target.value)}
                disabled={!!feedback}
              />
              {fb && (
                <div className={fb.is_correct ? styles.feedbackCorrect : styles.feedbackWrong}>
                  {fb.is_correct ? "Верно!" : `Ошибка. Правильный ответ: ${fb.correct_answer}`}
                </div>
              )}
            </div>
          );
        })}

        <div className={styles.buttonGroup}>
          {!feedback && (
            <button className={styles.actionButton} onClick={handleCheck}>
              Проверить
            </button>
          )}
          {feedback && !result && (
            <button className={styles.actionButton} onClick={handleFinish}>
              Завершить уровень
            </button>
          )}
        </div>

        {result && (
          <div className={styles.resultCard}>
            <div className={styles.resultTitle}>Уровень пройден!</div>
            <div className={styles.stars}>
              {"⭐".repeat(result.stars)}
            </div>
            <div className={styles.xp}>Получено XP: {result.common_xp}</div>
            <button
              className={styles.actionButton}
              onClick={() => navigate("/student/dashboard")}
              style={{ marginTop: "20px" }}
            >
              К карте уровней
            </button>
          </div>
        )}
      </div>
    </div>
  );
}

export default GamePage;