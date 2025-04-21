import TaskList from '../components/TaskList';
import CreateTaskPage from './CreateTaskPage';

export default function Dashboard() {
  return (
    <div className="p-6">
      <h1 className="text-xl mb-4">Your Tasks</h1>
      <CreateTaskPage onTaskCreated={() => window.location.reload()} />
      <TaskList />
    </div>
  );
}