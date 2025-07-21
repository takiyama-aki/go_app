import { Link } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { getCurrentUser } from "../api/auth";

export default function Header() {
  const { data } = useQuery({
    queryKey: ["me"],
    queryFn: getCurrentUser,
    retry: false,
  });

  return (
    <header className="h-14 flex items-center px-4 shadow">
      <h1 className="font-bold">Go&nbsp;Trade&nbsp;App</h1>
      <nav className="ml-auto space-x-4">
        <Link to="/" className="text-blue-600 hover:underline">
          Home
        </Link>
        <Link to="/about" className="text-blue-600 hover:underline">
          About
        </Link>
        <Link to="/trades" className="text-blue-600 hover:underline">
          Trades
        </Link>
        <Link to="/login" className="text-blue-600 hover:underline">
          Login
        </Link>
      </nav>
      <span className="ml-4 text-sm text-gray-600">
        {data?.email ?? "未ログイン"}
      </span>
    </header>
  );
}
