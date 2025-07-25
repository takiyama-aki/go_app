// 画面上部のナビゲーションバー
import { Link } from "react-router-dom";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import { getCurrentUser, logout } from "../api/auth";
import { useAuthStore } from "../store/auth";

// ヘッダーコンポーネント。ログイン状態に応じて表示を切り替える
export default function Header() {
  const queryClient = useQueryClient();
  const { logout: setLogout } = useAuthStore();

  // 現在のユーザー情報を取得
  const { data } = useQuery({
    queryKey: ["me"],
    queryFn: getCurrentUser,
    retry: false,
  });
  // ログアウト処理
  const { mutate } = useMutation({
    mutationFn: logout,
    onSuccess: () => {
      setLogout();
      queryClient.invalidateQueries({ queryKey: ["me"] });
    },
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
        {data ? (
          <button onClick={() => mutate()} className="text-blue-600 hover:underline">
            Logout
          </button>
        ) : (
          <Link to="/login" className="text-blue-600 hover:underline">
            Login
          </Link>
        )}
      </nav>
      <span className="ml-4 text-sm text-gray-600">
        {data?.email ?? "未ログイン"}
      </span>
    </header>
  );
}
