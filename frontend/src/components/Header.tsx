'use client';

import { useAuth } from '@/context/AuthContext';
import Link from 'next/link';

const Header = () => {
  const { user, logout } = useAuth();

  return (
    <header className="bg-surface border-b border-outline sticky top-0 z-10 px-4 py-2">
      <div className="flex justify-between items-center">
        <div className="flex-1 max-w-xl">
          <div className="relative">
            <span className="material-icons-outlined absolute left-3 top-1/2 -translate-y-1/2 text-on-surface-variant">
              search
            </span>
            <input
              className="bg-[#272729] border border-outline rounded-full pl-10 pr-4 py-2 w-full focus:outline-none focus:ring-2 focus:ring-primary/50 text-on-surface"
              placeholder="Search Geekhack"
              type="text"
            />
          </div>
        </div>
        <div className="flex items-center space-x-4">
          {user ? (
            <>
              <span>Welcome, {user.username}</span>
              <button
                onClick={logout}
                className="bg-primary hover:bg-primary/90 text-white font-bold py-1.5 px-6 rounded-full text-sm transition-colors"
              >
                Logout
              </button>
            </>
          ) : (
            <>
              <Link
                href="/auth/login"
                className="bg-transparent border border-primary text-primary font-bold py-1.5 px-6 rounded-full text-sm hover:bg-primary/10 transition-colors"
              >
                Login
              </Link>
              <Link
                href="/auth/register"
                className="bg-primary hover:bg-primary/90 text-white font-bold py-1.5 px-6 rounded-full text-sm transition-colors"
              >
                Sign Up
              </Link>
            </>
          )}
        </div>
      </div>
    </header>
  );
};

export default Header;
