import { useEffect, useState } from 'react';
import { fetchSchools } from '../services/SchoolServices';

export const SchoolSelect = ({ value, onChange }) => {
    const [schools, setSchools] = useState([]);

    useEffect(() => {
        fetchSchools().then(setSchools);
    }, []);

    return (
        <select value={value} onChange={(e) => onChange(e.target.value)}>
            <option value="">Выберите школу</option>
            {schools.map(school => (
                <option key={school.id} value={school.id}>
                    {school.name}
                </option>
            ))}
        </select>
    );
};

