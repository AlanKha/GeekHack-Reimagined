import { db } from "~/server/db";

export const dynamic = "force-dynamic";

export default async function HomePage() {
  const posts = await db.query.posts.findMany({
    orderBy: (model, { asc }) => asc(model.id),
  });

  return (
    <main className="flex min-h-screen flex-col items-center justify-center bg-gradient-to-b from-[#E9492B] to-black text-white">
      <div className="">
        {posts.map((post, index) => (
          <div key={post.id + "-" + index} className="p-2">
            <a href={post.url}>
              <h2 className="text-xl">{post.title}</h2>
            </a>
          </div>
        ))}
      </div>
    </main>
  );
}
