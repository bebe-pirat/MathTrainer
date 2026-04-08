import { useState } from 'react';
import { fetchSchools, fetchClasses } from '../services/schoolService';

export const useSchools = () => {
    const [schools, setSchools] = useState([]);
    const [classes, setClasses] = useState([]);
    const [loading, setLoading] = useState(false);

    const loadSchools = async () => {
        setLoading(true);
        try {
            const data = await fetchSchools();
            setSchools(data);
        } finally {
            setLoading(false);
        }
    };

    const loadClasses = async (schoolId) => {
        setLoading(true);
        try {
            const data = await fetchClasses(schoolId);
            setClasses(data);
        } finally {
            setLoading(false);
        }
    };

    return { schools, classes, loading, loadSchools, loadClasses };
};