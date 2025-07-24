// ログインとサインアップを行うページ
import { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { signup, login } from "../api/auth";
import { useNavigate } from "react-router-dom";
import { useAuthStore } from "../store/auth";

// フォーム入力を受け付けて認証 API を呼び出すコンポーネント
export default function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const navigate = useNavigate();
  const queryClient = useQueryClient();
  const { login: setLogin } = useAuthStore();

  // Google ログイン処理
  const handleGoogleLogin = () => {
    const baseUrl = import.meta.env.VITE_API_URL || "http://localhost:8081";
    window.location.href = `${baseUrl}/oauth/login`;
  };

  // 新規ユーザー作成処理
  const handleSignup = async () => {
    await signup(email, password);
    setLogin();
    await queryClient.invalidateQueries({ queryKey: ["me"] });
    navigate("/");
  };

  // ログイン処理
  const handleLogin = async () => {
    await login(email, password);
    setLogin();
    await queryClient.invalidateQueries({ queryKey: ["me"] });
    navigate("/");
  };

  return (
    <div className="space-y-6 max-w-md w-full mx-auto bg-white p-8 rounded-2xl shadow">
      <h2 className="text-3xl font-bold text-center">Login</h2>
      <div className="space-y-4">
        <input
          type="email"
          placeholder="メールアドレス入力"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          className="input w-full"
        />
        <input
          type="password"
          placeholder="パスワード入力"
          value={password}
          onChange={(e) => setPassword(e.target.value)}
          className="input w-full"
        />
      </div>
      <div className="flex justify-between">
        <button
          className="btn bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700"
          onClick={handleSignup}
          disabled={!email || !password}
        >
          新規作成
        </button>
        <button
          className="btn bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
          onClick={handleLogin}
          disabled={!email || !password}
        >
          ログイン
        </button>
        <button
          className="btn bg-red-600 text-white px-4 py-2 rounded hover:bg-red-700"
          onClick={handleGoogleLogin}
        >
          Googleでログイン
        </button>
      </div>
    </div>
  );
}
