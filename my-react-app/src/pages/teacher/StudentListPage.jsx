import { useEffect, useState } from "react";
import { BASE_URL } from "./../../constants";
import { SchoolSelect } from "../../components/SchoolSelect";
import { ClassSelect } from "../../components/ClassSelect";

function StudentsPage() {
    const [students, setStudents] = useState([]);
    const [selectedSchool, setSelectedSchool] = useState("");
    const [loading, setLoading] = useState(false);

    const [form, setForm] = useState({
        id: null,
        fullname: "",
        login: "",
        email: "",
        class_id: "", 
    });

    const fetchStudents = async () => {
        try {
            const res = await fetch(`${BASE_URL}/teacher/students`, {
                credentials: "include",
            });
            const data = await res.json();
            console.log(data);
            setStudents(data);
        } catch (err) {
            console.error("Ошибка загрузки студентов:", err);
            alert("Ошибка загрузки студентов");
        }
    };

    useEffect(() => {
        fetchStudents();
    }, []);

    const handleChange = (field, value) => {
        setForm((prev) => ({
            ...prev,
            [field]: value,
        }));
    };

    const handleSubmit = async (e) => {
        e.preventDefault();
        setLoading(true);

        const method = form.id ? "PUT" : "POST";
        
        let url = `${BASE_URL}/teacher/students`;
        if (form.id && method === "PUT") {
            url = `${BASE_URL}/teacher/students/${form.id}`;
        }

        const body = {
            fullname: form.fullname,
            login: form.login,
            email: form.email,
            class_id: form.class_id ? Number(form.class_id) : 0,
        };

        console.log("Sending request:", method, url, body);

        try {
            const res = await fetch(url, {
                method: method,
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify(body),
            });

            if (!res.ok) {
                const errorText = await res.text();
                console.error("Server error:", errorText);
                alert(`Ошибка: ${res.status}`);
                return;
            }

            const data = await res.json();
            
            alert(form.id ? "Студент обновлен!" : "Студент создан!");
            
            if (!form.id && data.password) {
                alert(`Логин: ${data.login}\nПароль: ${data.password}\nСохраните пароль!`);
            }
            
            resetForm();
            await fetchStudents();
        } catch (err) {
            console.error("Network error:", err);
            alert("Ошибка сети: " + err.message);
        } finally {
            setLoading(false);
        }
    };

    const handleDelete = async (id) => {
        if (!window.confirm("Вы уверены, что хотите удалить этого студента?")) {
            return;
        }

        try {
            const res = await fetch(`${BASE_URL}/teacher/students/${id}`, {
                method: "DELETE",
                credentials: "include",
            });

            if (!res.ok) {
                alert("Ошибка удаления");
                return;
            }

            alert("Студент удален!");
            await fetchStudents();
        } catch (err) {
            console.error("Error deleting:", err);
            alert("Ошибка при удалении");
        }
    };

    const handleEdit = (student) => {
        setForm({
            id: student.id,
            fullname: student.fullname,
            login: student.login,
            email: student.email,
            class_id: String(student.class_id), // Преобразуем в строку для select
        });
        
        if (student.school_id) {
            setSelectedSchool(String(student.school_id));
        }
    };

    const resetForm = () => {
        setForm({
            id: null,
            fullname: "",
            login: "",
            email: "",
            class_id: "",
        });
        setSelectedSchool("");
    };

    return (
        <div>
            <h2>Студенты</h2>

            <form onSubmit={handleSubmit}>
                <input
                    placeholder="ФИО"
                    value={form.fullname}
                    onChange={(e) => handleChange("fullname", e.target.value)}
                    required
                />

                <input
                    placeholder="Логин"
                    value={form.login}
                    onChange={(e) => handleChange("login", e.target.value)}
                    required
                />

                <input
                    placeholder="Почта"
                    type="email"
                    value={form.email}
                    onChange={(e) => handleChange("email", e.target.value)}
                    required
                />

                <SchoolSelect 
                    value={selectedSchool} 
                    onChange={setSelectedSchool}
                />

                <ClassSelect 
                    schoolId={selectedSchool}
                    value={form.class_id} 
                    onChange={(value) => handleChange("class_id", value)}
                />

                <button type="submit" disabled={loading}>
                    {loading ? "Загрузка..." : (form.id ? "Обновить" : "Создать")}
                </button>

                {form.id && (
                    <button type="button" onClick={resetForm}>
                        Отмена
                    </button>
                )}
            </form>

            <h3>Список студентов</h3>
            <table border="1" cellPadding="8" cellSpacing="0">
                <thead>
                    <tr>
                        <th>ID</th>
                        <th>ФИО</th>
                        <th>Логин</th>
                        <th>Почта</th>
                        <th>ID Класса</th>
                        <th>Действия</th>
                    </tr>
                </thead>
                <tbody>
                    {students.map((student) => (
                        <tr key={student.id}>
                            <td>{student.id}</td>
                            <td>{student.fullname}</td>
                            <td>{student.login}</td>
                            <td>{student.email}</td>
                            <td>{student.class_id}</td>
                            <td>
                                <button onClick={() => handleEdit(student)}>
                                    Редактировать
                                </button>
                                <button onClick={() => handleDelete(student.id)}>
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