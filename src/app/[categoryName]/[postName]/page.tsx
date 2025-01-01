export default async function PostPage({
  params,
}: {
  params: { categoryName: string; postName: string };
}) {
  // eslint-disable-next-line @typescript-eslint/await-thenable
  const { categoryName, postName } = await params;
  return (
    <div>
      <h1>Category: {categoryName}</h1>
      <h1>Post: {postName}</h1>
    </div>
  );
}
