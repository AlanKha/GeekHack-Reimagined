import { User } from './user';

export interface Post {
  id: number;
  content: string;
  user: User;
  createdAt: string;
  updatedAt: string;
}

export interface PostData {
  content?: string;
}
