import { useState, useEffect } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'
import axios from 'axios'

function App() {
  const [todos, setTodos] = useState([])
  const [task, setTask] = useState('')
  const [completed, setCompleted] = useState(false)

  useEffect(() => {
    axios
      .get("http://backend:8080/todos") // Changed from localhost
      .then((res) => setTodos(res.data))
      .catch((err) => console.error(err));
  }, [])

  const addTodo = () => {
    axios.post('http://backend:8080/todos', {task})
    .then(res => setTodos([...todos, res.data]))
    .catch(err => console.error(err))
  }
  
  return (
    <>
      <div>
        <h1>TODO LIST APP (Go + React + Docker)</h1>
        <input type="text" value={task} onChange={(e) => setTask(e.target.value)} />
        <button onClick={addTodo}>Add</button>
        <ul>
          {todos.map(todo => (
            <li key={todo.id}>{todo.task}</li>
          ))}
        </ul>
      </div>
    </>
  )
}

export default App
