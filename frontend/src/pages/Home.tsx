import { useQuery } from "@tanstack/react-query";
import client from "../api/client";

export default function Home() {
  const { data, isLoading, error } = useQuery({
    queryKey: ["ping"],
    queryFn: () => client.get("/ping").then((res) => res.data),
  });

  if (isLoading) return <p>Loading...</p>;
  if (error) return (
    <p className="text-red-600">Error: {(error as Error).message}</p>
  );

  return (
    <div className="text-center space-y-4">
      <h2 className="text-3xl font-bold">Home Page</h2>
      <pre className="bg-gray-100 p-4 rounded-xl">{JSON.stringify(data, null, 2)}</pre>
    </div>
  );
}
