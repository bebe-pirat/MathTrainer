import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../../constants";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import EquationTypeModal from "../../../components/EquationTypeModal/EquationTypeModal";
import OperandsModal from "../../../components/OperandsModal/OperandsModal";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./EquationTypesPage.module.css";

function EquationTypesPage() {
  const navigate = useNavigate();
  const [types, setTypes] = useState([]);
  const [filterClass, setFilterClass] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [sections, setSections] = useState([]);

  // Modal states
  const [showTypeModal, setShowTypeModal] = useState(false);
  const [editingType, setEditingType] = useState(null);
  const [showOperandsModal, setShowOperandsModal] = useState(false);
  const [currentOperands, setCurrentOperands] = useState([]);

  const fetchTypes = async () => {
    setLoading(true);
    try {
      let url = `${BASE_URL}/admin/equation-types`;
      if (filterClass) url += `?class=${filterClass}`;
      const res = await fetch(url, { credentials: "include" });
      if (!res.ok) throw new Error("Ошибка загрузки типов уравнений");
      const data = await res.json();
      setTypes(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  const fetchSections = async () => {
    try {
      const res = await fetch(`${BASE_URL}/admin/sections`, { credentials: "include" });
      if (!res.ok) throw new Error("Ошибка загрузки секций");
      const data = await res.json();
      setSections(data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    fetchTypes();
    fetchSections();
  }, [filterClass]);

  const handleSaveType = async (formData) => {
    try {
      const method = editingType ? "PUT" : "POST";
      const url = editingType
        ? `${BASE_URL}/admin/equation-types/${editingType.id}`
        : `${BASE_URL}/admin/equation-types`;
      const response = await fetch(url, {
        method,
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          class: parseInt(formData.class),
          name: formData.name,
          description: formData.description,
          operations: formData.operations,
          num_operands: parseInt(formData.num_operands),
          no_remainder: formData.no_remainder,
          max_result: parseInt(formData.max_result),
          section_id: parseInt(formData.section_id),
          operands: formData.operands.map((op, idx) => ({
            id: op.id || 0,
            operand_order: op.operand_order,
            min_value: parseInt(op.min_value),
            max_value: parseInt(op.max_value),
          })),
        }),
      });
      if (!response.ok) throw new Error("Ошибка сохранения");
      await fetchTypes();
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Удалить тип уравнения?")) return;
    try {
      const res = await fetch(`${BASE_URL}/admin/equation-types/${id}`, {
        method: "DELETE",
        credentials: "include",
      });
      if (!res.ok) throw new Error("Ошибка удаления");
      await fetchTypes();
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  const handleShowOperands = async (typeId) => {
    try {
      const res = await fetch(`${BASE_URL}/admin/operands?equation_type_id=${typeId}`, {
        credentials: "include",
      });
      if (!res.ok) throw new Error("Ошибка загрузки операндов");
      const data = await res.json();
      setCurrentOperands(data);
      setShowOperandsModal(true);
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  const openEditModal = (type) => {
    setEditingType(type);
    setShowTypeModal(true);
  };

  const openCreateModal = () => {
    setEditingType(null);
    setShowTypeModal(true);
  };

  if (loading) return <div className={sharedStyles.loader}>Загрузка...</div>;
  if (error) return (
    <div className={sharedStyles.errorBox}>
      <p>Ошибка: {error}</p>
      <button className={sharedStyles.formButton} onClick={() => navigate("/admin/dashboard")}>
        Назад
      </button>
    </div>
  );

  return (
    <div className={sharedStyles.adminPage}>
      <div className={sharedStyles.adminContainer}>
        <div className={sharedStyles.adminHeader}>
          <button className={sharedStyles.headerButton} onClick={() => navigate("/admin/dashboard")}>
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.adminTitle}>Типы уравнений</h1>

        <div className={styles.toolbar}>
          <div className={styles.filterGroup}>
            <label className={styles.filterLabel}>Класс</label>
            <select
              className={sharedStyles.formSelect}
              value={filterClass}
              onChange={(e) => setFilterClass(e.target.value)}
            >
              <option value="">Все классы</option>
              <option value="1">1 класс</option>
              <option value="2">2 класс</option>
              <option value="3">3 класс</option>
              <option value="4">4 класс</option>
            </select>
          </div>
          <button className={sharedStyles.formButton} onClick={openCreateModal}>
            + Добавить тип
          </button>
        </div>

        <div className={sharedStyles.formCard}>
          {types.length === 0 ? (
            <div className={sharedStyles.emptyMessage}>Нет типов уравнений</div>
          ) : (
            <div className={sharedStyles.dataTableWrapper}>
              <table className={sharedStyles.dataTable}>
                <thead>
                  <tr>
                    <th>ID</th>
                    <th>Название</th>
                    <th>Класс</th>
                    <th>Описание</th>
                    <th>Операции</th>
                    <th>Операндов</th>
                    <th>Без остатка</th>
                    <th>Max результат</th>
                    <th>Действия</th>
                  </tr>
                </thead>
                <tbody>
                  {types.map((type) => (
                    <tr key={type.id}>
                      <td>{type.id}</td>
                      <td>{type.name}</td>
                      <td>{type.class} класс</td>
                      <td>{type.description || "-"}</td>
                      <td>{type.operations}</td>
                      <td>{type.num_operands}</td>
                      <td>{type.no_remainder ? "Да" : "Нет"}</td>
                      <td>{type.max_result}</td>
                      <td>
                        <div className={styles.actionButtons}>
                          <button className={sharedStyles.smallButton} onClick={() => openEditModal(type)}>
                            Редактировать
                          </button>
                          <button className={sharedStyles.smallButton} onClick={() => handleDelete(type.id)}>
                            Удалить
                          </button>
                          <button className={sharedStyles.smallButton} onClick={() => handleShowOperands(type.id)}>
                            Посмотреть операнды
                          </button>
                        </div>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>
      </div>

      <EquationTypeModal
        isOpen={showTypeModal}
        onClose={() => setShowTypeModal(false)}
        onSave={handleSaveType}
        initialData={editingType}
        sections={sections}
      />

      <OperandsModal
        isOpen={showOperandsModal}
        onClose={() => setShowOperandsModal(false)}
        operands={currentOperands}
      />
    </div>
  );
}

export default EquationTypesPage;