// src/components/OperandsModal/OperandsModal.jsx
import sharedStyles from "../../styles/shared.module.css";
import styles from "./OperandsModal.module.css";

function OperandsModal({ isOpen, onClose, operands = [] }) {
  if (!isOpen) return null;

  return (
    <div className={styles.modalOverlay}>
      <div className={styles.modalContent}>
        <div className={styles.modalHeader}>
          <h3 className={styles.modalTitle}>Операнды</h3>
          <button className={styles.modalClose} onClick={onClose}>✕</button>
        </div>
        <div className={styles.modalBody}>
          {operands.length === 0 ? (
            <p>Нет операндов</p>
          ) : (
            <table className={sharedStyles.dataTable}>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Порядок</th>
                  <th>Мин. значение</th>
                  <th>Макс. значение</th>
                </tr>
              </thead>
              <tbody>
                {operands.map((op) => (
                  <tr key={op.id}>
                    <td>{op.id}</td>
                    <td>{op.operand_order}</td>
                    <td>{op.min_value}</td>
                    <td>{op.max_value}</td>
                  </tr>
                ))}
              </tbody>
            </table>
          )}
        </div>
        <div className={styles.modalFooter}>
          <button className={sharedStyles.formButton} onClick={onClose}>Закрыть</button>
        </div>
      </div>
    </div>
  );
}

export default OperandsModal;