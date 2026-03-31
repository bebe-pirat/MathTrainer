import { useEffect, useState } from "react";
import { BASE_URL } from "../constants";

function ClassesPage() {
    const [classes, setClasses] = useState([]);
    const [name, setName] = useState("");
    const [grade, setGrade] = useState("");
    const [schoolId, setSchoolId] = useState("");
    const [schools, setSchools] = useState([]);
   
    const fetchClasses = async () => {
        try {
            const response = await fetch(BASE_URL + "/admin/classes", {
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
                console.error("Ошибка загрузки школ");
                return;
            }

            const data = await response.json();
            setSchools(data);
        } catch (err) {
            console.error(err);
        }
    };

    useEffect(() => {
        fetchClasses();
        fetchSchools();
    }, []);

    const handleCreate = async (e) => {
        e.preventDefault();

        try {
            const response = await fetch(BASE_URL + "/admin/classes", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                },
                credentials: "include",
                body: JSON.stringify({
                    name: name,
                    grade: parseInt(grade),
                    school_id: parseInt(schoolId),
                }),
            });

            if (!response.ok) {
                alert("Ошибка создания");
                return;
            }

            const data = await response.json();
            
            setName("");
            setGrade("");
            setSchoolId("");

            fetchClasses();
        } catch (err) {
            console.error(err);
        }
    };

    return (
        <div>
            <h2>Классы</h2>

            <form onSubmit={handleCreate}>
                <input
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="Название класса"
                    required
                />

                <input
                    value={grade}
                    onChange={(e) => setGrade(e.target.value)}
                    placeholder="Степень"
                    type="number"
                    required
                />

                <select
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

                <button type="submit">Создать</button>
            </form>

            <table>
                <thead>
                    <tr>
                        <th>id</th>
                        <th>Название</th>
                        <th>Степень</th>
                        <th>ID школы</th>
                        <th>Создан в</th>
                    </tr>
                </thead>

                <tbody>
                    {classes.map((cls) => (
                        <tr key={cls.id}>
                            <td>{cls.id}</td>
                            <td>{cls.name}</td>
                            <td>{cls.grade}</td>
                            <td>{cls.school_id}</td>
                            <td>{cls.created_at}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default ClassesPage;