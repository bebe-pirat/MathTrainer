import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { BASE_URL } from "../../../constants";
import LogoutButton from "../../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../../styles/shared.module.css";
import styles from "./SectionsEquationTypesPage.module.css";

function SectionsEquationTypesPage() {
  const navigate = useNavigate();
  const [links, setLinks] = useState([]);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState(null);
  const [showModal, setShowModal] = useState(false);
  const [sections, setSections] = useState([]);
  const [equationTypes, setEquationTypes] = useState([]);
  const [selectedSectionId, setSelectedSectionId] = useState("");
  const [selectedEquationTypeId, setSelectedEquationTypeId] = useState("");
  const [submitting, setSubmitting] = useState(false);

  // Загрузка всех связей
  const fetchLinks = async () => {
    setLoading(true);
    try {
      const res = await fetch(`${BASE_URL}/admin/sections-equation-types`, {
        credentials: "include",
      });
      if (!res.ok) throw new Error("Ошибка загрузки связей");
      const data = await res.json();
      setLinks(data);
    } catch (err) {
      console.error(err);
      setError(err.message);
    } finally {
      setLoading(false);
    }
  };

  // Загрузка секций для селекта
  const fetchSections = async () => {
    try {
      const res = await fetch(`${BASE_URL}/admin/sections`, {
        credentials: "include",
      });
      if (!res.ok) throw new Error("Ошибка загрузки секций");
      const data = await res.json();
      setSections(data);
    } catch (err) {
      console.error(err);
    }
  };

  // Загрузка всех типов уравнений для селекта
  const fetchEquationTypes = async () => {
    try {
      const res = await fetch(`${BASE_URL}/admin/equation-types`, {
        credentials: "include",
      });
      if (!res.ok) throw new Error("Ошибка загрузки типов примеров");
      const data = await res.json();
      setEquationTypes(data);
    } catch (err) {
      console.error(err);
    }
  };

  useEffect(() => {
    fetchLinks();
    fetchSections();
    fetchEquationTypes();
  }, []);

  // Добавление связи
  const handleAddLink = async (e) => {
    e.preventDefault();
    if (!selectedSectionId || !selectedEquationTypeId) {
      alert("Выберите секцию и тип примера");
      return;
    }
    setSubmitting(true);
    try {
      const res = await fetch(`${BASE_URL}/admin/sections-equation-types`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          section_id: parseInt(selectedSectionId),
          equation_type_id: parseInt(selectedEquationTypeId),
        }),
      });
      if (!res.ok) throw new Error("Ошибка добавления связи");
      await fetchLinks();
      setShowModal(false);
      setSelectedSectionId("");
      setSelectedEquationTypeId("");
    } catch (err) {
      console.error(err);
      alert(err.message);
    } finally {
      setSubmitting(false);
    }
  };

  // Удаление связи
  const handleDeleteLink = async (sectionId, equationTypeId) => {
    if (!window.confirm("Удалить связь?")) return;
    try {
      const res = await fetch(
        `${BASE_URL}/admin/sections-equation-types?section_id=${sectionId}&equation_type_id=${equationTypeId}`,
        {
          method: "DELETE",
          credentials: "include",
        }
      );
      if (!res.ok) throw new Error("Ошибка удаления связи");
      await fetchLinks();
    } catch (err) {
      console.error(err);
      alert(err.message);
    }
  };

  if (loading) return <div className={sharedStyles.loader}>Загрузка...</div>;
  if (error)
    return (
      <div className={sharedStyles.errorBox}>
        <p>Ошибка: {error}</p>
        <button
          className={sharedStyles.formButton}
          onClick={() => navigate("/admin/dashboard")}
        >
          Назад
        </button>
      </div>
    );

  return (
    <div className={sharedStyles.adminPage}>
      <div className={sharedStyles.adminContainer}>
        <div className={sharedStyles.adminHeader}>
          <button
            className={sharedStyles.headerButton}
            onClick={() => navigate("/admin/dashboard")}
          >
            Назад
          </button>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.adminTitle}>Связи секций и типов примеров</h1>

        <div className={styles.toolbar}>
          <button
            className={sharedStyles.formButton}
            onClick={() => setShowModal(true)}
          >
            + Добавить связь
          </button>
        </div>

        <div className={sharedStyles.formCard}>
          {links.length === 0 ? (
            <div className={sharedStyles.emptyMessage}>Нет связей</div>
          ) : (
            <div className={sharedStyles.dataTableWrapper}>
              <table className={sharedStyles.dataTable}>
                <thead>
                  <tr>
                    <th>ID секции</th>
                    <th>Название секции</th>
                    <th>ID типа</th>
                    <th>Тип примера</th>
                    <th>Класс</th>
                    <th>Действия</th>
                  </tr>
                </thead>
                <tbody>
                  {links.map((link, idx) => (
                    <tr key={idx}>
                      <td>{link.section_id}</td>
                      <td>{link.section_name}</td>
                      <td>{link.equation_type_id}</td>
                      <td>{link.equation_type_name}</td>
                      <td>{link.class} класс</td>
                      <td>
                        <button
                          className={sharedStyles.smallButton}
                          onClick={() =>
                            handleDeleteLink(link.section_id, link.equation_type_id)
                          }
                        >
                          Удалить
                        </button>
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </div>
      </div>

      {/* Модальное окно добавления связи */}
      {showModal && (
        <div className={styles.modalOverlay}>
          <div className={styles.modalContent}>
            <div className={styles.modalHeader}>
              <h3 className={styles.modalTitle}>Добавить связь</h3>
              <button
                className={styles.modalClose}
                onClick={() => setShowModal(false)}
              >
                ✕
              </button>
            </div>
            <form onSubmit={handleAddLink}>
              <div className={styles.modalBody}>
                <div className={sharedStyles.formGroup}>
                  <label>Секция</label>
                  <select
                    className={sharedStyles.formSelect}
                    value={selectedSectionId}
                    onChange={(e) => setSelectedSectionId(e.target.value)}
                    required
                  >
                    <option value="">Выберите секцию</option>
                    {sections.map((sec) => (
                      <option key={sec.id} value={sec.id}>
                        {sec.name} (класс {sec.class})
                      </option>
                    ))}
                  </select>
                </div>
                <div className={sharedStyles.formGroup}>
                  <label>Тип примера</label>
                  <select
                    className={sharedStyles.formSelect}
                    value={selectedEquationTypeId}
                    onChange={(e) => setSelectedEquationTypeId(e.target.value)}
                    required
                  >
                    <option value="">Выберите тип</option>
                    {equationTypes.map((type) => (
                      <option key={type.id} value={type.id}>
                        {type.name} (класс {type.class})
                      </option>
                    ))}
                  </select>
                </div>
              </div>
              <div className={styles.modalFooter}>
                <button
                  type="button"
                  className={sharedStyles.smallButton}
                  onClick={() => setShowModal(false)}
                >
                  Отмена
                </button>
                <button
                  type="submit"
                  className={sharedStyles.formButton}
                  disabled={submitting}
                >
                  {submitting ? "Добавление..." : "Добавить"}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

export default SectionsEquationTypesPage;   