import { useState, useEffect } from "react";
import sharedStyles from "../../styles/shared.module.css";
import styles from "./SectionModal.module.css";

function SectionModal({ isOpen, onClose, onSave, initialData = null }) {
  const [formData, setFormData] = useState({
    name: "",
    class: "",
    section_order: "",
  });
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (initialData) {
      setFormData({
        name: initialData.name || "",
        class: initialData.class || "",
        section_order: initialData.section_order || "",
      });
    } else {
      setFormData({ name: "", class: "", section_order: "" });
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
            {initialData ? "Редактировать секцию" : "Добавить секцию"}
          </h3>
          <button className={styles.modalClose} onClick={onClose}>✕</button>
        </div>
        <form onSubmit={handleSubmit}>
          <div className={styles.modalBody}>
            <div className={sharedStyles.formGroup}>
              <label>Название секции</label>
              <input
                name="name"
                className={sharedStyles.formInput}
                value={formData.name}
                onChange={handleChange}
                required
              />
            </div>
            <div className={sharedStyles.formGroup}>
              <label>Класс (1-4)</label>
              <input
                name="class"
                type="number"
                min="1"
                max="4"
                className={sharedStyles.formInput}
                value={formData.class}
                onChange={handleChange}
                required
              />
            </div>
            <div className={sharedStyles.formGroup}>
              <label>Порядок секции</label>
              <input
                name="section_order"
                type="number"
                min="1"
                className={sharedStyles.formInput}
                value={formData.section_order}
                onChange={handleChange}
                required
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

export default SectionModal;