import { useAuth } from '../context/AuthContext';
import { Link } from 'react-router-dom';


const Header = () => {
  const { user, logout } = useAuth();

  return (
    <header >
      <h1 className="header-h1">📝 ToDo App {user.username ? `Hello ${user.username}` : `Please Login To See Your Task` }!</h1>

      <div className="header-button">
        {user.username ? (
          <button
            onClick={logout}
            className="link-button"
          >
            Logout
          </button>
        ) : (
          <>
            <Link
              to="/login"
              className="link-button"
            >
              Login
            </Link>
            <Link
              to="/signup"
              className="link-button"
            >
              SignUp
            </Link>
          </>
        )}
      </div>
    </header>
  );
};

export default Header;


