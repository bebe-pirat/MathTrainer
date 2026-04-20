import { useEffect, useState } from "react";
import { useLocation } from "react-router-dom";
import { BASE_URL } from "./../../constants";

function GamePage() {
    const location = useLocation();
    const { sectionId, levelOrder } = location.state || {};

    const [equations, setEquations] = useState([]);
    const [answers, setAnswers] = useState({});
    const [feedback, setFeedback] = useState(null);
    const [result, setResult] = useState(null);

    // загрузка уравнений
    useEffect(() => {
        console.log(sectionId);

        fetch(BASE_URL + "/game/equations-set", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify(sectionId),
        })
            .then(res => res.json())
            .then(setEquations)
            .catch(console.error);
    }, [sectionId]);

    // ввод ответа
    const handleChange = (id, value) => {
        setAnswers(prev => ({
            ...prev,
            [id]: value,
        }));
    };

    // проверка
    const handleCheck = async () => {
        if (Object.keys(answers).length !== equations.length) {
            alert("Ответь на все вопросы");
            return;
        }

        const payload = equations.map(eq => ({
            equation_id: eq.id,
            equation_text: eq.equation_text,
            correct_answer: eq.correct_answer,
            user_answer: Number(answers[eq.id]),
            equation_type_id: eq.equation_type_id,
        }));

        const res = await fetch(BASE_URL + "/game/check", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify(payload),
        });

        const data = await res.json();
        setFeedback(data);
    };

    // завершение уровня
    const handleFinish = async () => {
        const res = await fetch(BASE_URL + "/game/finish-level", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                feedback: feedback,
                section_id: sectionId,
                level_order: levelOrder,
            }),
        });

        const data = await res.json();
        setResult(data);
    };

    if (!equations.length) return <div>Loading...</div>;

    return (
        <div style={styles.page}>
            <h2>Решение уравнений</h2>

            {/* список уравнений */}
            {equations.map(eq => {
                const fb = feedback?.find(f => f.equation_id === eq.id);
                console.log('equations:', equations.map(eq => ({ id: eq.id, text: eq.equation_text })));
                return (
                    <div key={eq.id} style={styles.card}>
                        <div>{eq.equation_text}</div>

                        <input
                            type="number"
                            value={answers[eq.id] || ""}
                            onChange={(e) =>
                                handleChange(eq.id, e.target.value)
                            }
                        />

                        {/* feedback */}
                        {fb && (
                            <div
                                style={{
                                    color: fb.is_correct ? "green" : "red",
                                }}
                            >
                                {fb.is_correct
                                    ? "✔"
                                    : `✘ (правильно: ${fb.correct_answer})`}
                            </div>
                        )}
                    </div>
                );
            })}

            {/* кнопки */}
            {!feedback && (
                <button onClick={handleCheck}>
                    Проверить
                </button>
            )}

            {feedback && !result && (
                <button onClick={handleFinish}>
                    Завершить уровень
                </button>
            )}

            {/* результат */}
            {result && (
                <div style={styles.result}>
                    <h3>Результат</h3>
                    <div>⭐ {result.stars}</div>
                    <div>XP: {result.common_xp}</div>
                </div>
            )}
        </div>
    );
}

export default GamePage;

const styles = {
    page: {
        padding: "20px",
        backgroundColor: "white",
        minHeight: "100vh",
    },

    card: {
        border: "1px solid #ddd",
        padding: "10px",
        marginBottom: "10px",
    },

    result: {
        marginTop: "20px",
        padding: "10px",
        border: "2px solid #4da6ff",
    },
};
