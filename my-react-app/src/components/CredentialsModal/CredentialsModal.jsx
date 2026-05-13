import sharedStyles from "../../styles/shared.module.css";
import styles from "./../../styles/Modal.module.css";

function CredentialsModal({ isOpen, onClose, title, login, password }) {
  if (!isOpen) return null;

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modalContent}>
        <div className={styles.modalHeader}>
          <h3 className={styles.modalTitle}>{title}</h3>
          <button className={styles.modalClose} onClick={onClose}>✕</button>
        </div>
        <div className={styles.modalBody}>
          <div className={styles.credentialRow}>
            <span className={styles.credentialLabel}>Логин:</span>
            <span className={styles.credentialValue}>{login}</span>
          </div>
          <div className={styles.credentialRow}>
            <span className={styles.credentialLabel}>Пароль:</span>
            <span className={styles.credentialValue}>{password}</span>
          </div>
          <p className={styles.warning}>Сохраните эти данные. Пароль будет отображаться только один раз.</p>
        </div>
        <div className={styles.modalFooter}>
          <button className={sharedStyles.formButton} onClick={onClose}>Закрыть</button>
        </div>
      </div>
    </div>
  );
}

export default CredentialsModal;