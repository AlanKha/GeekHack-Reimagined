import { getPostsFromCategory } from "~/server/db/queries";
import { notFound } from "next/navigation";
import Link from "next/link";
import Image from "next/image";

const keepPath = "/keep.svg";

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
              <thead className="border-b border-gray-300 text-secondary">
                <tr>
                  <th className="w-3/5 border-r px-6 py-3 text-left text-sm font-semibold">
                    Subject / Started by
                  </th>
                  <th className="w-1/5 border-r px-6 py-3 text-left text-sm font-semibold">
                    Replies / Views
                  </th>
                  <th className="w-1/5 px-6 py-3 text-left text-sm font-semibold">
                    Last Post
                  </th>
                </tr>
              </thead>
              <tbody className="divide-y divide-gray-200">
                {posts.map((post) => (
                  <tr
                    key={post.id}
                    className={`group ${post.isSticky ? "bg-sticky" : ""}`}
                  >
                    <td className="w-3/5 whitespace-nowrap border-r px-6 py-4 text-sm">
                      <div className="flex justify-between">
                        <div className="flex flex-col">
                          <Link
                            href={`/${categoryName}/${post.url}`}
                            className="block text-primary hover:underline"
                          >
                            {post.title}
                          </Link>
                          <p>
                            Started by{" "}
                            {post.creator.displayName ?? post.creator.username}
                          </p>
                        </div>
                        {post.isSticky && (
                          <div className="flex-shrink-0 rotate-45 pl-2">
                            <Image
                              src={keepPath}
                              alt="Sticky Icon"
                              width={25}
                              height={0}
                            />
                          </div>
                        )}
                      </div>
                    </td>
                    <td className="w-1/5 whitespace-nowrap border-r px-6 py-4 text-sm">
                      <div className="flex flex-col">
                        <div>{post.commentCount} Replies</div>
                        <div>{post.viewCount} Views</div>
                      </div>
                    </td>
                    <td className="w-1/5 whitespace-nowrap px-6 py-4 text-sm">
                      <div className="flex flex-col">
                        <div>{post.updatedAt.toLocaleString()}</div>
                        <div>
                          by {post.updater.displayName ?? post.updater.username}
                        </div>
                      </div>
                    </td>
                  </tr>
                ))}
              </tbody>
            </table>
          </div>
        ) : (
          <p className="py-4 text-center text-xl text-primary underline">
            No posts found
          </p>
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
