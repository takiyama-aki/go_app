import { Routes, Route, Link } from "react-router-dom";

import Header from "./components/Header";
import Home from "./pages/Home";
import About from "./pages/About";
import './App.css'

function App() {
  return (
    <div className="grid grid-rows-[auto_1fr] min-h-screen">
      <Header />
      <main className="flex items-center justify-center bg-gray-50">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/about" element={<About />} />
        </Routes>
      </main>
    </div>
  );
}

export default App
