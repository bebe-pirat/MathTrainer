import { useEffect, useState } from "react";
import { BASE_URL } from "../constants";

function TeachersPage() {
    const [teachers, setTeachers] = useState([]);
    const [email, setEmail] = useState("");
    const [login, setLogin] = useState("");
    const [fullname, setFullname] = useState("");
    const [classId, setClassId] = useState(0);
    const [classes, setClasses] = useState([]);
    const [schools, setSchools] = useState([]);
    const [selectedSchoolId, setSelectedSchoolId] = useState("");
    const [showCredentialsModal, setShowCredentialsModal] = useState(false);
    const [generatedCredentials, setGeneratedCredentials] = useState({ login: "", password: "" });
   
    const fetchTeachers = async () => {
        try {
            const response = await fetch(BASE_URL + "/admin/teachers", {
                method: "GET",
                credentials: "include",
            });

            if (!response.ok) {
                console.error("Ошибка загрузки учитлеей");
                return;
            }

            const data = await response.json();
            setTeachers(data);
            console.log(data);
        } catch (err) {
            console.error(err);
        }
    };

    const fetchSchools = async () => {
        try {
            const response = await fetch(BASE_URL + "/admin/schools", {
                method: "GET", 
                credentials: "include",
            });

            if (!response.ok) {
                console.error("ОШибка загрузки школ");
                return;
            }

            const data = await response.json();
            setSchools(data);
        } catch (err) {
            console.error(err);
        }
    }

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

            if (!response.ok) {
                console.error("Ошибка загрузки классов");
                return;
            }

            const data = await response.json();
            setClasses(data);
            console.log(data);
            setClassId(0);
        } catch (err) {
            console.error(err);
        }
    };

    const handleSchoolChange = (e) => {
        const schoolId = e.target.value;
        setSelectedSchoolId(schoolId);
        fetchClasses(schoolId);
    };

    useEffect(() => {
        fetchTeachers();
        fetchSchools();
    }, []);

    const handleCreate = async (e) => {
        e.preventDefault();

        try {
            console.log(fullname);
            const response = await fetch(BASE_URL + "/admin/teachers", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({
                    email: email, 
                    login: login,
                    fullname: fullname, 
                    class_id: classId,
                }),
            });

            if (!response.ok) {
                alert("Ошибка создания");
                return;
            }

            const data = await response.json();
            
            setGeneratedCredentials({
                login: data.login || login,
                password: data.password || "Пароль сгенерирован"
            });
            setShowCredentialsModal(true);

            setEmail("");
            setLogin("");
            setFullname("");
            setClassId(0);
            setSelectedSchoolId("");

            fetchTeachers();
        } catch (err) {
            console.error(err);
        }
    };

    const closeModal = () => {
        setShowCredentialsModal(false);
        setGeneratedCredentials({ login: "", password: "" });
    };

    return (
        <div>
            <h2>Учителя</h2>

            <form onSubmit={handleCreate}>
                <input
                    value={fullname}
                    onChange={(e) => setFullname(e.target.value)}
                    placeholder="ФИО"
                    required
                />

                <input
                    value={email}
                    onChange={(e) => setEmail(e.target.value)}
                    placeholder="Почта"
                    type="email"
                    required
                />

                <input
                    value={login}
                    onChange={(e) => setLogin(e.target.value)}
                    placeholder="Логин"
                    required
                />

                <select
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

                <select
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

                <button type="submit">Создать</button>
            </form>

            {showCredentialsModal && (
                <div>
                    <div>
                        <button onClick={closeModal}>✕</button>
                        <h3>Учитель успешно создан!</h3>
                        <div>
                            <strong>Логин:</strong> {generatedCredentials.login}
                        </div>
                        <div>
                            <strong>Пароль:</strong> {generatedCredentials.password}
                        </div>
                        <div>
                            Сохраните эти данные. Пароль будет отображаться только один раз.
                        </div>
                        <button onClick={closeModal}>Закрыть</button>
                    </div>
                </div>
            )}

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
                             <td>{teacher.blocked ? "Да" : "Нет"}</td>
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