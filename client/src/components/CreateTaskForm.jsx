export default function CreateTaskForm ({title, setTitle, handleSubmit}){
    return (
      <div className="container">
        <form onSubmit={handleSubmit} className="">
      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="Add task"
      />
      <button type="submit" className="bg-blue-500 text-white px-4 py-1 rounded">Add</button>
    </form>
    </div>
    );
}