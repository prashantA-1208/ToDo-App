import { useState } from 'react';
import { createTask } from '../api/tasks';
import CreateTaskForm from '../components/CreateTaskForm';

export default function CreateTaskPage({ onTaskCreated }) {
  const [title, setTitle] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    await createTask({ title, completed: false });
    setTitle('');
    onTaskCreated();
  };

  return (
    <CreateTaskForm
        title={title}
        setTitle={setTitle}
        handleSubmit={handleSubmit}
    />
  );
}
