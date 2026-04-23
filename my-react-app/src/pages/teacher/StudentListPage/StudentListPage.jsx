import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "./../../../constants";
import { SchoolSelect } from "../../../components/SchoolSelect";
import { ClassSelect } from "../../../components/ClassSelect";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./StudentListPage.module.css";

function StudentsPage() {
  const navigate = useNavigate();
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
    setForm((prev) => ({ ...prev, [field]: value }));
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

    try {
      const res = await fetch(url, {
        method: method,
        headers: { "Content-Type": "application/json" },
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
      alert(form.id ? "Студент обновлён!" : "Студент создан!");

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
    if (!window.confirm("Вы уверены, что хотите удалить этого студента?")) return;

    try {
      const res = await fetch(`${BASE_URL}/teacher/students/${id}`, {
        method: "DELETE",
        credentials: "include",
      });

      if (!res.ok) {
        alert("Ошибка удаления");
        return;
      }

      alert("Студент удалён!");
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
      class_id: String(student.class_id),
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
    <div className={styles.page}>
      <div className={styles.container}>
        <div className={styles.header}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/teacher/dashboard")}
          >
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={styles.title}>Управление учениками</h1>

        {/* Форма добавления/редактирования */}
        <div className={styles.formCard}>
          <h3 className={styles.formTitle}>
            {form.id ? "Редактирование ученика" : "Добавить нового ученика"}
          </h3>
          <form onSubmit={handleSubmit} className={styles.form}>
            <div className={styles.inputGroup}>
              <label>ФИО</label>
              <input
                className={styles.inputField}
                placeholder="Иванов Иван Иванович"
                value={form.fullname}
                onChange={(e) => handleChange("fullname", e.target.value)}
                required
              />
            </div>

            <div className={styles.inputGroup}>
              <label>Логин</label>
              <input
                className={styles.inputField}
                placeholder="ivan.ivanov"
                value={form.login}
                onChange={(e) => handleChange("login", e.target.value)}
                required
              />
            </div>

            <div className={styles.inputGroup}>
              <label>Почта</label>
              <input
                className={styles.inputField}
                type="email"
                placeholder="ivan@school.ru"
                value={form.email}
                onChange={(e) => handleChange("email", e.target.value)}
                required
              />
            </div>

            <div className={styles.inputGroup}>
              <label>Школа</label>
              <SchoolSelect
                value={selectedSchool}
                onChange={setSelectedSchool}
                className={styles.selectField}
              />
            </div>

            <div className={styles.inputGroup}>
              <label>Класс</label>
              <ClassSelect
                schoolId={selectedSchool}
                value={form.class_id}
                onChange={(value) => handleChange("class_id", value)}
                className={styles.selectField}
              />
            </div>

            <div className={styles.formActions}>
              <button
                type="submit"
                disabled={loading}
                className={styles.formButton}
              >
                {loading ? "Сохранение..." : form.id ? "Обновить" : "Создать"}
              </button>
              {form.id && (
                <button
                  type="button"
                  onClick={resetForm}
                  className={`${styles.formButton} ${styles.cancelButton}`}
                >
                  Отмена
                </button>
              )}
            </div>
          </form>
        </div>

        {/* Список студентов */}
        <h3 className={styles.formTitle}>Список учеников</h3>
        {students.length === 0 ? (
          <div className={styles.emptyMessage}>
            Нет студентов. Создайте первого студента!
          </div>
        ) : (
          <div className={styles.tableWrapper}>
            <table className={styles.studentsTable}>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>ФИО</th>
                  <th>Логин</th>
                  <th>Почта</th>
                  <th>ID класса</th>
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
                      <div className={styles.actionButtons}>
                        <button
                          className={styles.editButton}
                          onClick={() => handleEdit(student)}
                        >
                          Редактировать
                        </button>
                        <button
                          className={styles.deleteButton}
                          onClick={() => handleDelete(student.id)}
                        >
                          Удалить
                        </button>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>
    </div>
  );
}

export default StudentsPage;