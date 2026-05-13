import { useState } from "react";
import sharedStyles from "../../styles/shared.module.css";
import styles from "./../../styles/Modal.module.css";
import { BASE_URL } from "../../constants";

function PasswordModal({ isOpen, onClose, userId, onPasswordUpdated }) {
  const [loading, setLoading] = useState(false);

  const handleSubmit = async () => {
    setLoading(true);
    try {
      const response = await fetch(`${BASE_URL}/admin/users/${userId}`, {
        method: "PATCH",
        credentials: "include",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({}),
      });
      console.log(`${BASE_URL}/admin/users/${userId}`);
      if (!response.ok) throw new Error("Ошибка смены пароля");
      const data = await response.json();
      onPasswordUpdated(data.login, data.password);
      onClose();
    } catch (err) {
      console.error(err);
      alert(err.message);
    } finally {
      setLoading(false);
    }
  };

  if (!isOpen) return null;

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modalContent}>
        <div className={styles.modalHeader}>
          <h3 className={styles.modalTitle}>Смена пароля</h3>
          <button className={styles.modalClose} onClick={onClose}>✕</button>
        </div>
        <div className={styles.modalBody}>
          <p>Вы действительно хотите сгенерировать новый пароль для пользователя?</p>
        </div>
        <div className={styles.modalFooter}>
          <button className={sharedStyles.smallButton} onClick={onClose}>Отмена</button>
          <button className={sharedStyles.formButton} onClick={handleSubmit} disabled={loading}>
            {loading ? "Генерация..." : "Да, сбросить пароль"}
          </button>
        </div>
      </div>
    </div>
  );
}

export default PasswordModal;