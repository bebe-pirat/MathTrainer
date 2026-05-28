import { useState, useEffect } from "react";
import sharedStyles from "../../styles/shared.module.css";
import styles from "./EquationTypeModal.module.css";
import { BASE_URL } from "../../constants";

function EquationTypeModal({ isOpen, onClose, onSave, initialData = null }) {
  const [formData, setFormData] = useState({
    class: "",
    name: "",
    description: "",
    operations: "",
    num_operands: 2,
    no_remainder: false,
    max_result: 100,
    operands: [{ operand_order: 1, min_value: 0, max_value: 10 }],
  });
  const [loading, setLoading] = useState(false);
  const [loadingOperands, setLoadingOperands] = useState(false);

  useEffect(() => {
    if (!isOpen) return;

    if (initialData && initialData.id) {
      setLoadingOperands(true);
      fetch(`${BASE_URL}/admin/operands?equation_type_id=${initialData.id}`, {
        credentials: "include",
      })
        .then((res) => {
          if (!res.ok) throw new Error("Ошибка загрузки операндов");
          return res.json();
        })
        .then((operands) => {
          setFormData((prev) => ({
            ...prev,
            operands: operands.length
              ? operands.map((op) => ({
                  operand_order: op.operand_order,
                  min_value: op.min_value,
                  max_value: op.max_value,
                }))
              : [{ operand_order: 1, min_value: 0, max_value: 10 }],
          }));
        })
        .catch((err) => console.error(err))
        .finally(() => setLoadingOperands(false));
    } else {
      // Новый тип: сброс операндов
      setFormData((prev) => ({
        ...prev,
        operands: [{ operand_order: 1, min_value: 0, max_value: 10 }],
      }));
    }
  }, [initialData, isOpen]);

  // Заполнение остальных полей при открытии
  useEffect(() => {
    if (initialData) {
      setFormData((prev) => ({
        ...prev,
        class: initialData.class || "",
        name: initialData.name || "",
        description: initialData.description || "",
        operations: initialData.operations || "",
        num_operands: initialData.num_operands || 2,
        no_remainder: initialData.no_remainder || false,
        max_result: initialData.max_result || 100,
      }));
    } else {
      setFormData((prev) => ({
        ...prev,
        class: "",
        name: "",
        description: "",
        operations: "",
        num_operands: 2,
        no_remainder: false,
        max_result: 100,
      }));
    }
  }, [initialData, isOpen]);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: type === "checkbox" ? checked : value,
    }));
  };

  const handleNumOperandsChange = (e) => {
    const num = parseInt(e.target.value) || 1;
    let newOperands = [...formData.operands];
    if (num > newOperands.length) {
      for (let i = newOperands.length; i < num; i++) {
        newOperands.push({ operand_order: i + 1, min_value: 0, max_value: 10 });
      }
    } else {
      newOperands = newOperands.slice(0, num);
    }
    setFormData((prev) => ({
      ...prev,
      num_operands: num,
      operands: newOperands,
    }));
  };

  const handleOperandChange = (index, field, value) => {
    const updated = [...formData.operands];
    updated[index][field] = value;
    setFormData((prev) => ({ ...prev, operands: updated }));
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
            {initialData ? "Редактировать тип примера" : "Создать тип примера"}
          </h3>
          <button className={styles.modalClose} onClick={onClose}>✕</button>
        </div>
        <form onSubmit={handleSubmit}>
          <div className={styles.modalBody}>
            <div className={sharedStyles.formGrid}>
              <div className={sharedStyles.formGroup}>
                <label>Класс</label>
                <input
                  type="number"
                  name="class"
                  className={sharedStyles.formInput}
                  value={formData.class}
                  onChange={handleChange}
                  min="1"
                  max="4"
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Название</label>
                <input
                  name="name"
                  className={sharedStyles.formInput}
                  value={formData.name}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Описание</label>
                <input
                  name="description"
                  className={sharedStyles.formInput}
                  value={formData.description}
                  onChange={handleChange}
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Операции (например: +, -)</label>
                <input
                  name="operations"
                  className={sharedStyles.formInput}
                  value={formData.operations}
                  onChange={handleChange}
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Количество операндов</label>
                <input
                  type="number"
                  name="num_operands"
                  className={sharedStyles.formInput}
                  value={formData.num_operands}
                  onChange={handleNumOperandsChange}
                  min="1"
                  max="5"
                  required
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Без остатка (для деления)</label>
                <input
                  type="checkbox"
                  name="no_remainder"
                  checked={formData.no_remainder}
                  onChange={handleChange}
                />
              </div>
              <div className={sharedStyles.formGroup}>
                <label>Максимальный результат</label>
                <input
                  type="number"
                  name="max_result"
                  className={sharedStyles.formInput}
                  value={formData.max_result}
                  onChange={handleChange}
                  min="1"
                  required
                />
              </div>
            </div>

            <div className={styles.operandsSection}>
              <h4>Операнды</h4>
              {loadingOperands ? (
                <div className={sharedStyles.loader}>Загрузка операндов...</div>
              ) : (
                <table className={styles.operandsTable}>
                  <thead>
                    <tr>
                      <th>Порядок</th>
                      <th>Мин. значение</th>
                      <th>Макс. значение</th>
                    </tr>
                  </thead>
                  <tbody>
                    {formData.operands.map((op, idx) => (
                      <tr key={idx}>
                        <td>{op.operand_order}</td>
                        <td>
                          <input
                            type="number"
                            value={op.min_value}
                            onChange={(e) =>
                              handleOperandChange(idx, "min_value", parseInt(e.target.value))
                            }
                            className={sharedStyles.formInput}
                          />
                        </td>
                        <td>
                          <input
                            type="number"
                            value={op.max_value}
                            onChange={(e) =>
                              handleOperandChange(idx, "max_value", parseInt(e.target.value))
                            }
                            className={sharedStyles.formInput}
                          />
                        </td>
                      </tr>
                    ))}
                  </tbody>
                </table>
              )}
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

export default EquationTypeModal;