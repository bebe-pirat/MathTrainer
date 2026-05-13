import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../../constants";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import UserModal from "../../../components/UserModal/UserModal";
import PasswordModal from "../../../components/PasswordModal/PasswordModal";
import CredentialsModal from "../../../components/CredentialsModal/CredentialsModal";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./UsersPage.module.css";

function UsersPage() {
  const navigate = useNavigate();
  const [users, setUsers] = useState([]);
  const [loading, setLoading] = useState(false);
  const [updatingUserId, setUpdatingUserId] = useState(null);
  const [error, setError] = useState(null);

  // состояния модальных окон
  const [showUserModal, setShowUserModal] = useState(false);
  const [editingUser, setEditingUser] = useState(null);
  const [showPasswordModal, setShowPasswordModal] = useState(false);
  const [passwordUserId, setPasswordUserId] = useState(null);
  const [showCredentialsModal, setShowCredentialsModal] = useState(false);
  const [credentials, setCredentials] = useState({ login: "", password: "", title: "" });

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

  // Создание/обновление пользователя
  const handleSaveUser = async (formData) => {
    try {
      const method = editingUser ? "PUT" : "POST";
      const url = editingUser
        ? `${BASE_URL}/admin/users/${editingUser.id}`
        : `${BASE_URL}/admin/users`;
      const response = await fetch(url, {
        method,
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          email: formData.email,
          login: formData.login,
          role_id: parseInt(formData.role_id),
          fullname: formData.fullname,
          class_id: formData.class_id ? parseInt(formData.class_id) : null,
        }),
      });
      if (!response.ok) throw new Error("Ошибка сохранения пользователя");
      const data = await response.json();
      if (!editingUser) {
        // после создания показываем логин и пароль
        setCredentials({
          login: data.login || formData.login,
          password: data.password,
          title: "Пользователь создан",
        });
        setShowCredentialsModal(true);
      }
      await fetchUsers();
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  // Смена пароля
  const handlePasswordUpdated = async (login, newPassword) => {
    setCredentials({
      login: login,
      password: newPassword,
      title: "Пароль обновлён",
    });
    setShowCredentialsModal(true);
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Удалить пользователя?")) return;
    try {
      const response = await fetch(`${BASE_URL}/admin/users/${id}`, {
        method: "DELETE",
        credentials: "include",
      });
      if (!response.ok) throw new Error("Ошибка удаления");
      await fetchUsers();
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  const openEditModal = (user) => {
    setEditingUser(user);
    setShowUserModal(true);
  };

  const openCreateModal = () => {
    setEditingUser(null);
    setShowUserModal(true);
  };

  const openPasswordModal = (userId) => {
    setPasswordUserId(userId);
    setShowPasswordModal(true);
  };

  const formatDate = (dateString) => {
    if (!dateString) return "-";
    const date = new Date(dateString);
    return date.toLocaleString("ru-RU");
  };

  if (loading) return <div className={sharedStyles.loader}>Загрузка пользователей...</div>;
  if (error) return (
    <div className={sharedStyles.errorBox}>
      <p>Ошибка: {error}</p>
      <button className={sharedStyles.formButton} onClick={() => navigate("/admin/dashboard")}>
        Назад
      </button>
    </div>
  );

  return (
    <div className={sharedStyles.adminPage}>
      <div className={sharedStyles.adminContainer}>
        <div className={sharedStyles.adminHeader}>
          <button className={sharedStyles.headerButton} onClick={() => navigate("/admin/dashboard")}>
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.adminTitle}>Управление пользователями</h1>

        <div className={styles.toolbar}>
          <button className={sharedStyles.formButton} onClick={openCreateModal}>
            + Добавить пользователя
          </button>
        </div>

        <div className={sharedStyles.formCard}>
          <h3>Список пользователей системы</h3>
          {users.length === 0 ? (
            <div className={sharedStyles.emptyMessage}>Нет данных о пользователях</div>
          ) : (
            <div className={sharedStyles.dataTableWrapper}>
              <table className={sharedStyles.dataTable}>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>ФИО</th>
                    <th>Логин</th>
                    <th>Email</th>
                    <th>Роль</th>
                    <th>Создан</th>
                    <th>Статус</th>
                    <th>Действия</th>
                  </tr>
                </thead>
                <tbody>
                  {users.map((user) => (
                    <tr key={user.id} className={user.blocked ? sharedStyles.blockedRow : ""}>
                      <td>{user.id}</td>
                      <td>{user.fullname || "-"}</td>
                      <td>{user.login || "-"}</td>
                      <td>{user.email || "-"}</td>
                      <td>{user.role_name || user.role || "-"}</td>
                      <td>{formatDate(user.created_at)}</td>
                      <td>
                        <label style={{ display: "flex", alignItems: "center", gap: "8px" }}>
                          <input
                            type="checkbox"
                            className={sharedStyles.checkbox}
                            checked={user.blocked === true}
                            onChange={() => {/* блокировка через отдельный обработчик */}}
                            disabled
                          />
                          <span className={user.blocked ? sharedStyles.statusBlocked : sharedStyles.statusActive}>
                            {user.blocked ? "Заблокирован" : "Активен"}
                          </span>
                        </label>
                      </td>
                      <td>
                        <div className={styles.actionButtons}>
                          <button className={sharedStyles.smallButton} onClick={() => openEditModal(user)}>
                            Редактировать
                          </button>
                          <button className={sharedStyles.smallButton} onClick={() => handleDelete(user.id)}>
                            Удалить
                          </button>
                          <button className={sharedStyles.smallButton} onClick={() => openPasswordModal(user.id)}>
                            Обновить пароль
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

      <UserModal
        isOpen={showUserModal}
        onClose={() => setShowUserModal(false)}
        onSave={handleSaveUser}
        initialData={editingUser}
      />

      <PasswordModal
        isOpen={showPasswordModal}
        onClose={() => setShowPasswordModal(false)}
        userId={passwordUserId}
        onPasswordUpdated={handlePasswordUpdated}
      />

      <CredentialsModal
        isOpen={showCredentialsModal}
        onClose={() => setShowCredentialsModal(false)}
        title={credentials.title}
        login={credentials.login}
        password={credentials.password}
      />
    </div>
  );
}

export default UsersPage;