export default function LoginForm({ email, password, onEmailChange, onPasswordChange, onSubmit }) {
    return (
      <form onSubmit={onSubmit} className="p-4">
        <input placeholder="Email" value={email} onChange={(e) => onEmailChange(e.target.value)} className="block mb-2 border px-2 py-1" />
        <input placeholder="Password" type="password" value={password} onChange={(e) => onPasswordChange(e.target.value)} className="block mb-2 border px-2 py-1" />
        <button type="submit" className="bg-blue-500 text-white px-4 py-1">Login</button>
      </form>
    );
  }