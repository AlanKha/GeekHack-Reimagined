import Link from 'next/link';
import Image from 'next/image';

const Sidebar = () => {
  return (
    <aside className="w-64 sticky top-0 h-screen bg-black p-4 border-r border-outline flex-shrink-0 hidden lg:block">
      <div className="flex items-center space-x-3 mb-8">
        <Image
          alt="Geekhack logo"
          className="h-10 w-10"
          src="https://lh3.googleusercontent.com/aida-public/AB6AXuBKbeAccgf30IJE6YTboPFBBJoHn0Ld2eQRyjZbX25A_BYKu8bwCiroqBwz4kFPoT7gRfoAu8sjO8yh-2ciJBatSQC__ir3jafHsUKG63MHU4DJjoaHBYtQb3AawITEN_OaAzOMHIMvzVhY-hLGItCgIF2k2xatzF5H5oiqGjStMPdLC3FMdST1aovd26jYNDvw9K0wHkJjdIDV7I9lrQR11eHYZsqq-5Rk4U8XhqBFdjri3CaDNt26hlovvfEL7hzaJV-hPGXsJ8Jj"
          width={40}
          height={40}
        />
        <h1 className="text-xl font-bold text-white">geekhack</h1>
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
