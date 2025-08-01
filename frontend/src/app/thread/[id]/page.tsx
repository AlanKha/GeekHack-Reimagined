'use client';

import { useEffect, useState } from 'react';
import { useParams } from 'next/navigation';
import api from '@/lib/api';
import PostCard from '@/components/PostCard';

interface Post {
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

interface Thread {
  id: number;
  title: string;
  posts: Post[];
}

const ThreadPage = () => {
  const params = useParams();
  const { id } = params;
  const [thread, setThread] = useState<Thread | null>(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    if (id) {
      api
        .get(`/threads/${id}`)
        .then((response) => {
          setThread(response.data);
        })
        .catch((error) => {
          console.error('Failed to fetch thread:', error);
        })
        .finally(() => {
          setLoading(false);
        });
    }
  }, [id]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!thread) {
    return <div>Thread not found</div>;
  }

  return (
    <div>
      <h1 className="text-2xl font-bold mb-4">{thread.title}</h1>
      <div className="space-y-4">
        {thread.posts.map((post) => (
          <PostCard key={post.id} post={post} />
        ))}
      </div>
    </div>
  );
};

export default ThreadPage;
