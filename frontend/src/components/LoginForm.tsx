// src/components/LoginForm.tsx
//-------------------------------------
// ログインフォーム（/login を呼び出す）
//-------------------------------------
import { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { login } from "../api/auth";
import { useAuthStore } from "../store/auth";

// ログインフォーム
export default function LoginForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [result, setResult] = useState<string | null>(null);
  const [error, setError]   = useState<string | null>(null);

  const queryClient = useQueryClient();
  const { login: setLogin } = useAuthStore();

  // ログインボタン押下時の処理
  const handleLogin = async () => {
    setResult(null);
    setError(null);
    try {
      const { message } = await login(email, password);
      setResult(message);
      setLogin();
      await queryClient.invalidateQueries({ queryKey: ["me"] });
    } catch (e: unknown) {
      const err = e as { response?: { data?: { message?: string } }; message?: string };
      setError(err.response?.data?.message || err.message || String(e));
    }
  };

  return (
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

      <button
        className="btn bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 w-full"
        onClick={handleLogin}
        disabled={!email || !password}
      >
        ログイン
      </button>

      {result && <p className="text-green-700 text-sm">{result}</p>}
      {error  && <p className="text-red-600 text-sm">{error}</p>}
    </div>
  );
}
