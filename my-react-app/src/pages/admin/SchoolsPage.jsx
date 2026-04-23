import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../constants";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../styles/shared.module.css";

function SchoolsPage() {
  const navigate = useNavigate();
  const [schools, setSchools] = useState([]);
  const [name, setName] = useState("");
  const [address, setAddress] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);

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
    fetchSchools();
  }, []);

  const handleCreate = async (e) => {
    e.preventDefault();
    setLoading(true);
    try {
      const response = await fetch(BASE_URL + "/admin/schools", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ name, address }),
      });
      if (!response.ok) throw new Error("Ошибка создания школы");
      await response.json();
      setName("");
      setAddress("");
      await fetchSchools();
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

        <h1 className={sharedStyles.adminTitle}>Управление школами</h1>

        {/* Форма создания школы */}
        <div className={sharedStyles.formCard}>
          <h3>Добавить новую школу</h3>
          <form onSubmit={handleCreate}>
            <div className={sharedStyles.formGrid}>
              <div className={sharedStyles.formGroup}>
                <label>Название школы</label>
                <input
                  className={sharedStyles.formInput}
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  placeholder="МБОУ СОШ №1"
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Адрес</label>
                <input
                  className={sharedStyles.formInput}
                  value={address}
                  onChange={(e) => setAddress(e.target.value)}
                  placeholder="г. Москва, ул. Ленина, 1"
                  required
                />
              </div>
              <div className={sharedStyles.formActions}>
                <button
                  type="submit"
                  className={sharedStyles.formButton}
                  disabled={loading}
                >
                  {loading ? "Создание..." : "Создать школу"}
                </button>
              </div>
            </div>
          </form>
        </div>

        {/* Список школ */}
        <h3 className={sharedStyles.adminTitle} style={{ fontSize: "1.5rem", marginBottom: "20px" }}>
          Существующие школы
        </h3>
        {schools.length === 0 ? (
          <div className={sharedStyles.emptyMessage}>
            Нет школ. Создайте первую школу!
          </div>
        ) : (
          <div className={sharedStyles.dataTableWrapper}>
            <table className={sharedStyles.dataTable}>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Название</th>
                  <th>Адрес</th>
                  <th>Создана</th>
                </tr>
              </thead>
              <tbody>
                {schools.map((school) => (
                  <tr key={school.id}>
                    <td>{school.id}</td>
                    <td>{school.name}</td>
                    <td>{school.address}</td>
                    <td>{new Date(school.created_at).toLocaleDateString()}</td>
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

export default SchoolsPage;