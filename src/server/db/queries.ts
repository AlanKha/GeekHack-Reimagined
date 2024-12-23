import { db } from "~/server/db";

export async function getPostsFromCategory(categoryUrl: string) {
  const category = await db.query.categories.findFirst({
    where: (category, { eq }) => eq(category.url, categoryUrl),
  });

  if (!category) {
    throw new Error(`Category "${categoryUrl}" not found.`);
  }

  const posts = await db.query.posts.findMany({
    where: (post, { eq }) => eq(post.categoryId, category.id),
    orderBy: (post, { desc }) => desc(post.createdAt),
  });

  return posts;
}
