import { BASE_URL } from '../constants';

export const fetchSchools = async () => {
    try {
        const response = await fetch(`${BASE_URL}/admin/schools`, {
            method: "GET", 
            credentials: "include",
        });

        if (!response.ok) {
            throw new Error("Ошибка загрузки школ");
        }

        return await response.json();
    } catch (err) {
        console.error(err);
        throw err;
    }
};

export const fetchClasses = async (schoolId) => {
    if (!schoolId) {
        return [];
    }
    
    try {
        const response = await fetch(`${BASE_URL}/admin/classes?school_id=${schoolId}`, {
            method: "GET",
            credentials: "include",
        });

        if (!response.ok) {
            throw new Error("Ошибка загрузки классов");
        }

        return await response.json();
    } catch (err) {
        console.error(err);
        throw err;
    }
};