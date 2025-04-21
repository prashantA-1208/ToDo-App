import { useEffect, useState } from 'react';
import { getTasks, deleteTask, updateTask } from '../api/tasks';

export default function TaskList() {
  const [tasks, setTasks] = useState([]);

  const loadTasks = async () => {
    try {
      const res = await getTasks();
      setTasks(res.data || []); // fallback to empty array
    } catch (error) {
      console.error('Failed to load tasks:', error);
      setTasks([]); // fallback to avoid crash
    }
  };
  

  useEffect(() => {
    loadTasks();
  }, []);

  const handleDelete = async (id) => {
    await deleteTask(id);
    loadTasks();
  };

  const handleToggle = async (task) => {
    await updateTask(task.id, { ...task, completed: !task.completed });
    loadTasks();
  };

  return (
    <ul>
      {tasks.map((task) => (
        <li key={task.id} className="flex justify-between items-center py-1">
          <span className={task.completed ? 'line-through' : ''}>{task.title}</span>
          <div className="flex gap-2">
            <button onClick={() => handleToggle(task)} className="text-sm text-green-600"> {task.completed ? 'Completed' : 'Pending'}</button>
            <button onClick={() => handleDelete(task.id)} className="text-sm text-red-600">Delete</button>
          </div>
        </li>
      ))}
    </ul>
  );
}
