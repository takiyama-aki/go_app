import axios from "axios";

export const client = axios.create({
  // .env で VITE_API_URL を上書き可能にしておくと環境切り替えが楽
  baseURL: import.meta.env.VITE_API_URL || "http://localhost:8081",
  timeout: 10000,
});

export default client;