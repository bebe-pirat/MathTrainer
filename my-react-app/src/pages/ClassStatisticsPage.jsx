import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../constants";

function ClassStatistics() {
    const navigate = useNavigate();
    const [stats, setStats] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        fetchClassStats();
    }, []);

    const fetchClassStats = async () => {
        try {
            const response = await fetch(BASE_URL + "/teacher/class/stats", {
                method: "GET",
                credentials: "include",
            });

            if (!response.ok) {
                throw new Error("Ошибка загрузки статистики");
            }

            const data = await response.json();
            setStats(data);
        } catch (err) {
            setError(err.message);
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    if (loading) {
        return <div>Загрузка статистики...</div>;
    }

    if (error) {
        return (
            <div>
                <p>Ошибка: {error}</p>
                <button onClick={() => navigate("/teacher/dashboard")}>Назад</button>
            </div>
        );
    }

    if (!stats) {
        return <div>Нет данных</div>;
    }

    // const maxAttempts = Math.max(...stats.students.map(s => s.total_attempts), 0);
    // const maxCorrect = Math.max(...stats.students.map(s => s.correct_answers), 0);

    return (
        <div>
            <div>
                <button onClick={() => navigate("/teacher/dashboard")}>Назад</button>
            </div>

            <h1>Статистика класса</h1>

            {/* Общая статистика */}
            <div>
                <h3>Общие показатели</h3>
                <div>
                    <div>Количество учеников: {stats.students_count}</div>
                    <div>Всего попыток: {stats.total_attempts}</div>
                    <div>Правильных ответов: {stats.correct_answers}</div>
                    <div>Неправильных ответов: {stats.wrong_answers}</div>
                    <div>Общая точность: {stats.accuracy_percent.toFixed(1)}%</div>
                </div>
            </div>

            {/* График точности по типам уравнений */}
            {stats.equation_types_stats && stats.equation_types_stats.length > 0 && (
                <div>
                    <h3>Точность по типам уравнений</h3>
                    <div>
                        {stats.equation_types_stats.map((type, index) => (
                            <div key={index} style={{ marginBottom: "15px" }}>
                                <div>{type.type}</div>
                                <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
                                    <div style={{ 
                                        width: `${type.accuracy_percent}%`, 
                                        height: "30px", 
                                        backgroundColor: type.accuracy_percent >= 70 ? "#4caf50" : 
                                                       type.accuracy_percent >= 40 ? "#ff9800" : "#f44336",
                                        minWidth: "5px"
                                    }} />
                                    <span>{type.accuracy_percent}%</span>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            )}

            {/* Таблица учеников */}
            {stats.students && stats.students.length > 0 && (
                <div>
                    <h3>Успеваемость учеников</h3>
                    <table border="1" cellPadding="8" cellSpacing="0">
                        <thead>
                            <tr>
                                <th>ID ученика</th>
                                <th>Имя</th>
                                <th>Точность</th>
                                <th>Пройдено уровней</th>
                                <th>Всего попыток</th>
                                <th>Правильно</th>
                                <th>Неправильно</th>
                            </tr>
                        </thead>
                        <tbody>
                            {stats.students.map((student) => (
                                <tr key={student.student_id}>
                                    <td>{student.student_id}</td>
                                    <td>{student.name}</td>
                                    <td>
                                        <div style={{ display: "flex", alignItems: "center", gap: "5px" }}>
                                            <div style={{ 
                                                width: "100px", 
                                                height: "20px", 
                                                backgroundColor: "#f0f0f0"
                                            }}>
                                                <div style={{ 
                                                    width: `${student.accuracy}%`, 
                                                    height: "100%", 
                                                    backgroundColor: student.accuracy >= 70 ? "#4caf50" : 
                                                                   student.accuracy >= 40 ? "#ff9800" : "#f44336"
                                                }} />
                                            </div>
                                            <span>{student.accuracy}%</span>
                                        </div>
                                    </td>
                                    <td>{student.levels_complited}</td>
                                    <td>{student.total_attempts}</td>
                                    <td style={{ color: "#4caf50" }}>{student.correct_answers}</td>
                                    <td style={{ color: "#f44336" }}>{student.wrong_answers}</td>
                                </tr>
                            ))}
                        </tbody>
                    </table>
                </div>
            )}

            {/* График сравнения учеников */}
            {stats.students && stats.students.length > 0 && (
                <div>
                    <h3>Сравнение точности учеников</h3>
                    <div>
                        {stats.students.map((student) => (
                            <div key={student.student_id} style={{ marginBottom: "10px" }}>
                                <div>{student.name}</div>
                                <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
                                    <div style={{ 
                                        width: `${student.accuracy}%`, 
                                        height: "25px", 
                                        backgroundColor: student.accuracy >= 70 ? "#4caf50" : 
                                                       student.accuracy >= 40 ? "#ff9800" : "#f44336",
                                        minWidth: "5px"
                                    }} />
                                    <span>{student.accuracy}%</span>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            )}

            {/* График пройденных уровней */}
            {stats.students && stats.students.length > 0 && (
                <div>
                    <h3>Пройденные уровни</h3>
                    <div>
                        {stats.students.map((student) => (
                            <div key={student.student_id} style={{ marginBottom: "10px" }}>
                                <div>{student.name}</div>
                                <div style={{ display: "flex", alignItems: "center", gap: "10px" }}>
                                    <div style={{ 
                                        width: `${(student.levels_complited / Math.max(...stats.students.map(s => s.levels_complited))) * 100}%`, 
                                        height: "25px", 
                                        backgroundColor: "#2196f3",
                                        minWidth: "5px"
                                    }} />
                                    <span>{student.levels_complited}</span>
                                </div>
                            </div>
                        ))}
                    </div>
                </div>
            )}
        </div>
    );
}

export default ClassStatistics;