import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../../AuthContext";
import { BASE_URL } from "../../constants";
import { ROLES } from "../../constants";
import styles from "./Login.module.css";

function Login()  {
  const [loginInput, setLoginInput] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate();
  const { login } = useAuth();

  const handleSubmit = async (e) => {
    e.preventDefault();

    const response = await fetch(BASE_URL + "/auth/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      credentials: "include",
      body: JSON.stringify({ login: loginInput, password }),
    });

    if (!response.ok) {
      alert("Ошибка входа. Проверьте логин и пароль.");
      return;
    }

    const data = await response.json();
    login(data);
    redirectByRole(data.role, navigate);
  };

  return (
    <div className={styles.page}>
      <div className={styles.card}>
        <div className={styles.title}>Математический тренажёр</div>
        <div className={styles.subtitle}>Войдите, чтобы продолжить</div>

        <form onSubmit={handleSubmit} className={styles.form}>
          <div className={styles.inputGroup}>
            <label>Логин</label>
            <input
              className={styles.input}
              value={loginInput}
              onChange={(e) => setLoginInput(e.target.value)}
              placeholder="Введите логин"
              required
            />
          </div>

          <div className={styles.inputGroup}>
            <label>Пароль</label>
            <input
              className={styles.input}
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="Введите пароль"
              required
            />
          </div>

          <button type="submit" className={styles.button}>
            Войти
          </button>
        </form>
      </div>
    </div>
  );
}

function redirectByRole(role_id, navigate) {
  switch (role_id) {
    case ROLES.ADMIN:
      navigate("/admin/dashboard");
      break;
    case ROLES.STUDENT:
      navigate("/student/dashboard");
      break;
    case ROLES.TEACHER:
      navigate("/teacher/dashboard");
      break;
    case ROLES.HEAD:
      navigate("/director/dashboard");
      break;
    default:
      navigate("/");
  }
}

export default Login;