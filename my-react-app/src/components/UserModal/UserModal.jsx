import { useState, useEffect } from "react";
import sharedStyles from "../../styles/shared.module.css";
import styles from "./../../styles/Modal.module.css";
import { ROLES } from "./../../constants";

function UserModal({ isOpen, onClose, onSave, initialData = null }) {
  const [formData, setFormData] = useState({
    email: "",
    login: "",
    role_id: "",
    fullname: "",
    class_id: "",
  });
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (initialData) {
      setFormData({
        email: initialData.email || "",
        login: initialData.login || "",
        role_id: initialData.role_id || "",
        fullname: initialData.fullname || "",
        class_id: initialData.class_id || "",
      });
    } else {
      setFormData({ email: "", login: "", role_id: "", fullname: "", class_id: "" });
    }
  }, [initialData, isOpen]);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    setLoading(true);
    await onSave(formData);
    setLoading(false);
    onClose();
  };

  if (!isOpen) return null;

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modalContent}>
        <div className={styles.modalHeader}>
          <h3 className={styles.modalTitle}>
            {initialData ? "Редактировать пользователя" : "Добавить пользователя"}
          </h3>
          <button className={styles.modalClose} onClick={onClose}>✕</button>
        </div>
        <form onSubmit={handleSubmit}>
          <div className={styles.modalBody}>
            <div className={sharedStyles.formGroup}>
              <label>ФИО</label>
              <input
                name="fullname"
                className={sharedStyles.formInput}
                value={formData.fullname}
                onChange={handleChange}
                required
              />
            </div>
            <div className={sharedStyles.formGroup}>
              <label>Логин</label>
              <input
                name="login"
                className={sharedStyles.formInput}
                value={formData.login}
                onChange={handleChange}
                required
              />
            </div>
            <div className={sharedStyles.formGroup}>
              <label>Email</label>
              <input
                name="email"
                type="email"
                className={sharedStyles.formInput}
                value={formData.email}
                onChange={handleChange}
                required
              />
            </div>
            <div className={sharedStyles.formGroup}>
              <label>Роль</label>
              <select
                name="role_id"
                className={sharedStyles.formSelect}
                value={formData.role_id}
                onChange={handleChange}
                required
              >
                <option value="">Выберите роль</option>
                <option value={[ROLES.ADMIN]}>Администратор</option>
                <option value={[ROLES.TEACHER]}>Учитель</option>
                <option value={[ROLES.STUDENT]}>Ученик</option>
                <option value={[ROLES.HEAD]}>Завуч</option>
              </select>
            </div>
            <div className={sharedStyles.formGroup}>
              <label>ID класса (только для ученика/учителя)</label>
              <input
                name="class_id"
                type="number"
                className={sharedStyles.formInput}
                value={formData.class_id}
                onChange={handleChange}
                placeholder="Необязательно"
              />
            </div>
          </div>
          <div className={styles.modalFooter}>
            <button type="button" className={sharedStyles.smallButton} onClick={onClose}>
              Отмена
            </button>
            <button type="submit" className={sharedStyles.formButton} disabled={loading}>
              {loading ? "Сохранение..." : "Сохранить"}
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}

export default UserModal;