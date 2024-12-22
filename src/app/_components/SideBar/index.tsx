import Link from "next/link";
import { db } from "~/server/db";

export const dynamic = "force-dynamic";

export default async function SideBar() {
  const categories = await db.query.categories.findMany({
    orderBy: (model, { asc }) => asc(model.id),
  });

  return (
    <aside className="bh-screen border-primary min-w-64 overflow-y-auto border-r p-4">
      <nav>
        <h2 className="mb-4 text-xl font-bold underline">Categories</h2>
        <ul className="space-y-2">
          {categories.map((category) => (
            <li key={category.id}>
              <Link
                href={`/${category.url}`}
                className={
                  "block rounded-md px-3 py-2 transition-colors duration-200 ease-in-out hover:bg-secondary"
                }
              >
                {category.name}
              </Link>
            </li>
          ))}
        </ul>
      </nav>
    </aside>
  );
}
