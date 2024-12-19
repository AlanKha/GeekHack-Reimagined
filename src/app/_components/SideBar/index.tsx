import Link from "next/link";
import { db } from "~/server/db";

export const dynamic = "force-dynamic";

export default async function SideBar() {
  const categories = await db.query.categories.findMany({
    orderBy: (model, { asc }) => asc(model.id),
  });

  return (
    <aside className="h-screen w-64 overflow-y-auto border-r p-4">
      <nav>
        <h2 className="mb-4 text-xl font-bold">Categories</h2>
        <ul className="space-y-2">
          {categories.map((category) => (
            <li key={category.id}>
              <Link
                href={`/${category.name.toLowerCase()}`}
                className={"hover:bg-gray-500 block rounded-md px-3 py-2 transition-colors duration-200 ease-in-out"}
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
