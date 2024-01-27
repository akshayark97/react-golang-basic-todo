import { useEffect, useState } from "react";
import "./App.css";

interface IItems {
  id: string;
  name: string;
  description: string;
}
function App() {
  const [items, setItems] = useState<IItems[]>([]);
  const [newTodo, setNewTodo] = useState("");

  useEffect(() => {
    const renderItems = async () => {
      await fetch(`http://localhost:8000/api/items`, {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
          accept: "application/json",
        },
      })
        .then((res) => res.json())
        .then((data) => setItems(data));
    };

    renderItems();
  }, []);

  const handleAddTodo = async () => {
    if (newTodo.length !== 0) {
      await fetch("http://localhost:8000/api/items", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ description: newTodo }),
      })
        .then((res) => res.json())
        .then((data) => setItems(data));

      setNewTodo("");
    }
  };

  const deleteItem = async (id: string) => {
    // await fetch(`http://localhost:8000/api/items/${id}`, {
    //   method: "DELETE",
    //   headers: {
    //     "Content-Type": "application/json",
    //   },
    // })
    //   .then((res) => res.json())
    //   .then((data) => setItems(data));

    // Comment below code if above api works
    setItems((prevData) => {
      return [...prevData.filter((data) => data.id !== id)];
    });
  };

  return (
    <>
      <h1>My Todo App</h1>
      <div>
        <input
          type="text"
          value={newTodo}
          onChange={(e) => setNewTodo(e.target.value)}
        />
        <button onClick={handleAddTodo}>Add Todo</button>
      </div>
      <ul style={{ listStyle: "none" }}>
        {items.map((item) => (
          <li key={item.id}>
            <strong onClick={() => deleteItem(item.id)}>{item.name}</strong>:{" "}
            {item.description}
          </li>
        ))}
      </ul>
    </>
  );
}

export default App;
