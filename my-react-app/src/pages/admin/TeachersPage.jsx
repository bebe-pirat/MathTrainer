import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../constants";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../styles/shared.module.css";

function TeachersPage() {
  const navigate = useNavigate();
  const [teachers, setTeachers] = useState([]);
  const [email, setEmail] = useState("");
  const [login, setLogin] = useState("");
  const [fullname, setFullname] = useState("");
  const [classId, setClassId] = useState(0);
  const [classes, setClasses] = useState([]);
  const [schools, setSchools] = useState([]);
  const [selectedSchoolId, setSelectedSchoolId] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [showCredentialsModal, setShowCredentialsModal] = useState(false);
  const [generatedCredentials, setGeneratedCredentials] = useState({ login: "", password: "" });

  const fetchTeachers = async () => {
    try {
      const response = await fetch(BASE_URL + "/admin/teachers", {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка загрузки учителей");
      const data = await response.json();
      setTeachers(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    }
  };

  const fetchSchools = async () => {
    try {
      const response = await fetch(BASE_URL + "/admin/schools", {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка загрузки школ");
      const data = await response.json();
      setSchools(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    }
  };

  const fetchClasses = async (schoolId) => {
    if (!schoolId) {
      setClasses([]);
      setClassId(0);
      return;
    }
    try {
      const response = await fetch(`${BASE_URL}/admin/classes?school_id=${schoolId}`, {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка загрузки классов");
      const data = await response.json();
      setClasses(data);
      setClassId(0);
    } catch (err) {
      console.error(err);
      setError(err.message);
    }
  };

  const handleSchoolChange = (e) => {
    const schoolId = e.target.value;
    setSelectedSchoolId(schoolId);
    fetchClasses(schoolId);
  };

  useEffect(() => {
    Promise.all([fetchTeachers(), fetchSchools()]);
  }, []);

  const handleCreate = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      const response = await fetch(BASE_URL + "/admin/teachers", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          email,
          login,
          fullname,
          class_id: classId,
        }),
      });
      if (!response.ok) throw new Error("Ошибка создания учителя");
      const data = await response.json();

      setGeneratedCredentials({
        login: data.login || login,
        password: data.password || "Пароль сгенерирован",
      });
      setShowCredentialsModal(true);

      // сброс формы
      setEmail("");
      setLogin("");
      setFullname("");
      setClassId(0);
      setSelectedSchoolId("");
      setClasses([]);

      await fetchTeachers();
    } catch (err) {
      console.error(err);
      alert(err.message);
    } finally {
      setLoading(false);
    }
  };

  const closeModal = () => {
    setShowCredentialsModal(false);
    setGeneratedCredentials({ login: "", password: "" });
  };

  if (error) {
    return (
      <div className={sharedStyles.errorBox}>
        <p>Ошибка: {error}</p>
        <button
          className={sharedStyles.formButton}
          onClick={() => navigate("/admin/dashboard")}
        >
          Назад
        </button>
      </div>
    );
  }

  return (
    <div className={sharedStyles.adminPage}>
      <div className={sharedStyles.adminContainer}>
        <div className={sharedStyles.adminHeader}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/admin/dashboard")}
          >
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.adminTitle}>Управление учителями</h1>

        {/* Форма создания учителя */}
        <div className={sharedStyles.formCard}>
          <h3>Добавить нового учителя</h3>
          <form onSubmit={handleCreate}>
            <div className={sharedStyles.formGrid}>
              <div className={sharedStyles.formGroup}>
                <label>ФИО</label>
                <input
                  className={sharedStyles.formInput}
                  value={fullname}
                  onChange={(e) => setFullname(e.target.value)}
                  placeholder="Иванов Иван Иванович"
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Почта</label>
                <input
                  className={sharedStyles.formInput}
                  type="email"
                  value={email}
                  onChange={(e) => setEmail(e.target.value)}
                  placeholder="ivan@school.ru"
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Логин</label>
                <input
                  className={sharedStyles.formInput}
                  value={login}
                  onChange={(e) => setLogin(e.target.value)}
                  placeholder="ivan.ivanov"
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Школа</label>
                <select
                  className={sharedStyles.formSelect}
                  value={selectedSchoolId}
                  onChange={handleSchoolChange}
                  required
                >
                  <option value="">Выберите школу</option>
                  {schools.map((school) => (
                    <option key={school.id} value={school.id}>
                      {school.name || school.fullname || `Школа ${school.id}`}
                    </option>
                  ))}
                </select>
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Класс</label>
                <select
                  className={sharedStyles.formSelect}
                  value={classId}
                  onChange={(e) => setClassId(Number(e.target.value))}
                  disabled={!selectedSchoolId}
                  required
                >
                  <option value="0">Выберите класс</option>
                  {classes.map((cls) => (
                    <option key={cls.id} value={cls.id}>
                      {cls.name || `Класс ${cls.id}`}
                    </option>
                  ))}
                </select>
              </div>
              <div className={sharedStyles.formActions}>
                <button
                  type="submit"
                  className={sharedStyles.formButton}
                  disabled={loading}
                >
                  {loading ? "Создание..." : "Создать учителя"}
                </button>
              </div>
            </div>
          </form>
        </div>

        {/* Список учителей */}
        <h3 className={sharedStyles.adminTitle} style={{ fontSize: "1.5rem", marginBottom: "20px" }}>
          Существующие учителя
        </h3>
        {teachers.length === 0 ? (
          <div className={sharedStyles.emptyMessage}>
            Нет учителей. Создайте первого учителя!
          </div>
        ) : (
          <div className={sharedStyles.dataTableWrapper}>
            <table className={sharedStyles.dataTable}>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>ФИО</th>
                  <th>Почта</th>
                  <th>Логин</th>
                  <th>Заблокирован</th>
                  <th>ID класса</th>
                  <th>Создан</th>
                  <th>Последний вход</th>
                </tr>
              </thead>
              <tbody>
                {teachers.map((teacher) => (
                  <tr key={teacher.id}>
                    <td>{teacher.id}</td>
                    <td>{teacher.fullname}</td>
                    <td>{teacher.email}</td>
                    <td>{teacher.login}</td>
                    <td>{teacher.blocked ? "Да" : "Нет"}</td>
                    <td>{teacher.class_id || "—"}</td>
                    <td>{new Date(teacher.created_at).toLocaleDateString()}</td>
                    <td>{teacher.last_login ? new Date(teacher.last_login).toLocaleString() : "—"}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        )}
      </div>

      {/* Модальное окно с учётными данными */}
      {showCredentialsModal && (
        <div className={sharedStyles.modalOverlay}>
          <div className={sharedStyles.modalContent}>
            <div className={sharedStyles.modalHeader}>
              <button className={sharedStyles.modalClose} onClick={closeModal}>
                ✕
              </button>
            </div>
            <div className={sharedStyles.modalTitle}>Учитель успешно создан!</div>
            <div className={sharedStyles.modalBody}>
              <div className={sharedStyles.credentialRow}>
                <span className={sharedStyles.credentialLabel}>Логин:</span>
                <span className={sharedStyles.credentialValue}>{generatedCredentials.login}</span>
              </div>
              <div className={sharedStyles.credentialRow}>
                <span className={sharedStyles.credentialLabel}>Пароль:</span>
                <span className={sharedStyles.credentialValue}>{generatedCredentials.password}</span>
              </div>
              <p style={{ marginTop: "15px", color: "#e74c3c", fontSize: "0.9rem" }}>
                Сохраните эти данные. Пароль будет отображаться только один раз.
              </p>
            </div>
            <div className={sharedStyles.modalFooter}>
              <button className={sharedStyles.formButton} onClick={closeModal}>
                Закрыть
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default TeachersPage;