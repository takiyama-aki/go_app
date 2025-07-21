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
export const login = (email: string, password: string) =>
  client
    .post<LoginResponse>("/login", { email, password })
    .then((r) => r.data);

export interface MeResponse {
  id: number;
  email: string;
}

export const getCurrentUser = () =>
  client.get<MeResponse>("/me").then((r) => r.data);

