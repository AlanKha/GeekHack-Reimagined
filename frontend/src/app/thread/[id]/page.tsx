'use client';

import PostCard from '@/components/PostCard';
import { useThread } from '@/lib/useThread';

const ThreadPage = () => {
  const { thread, loading } = useThread();

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
