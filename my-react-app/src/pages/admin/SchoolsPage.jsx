import { useEffect, useState } from "react";
import { BASE_URL } from "../../constants";

function SchoolsPage() {
    const [schools, setSchools] = useState([]);
    const [name, setName] = useState("");
    const [address, setAddress] = useState("");

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
            console.log(data);
        } catch (err) {
            console.error(err);
        }
    };

    useEffect(() => {
        fetchSchools();
        console.log(schools.length);
    }, []);

    const handleCreate = async (e) => {
    e.preventDefault();

    try {
        const response = await fetch(BASE_URL + "/admin/schools", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            credentials: "include",
            body: JSON.stringify({
                name: name,
                address: address,
            }),
        });

            if (!response.ok) {
                alert("Ошибка создания");
                return;
            }

            setName("");
            setAddress("");

            fetchSchools();
        } catch (err) {
            console.error(err);
        }
    };

    return (
        <div>
            <h2>Школы</h2>

            {/* создание школы */}
            <form onSubmit={handleCreate}>
                <input
                    value={name}
                    onChange={(e) => setName(e.target.value)}
                    placeholder="Название"
                />

                <input
                    value={address}
                    onChange={(e) => setAddress(e.target.value)}
                    placeholder="Адрес"
                />

                <button type="submit">Создать</button>
            </form>

            {/* таблица */}
            <table>
                <thead>
                    <tr>
                        <th>id</th>
                        <th>название</th>
                        <th>адрес</th>
                        <th>создана в</th>
                    </tr>
                </thead>

                <tbody>
                    {schools.map((school) => (
                        <tr key={school.id}>
                            <td>{school.id}</td>
                            <td>{school.name}</td>
                            <td>{school.address}</td>
                            <td>{school.created_at}</td>
                        </tr>
                    ))}
                </tbody>
            </table>
        </div>
    );
}

export default SchoolsPage;
