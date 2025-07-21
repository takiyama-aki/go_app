// src/components/SignupForm.tsx
//-------------------------------------
// 新規作成フォーム（/signup を呼び出す）
//-------------------------------------
import { useState } from "react";
import { useQueryClient } from "@tanstack/react-query";
import { signup } from "../api/auth";

export default function SignupForm() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [result, setResult] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);
  const queryClient = useQueryClient();

  const handleSignup = async () => {
    setResult(null);
    setError(null);
    try {
      const { message } = await signup(email, password);
      setResult(message);
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
        className="btn bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700 w-full"
        onClick={handleSignup}
        disabled={!email || !password}
      >
        新規作成
      </button>

      {result && <p className="text-green-700 text-sm">{result}</p>}
      {error  && <p className="text-red-600 text-sm">{error}</p>}
    </div>
  );
}
