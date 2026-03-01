import { useEffect, useState } from "react"
import "./App.css"

function App() {
  const [tasks, setTasks] = useState([])
  const [title, setTitle] = useState("")
  const [error, setError] = useState("")

  const loadTasks = async () => {
    const res = await fetch("http://localhost:8080/tasks")
    const data = await res.json()
    setTasks(data)
  }

  useEffect(() => {
    loadTasks()
  }, [])

  const createTask = async () => {
    if (!title.trim()) {
      setError("Task cannot be empty")
      return
    }

    setError("")

    await fetch("http://localhost:8080/tasks", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ title }),
    })

    setTitle("")
    loadTasks()
  }

    const completeTask = async (id) => {
    await fetch(`http://localhost:8080/tasks/${id}`, {
      method: "PATCH",
    })
    loadTasks()
  }

  const deleteTask = async (id) => {
    await fetch(`http://localhost:8080/tasks/${id}`, {
      method: "DELETE",
    })
    loadTasks()
  }

  return (
    <div style={{ padding: 30 }}>
      <h1>Todo list</h1>
      <p className="version">version 1.0</p>

      <input
        value={title}
        onChange={(e) => {
          setTitle(e.target.value)
          if (error) setError("")
        }}
        placeholder="new task"
      />

      <button onClick={createTask}>Add</button>

      {error && (
        <p style={{ color: "red", marginTop: 5 }}>{error}</p>
      )}

      <ul>
        {tasks.map((t) => (
          <li key={t.id} style={{ marginTop: 10 }}>
            {t.title} — {t.completed ? "done" : "open"}

            <button onClick={() => completeTask(t.id)} style={{ marginLeft: 10 }}>
              ✔
            </button>

            <button onClick={() => deleteTask(t.id)} style={{ marginLeft: 6 }}>
              🗑
            </button>

            <div style={{ fontSize: 12, opacity: 0.7 }}>
              created: {
                new Date(t.created_at).toLocaleString()
              }
            </div>
            
            {t.completed_at && (
              <div style={{ fontSize: 12, opacity: 0.7 }}>
                completed: {new Date(t.completed_at).toLocaleString()}
              </div>
            )}
            
          </li>
        ))}
      </ul> 
    </div>
  )
}

export default App