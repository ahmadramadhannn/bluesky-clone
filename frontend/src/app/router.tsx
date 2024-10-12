import { useQuery } from "@tanstack/react-query";
import { createBrowserRouter, RouterProvider } from "react-router-dom";


const Example = () => {
  const { data, isLoading } = useQuery({
    queryKey: ["todos"], queryFn: async () => {
      const response = await fetch("https://jsonplaceholder.typicode.com/todos")
      const todos = await response.json()
      return todos
    }
  })

  if (isLoading) return <h1>Loading ....</h1>


  return (
    <>
      <h1>{JSON.stringify(data)}</h1>
    </>
  )
}


const router = createBrowserRouter([
  {
    path: "/",
    element: <Example />
  }
])

export const AppRouter = () => {
  return <RouterProvider router={router} />
}
