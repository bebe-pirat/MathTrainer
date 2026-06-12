import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../../constants";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../../styles/shared.module.css";

function DirectorClassesList() {
  const navigate = useNavigate();
  const [classes, setClasses] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchClasses();
  }, []);

  const fetchClasses = async () => {
    try {
      const res = await fetch(`${BASE_URL}/director/classes`, {
        credentials: "include",
      });
      if (!res.ok) throw new Error("Ошибка загрузки списка классов");
      const data = await res.json();
      setClasses(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const handleClassStats = (classId) => {
    navigate(`/director/class-stats/${classId}`);
  };

  const formatDate = (dateString) => {
    if (!dateString) return "-";
    const date = new Date(dateString);
    return date.toLocaleDateString("ru-RU");
  };

  if (loading) return <div className={sharedStyles.loader}>Загрузка списка классов...</div>;
  if (error)
    return (
      <div className={sharedStyles.errorBox}>
        <p>Ошибка: {error}</p>
        <button
          className={sharedStyles.formButton}
          onClick={() => navigate("/director/dashboard")}
        >
          Назад
        </button>
      </div>
    );

  return (
    <div className={sharedStyles.dashboardPage}>
      <div className={sharedStyles.dashboardContainer}>
        <div className={sharedStyles.dashboardHeader}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/director/dashboard")}
          >
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.dashboardTitle}>Классы школы</h1>

        <div className={sharedStyles.formCard}>
          {classes.length === 0 ? (
            <div className={sharedStyles.emptyMessage}>Нет классов</div>
          ) : (
            <div className={sharedStyles.dataTableWrapper}>
              <table className={sharedStyles.dataTable}>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Название</th>
                    <th>Степень (класс)</th>
                    <th>Дата создания</th>
                    <th>Действия</th>
                  </tr>
                </thead>
                <tbody>
                  {classes.map((cls) => (
                    <tr key={cls.id}>
                      <td>{cls.id}</td>
                      <td>{cls.name}</td>
                      <td>{cls.grade} класс</td>
                      <td>{formatDate(cls.created_at)}</td>
                      <td>
                        <button
                          className={sharedStyles.smallButton}
                          onClick={() => handleClassStats(cls.id)}
                        >
                          Статистика
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>
      </div>
    </div>
  );
}

export default DirectorClassesList;