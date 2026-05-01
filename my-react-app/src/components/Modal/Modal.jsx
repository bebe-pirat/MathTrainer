import { useEffect } from "react";
import sharedStyles from "../../styles/shared.module.css";
import styles from "./Modal.module.css";

function Modal({ isOpen, onClose, title, message, onConfirm, confirmText = "OK", showCancel = false }) {
  useEffect(() => {
    if (isOpen) {
      document.body.style.overflow = "hidden";
    } else {
      document.body.style.overflow = "unset";
    }
    return () => {
      document.body.style.overflow = "unset";
    };
  }, [isOpen]);

  if (!isOpen) return null;

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modalContent}>
        <div className={styles.modalHeader}>
          <h3 className={styles.modalTitle}>{title}</h3>
          <button className={styles.modalClose} onClick={onClose}>✕</button>
        </div>
        <div className={styles.modalBody}>
          <p>{message}</p>
        </div>
        <div className={styles.modalFooter}>
          {showCancel && (
            <button className={sharedStyles.smallButton} onClick={onClose}>
              Отмена
            </button>
          )}
          <button className={sharedStyles.formButton} onClick={onConfirm || onClose}>
            {confirmText}
          </button>
        </div>
      </div>
    </div>
  );
}

export default Modal;