import { Post } from './post';
import { User } from './user';

export interface Thread {
  id: number;
  title: string;
  content: string;
  user: User;
  posts: Post[];
  createdAt: string;
  updatedAt: string;
}

export interface ThreadData {
  title?: string;
  content?: string;
}
