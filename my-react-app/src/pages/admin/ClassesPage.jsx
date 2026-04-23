import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../constants";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../styles/shared.module.css";

function ClassesPage() {
  const navigate = useNavigate();
  const [classes, setClasses] = useState([]);
  const [schools, setSchools] = useState([]);
  const [name, setName] = useState("");
  const [grade, setGrade] = useState("");
  const [schoolId, setSchoolId] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

  const fetchClasses = async () => {
    try {
      const response = await fetch(BASE_URL + "/admin/classes", {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка загрузки классов");
      const data = await response.json();
      setClasses(data);
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

  useEffect(() => {
    Promise.all([fetchClasses(), fetchSchools()]);
  }, []);

  const handleCreate = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      const response = await fetch(BASE_URL + "/admin/classes", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          name: name,
          grade: parseInt(grade),
          school_id: parseInt(schoolId),
        }),
      });
      if (!response.ok) throw new Error("Ошибка создания класса");
      await response.json();
      setName("");
      setGrade("");
      setSchoolId("");
      await fetchClasses();
    } catch (err) {
      console.error(err);
      alert(err.message);
    } finally {
      setLoading(false);
    }
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

        <h1 className={sharedStyles.adminTitle}>Управление классами</h1>

        {/* Форма создания класса */}
        <div className={sharedStyles.formCard}>
          <h3>Добавить новый класс</h3>
          <form onSubmit={handleCreate}>
            <div className={sharedStyles.formGrid}>
              <div className={sharedStyles.formGroup}>
                <label>Название класса</label>
                <input
                  className={sharedStyles.formInput}
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  placeholder="Например: 11А"
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Степень (число)</label>
                <input
                  className={sharedStyles.formInput}
                  type="number"
                  value={grade}
                  onChange={(e) => setGrade(e.target.value)}
                  placeholder="11"
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Школа</label>
                <select
                  className={sharedStyles.formSelect}
                  value={schoolId}
                  onChange={(e) => setSchoolId(e.target.value)}
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
              <div className={sharedStyles.formActions}>
                <button
                  type="submit"
                  className={sharedStyles.formButton}
                  disabled={loading}
                >
                  {loading ? "Создание..." : "Создать класс"}
                </button>
              </div>
            </div>
          </form>
        </div>

        {/* Список классов */}
        <h3 className={sharedStyles.adminTitle} style={{ fontSize: "1.5rem", marginBottom: "20px" }}>
          Существующие классы
        </h3>
        {classes.length === 0 ? (
          <div className={sharedStyles.emptyMessage}>
            Нет классов. Создайте первый класс!
          </div>
        ) : (
          <div className={sharedStyles.dataTableWrapper}>
            <table className={sharedStyles.dataTable}>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Название</th>
                  <th>Степень</th>
                  <th>ID школы</th>
                  <th>Создан</th>
                </tr>
              </thead>
              <tbody>
                {classes.map((cls) => (
                  <tr key={cls.id}>
                    <td>{cls.id}</td>
                    <td>{cls.name}</td>
                    <td>{cls.grade}</td>
                    <td>{cls.school_id}</td>
                    <td>{new Date(cls.created_at).toLocaleDateString()}</td>
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

export default ClassesPage;