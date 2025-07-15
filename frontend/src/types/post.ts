export interface Post {
  id: number;
  author: string;
  authorUrl: string;
  createdAt: string;
  community: string;
  communityUrl: string;
  title: string;
  description: string;
  votes: string;
  comments: string;
}

export interface PostData {
  content?: string;
}
