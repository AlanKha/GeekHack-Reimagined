'use client';

import { useEffect, useState } from 'react';
import PostCard from '@/components/PostCard';
import TopCommunities from '@/components/TopCommunities';
import CreatePost from '@/components/CreatePost';
import { Post } from '../types/post';
import { posts as initialPosts } from './_data';

export default function Home() {
  const [posts, setPosts] = useState<Post[]>([]);

  useEffect(() => {
    // Simulate fetching data from an API
    setPosts(initialPosts);
  }, []);

  return (
    <div className="grid grid-cols-12 gap-8">
      <div className="col-span-12 lg:col-span-8">
        <div className="flex items-center mb-4 space-x-2">
          <button className="bg-surface hover:bg-surface/80 text-on-surface font-semibold py-2 px-4 rounded-full flex items-center space-x-2">
            <span className="material-icons-outlined">whatshot</span>
            <span>Hot</span>
          </button>
          <button className="bg-transparent text-on-surface-variant font-semibold py-2 px-4 rounded-full flex items-center space-x-2">
            <span>New</span>
          </button>
          <button className="bg-transparent text-on-surface-variant font-semibold py-2 px-4 rounded-full flex items-center space-x-2">
            <span>Top</span>
          </button>
        </div>
        <div className="space-y-4">
          {posts.map((post) => (
            <PostCard key={post.id} post={post} />
          ))}
        </div>
      </div>
      <div className="col-span-12 lg:col-span-4 space-y-4">
        <TopCommunities />
        <CreatePost />
      </div>
    </div>
  );
}
