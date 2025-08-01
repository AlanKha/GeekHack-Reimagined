import PostCard from '@/components/PostCard';
import TopCommunities from '@/components/TopCommunities';
import CreatePost from '@/components/CreatePost';
import { Post } from '../types/post';

const posts: Post[] = [
  {
    id: 1,
    author: 'u/S&M Designs',
    authorUrl: '#',
    createdAt: '2 hours ago',
    community: 'g/Announcements',
    communityUrl: '#',
    title: 'Re: Keep getting Error 4...',
    description:
      "Hey everyone, we've seen the reports and are actively investigating the Error 404 issues some of you are experiencing. We appreciate your patience!",
    votes: '1.8k',
    comments: '894',
  },
  {
    id: 2,
    author: 'u/HoffmanHyster',
    authorUrl: '#',
    createdAt: '5 hours ago',
    community: 'g/Keyboards',
    communityUrl: '#',
    title: 'Re: Lucky65v2 Pre-Sale',
    description:
      "The pre-sale for the Lucky65v2 is now LIVE! Limited quantities available. Don't miss out on this fantastic board. Link in the post.",
    votes: '1.2k',
    comments: '28k',
  },
  {
    id: 3,
    author: 'u/easygrader',
    authorUrl: '#',
    createdAt: '7 hours ago',
    community: 'g/KeyCon2025',
    communityUrl: '#',
    title: 'Re: KeyCon 2025 Call for...',
    description:
      "We are officially opening the call for presenters and vendors for KeyCon 2025! If you have a talk you'd like to give or a booth you'd like to run, please see the application details inside.",
    votes: '986',
    comments: '125',
  },
  {
    id: 4,
    author: 'u/papayanana',
    authorUrl: '#',
    createdAt: '1 day ago',
    community: 'g/NewMembers',
    communityUrl: '#',
    title: 'Re: Welcome new members!',
    description:
      "Just wanted to say hello! I'm new to the mechanical keyboard world and this community seems amazing. Excited to learn from you all!",
    votes: '542',
    comments: '7.4k',
  },
];

export default function Home() {
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
