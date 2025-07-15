import axios from 'axios';
import { LoginData, RegisterData } from '../types/auth';
import { Thread, ThreadData } from '../types/thread';
import { PostData } from '../types/post';

const api = axios.create({
  baseURL: process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api',
  headers: {
    'Content-Type': 'application/json',
  },
  withCredentials: true, // Send cookies with requests
});

// Auth
export const register = (data: RegisterData) => api.post('/register', data);
export const login = (data: LoginData) => api.post('/login', data);

// Threads
export const getThreads = (): Promise<{ data: { threads: Thread[] } }> =>
  api.get('/threads');
export const getThread = (id: string): Promise<{ data: { thread: Thread } }> =>
  api.get(`/threads/${id}`);
export const createThread = (data: ThreadData) => api.post('/threads', data);
export const updateThread = (id: string, data: ThreadData) =>
  api.put(`/threads/${id}`, data);
export const deleteThread = (id: string) => api.delete(`/threads/${id}`);

// Posts
export const createPost = (threadId: string, data: PostData) =>
  api.post(`/threads/${threadId}/posts`, data);
export const updatePost = (id: string, data: PostData) =>
  api.put(`/posts/${id}`, data);
export const deletePost = (id: string) => api.delete(`/posts/${id}`);

export default api;
