import { getPost } from "~/server/db/queries";
import Image from "next/image";

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
      {/* Post */}
      <div className="flex bg-content p-4 my-8">
        {/* Profile */}
        <div className="w-40">
          <p className="text-xl font-semibold text-primary">
            {post.creator.displayName}
          </p>
          <p className="text-xs font-semibold">Thread Starter</p>
          {post.creator.avatarUrl && (
            <div className="relative h-32 w-32">
              <Image
                src={post.creator.avatarUrl}
                fill
                alt="Profile"
                className="object-cover"
              />
            </div>
          )}
        </div>
        {/* Content */}
        <div className="grow">blue</div>
      </div>

      {/* Comments */}
      {post.comments.map((comment) => (
        <div className="flex flex-col bg-content p-4" key={comment.id}>
          <p className="font-semibold">{comment.user.displayName}:</p>
          <p className="pl-4">{comment.content}</p>
        </div>
      ))}
    </div>
  );
}
