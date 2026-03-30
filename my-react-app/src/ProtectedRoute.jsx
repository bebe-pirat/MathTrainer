import { Navigate } from "react-router-dom";
import { useAuth } from "./AuthContext";

function ProtectedRoute({ children, allowedRoles }) {
    const { user } = useAuth();

    if (!user) {
        return <Navigate to="/login"/>;
    }

    if (allowedRoles && !allowedRoles.includes(user.role_id)) {
        return <Navigate to="/login"/>;
    }
    return children;
}

export default ProtectedRoute;