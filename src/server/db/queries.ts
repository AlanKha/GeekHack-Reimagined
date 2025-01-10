import { db } from "~/server/db";
import { desc } from "drizzle-orm";
import { posts } from "~/server/db/schema";

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

export async function getPost(categoryUrl: string, postUrl: string) {
  const category = await db.query.categories.findFirst({
    where: (category, { eq }) => eq(category.url, categoryUrl),
  });

  if (!category) {
    throw new Error(`Category "${categoryUrl}" not found.`);
  }

  const post = await db.query.posts.findFirst({
    where: (post, { and, eq }) => and(eq(post.categoryId, category.id), eq(post.url, postUrl)),
    with: {
      creator: {
        columns: {
          username: true,
          displayName: true,
          avatarUrl: true,
        },
      },
      updater: {
        columns: {
          username: true,
          displayName: true,
        },
      },
      comments: {
        with: {
          user: {
            columns: {
              username: true,
              displayName: true,
              avatarUrl: true,
            },
          },
          child: {
            with: {
              user: {
                columns: {
                  username: true,
                  displayName: true,
                  avatarUrl: true,
                }
              }
            }
          }
        },
        orderBy: (comments, { desc }) => [desc(comments.createdAt)]
      }
    }
  });

  if (!post) {
    throw new Error(`Post "${postUrl}" not found.`);
  }

  return post;  
}