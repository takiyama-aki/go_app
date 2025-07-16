// src/store/auth.ts
//-------------------------------------
// ログイン状態を保持する zustand ストア
//-------------------------------------
import { create } from "zustand";

interface AuthState {
  token: string | null;        // JWT など
  isLoggedIn: boolean;         // 認証フラグ
  login:  (token: string) => void;
  logout: () => void;
}

export const useAuthStore = create<AuthState>((set) => ({
  token:      localStorage.getItem("token"),
  isLoggedIn: !!localStorage.getItem("token"),

  /** トークン保存＆ログイン状態 ON */
  login: (token) => {
    localStorage.setItem("token", token);
    set({ token, isLoggedIn: true });
  },

  /** トークン削除＆ログイン状態 OFF */
  logout: () => {
    localStorage.removeItem("token");
    set({ token: null, isLoggedIn: false });
  },
}));
