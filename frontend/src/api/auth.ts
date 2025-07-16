// src/api/auth.ts
//-------------------------------------
// ※レスポンス内に必ず `message` フィールドがある想定に変更
//-------------------------------------
import { client } from "./client";

export interface SignupResponse {
  message: string;
}

export interface LoginResponse {
  message: string;
  token?: string;
}

/** 新規作成 */
export const signup = (email: string, password: string) =>
  client
    .post<SignupResponse>("/signup", { email, password })
    .then((r) => r.data);

/** ログイン */
export const login = (email: string, password: string) =>
  client
    .post<LoginResponse>("/login", { email, password })
    .then((r) => r.data);
