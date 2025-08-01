import Link from 'next/link';

const Sidebar = () => {
  return (
    <aside className="w-64 sticky top-0 h-screen bg-black p-4 border-r border-outline flex-shrink-0 hidden lg:block">
      <div className="flex justify-center items-center space-x-3 mb-8">
        <h1 className="text-xl font-bold text-white">GeekHack</h1>
      </div>
      <nav className="space-y-2">
        <Link
          className="flex items-center space-x-3 px-3 py-2 rounded-md text-on-surface hover:bg-surface transition-colors"
          href="#"
        >
          <span className="material-icons-outlined">home</span>
          <span>Home</span>
        </Link>
        <Link
          className="flex items-center space-x-3 px-3 py-2 rounded-md text-on-surface hover:bg-surface transition-colors"
          href="#"
        >
          <span className="material-icons-outlined">trending_up</span>
          <span>Popular</span>
        </Link>
        <Link
          className="flex items-center space-x-3 px-3 py-2 rounded-md text-on-surface hover:bg-surface transition-colors"
          href="#"
        >
          <span className="material-icons-outlined">explore</span>
          <span>All</span>
        </Link>
      </nav>
      <div className="mt-6 pt-4 border-t border-outline">
        <h2 className="text-xs font-semibold text-on-surface-variant px-3 mb-2 uppercase">
          Forums
        </h2>
        <nav className="space-y-1">
          <Link
            className="flex items-center space-x-3 px-3 py-2 rounded-md text-on-surface hover:bg-surface transition-colors"
            href="#"
          >
            <span className="material-icons-outlined text-blue-400">
              campaign
            </span>
            <span>Announcements</span>
          </Link>
          <Link
            className="flex items-center space-x-3 px-3 py-2 rounded-md text-on-surface hover:bg-surface transition-colors"
            href="#"
          >
            <span className="material-icons-outlined text-green-400">
              event
            </span>
            <span>KeyCon 2025</span>
          </Link>
          <Link
            className="flex items-center space-x-3 px-3 py-2 rounded-md text-on-surface hover:bg-surface transition-colors"
            href="#"
          >
            <span className="material-icons-outlined text-purple-400">
              groups
            </span>
            <span>Community</span>
          </Link>
          <Link
            className="flex items-center space-x-3 px-3 py-2 rounded-md text-on-surface hover:bg-surface transition-colors"
            href="#"
          >
            <span className="material-icons-outlined text-red-400">
              keyboard
            </span>
            <span>Keyboards</span>
          </Link>
        </nav>
      </div>
    </aside>
  );
};

export default Sidebar;
