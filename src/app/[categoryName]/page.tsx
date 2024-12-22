import { getPostsFromCategory, getCategoryObj } from "~/server/db/queries";
import { notFound } from "next/navigation";

export default async function CategoryPage({
  params,
}: {
  params: { categoryName: string };
}) {
  try {
    // eslint-disable-next-line @typescript-eslint/await-thenable
    const { categoryName } = await params;

    const categoryId = await getCategoryObj(categoryName);

    const posts = await getPostsFromCategory(categoryId);

    return (
      <div>
        <h1 className="underline">Posts</h1>
        {posts.map((post) => (
          <div key={post.id} className="flex">
            <h2 className="border p-2">{post.title}</h2>
            <p className="border p-2">{post.content}</p>
          </div>
        ))}
        {posts.length === 0 && <p>No posts found</p>}
      </div>
    );
  } catch (error) {
    if (error instanceof Error && error.message.includes("not found")) {
      notFound();
    }

    console.error(error);

    return <div>Error loading the page</div>;
  }
}
