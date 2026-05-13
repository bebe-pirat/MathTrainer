import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../../constants";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import SectionModal from "../../../components/SectionModal/SectionModal";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./SectionsPage.module.css";

function SectionsPage() {
  const navigate = useNavigate();
  const [sections, setSections] = useState([]);
  const [filterClass, setFilterClass] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [modalOpen, setModalOpen] = useState(false);
  const [editingSection, setEditingSection] = useState(null);

  const fetchSections = async () => {
    setLoading(true);
    try {
      let url = `${BASE_URL}/admin/sections`;
      if (filterClass) {
        url += `?class=${filterClass}`;
      }
      const res = await fetch(url, { credentials: "include" });
      if (!res.ok) throw new Error("Ошибка загрузки секций");
      const data = await res.json();
      setSections(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchSections();
  }, [filterClass]);

  const handleCreate = async (formData) => {
    try {
      const res = await fetch(`${BASE_URL}/admin/sections`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          name: formData.name,
          class: parseInt(formData.class),
          section_order: parseInt(formData.section_order),
        }),
      });
      if (!res.ok) throw new Error("Ошибка создания секции");
      await fetchSections();
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  const handleUpdate = async (formData) => {
    try {
      const res = await fetch(`${BASE_URL}/admin/sections/${editingSection.id}`, {
        method: "PUT",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          name: formData.name,
          class: parseInt(formData.class),
          section_order: parseInt(formData.section_order),
        }),
      });
      if (!res.ok) throw new Error("Ошибка обновления секции");
      await fetchSections();
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm("Вы уверены, что хотите удалить эту секцию?")) return;
    try {
      const res = await fetch(`${BASE_URL}/admin/sections/${id}`, {
        method: "DELETE",
        credentials: "include",
      });
      if (!res.ok) throw new Error("Ошибка удаления секции");
      await fetchSections();
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  const openEditModal = (section) => {
    setEditingSection(section);
    setModalOpen(true);
  };

  const openCreateModal = () => {
    setEditingSection(null);
    setModalOpen(true);
  };

  const closeModal = () => {
    setModalOpen(false);
    setEditingSection(null);
  };

  const handleSave = async (formData) => {
    if (editingSection) {
      await handleUpdate(formData);
    } else {
      await handleCreate(formData);
    }
    closeModal();
  };

  if (error) {
    return (
      <div className={sharedStyles.errorBox}>
        <p>Ошибка: {error}</p>
        <button className={sharedStyles.headerButton} onClick={() => navigate("/admin/dashboard")}>
          Назад
        </button>
      </div>
    );
  }

  return (
    <div className={sharedStyles.adminPage}>
      <div className={sharedStyles.adminContainer}>
        <div className={sharedStyles.adminHeader}>
          <button className={sharedStyles.headerButton} onClick={() => navigate("/admin/dashboard")}>
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.adminTitle}>Управление секциями</h1>

        <div className={styles.filterBar}>
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
            + Добавить секцию
          </button>
        </div>

        {loading ? (
          <div className={sharedStyles.loader}>Загрузка...</div>
        ) : sections.length === 0 ? (
          <div className={sharedStyles.emptyMessage}>Нет секций</div>
        ) : (
          <div className={sharedStyles.dataTableWrapper}>
            <table className={sharedStyles.dataTable}>
              <thead>
                <tr>
                  <th>ID</th>
                  <th>Название</th>
                  <th>Класс</th>
                  <th>Порядок</th>
                  <th>Действия</th>
                </tr>
              </thead>
              <tbody>
                {sections.map((section) => (
                  <tr key={section.id}>
                    <td>{section.id}</td>
                    <td>{section.name}</td>
                    <td>{section.class} класс</td>
                    <td>{section.section_order}</td>
                    <td>
                      <div className={styles.actionButtons}>
                        <button className={sharedStyles.smallButton} onClick={() => openEditModal(section)}>
                          Редактировать
                        </button>
                        <button className={sharedStyles.smallButton} onClick={() => handleDelete(section.id)}>
                          Удалить
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

      <SectionModal
        isOpen={modalOpen}
        onClose={closeModal}
        onSave={handleSave}
        initialData={editingSection}
      />
    </div>
  );
}

export default SectionsPage;