import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../constants";
import { SchoolSelect } from "../../components/SchoolSelect";
import { ClassSelect } from "../../components/ClassSelect";

function StudentsPage() {
    const navigate = useNavigate();
    const [students, setStudents] = useState([]);
    const [email, setEmail] = useState("");
    const [login, setLogin] = useState("");
    const [fullname, setFullname] = useState("");
    const [classId, setClassId] = useState(0);
    const [selectedSchoolId, setSelectedSchoolId] = useState("");
    const [showCredentialsModal, setShowCredentialsModal] = useState(false);
    const [generatedCredentials, setGeneratedCredentials] = useState({ login: "", password: "" });
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(true);

    const fetchStudents = async () => {
        try {
            const response = await fetch(BASE_URL + "/admin/students", {
                method: "GET",
                credentials: "include",
            });

            if (!response.ok) {
                throw new Error("Ошибка загрузки студентов");
            }

            const data = await response.json();
            setStudents(data);
            console.log(data);
        } catch (err) {
            setError(err.message);
            console.error(err);
        } finally {
            setLoading(false);
        }
    };

    const handleCreate = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch(BASE_URL + "/admin/students", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({
                    email: email,
                    login: login,
                    fullname: fullname,
                    class_id: classId,
                }),
            });

            if (!response.ok) {
                const errorData = await response.json();
                alert(errorData.message || "Ошибка создания студента");
                return;
            }

            const data = await response.json();
            
            setGeneratedCredentials({
                login: data.login || login,
                password: data.password || "Пароль сгенерирован"
            });
            setShowCredentialsModal(true);

            // Очищаем форму
            setEmail("");
            setLogin("");
            setFullname("");
            setClassId(0);
            setSelectedSchoolId("");

            // Обновляем список
            fetchStudents();
        } catch (err) {
            console.error(err);
            alert("Ошибка при создании студента");
        }
    };

    const handleDelete = async (studentId) => {
        if (!window.confirm("Вы уверены, что хотите удалить этого студента?")) {
            return;
        }

        try {
            const response = await fetch(`${BASE_URL}/admin/students/${studentId}`, {
                method: "DELETE",
                credentials: "include",
            });

            if (!response.ok) {
                throw new Error("Ошибка удаления студента");
            }

            // Обновляем список
            fetchStudents();
        } catch (err) {
            console.error(err);
            alert("Ошибка при удалении студента");
        }
    };

    const handleBlock = async (studentId, isBlocked) => {
        try {
            const response = await fetch(`${BASE_URL}/admin/students/${studentId}/block`, {
                method: "PUT",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({ blocked: !isBlocked }),
            });

            if (!response.ok) {
                throw new Error("Ошибка изменения статуса");
            }

            // Обновляем список
            fetchStudents();
        } catch (err) {
            console.error(err);
            alert("Ошибка при изменении статуса студента");
        }
    };

    const closeModal = () => {
        setShowCredentialsModal(false);
        setGeneratedCredentials({ login: "", password: "" });
    };

    useEffect(() => {
        fetchStudents();
    }, []);

    if (loading) {
        return <div>Загрузка студентов...</div>;
    }

    if (error) {
        return (
            <div>
                <p>Ошибка: {error}</p>
                <button onClick={() => navigate("/admin/dashboard")}>Вернуться в панель</button>
            </div>
        );
    }

    return (
        <div>
            <h2>Студенты</h2>

            {/* Форма создания студента */}
            <form onSubmit={handleCreate}>
                <h3>Создать нового студента</h3>
                
                <input
                    value={fullname}
                    onChange={(e) => setFullname(e.target.value)}
                    placeholder="ФИО"
                    required
                />

                <input
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="Почта"
                    type="email"
                    required
                />

                <input
                    value={login}
                    onChange={(e) => setLogin(e.target.value)}
                    placeholder="Логин"
                    required
                />

                <SchoolSelect 
                    value={selectedSchoolId} 
                    onChange={setSelectedSchoolId}
                />

                <ClassSelect 
                    schoolId={selectedSchoolId}
                    value={classId} 
                    onChange={setClassId}
                />

                <button type="submit">Создать студента</button>
            </form>

            {/* Модальное окно с учетными данными */}
            {showCredentialsModal && (
                <div style={{
                    position: "fixed",
                    top: 0,
                    left: 0,
                    right: 0,
                    bottom: 0,
                    backgroundColor: "rgba(0,0,0,0.5)",
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                    zIndex: 1000
                }}>
                    <div style={{
                        backgroundColor: "white",
                        padding: "20px",
                        borderRadius: "5px",
                        maxWidth: "400px",
                        width: "100%"
                    }}>
                        <button 
                            onClick={closeModal}
                            style={{ float: "right" }}
                        >
                            ✕
                        </button>
                        <h3>Студент успешно создан!</h3>
                        <div>
                            <strong>Логин:</strong> {generatedCredentials.login}
                        </div>
                        <div>
                            <strong>Пароль:</strong> {generatedCredentials.password}
                        </div>
                        <div style={{ marginTop: "10px", color: "red" }}>
                            Сохраните эти данные. Пароль будет отображаться только один раз!
                        </div>
                        <button onClick={closeModal}>Закрыть</button>
                    </div>
                </div>
            )}

            {/* Таблица со студентами */}
            <h3>Список студентов</h3>
            <table border="1" cellPadding="8" cellSpacing="0">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>ФИО</th>
                        <th>Почта</th>
                        <th>Логин</th>
                        <th>Заблокирован</th>
                        <th>ID Класса</th>
                        <th>Создан</th>
                        <th>Последний вход</th>
                        <th>Действия</th>
                    </tr>
                </thead>
                <tbody>
                    {students.map((student) => (
                        <tr key={student.id}>
                            <td>{student.id}</td>
                            <td>{student.fullname}</td>
                            <td>{student.email}</td>
                            <td>{student.login}</td>
                            <td>{student.blocked ? "Да" : "Нет"}</td>
                            <td>{student.class_id}</td>
                            <td>{student.created_at}</td>
                            <td>{student.last_login || "Никогда"}</td>
                            <td>
                                <button 
                                    onClick={() => handleBlock(student.id, student.blocked)}
                                    style={{
                                        backgroundColor: student.blocked ? "#4CAF50" : "#ff9800",
                                        color: "white",
                                        marginRight: "5px"
                                    }}
                                >
                                    {student.blocked ? "Разблокировать" : "Заблокировать"}
                                </button>
                                <button 
                                    onClick={() => handleDelete(student.id)}
                                    style={{ backgroundColor: "#f44336", color: "white" }}
                                >
                                    Удалить
                                </button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </table>

            {students.length === 0 && (
                <p>Нет студентов. Создайте первого студента!</p>
            )}
        </div>
    );
}

export default StudentsPage;