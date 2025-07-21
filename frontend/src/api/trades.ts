import { client } from "./client";

export interface Trade {
  id: number;
  date: string;
  symbolName: string;
  symbolCode: string;
  price: number;
  quantity: number;
  side: string;
  profitLoss: number;
  manualEntry: boolean;
}

export interface TradeListResponse {
  trades: Trade[];
}

export interface GetTradeResponse {
  trade: Trade;
}

export interface TradeRequest {
  date: string;
  symbolName: string;
  symbolCode: string;
  price: number;
  quantity: number;
  side: string;
  profitLoss: number;
  manualEntry: boolean;
}

export const listTrades = (month?: string) =>
  client
    .get<TradeListResponse>("/trades", { params: month ? { month } : {} })
    .then((r) => r.data);

export const getTrade = (id: number) =>
  client.get<GetTradeResponse>(`/trades/${id}`).then((r) => r.data);

export const createTrade = (data: TradeRequest) =>
  client.post<{ id: number }>("/trades", data).then((r) => r.data);

export const updateTrade = (id: number, data: TradeRequest) =>
  client.put<void>(`/trades/${id}`, data).then((r) => r.data);