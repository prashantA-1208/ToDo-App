import axios from './axios';

export const getTasks = () => axios.get('/tasks');
export const getTaskById = (id) => axios.get(`/tasks/${id}`);
export const createTask = (task) => axios.post('/tasks', task);
export const updateTask = (id, task) => axios.put(`/tasks/${id}`, task);
export const deleteTask = (id) => axios.delete(`/tasks/${id}`);
export const getUser = () => axios.get(`/user`);