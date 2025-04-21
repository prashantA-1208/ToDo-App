export default function CreateTaskForm ({title, setTitle, handleSubmit}){
    return (
        <form onSubmit={handleSubmit} className="flex gap-2 mb-4">
      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Add task"
        className="border rounded px-2 py-1"
      />
      <button type="submit" className="bg-blue-500 text-white px-4 py-1 rounded">Add</button>
    </form>
    );
}