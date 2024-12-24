import { sql } from "drizzle-orm";
import {
  index,
  integer,
  pgTableCreator,
  timestamp,
  varchar,
  text,
  boolean,
} from "drizzle-orm/pg-core";
import { relations } from "drizzle-orm";

export const createTable = pgTableCreator((name) => `geekhack_${name}`);

export const categories = createTable("category", {
  id: integer("id").primaryKey().generatedByDefaultAsIdentity(),
  name: varchar("name", { length: 256 }).notNull().unique(),
  url: varchar("url", { length: 1024 }).notNull(),
  description: text("description"),
  postCount: integer("post_count").default(0).notNull(),
  createdAt: timestamp("created_at", { withTimezone: true })
    .default(sql`CURRENT_TIMESTAMP`)
    .notNull(),
  updatedAt: timestamp("updated_at", { withTimezone: true })
    .$onUpdate(() => new Date())
    .default(sql`CURRENT_TIMESTAMP`)
    .notNull(),
});

export const posts = createTable(
  "post",
  {
    id: integer("id").primaryKey().generatedByDefaultAsIdentity(),
    title: varchar("title", { length: 256 }).notNull(),
    description: text("description"),
    content: text("content").notNull(),
    categoryId: integer("category_id")
      .references(() => categories.id)
      .notNull(),
    userId: integer("user_id")
      .references(() => users.id)
      .notNull(), // Link to user who created the post
    url: varchar("url", { length: 1024 }).notNull(),
    commentCount: integer("comment_count").default(0).notNull(),
    isSticky: boolean("is_sticky").default(false),
    isClosed: boolean("is_closed").default(false),
    createdAt: timestamp("created_at", { withTimezone: true })
      .default(sql`CURRENT_TIMESTAMP`)
      .notNull(),
    updatedAt: timestamp("updated_at", { withTimezone: true })
      .$onUpdate(() => new Date())
      .default(sql`CURRENT_TIMESTAMP`)
      .notNull(),
  },
  (posts) => ({
    nameIndex: index("name_idx").on(posts.title),
    userIndex: index("user_idx").on(posts.userId),
  }),
);

export const users = createTable("user", {
  id: integer("id").primaryKey().generatedByDefaultAsIdentity(),
  username: varchar("username", { length: 256 }).notNull().unique(),
  userId: varchar("user_id", { length: 256 }).notNull().unique(), // Clerk UID
  email: varchar("email", { length: 256 }).unique(),
  displayName: varchar("display_name", { length: 256 }),
  bio: text("bio"),
  avatarUrl: varchar("avatar_url", { length: 1024 }),
  isAdmin: boolean("is_admin").default(false),
  isBanned: boolean("is_banned").default(false),
  createdAt: timestamp("created_at", { withTimezone: true })
    .default(sql`CURRENT_TIMESTAMP`)
    .notNull(),
  updatedAt: timestamp("updated_at", { withTimezone: true })
    .$onUpdate(() => new Date())
    .default(sql`CURRENT_TIMESTAMP`)
    .notNull(),
});

export const comments = createTable(
  "comment",
  {
    id: integer("id").primaryKey().generatedByDefaultAsIdentity(),
    content: text("content").notNull(),
    postId: integer("post_id")
      .references(() => posts.id)
      .notNull(),
    userId: integer("user_id")
      .references(() => users.id)
      .notNull(),
    parentCommentId: integer("parent_comment_id"),
    createdAt: timestamp("created_at", { withTimezone: true })
      .default(sql`CURRENT_TIMESTAMP`)
      .notNull(),
    updatedAt: timestamp("updated_at", { withTimezone: true })
      .$onUpdate(() => new Date())
      .default(sql`CURRENT_TIMESTAMP`)
      .notNull(),
  },
  (comments) => ({
    postIndex: index("comment_post_idx").on(comments.postId),
    userIndex: index("comment_user_idx").on(comments.userId),
  }),
);

export const commentRelations = relations(comments, ({ one }) => ({
  child: one(comments, {
    fields: [comments.parentCommentId],
    references: [comments.id],
  }),
}));
