import { useEffect, useState } from "react";
import { BASE_URL } from "../constants";

function UsersPage() {
    const [users, setUsers] = useState([]);
    const [loading, setLoading] = useState(false);
    const [updatingUserId, setUpdatingUserId] = useState(null);

    const fetchUsers = async () => {
        setLoading(true);
        try {
            const response = await fetch(BASE_URL + "/admin/users", {
                method: "GET",
                credentials: "include",
            });

            if (!response.ok) {
                console.error("Ошибка загрузки пользователей");
                return;
            }

            const data = await response.json();
            setUsers(data);
            console.log(data);
        } catch (err) {
            console.error(err);
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
                headers: {
                    "Content-Type": "application/json",
                },
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

            setUsers(prevUsers => 
                prevUsers.map(user => 
                    user.id === userId 
                        ? { ...user, blocked: !currentBlockedStatus }
                        : user
                )
            );
            
            const newStatus = !currentBlockedStatus;
            alert(`Пользователь ${userId} ${newStatus ? 'заблокирован' : 'разблокирован'}`);
            
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
        return date.toLocaleString('ru-RU');
    };

    if (loading) {
        return <div>Загрузка...</div>;
    }

    return (
        <div>
            <h2>Пользователи</h2>

            {users.length === 0 ? (
                <p>Нет данных о пользователях</p>
            ) : (
                <table border="1" cellPadding="8" cellSpacing="0" style={{ width: "100%", borderCollapse: "collapse" }}>
                    <thead>
                        <tr style={{ backgroundColor: "#f2f2f2" }}>
                            <th>ID</th>
                            <th>Имя пользователя</th>
                            <th>Email</th>
                            <th>Роль</th>
                            <th>Создан в</th>
                            <th>Статус блокировки</th>
                        </tr>
                    </thead>

                    <tbody>
                        {users.map((user) => (
                            <tr key={user.id} style={{ backgroundColor: user.blocked ? "#ffe6e6" : "white" }}>
                                <td>{user.id}</td>
                                <td>{user.login || "-"}</td>
                                <td>{user.email || "-"}</td>
                                <td>{user.role || "user"}</td>
                                <td>{formatDate(user.created_at)}</td>
                                <td>
                                    <label style={{ display: "flex", alignItems: "center", gap: "8px" }}>
                                        <input
                                            type="checkbox"
                                            checked={user.blocked === true}
                                            onChange={() => handleBlockToggle(user.id, user.blocked)}
                                            disabled={updatingUserId === user.id}
                                        />
                                        <span style={{ 
                                            color: user.blocked ? "red" : "green",
                                            fontWeight: "bold"
                                        }}>
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
            )}
        </div>
    );
}

export default UsersPage;