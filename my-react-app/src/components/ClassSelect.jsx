import { useState, useEffect } from 'react';
import { fetchClasses } from '../services/schoolService';

export const ClassSelect = ({ schoolId, value, onChange }) => {
    const [classes, setClasses] = useState([]);
    const [loading, setLoading] = useState(false);

    useEffect(() => {
        if (!schoolId) {
            setClasses([]);
            onChange?.(''); 
            return;
        }

        const loadClasses = async () => {
            setLoading(true);
            try {
                const data = await fetchClasses(schoolId);
                setClasses(data);
            } finally {
                setLoading(false);
            }
        };

        loadClasses();
    }, [schoolId, onChange]); 

    return (
        <select 
            value={value} 
            onChange={(e) => onChange(e.target.value)}
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