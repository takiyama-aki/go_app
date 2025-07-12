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

    // <div className="flex items-center justify-center min-h-screen bg-gray-50">
    //   <h1 className="text-4xl font-bold">Hello, React + Tailwind!</h1>
    //   <h2> test</h2>
    // </div>
  );
}

export default App
