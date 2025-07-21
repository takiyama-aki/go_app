import { useState } from "react";
import { useQuery, useMutation, useQueryClient } from "@tanstack/react-query";
import {
  listTrades,
  createTrade,
  updateTrade,
} from "../api/trades";
import type { Trade, TradeRequest } from "../api/trades";

const emptyForm: TradeRequest = {
  date: new Date().toISOString().slice(0, 10),
  symbolName: "",
  symbolCode: "",
  price: 0,
  quantity: 1,
  side: "LONG",
  profitLoss: 0,
  manualEntry: false,
};

export default function Trades() {
  const queryClient = useQueryClient();
  const [form, setForm] = useState<TradeRequest>(emptyForm);
  const [editingId, setEditingId] = useState<number | null>(null);

  const { data } = useQuery({
    queryKey: ["trades"],
    queryFn: () => listTrades(),
  });

  const createMutation = useMutation({
    mutationFn: createTrade,
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["trades"] }),
  });

  const updateMutation = useMutation({
    mutationFn: ({ id, data }: { id: number; data: TradeRequest }) =>
      updateTrade(id, data),
    onSuccess: () => {
      setEditingId(null);
      queryClient.invalidateQueries({ queryKey: ["trades"] });
    },
  });

  const handleChange = (
    e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement>,
  ) => {
    const { name, value } = e.target;
    setForm({ ...form, [name]: value });
  };

  const handleSubmit = () => {
    if (editingId) {
      updateMutation.mutate({ id: editingId, data: form });
    } else {
      createMutation.mutate(form);
    }
    setForm(emptyForm);
  };

  const startEdit = (t: Trade) => {
    setEditingId(t.id);
    setForm({
      date: t.date.slice(0, 10),
      symbolName: t.symbolName,
      symbolCode: t.symbolCode,
      price: t.price,
      quantity: t.quantity,
      side: t.side,
      profitLoss: t.profitLoss,
      manualEntry: t.manualEntry,
    });
  };

  return (
    <div className="space-y-6 max-w-xl w-full mx-auto bg-white p-8 rounded-2xl shadow">
      <h2 className="text-3xl font-bold text-center">Trades</h2>
      <div className="space-y-4">
        <input
          type="date"
          name="date"
          value={form.date}
          onChange={handleChange}
          className="input w-full"
        />
        <input
          type="text"
          name="symbolName"
          placeholder="Symbol Name"
          value={form.symbolName}
          onChange={handleChange}
          className="input w-full"
        />
        <input
          type="text"
          name="symbolCode"
          placeholder="Symbol Code"
          value={form.symbolCode}
          onChange={handleChange}
          className="input w-full"
        />
        <input
          type="number"
          name="price"
          placeholder="Price"
          value={form.price}
          onChange={handleChange}
          className="input w-full"
        />
        <input
          type="number"
          name="quantity"
          placeholder="Quantity"
          value={form.quantity}
          onChange={handleChange}
          className="input w-full"
        />
        <select
          name="side"
          value={form.side}
          onChange={handleChange}
          className="input w-full"
        >
          <option value="LONG">LONG</option>
          <option value="SHORT">SHORT</option>
        </select>
        <button
          className="btn bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 w-full"
          onClick={handleSubmit}
        >
          {editingId ? "更新" : "登録"}
        </button>
      </div>

      <table className="w-full text-left border-t mt-8">
        <thead>
          <tr>
            <th className="py-2">Date</th>
            <th>Symbol</th>
            <th>Price</th>
            <th>Qty</th>
            <th>Side</th>
            <th></th>
          </tr>
        </thead>
        <tbody>
          {data?.trades.map((t) => (
            <tr key={t.id} className="border-t">
              <td className="py-2">{t.date.slice(0, 10)}</td>
              <td>{t.symbolName}</td>
              <td>{t.price}</td>
              <td>{t.quantity}</td>
              <td>{t.side}</td>
              <td>
                <button
                  className="text-blue-600 hover:underline"
                  onClick={() => startEdit(t)}
                >
                  編集
                </button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}