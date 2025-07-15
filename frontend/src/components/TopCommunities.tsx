import Link from "next/link";

const TopCommunities = () => {
  return (
    <div className="bg-surface border border-outline rounded-lg">
      <div className="p-4 border-b border-outline">
        <h3 className="font-bold text-on-surface">Top Communities</h3>
      </div>
      <div className="p-4 space-y-4">
        <div className="flex items-center space-x-4">
          <span className="font-bold">1</span>
          <span className="material-icons-outlined text-red-400">keyboard</span>
          <Link className="text-sm font-semibold hover:underline flex-1" href="#">g/Keyboards</Link>
        </div>
        <div className="flex items-center space-x-4">
          <span className="font-bold">2</span>
          <span className="material-icons-outlined text-blue-400">view_module</span>
          <Link className="text-sm font-semibold hover:underline flex-1" href="#">g/KeyboardKeycaps</Link>
        </div>
        <div className="flex items-center space-x-4">
          <span className="font-bold">3</span>
          <span className="material-icons-outlined text-purple-400">groups</span>
          <Link className="text-sm font-semibold hover:underline flex-1" href="#">g/Meetups</Link>
        </div>
        <div className="flex items-center space-x-4">
          <span className="font-bold">4</span>
          <span className="material-icons-outlined text-green-400">person_add</span>
          <Link className="text-sm font-semibold hover:underline flex-1" href="#">g/NewMembers</Link>
        </div>
      </div>
      <div className="p-4">
        <button className="w-full bg-primary hover:bg-primary/90 text-white font-bold py-2 px-4 rounded-full text-sm">View All</button>
      </div>
    </div>
  );
};

export default TopCommunities;