import { useNavigate } from "react-router-dom";
import LogoutButton from "../../components/LogoutButton/LogoutButton";
import sharedStyles from "../../styles/shared.module.css";
import { ReactComponent as UserIcon } from "../../assets/user-list.svg";
import { ReactComponent as EquationTypesIcon } from "../../assets/pen.svg"
import { ReactComponent as SchoolIcon } from "../../assets/building.svg"
import { ReactComponent as TeacherIcon } from "../../assets/building-user.svg"
import { ReactComponent as SectionsIcon } from "../../assets/book-open.svg"
import { ReactComponent as ClassesIcon } from "../../assets/group-arrows-rotate.svg"
import { ReactComponent as SectionEquationTypeIcon } from "../../assets/book.svg"

function AdminDashboard() {
  const navigate = useNavigate();

  return (
    <div className={sharedStyles.dashboardPage}>
      <div className={sharedStyles.dashboardContainer}>
        <div className={sharedStyles.dashboardHeader}>
          <LogoutButton />
        </div>

        <h1 className={sharedStyles.dashboardTitle}>Панель администратора</h1>
        <div className={sharedStyles.dashboardSubtitle}>
          Управление школами, учителями, классами и пользователями
        </div>

        <div className={sharedStyles.dashboardCardsGrid}>
          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/schools")}
          >
            <SchoolIcon className={sharedStyles.iconBlue}/>
            <div className={sharedStyles.dashboardCardTitle}>Школы</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Добавление школ
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/teachers")}
          >
            <TeacherIcon className={sharedStyles.iconBlue}/>
            <div className={sharedStyles.dashboardCardTitle}>Учителя</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Управление учителями и их данными
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/classes")}
          >
            <ClassesIcon className={sharedStyles.iconBlue}/>
            <div className={sharedStyles.dashboardCardTitle}>Классы</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Создание и настройка классов
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/users")}
          >
            <UserIcon className={sharedStyles.iconBlue} />            
            <div className={sharedStyles.dashboardCardTitle}>Пользователи</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Управление пользователями системы
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/sections")}
          >
            <SectionsIcon className={sharedStyles.iconBlue}/>
            <div className={sharedStyles.dashboardCardTitle}>Секции</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Управление секциями для примеров школьников
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/sections-equation-types")}
          >
            <SectionEquationTypeIcon className={sharedStyles.iconBlue}/>
            <div className={sharedStyles.dashboardCardTitle}>Секции и типы примеров</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Управление содержимым секций  
            </div>
          </div>

          <div
            className={sharedStyles.dashboardCard}
            onClick={() => navigate("/admin/equation-types")}
          >
            <EquationTypesIcon className={sharedStyles.iconBlue} />            
            <div className={sharedStyles.dashboardCardTitle}>Типы примеров и операнды</div>
            <div className={sharedStyles.dashboardCardDesc}>
              Управление типами примеров и их операндами
            </div>
          </div>

          </div>
        <div className={sharedStyles.dashboardFooter} />
      </div>
    </div>
  );
}

export default AdminDashboard;