export const ROLES = {
  ADMIN: 2,
  TEACHER: 3,
  STUDENT: 1,
  HEAD: 4
};

export const API_ENDPOINTS = {
  AUTH: {
    LOGIN: '/auth/login',
    REGISTER: '/auth/register',
    LOGOUT: '/auth/logout',
  },
};

export const APP_CONFIG = {
  STORAGE_KEYS: {
    ROLE: 'role',
    SESSION: 'session_id',
    USER: 'user_id',
  },
};

export const BASE_URL = "http://localhost:8080"