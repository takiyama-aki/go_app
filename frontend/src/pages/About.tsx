// src/pages/About.tsx
// アプリの概要を説明する静的ページ

// アプリ紹介コンテンツを表示するコンポーネント
export default function About() {
  return (
    <div className="space-y-6 max-w-md w-full mx-auto bg-white p-8 rounded-2xl shadow">
      <h2 className="text-3xl font-bold text-center">About Page</h2>
      <p>
        Go Trade App is a sample application demonstrating a simple authentication
        flow using a Go backend and a React&nbsp;frontend.
      </p>
      <p>
        The frontend is built with Vite, React Router and Zustand for state
        management. Use the navigation links above to explore the login flow and
        API responses.
      </p>
    </div>
  );
}

