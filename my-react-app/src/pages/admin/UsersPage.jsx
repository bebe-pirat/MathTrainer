import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../constants";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../styles/shared.module.css";

function UsersPage() {
  const navigate = useNavigate();
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(false);
  const [updatingUserId, setUpdatingUserId] = useState(null);
  const [error, setError] = useState(null);

  const fetchUsers = async () => {
    setLoading(true);
    try {
      const response = await fetch(BASE_URL + "/admin/users", {
        method: "GET",
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка загрузки пользователей");
      const data = await response.json();
      setUsers(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchUsers();
  }, []);

  const handleBlockToggle = async (userId, currentBlockedStatus) => {
    setUpdatingUserId(userId);
    try {
      const response = await fetch(BASE_URL + "/admin/users/block", {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          user_id: userId,
          blocked: !currentBlockedStatus,
        }),
      });
      if (!response.ok) {
        const errorData = await response.text();
        alert(`Ошибка изменения статуса блокировки: ${errorData}`);
        return;
      }
      setUsers((prevUsers) =>
        prevUsers.map((user) =>
          user.id === userId ? { ...user, blocked: !currentBlockedStatus } : user
        )
      );
      const newStatus = !currentBlockedStatus;
      alert(`Пользователь ${userId} ${newStatus ? "заблокирован" : "разблокирован"}`);
    } catch (err) {
      console.error(err);
      alert("Ошибка при изменении статуса блокировки");
    } finally {
      setUpdatingUserId(null);
    }
  };

  const formatDate = (dateString) => {
    if (!dateString) return "-";
    const date = new Date(dateString);
    return date.toLocaleString("ru-RU");
  };

  if (loading) {
    return <div className={sharedStyles.loader}>Загрузка пользователей...</div>;
  }

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

        <h1 className={sharedStyles.adminTitle}>Управление пользователями</h1>

        <div className={sharedStyles.formCard}>
          <h3>Список пользователей системы</h3>
          {users.length === 0 ? (
            <div className={sharedStyles.emptyMessage}>
              Нет данных о пользователях
            </div>
          ) : (
            <div className={sharedStyles.dataTableWrapper}>
              <table className={sharedStyles.dataTable}>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Логин</th>
                    <th>Email</th>
                    <th>Роль</th>
                    <th>Создан</th>
                    <th>Статус блокировки</th>
                  </tr>
                </thead>
                <tbody>
                  {users.map((user) => (
                    <tr
                      key={user.id}
                      className={user.blocked ? sharedStyles.blockedRow : ""}
                    >
                      <td>{user.id}</td>
                      <td>{user.login || "-"}</td>
                      <td>{user.email || "-"}</td>
                      <td>{user.role || "user"}</td>
                      <td>{formatDate(user.created_at)}</td>
                      <td>
                        <label style={{ display: "flex", alignItems: "center", gap: "8px" }}>
                          <input
                            type="checkbox"
                            className={sharedStyles.checkbox}
                            checked={user.blocked === true}
                            onChange={() => handleBlockToggle(user.id, user.blocked)}
                            disabled={updatingUserId === user.id}
                          />
                          <span
                            className={
                              user.blocked
                                ? sharedStyles.statusBlocked
                                : sharedStyles.statusActive
                            }
                          >
                            {updatingUserId === user.id
                              ? "Обновление..."
                              : user.blocked
                              ? "Заблокирован"
                              : "Активен"}
                          </span>
                        </label>
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

export default UsersPage;