import { db } from "~/server/db";
import { desc } from "drizzle-orm";
import { posts } from "~/server/db/schema"; // Make sure to import your schema

export async function getPostsFromCategory(categoryUrl: string) {
  const category = await db.query.categories.findFirst({
    where: (category, { eq }) => eq(category.url, categoryUrl),
  });

  if (!category) {
    throw new Error(`Category "${categoryUrl}" not found.`);
  }

  const result = await db.query.posts.findMany({
    where: (post, { eq }) => eq(post.categoryId, category.id),
    with: {
      creator: {
        columns: {
          username: true,
          displayName: true,
        },
      },
      updater: {
        columns: {
          username: true,
          displayName: true,
        },
      },
    },
    orderBy: [desc(posts.isSticky), desc(posts.updatedAt)],
  });

  return result;
}
