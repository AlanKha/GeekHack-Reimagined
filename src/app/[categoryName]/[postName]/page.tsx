import { getPost } from "~/server/db/queries";

export default async function PostPage({
  params,
}: {
  params: { categoryName: string; postName: string };
}) {
  // eslint-disable-next-line @typescript-eslint/await-thenable
  const { categoryName, postName } = await params;
  const post = await getPost(categoryName, postName);
  
  return (
    <div>
      <h1>Category: {post.categoryId}</h1>
      <h1>Post: {post.title}</h1>
      <p>Content: {post.content}</p>
      <p>Comments: {post.commentCount}</p>
      {post.comments.map((comment) => (
        <div className="flex" key={comment.id}>
          <p>{comment.content}</p>
          <p>By: {comment.userId}</p>
        </div>
      ))}
    </div>
  );
}
