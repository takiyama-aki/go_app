// ルーティング設定と共通レイアウトを定義する
import { Routes, Route} from "react-router-dom";

import Header from "./components/Header";
import Home from "./pages/Home";
import About from "./pages/About";
import Login from "./pages/Login";
import Trades from "./pages/Trades";
import './App.css'

// アプリケーション全体のルーティングを返すコンポーネント
function App() {
  return (
    <div className="grid grid-rows-[auto_1fr] min-h-screen">
      <Header />
      <main className="flex items-center justify-center bg-gray-50">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
          <Route path="/trades" element={<Trades />} />
          <Route path="/login" element={<Login />} />
          
        </Routes>
      </main>
    </div>
  );
}

export default App
