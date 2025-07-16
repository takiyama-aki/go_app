// src/pages/About.tsx
//-------------------------------------
// About ページ
// ・メール／パスワードの 2 テキストボックス
// ・「新規作成」「ログイン」ボタン
// ・API 返却メッセージをその場で表示
//-------------------------------------
import { useState } from "react";
import { signup, login } from "../api/auth";
import { useAuthStore } from "../store/auth";

export default function About() {
  const [email, setEmail]       = useState("");
  const [password, setPassword] = useState("");
  const [msg, setMsg]           = useState<string | null>(null);
  const [err, setErr]           = useState<string | null>(null);

  const saveToken = useAuthStore((s) => s.login); // ログイン成功時だけ使用

  /** 新規作成 */
  const handleSignup = async () => {
    setMsg(null); setErr(null);
    try {
      const { message } = await signup(email, password);
      setMsg(message);                  // ← 常に表示される
    } catch (e: any) {
      setErr(e.response?.data?.message || e.message);
    }
  };

  /** ログイン */
  const handleLogin = async () => {
    setMsg(null); setErr(null);
    try {
      const { message, token } = await login(email, password);
      if (token) saveToken(token);        // トークン保存
      setMsg(message);
    } catch (e: any) {
      setErr(e.response?.data?.message || e.message);
    }
  };

  return (
    <div className="space-y-6 max-w-md w-full mx-auto bg-white p-8 rounded-2xl shadow">
      <h2 className="text-3xl font-bold text-center">About Page</h2>

      {/* テキストボックス */}
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

      {/* ボタン */}
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
      </div>

      {/* メッセージ表示 */}
      {msg && <p className="text-green-700 text-sm break-words">{msg}</p>}
      {err && <p className="text-red-600 text-sm break-words">{err}</p>}
    </div>
  );
}
