import Link from 'next/link';

import { Post } from '../types/post';

interface PostCardProps {
  post: Post;
}

const PostCard = ({ post }: PostCardProps) => {
  return (
    <div className="bg-surface border border-outline rounded-lg flex">
      <div className="bg-black/20 p-2 flex flex-col items-center justify-start space-y-2 w-12">
        <button>
          <span className="material-icons-outlined text-on-surface-variant hover:text-primary">
            arrow_upward
          </span>
        </button>
        <span className="font-bold text-sm">{post.votes}</span>
        <button>
          <span className="material-icons-outlined text-on-surface-variant hover:text-blue-500">
            arrow_downward
          </span>
        </button>
      </div>
      <div className="p-4 flex-1">
        <p className="text-xs text-on-surface-variant mb-2">
          Posted by{' '}
          <Link className="hover:underline" href={post.authorUrl}>
            {post.author}
          </Link>{' '}
          {post.createdAt} in{' '}
          <Link
            className="font-bold text-on-surface hover:underline"
            href={post.communityUrl}
          >
            {post.community}
          </Link>
        </p>
        <Link
          className="text-lg font-semibold text-on-surface hover:text-primary transition-colors block"
          href="#"
        >
          {post.title}
        </Link>
        <p className="text-on-surface-variant text-sm mt-1">
          {post.description}
        </p>
        <div className="flex items-center space-x-4 mt-3 text-sm text-on-surface-variant">
          <Link
            className="flex items-center space-x-1 hover:bg-zinc-800 p-1 rounded"
            href="#"
          >
            <span className="material-icons-outlined text-base">
              chat_bubble_outline
            </span>
            <span>{post.comments} Comments</span>
          </Link>
          <Link
            className="flex items-center space-x-1 hover:bg-zinc-800 p-1 rounded"
            href="#"
          >
            <span className="material-icons-outlined text-base">share</span>
            <span>Share</span>
          </Link>
          <Link
            className="flex items-center space-x-1 hover:bg-zinc-800 p-1 rounded"
            href="#"
          >
            <span className="material-icons-outlined text-base">
              bookmark_border
            </span>
            <span>Save</span>
          </Link>
        </div>
      </div>
    </div>
  );
};

export default PostCard;
