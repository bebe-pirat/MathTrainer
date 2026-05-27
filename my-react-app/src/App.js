import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import { useAuth, AuthProvider } from "./AuthContext";
import Login from "./pages/Login/Login";
import ProtectedRoute from "./ProtectedRoute";
import { ROLES } from "./constants";
import AdminDashboard from "./pages/admin/AdminDashboard";
import TeacherDashboard from "./pages/teacher/TeacherDashboard/TeacherDashboard";
import TeachersPage from "./pages/admin/TeachersPage";
import SchoolsPage from "./pages/admin/SchoolsPage";
import UsersPage from "./pages/admin/UsersPage/UsersPage";
import ClassesPage from "./pages/admin/ClassesPage";
import ClassStatistics from "./pages/teacher/ClassStatisticsPage/ClassStatisticsPage";
import StudentsPage from "./pages/teacher/StudentListPage/StudentListPage";
import HomePage from "./pages/HomePage";
import LevelMapPage from "./pages/student/LevelsMapPage/LevelsMapPage";
import ProfilePage from "./pages/student/ProfilePage/ProfilePage";
import GamePage from "./pages/GamePage/GamePage";
import StatsPage from "./pages/student/StatsPage/StatsPage";
import StudentAttemptsPage from "./pages/teacher/StudentAttemptsPage/StudentAttemptsPage";
import SectionsPage from "./pages/admin/SectionsPage/SectionsPage";
import EquationTypesPage from "./pages/admin/EquationTypesPage/EquationTypesPage";
import SectionsEquationTypesPage from "./pages/admin/SectionsEquationTypesPage/SectionsEquationTypesPage";
import DirectorSchoolStats from "./pages/director/DirectorStatsPage/DirectorStatsPage";
import DirectorDashboard from "./pages/director/DirectorDashboard";

function AppRoutes() {
    const { user, loading } = useAuth();
    
    console.log("AppRoutes - user:", user); 
    console.log("AppRoutes - loading:", loading);
    if (loading) {
        return <div>Загрузка...</div>;
    }

    console.log("User role:", user?.role);

    return (
        <Routes>
            <Route path="/" element={<HomePage />} />
            <Route path="/login" element={<Login />} />

            <Route
                path="/admin/dashboard"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                        <AdminDashboard/>
                    </ProtectedRoute>
                }
            />

            <Route
                path="/admin/teachers"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                        <TeachersPage/>
                    </ProtectedRoute>
                }
            />
            <Route
                path="/admin/schools"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                        <SchoolsPage/>
                    </ProtectedRoute>
                }
            />
            <Route
                path="/admin/users"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                        <UsersPage/>
                    </ProtectedRoute>
                }
            />
            <Route
                path="/admin/classes"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                        <ClassesPage/>
                    </ProtectedRoute>
                }
            />
            <Route
                path="/admin/sections"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                        <SectionsPage/>
                    </ProtectedRoute>
                }
            />
            <Route
                path="/admin/equation-types"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                        <EquationTypesPage/>
                    </ProtectedRoute>
                }
            />
            <Route
                path="/admin/sections-equation-types"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                        <SectionsEquationTypesPage/>
                    </ProtectedRoute>
                }
            />
            <Route
                path="/student/dashboard"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.STUDENT]}>
                        <LevelMapPage/>
                    </ProtectedRoute>
                }
            />

            <Route
                path="/teacher/dashboard"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.TEACHER]}>
                        <TeacherDashboard/>
                    </ProtectedRoute>
                }
            />

            <Route
                path="/teacher/students"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.TEACHER]}>
                        <StudentsPage/>
                    </ProtectedRoute>
                }
            />

            <Route 
                path="/director/school-stats/" 
                element={
                    <ProtectedRoute allowedRoles={[ROLES.HEAD]}>
                        <DirectorSchoolStats/>
                    </ProtectedRoute>
                }
            />
        
            <Route 
                path="/director/dashboard/" 
                element={
                    <ProtectedRoute allowedRoles={[ROLES.HEAD]}>
                        <DirectorDashboard/>
                    </ProtectedRoute>
                }
            />
        

            <Route
                path="/game"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.STUDENT]}>
                        <GamePage/>
                    </ProtectedRoute>
                }
            />

              <Route
                path="/student/stats"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.STUDENT]}>
                        <StatsPage/>
                    </ProtectedRoute>
                }
            />

            <Route
                path="/student/profile"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.STUDENT]}>
                        <ProfilePage />
                    </ProtectedRoute>
                }
            />

            <Route
                path="/teacher/class-statistics"
                element={
                    <ProtectedRoute allowedRoles={[ROLES.TEACHER]}>
                        <ClassStatistics/>
                    </ProtectedRoute>
                }
            />

            <Route 
                path="/teacher/student-attempts/:studentId" 
                element={
                    <ProtectedRoute allowedRoles={[ROLES.TEACHER]}>
                        <StudentAttemptsPage />
                    </ProtectedRoute>
                } 
            />
         
        </Routes>
    );
}

function App() {
    return (
        <AuthProvider>
            <BrowserRouter>
                <AppRoutes />
            </BrowserRouter>
        </AuthProvider>
    );
}

export default App;