// src/pages/About.tsx
//-------------------------------------
// About ページ – 上の 2 フォームを読み込むだけ
//-------------------------------------
import SignupForm from "../components/SignupForm";
import LoginForm  from "../components/LoginForm";

export default function About() {
  return (
    <div className="space-y-6 max-w-md w-full mx-auto bg-white p-8 rounded-2xl shadow">
      <h2 className="text-3xl font-bold text-center">About Page</h2>

      {/* 新規作成フォーム */}
      <SignupForm />

      {/* ログインフォーム */}
      <LoginForm />
    </div>
  );
}
