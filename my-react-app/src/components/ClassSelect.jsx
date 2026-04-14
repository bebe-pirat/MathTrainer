import { useState, useEffect } from 'react';
import { fetchClasses } from '../services/SchoolServices';

export const ClassSelect = ({ schoolId, value, onChange }) => {
    const [classes, setClasses] = useState([]);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        if (!schoolId) {
            setClasses([]);
            return;
        }

        setLoading(true);
        fetchClasses(schoolId)
            .then(data => {
                setClasses(data);
            })
            .catch(error => {
                console.error("Error loading classes:", error);
                setClasses([]);
            })
            .finally(() => {
                setLoading(false);
            });
    }, [schoolId]); 

    const handleChange = (e) => {
        onChange(e.target.value);
    };

    return (
        <select 
            value={value} 
            onChange={handleChange}
            disabled={!schoolId || loading}
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