// src/api/auth.ts
// 認証系 API をまとめたモジュール
// ※レスポンス内に必ず `message` フィールドがある想定に変更
import { client } from "./client";

export interface SignupResponse {
  message: string;
}

export interface LoginResponse {
  message: string;
}

/** 新規作成 */
// アカウント作成 API
export const signup = async (
  email: string,
  password: string,
): Promise<SignupResponse> => {
  await client.post("/signup", { email, password });
  // 本文が無ければデフォルトメッセージを付与
  // const res = await client.post("/signup", { email, password });
  // return res.data ?? { message: "アカウントを作成しました" };
  return { message: "アカウントを作成しました。" };
};

/** ログイン */
// ログイン API
export const login = (email: string, password: string) =>
  client
    .post<LoginResponse>("/login", { email, password })
    .then((r) => r.data);

// ログアウト API
export const logout = () =>
  client.post<{ message: string }>("/logout").then((r) => r.data);

export interface MeResponse {
  id: number;
  email: string;
}

// セッション中のユーザー情報取得 API
export const getCurrentUser = () =>
  client.get<MeResponse>("/me").then((r) => r.data);

