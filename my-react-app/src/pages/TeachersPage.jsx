import { useEffect, useState } from "react";
import { BASE_URL } from "../constants";

function TeachersPage() {
    const [teachers, setTeachers] = useState([]);
    const [email, setEmail] = useState("");
    const [login, setLogin] = useState("");
    const [password, setPassword] = useState("");
    const [fullname, setFullname] = useState("");
    const [classId, SetClassId] = useState(0);
    
    const fetchTeachers = async () => {
        try {
            const response = await fetch(BASE_URL + "/admin/teachers", {
                method: "GET",
                credentials: "include",
            });

            if (!response.ok) {
                console.error("Ошибка загрузки школ");
                return;
            }

            const data = await response.json();
            setTeachers(data);
            console.log(data);
        } catch (err) {
            console.error(err);
        }
    };

    useEffect(() => {
        fetchTeachers();
    }, []);

    const handleCreate = async (e) => {
    e.preventDefault();

    try {
        const response = await fetch(BASE_URL + "/admin/teachers", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                email: email, 
                login: login, 
                password: password, 
                fullname: fullname, 
                classId: classId,
            }),
        });

            if (!response.ok) {
                alert("Ошибка создания");
                return;
            }

            setEmail("");
            setLogin("");
            setPassword("");
            setFullname("");
            SetClassId("");

            fetchTeachers();
        } catch (err) {
            console.error(err);
        }
    };

    return (
        <div>
            <h2>Учителя</h2>

            {/* создание школы */}
            <form onSubmit={handleCreate}>
                <input
                    value={fullname}
                    onChange={(e) => setFullname(e.target.value)}
                    placeholder="ФИО"
                />

                <input
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="Почта"
                />

                <input
                    value={login}
                    onChange={(e) => setLogin(e.target.value)}
                    placeholder="Логин"
                />

                <input
                    value={password}
                    onChange={(e) => setPassword(e.target.value)}
                    placeholder="Пароль"
                />

                <input
                    value={classId}
                    onChange={(e) => SetClassId(e.target.value)}
                    placeholder="Класс"
                />

                <button type="submit">Создать</button>
            </form>

            {/* таблица */}
            <table>
                <thead>
                    <tr>
                        <th>id</th>
                        <th>ФИО</th>
                        <th>почта</th>
                        <th>логин</th>
                        <th>заблокирован</th>
                        <th>класс id</th>
                        <th>создан в</th>
                        <th>last login</th>
                    </tr>
                </thead>

                <tbody>
                    {teachers.map((teacher) => (
                        <tr key={teacher.id}>
                            <td>{teacher.id}</td>
                            <td>{teacher.fullname}</td>
                            <td>{teacher.email}</td>
                            <td>{teacher.login}</td>
                            <td>{teacher.blocked}</td>
                            <td>{teacher.class_id}</td>
                            <td>{teacher.created_at}</td>
                            <td>{teacher.last_login}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default TeachersPage;
