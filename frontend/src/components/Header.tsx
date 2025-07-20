import { Link } from "react-router-dom";

export default function Header() {
  return (
    <header className="h-14 flex items-center px-4 shadow">
      <h1 className="font-bold">Go&nbsp;Trade&nbsp;App</h1>
              <nav className="space-x-4">
          <Link to="/" className="text-blue-600 hover:underline">
            Home
          </Link>
          <Link to="/about" className="text-blue-600 hover:underline">
            About
          </Link>
          <Link to="/login" className="text-blue-600 hover:underline">
            Login
          </Link>
        </nav>

    </header>
  );
}
