import React, { useEffect, useState } from "react";
import axios from "axios";
import "./App.css";
import { ToastContainer, toast } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";

const API_URL = "http://192.168.49.2:32053/todos";

function App() {
  const [todos, setTodos] = useState([]);
  const [task, setTask] = useState("");
  const [description, setDescription] = useState("");

  const fetchTodos = async () => {
    try {
      const response = await axios.get(API_URL);
      setTodos(response.data);
    } catch (err) {
      toast.error("Failed to fetch todos!");
      console.error(err);
    }
  };

  const addTodo = async () => {
    if (!task.trim() || !description.trim()) {
      toast.info("Please fill in both Task and Description!");
      return;
    }
    try {
      await axios.post(API_URL, {
        task,
        description,
        completed: false,
      });
      setTask("");
      setDescription("");
      toast.success("Todo added successfully!");
      fetchTodos();
    } catch (err) {
      toast.error("Failed to add todo!");
      console.error(err);
    }
  };

  const deleteTodo = async (id) => {
    try {
      await axios.delete(`${API_URL}/${id}`);
      toast.success("Todo deleted!");
      fetchTodos();
    } catch (err) {
      toast.error("Failed to delete todo!");
      console.error(err);
    }
  };

  const toggleCompleted = async (todo) => {
    try {
      await axios.put(`${API_URL}/${todo.id}`, {
        ...todo,
        completed: !todo.completed,
      });
      toast.success(
        `Marked as ${!todo.completed ? "completed" : "incomplete"}!`
      );
      fetchTodos();
    } catch (err) {
      toast.error("Failed to update todo status!");
      console.error(err);
    }
  };

  useEffect(() => {
    fetchTodos();
  }, []);

  return (
    <>
      <ToastContainer position="top-right" autoClose={2500} />
      <div className="glass-container">
        <div className="title">📝 My TODO List</div>

        <div className="input-section">
          <input
            type="text"
            placeholder="Task"
            value={task}
            onChange={(e) => setTask(e.target.value)}
            className="task-input"
          />
        </div>

        <div className="input-section">
          <input
            type="text"
            placeholder="Description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="task-input"
          />
          <button onClick={addTodo} className="add-button">
            Add
          </button>
        </div>

        <ul className="todo-list">
          {todos.map((todo) => (
            <li key={todo.id} className="todo-item">
              <div>
                <strong>{todo.task}</strong>
                <p>{todo.description}</p>
                <small>
                  Status: {todo.completed ? "✅ Done" : "⏳ Pending"}
                </small>
              </div>
              <div style={{ marginTop: "0.5rem" }}>
                <button
                  className="add-button"
                  onClick={() => toggleCompleted(todo)}
                >
                  {todo.completed ? "Mark Incomplete" : "Mark Complete"}
                </button>
                <button
                  className="add-button"
                  style={{ marginLeft: "0.5rem", background: "#ff4d4d" }}
                  onClick={() => deleteTodo(todo.id)}
                >
                  Delete
                </button>
              </div>
            </li>
          ))}
        </ul>
      </div>
    </>
  );
}

export default App;
