import { getPostsFromCategory } from "~/server/db/queries";
import { notFound } from "next/navigation";
import Link from "next/link";

export default async function CategoryPage({
  params,
}: {
  params: { categoryName: string };
}) {
  try {
    // eslint-disable-next-line @typescript-eslint/await-thenable
    const { categoryName } = await params;
    const posts = await getPostsFromCategory(categoryName);

    return (
      <div className="p-4">
        <h1 className="mb-6 text-2xl font-bold">Posts</h1>
        {posts.length > 0 ? (
          <div className="overflow-x-auto rounded-lg">
            <table className="w-full border-collapse bg-content">
              <thead className="border-b border-gray-300">
                <tr>
                  <th className="border-r px-6 py-3 text-left text-sm font-semibold">
                    Title
                  </th>
                  <th className="border-r px-6 py-3 text-left text-sm font-semibold">
                    Description
                  </th>
                  <th className="border-r px-6 py-3 text-left text-sm font-semibold">
                    Replies
                  </th>
                  <th className="px-6 py-3 text-left text-sm font-semibold">
                    Last Updated
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {posts.map((post) => (
                  <tr key={post.id} className="">
                    <Link href={`/${categoryName}/${post.url}`}>
                      <td className="whitespace-nowrap border-r px-6 py-4 text-sm">
                        {post.title}
                      </td>
                      <td className="border-r px-6 py-4 text-sm">
                        {post.content}
                      </td>
                      <td className="whitespace-nowrap border-r px-6 py-4 text-sm">
                        {post.commentCount}
                      </td>
                      <td className="whitespace-nowrap px-6 py-4 text-sm">
                        {post.updatedAt.toLocaleString()}
                      </td>
                    </Link>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        ) : (
          <p className="py-4 text-center text-gray-500">No posts found</p>
        )}
      </div>
    );
  } catch (error) {
    if (error instanceof Error && error.message.includes("not found")) {
      notFound();
    }
    console.error(error);
    return <div className="p-4 text-red-500">Error loading the page</div>;
  }
}
