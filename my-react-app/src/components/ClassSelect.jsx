import { useState, useEffect } from 'react';
import { fetchClasses } from '../services/SchoolServices';
import sharedStyles from '../styles/shared.module.css';

export const ClassSelect = ({ schoolId, value, onChange, className = '' }) => {
  const [classes, setClasses] = useState([]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (!schoolId) {
      setClasses([]);
      return;
    }
    setLoading(true);
    fetchClasses(schoolId)
      .then(data => setClasses(data))
      .catch(error => {
        console.error("Error loading classes:", error);
        setClasses([]);
      })
      .finally(() => setLoading(false));
  }, [schoolId]);

  const isDisabled = !schoolId || loading;
  const selectClass = (isDisabled ? sharedStyles.selectDisabled : sharedStyles.select) +
                      (className ? ` ${className}` : '');

  const handleChange = (e) => {
    onChange(e.target.value);
  };

  return (
    <select
      className={selectClass}
      value={value}
      onChange={handleChange}
      disabled={isDisabled}
    >
      <option value="">
        {!schoolId ? 'Сначала выберите школу' : loading ? 'Загрузка...' : 'Выберите класс'}
      </option>
      {classes.map(classItem => (
        <option key={classItem.id} value={classItem.id}>
          {classItem.name}
        </option>
      ))}
    </select>
  );
};