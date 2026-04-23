import { useEffect, useState } from 'react';
import { fetchSchools } from '../services/SchoolServices';
import sharedStyles from '../styles/shared.module.css'; 

export const SchoolSelect = ({ value, onChange, className = '' }) => {
  const [schools, setSchools] = useState([]);

  useEffect(() => {
    fetchSchools().then(setSchools);
  }, []);

  const selectClass = sharedStyles.select + (className ? ` ${className}` : '');

  return (
    <select
      className={selectClass}
      value={value}
      onChange={(e) => onChange(e.target.value)}
    >
      <option value="">Выберите школу</option>
      {schools.map(school => (
        <option key={school.id} value={school.id}>
          {school.name}
        </option>
      ))}
    </select>
  );
};