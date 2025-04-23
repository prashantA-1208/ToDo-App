import { useAuth } from '../context/AuthContext';


const Header = () => {
  const { user, logout } = useAuth();

  return (
    <header >
      <h1 className="header-h1">📝 ToDo App {user.username ? `Hello ${user.username}` : `Please Login To See Your Task` }!</h1>

      <div className="header-button">
        <button
          onClick={logout}
        >
          Logout
        </button>
        
      </div>
    </header>
  );
};

export default Header;


