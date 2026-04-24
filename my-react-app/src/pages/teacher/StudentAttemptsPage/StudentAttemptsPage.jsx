import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import { BASE_URL } from "../../../constants";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./StudentAttemptsPage.module.css";

function StudentAttemptsPage() {
  const navigate = useNavigate();
  const { studentId } = useParams();
  const [attempts, setAttempts] = useState([]);
  const [equationTypes, setEquationTypes] = useState([]);
  const [selectedTypeId, setSelectedTypeId] = useState("");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);
  const [studentName, setStudentName] = useState("");

  // Загрузка типов уравнений (для фильтра)
  const fetchEquationTypes = async () => {
    try {
      const res = await fetch(`${BASE_URL}/teacher/equation-types?student_id=${studentId}`, {
        credentials: "include",
      });
      if (!res.ok) throw new Error("Ошибка загрузки типов уравнений");
      const data = await res.json();
      setEquationTypes(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    }
  };

  // Загрузка попыток с учётом фильтра
  const fetchAttempts = async () => {
    setLoading(true);
    try {
      let url = `${BASE_URL}/teacher/students/attempts?student_id=${studentId}`;
      if (selectedTypeId) {
        url += `&equation_type_id=${selectedTypeId}`;
      }
      const res = await fetch(url, { credentials: "include" });
      if (!res.ok) throw new Error("Ошибка загрузки попыток");
      const data = await res.json();
      setAttempts(data);
      // Если есть имя ученика (можно из первого элемента или отдельного API)
      if (data.length > 0 && data[0].student_name) {
        setStudentName(data[0].student_name);
      }
    } catch (err) {
      console.error(err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    if (studentId) {
      fetchEquationTypes();
      fetchAttempts();
    }
  }, [studentId]);

  useEffect(() => {
    if (studentId) {
      fetchAttempts();
    }
  }, [selectedTypeId]);

  const formatDateTime = (dateString) => {
    if (!dateString) return "-";
    const date = new Date(dateString);
    return date.toLocaleString("ru-RU");
  };

  if (loading && attempts.length === 0) {
    return <div className={sharedStyles.loader}>Загрузка попыток...</div>;
  }

  if (error) {
    return (
      <div className={sharedStyles.errorBox}>
        <p>Ошибка: {error}</p>
        <button
          className={sharedStyles.headerButton}
          onClick={() => navigate("/teacher/class-statistics")}
        >
          Назад
        </button>
      </div>
    );
  }

  return (
    <div className={sharedStyles.dashboardPage}>
      <div className={sharedStyles.dashboardContainer}>
        <div className={sharedStyles.dashboardHeader}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/teacher/class-statistics")}
          >
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.dashboardTitle}>
          Попытки ученика {studentName && `— ${studentName}`}
        </h1>

        {/* Фильтр по типам уравнений */}
        <div className={styles.filterBar}>
          <div className={styles.filterGroup}>
            <label className={styles.filterLabel}>Тип уравнения</label>
            <select
              className={sharedStyles.formSelect}
              value={selectedTypeId}
              onChange={(e) => setSelectedTypeId(e.target.value)}
            >
              <option value="">Все типы</option>
              {equationTypes.map((type) => (
                <option key={type.id} value={type.id}>
                  {type.name}
                </option>
              ))}
            </select>
          </div>
          <div className={styles.filterGroup}>
            <button
              className={sharedStyles.smallButton}
              onClick={() => setSelectedTypeId("")}
              style={{ marginTop: "28px" }}
            >
              Сбросить фильтр
            </button>
          </div>
        </div>

        {/* Таблица попыток */}
        {attempts.length === 0 ? (
          <div className={sharedStyles.emptyMessage}>
            Нет попыток по выбранному фильтру
          </div>
        ) : (
          <div className={sharedStyles.dataTableWrapper}>
            <table className={styles.attemptsTable}>
              <thead>
                <tr>
                  <th>Дата и время</th>
                  <th>Уравнение</th>
                  <th>Тип</th>
                  <th>Ответ ученика</th>
                  <th>Правильный ответ</th>
                  <th>Результат</th>
                </tr>
              </thead>
              <tbody>
                {attempts.map((attempt, idx) => {
                  const isCorrect = attempt.given_answer === attempt.correct_answer;
                  return (
                    <tr key={idx}>
                      <td>{formatDateTime(attempt.answered_at)}</td>
                      <td>{attempt.equation_text}</td>
                      <td>{attempt.equation_type_name}</td>
                      <td>{attempt.given_answer}</td>
                      <td>{attempt.correct_answer}</td>
                      <td>
                        {isCorrect ? (
                          <span className={styles.correctBadge}>Верно</span>
                        ) : (
                          <span className={styles.incorrectBadge}>Неверно</span>
                        )}
                      </td>
                    </tr>
                  );
                })}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}

export default StudentAttemptsPage;