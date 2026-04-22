import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import { BASE_URL } from "../../../constants";
import styles from "./ProfilePage.module.css";
import sharedStyles from "./../../../styles/shared.module.css"

function ProfilePage() {
  const [profile, setProfile] = useState(null);
  const navigate = useNavigate();

  useEffect(() => {
    fetch(BASE_URL + "/student/profile", { credentials: "include" })
      .then((res) => {
        if (!res.ok) throw new Error("Ошибка загрузки профиля");
        return res.json();
      })
      .then(setProfile)
      .catch(console.error);
  }, []);

  if (!profile) return <div className={sharedStyles.loader}>Загрузка...</div>;

  return (
    <div className={styles.page}>
      <div className={styles.header}>
        <button
          className={sharedStyles.headerButton}
          onClick={() => navigate("/student/dashboard")}
        >
          Назад
        </button>
        <LogoutButton />
      </div>

      <div className={sharedStyles.card}>
        <div className={styles.avatar}>👤</div>
        <h2 className={styles.fullname}>{profile.fullname}</h2>

        <div className={styles.infoRow}>
          <span>Школа:</span>
          <span>{profile.school_name}</span>
        </div>

        <div className={styles.infoRow}>
          <span>Класс:</span>
          <span>{profile.class_name}</span>
        </div>

        <div className={styles.xp}>
          XP: {profile.xp}
        </div>
      </div>
    </div>
  );
}

export default ProfilePage;