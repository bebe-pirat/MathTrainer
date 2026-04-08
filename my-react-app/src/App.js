import { BrowserRouter, Routes, Route } from "react-router-dom";
import { AuthProvider } from "./AuthContext";
import Login from "./pages/Login";
import ProtectedRoute from "./ProtectedRoute";
import { ROLES } from "./constants";
import AdminDashboard from "./pages/admin/AdminDashboard";
import TeacherDashboard from "./pages/teacher/TeacherDashboard";
import TeachersPage from "./pages/admin/TeachersPage";
import SchoolsPage from "./pages/admin/SchoolsPage";
import UsersPage from "./pages/admin/UsersPage";
import ClassesPage from "./pages/admin/ClassesPage";
import ClassStatistics from "./pages/teacher/ClassStatisticsPage";

function App() {
    return (
        <AuthProvider>
            <BrowserRouter>
                <Routes>
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
                        path="/admin/dashboard"
                        element={
                            <ProtectedRoute allowedRoles={[ROLES.ADMIN]}>
                                <AdminDashboard/>
                            </ProtectedRoute>
                        }
                    />

                    <Route
                        path="/student/dashboard"
                        element={
                            <ProtectedRoute allowedRoles={[ROLES.STUDENT]}>
                                <div>Student</div>
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
                        path="/teacher/class-statistics"
                        element={
                            <ProtectedRoute allowedRoles={[ROLES.TEACHER]}>
                                <ClassStatistics/>
                            </ProtectedRoute>
                        }
                    />

                    <Route
                        path="/director/dashboard"
                        element={
                            <ProtectedRoute allowedRoles={[ROLES.HEAD]}>
                                <div>Director</div>
                            </ProtectedRoute>
                        }
                    />
                </Routes>
            </BrowserRouter>
        </AuthProvider>
    );
}

export default App;
