import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import LogoutButton from "./../../components/LogoutButton";
import { BASE_URL } from "../../constants";

function ProfilePage() {
    const [profile, setProfile] = useState(null);
    const navigate = useNavigate();

    useEffect(() => {
        fetch(BASE_URL + "/student/profile", {
            credentials: "include",
        })
            .then((res) => {
                if (!res.ok) throw new Error("Ошибка загрузки профиля");
                return res.json();
            })
            .then(setProfile)
            .catch(console.error);
    }, []);

    if (!profile) return <div>Loading...</div>;

    return (
        <div style={styles.page}>
            {/* верхняя панель */}
            <div style={styles.header}>
                <button onClick={() => navigate("/student/dashboard")}>
                    Назад
                </button>

                <LogoutButton />
            </div>

            {/* профиль */}
            <div style={styles.card}>
                <div style={{ fontSize: "40px" }}>
                    👤
                </div>

                <h2>{profile.fullname}</h2>

                <div style={styles.row}>
                    <span>Школа:</span>
                    <span>{profile.school_name}</span>
                </div>

                <div style={styles.row}>
                    <span>Класс:</span>
                    <span>{profile.class_name}</span>
                </div>

                <div style={styles.xp}>
                    XP: {profile.xp}
                </div>
            </div>
        </div>
    );
}

export default ProfilePage;

const styles = {
    page: {
        backgroundColor: "white",
        minHeight: "100vh",
        padding: "20px",
    },

    header: {
        display: "flex",
        justifyContent: "space-between",
        marginBottom: "20px",
    },

    card: {
        border: "1px solid #ddd",
        padding: "20px",
        maxWidth: "400px",
    },

    row: {
        display: "flex",
        justifyContent: "space-between",
        marginTop: "10px",
    },

    xp: {
        marginTop: "20px",
        fontWeight: "bold",
        color: "#4da6ff",
    },
};
