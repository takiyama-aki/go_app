// src/store/auth.ts
//-------------------------------------
// ログイン状態を保持する zustand ストア
//-------------------------------------
import { create } from "zustand";

interface AuthState {
  isLoggedIn: boolean;         // 認証フラグ
  login: () => void;           // ログイン時に呼び出す
  logout: () => void;          // ログアウト時に呼び出す
}

// 認証状態を管理する zustand ストア
export const useAuthStore = create<AuthState>((set) => ({
  isLoggedIn: false,

  /** ログイン状態 ON */
  login: () => {
    set({ isLoggedIn: true });
  },

  /** ログアウトで状態 OFF */
  logout: () => {
    set({ isLoggedIn: false });
  },
}));
