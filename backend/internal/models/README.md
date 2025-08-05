# Database Models

This document outlines the GORM models used for the forum's database schema.

## Core Models

### `User` (`user.go`)

Represents a user account in the system.

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `Username` | `string` | Unique username for login and display. |
| `Email` | `string` | Unique email for login and notifications. |
| `Password` | `string` | Hashed password. |
| `AvatarURL` | `string` | URL to the user's avatar image. |
| `Signature` | `string` | User's forum signature, displayed below their posts. |
| `Role` | `string` | User role (e.g., 'user', 'moderator', 'admin'). Defaults to 'user'. |
| `LastSeen` | `time.Time` | Timestamp of the user's last activity. |
| `JoinedAt` | `time.Time` | Timestamp of user registration. |
| `IsActive` | `bool` | Whether the user account is active. Used for soft deletes/bans. |
| `PostCount` | `int` | Denormalized count of the user's total posts. |
| `ThreadCount` | `int` | Denormalized count of the user's total threads. |
| `ReputationScore` | `int` | Score based on post reactions, used for ranking. |

### `Category` (`category.go`)

Represents a forum category, which is a top-level container for threads.

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `Name` | `string` | The display name of the category. |
| `Description` | `string` | A short description of the category's topic. |
| `Slug` | `string` | A URL-friendly version of the name. |
| `ThreadCount` | `int` | Denormalized count of threads in the category. |
| `PostCount` | `int` | Denormalized count of posts in the category. |
| `LastActivity` | `time.Time` | Timestamp of the most recent post or thread. |
| `DisplayOrder` | `int` | Determines the order in which categories are listed. |
| `IsActive` | `bool` | Whether the category is visible to users. |

### `Thread` (`thread.go`)

Represents a discussion thread within a category.

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `Title` | `string` | The title of the thread. |
| `Content` | `string` | The original post content of the thread. |
| `UserID` | `uint` | Foreign key for the user who created the thread. |
| `CategoryID` | `uint` | Foreign key for the category this thread belongs to. |
| `Views` | `int` | The number of times the thread has been viewed. |
| `PostCount` | `int` | Denormalized count of total posts in the thread (includes OP). |
| `ReplyCount` | `int` | Denormalized count of replies (excludes OP). |
| `Locked` | `bool` | If true, users cannot reply to the thread. |
| `Pinned` | `bool` | If true, the thread is "stuck" to the top of the category list. |
| `IsActive` | `bool` | Whether the thread is active. Used for soft deletes. |
| `LastActivity` | `time.Time` | Timestamp of the most recent reply. |
| `LastPostID` | `uint` | Foreign key for the last post in the thread. |
| `LastPostUserID` | `uint` | Foreign key for the user who made the last post. |

### `Post` (`post.go`)

Represents a single post within a thread.

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `Content` | `string` | The text content of the post. |
| `UserID` | `uint` | Foreign key for the user who created the post. |
| `ThreadID` | `uint` | Foreign key for the thread this post belongs to. |
| `EditedAt` | `*time.Time` | Timestamp of the last edit. Null if not edited. |
| `EditedByID` | `*uint` | Foreign key for the user who last edited the post. |
| `IsDeleted` | `bool` | Whether the post is deleted. Used for soft deletes. |
| `PostNumber` | `int` | The sequential number of the post within its thread. |
| `ReactionCount` | `int` | Denormalized count of all reactions on the post. |
| `LikeCount` | `int` | Denormalized count of 'like' reactions on the post. |

## Supporting Models

### `Reaction` (`reaction.go`)

Represents a user's reaction to a post (e.g., like, love).

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `PostID` | `uint` | Foreign key for the post that was reacted to. |
| `UserID` | `uint` | Foreign key for the user who reacted. |
| `ReactionType` | `ReactionType` | The type of reaction (e.g., 'like', 'love', 'angry'). |

### `UserSession` (`user_session.go`)

Represents an active user session for authentication and tracking.

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `UserID` | `uint` | Foreign key for the user associated with the session. |
| `Token` | `string` | The session token (e.g., JWT). |
| `ExpiresAt` | `time.Time` | The expiration time of the session. |
| `IPAddress` | `string` | The IP address from which the session was initiated. |
| `UserAgent` | `string` | The user agent of the client. |
| `IsActive` | `bool` | Whether the session is currently active. |
| `LastActivity` | `time.Time` | Timestamp of the last activity within the session. |

### `Notification` (`notification.go`)

Represents a notification sent to a user.

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `UserID` | `uint` | Foreign key for the user receiving the notification. |
| `Type` | `NotificationType` | The type of notification (e.g., 'new_post', 'mention'). |
| `Title` | `string` | The title of the notification. |
| `Message` | `string` | The content of the notification. |
| `IsRead` | `bool` | Whether the user has read the notification. |
| `RelatedThreadID` | `*uint` | Optional foreign key to a related thread. |
| `RelatedPostID` | `*uint` | Optional foreign key to a related post. |
| `RelatedUserID` | `*uint` | Optional foreign key to a related user (e.g., for mentions). |

### `ThreadSubscription` (`notification.go`)

Represents a user's subscription to a thread to receive notifications for new posts.

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `UserID` | `uint` | Foreign key for the subscribing user. |
| `ThreadID` | `uint` | Foreign key for the subscribed-to thread. |
| `IsActive` | `bool` | Whether the subscription is active. |

### `ModerationLog` (`moderation_log.go`)

Records actions taken by moderators for accountability.

| Field | Type | Description |
| :--- | :--- | :--- |
| `ID` | `uint` | Primary key. |
| `ModeratorID` | `uint` | Foreign key for the moderator who performed the action. |
| `TargetUserID` | `*uint` | Optional foreign key for the user who was the subject of the action. |
| `Action` | `ModerationAction` | The type of action taken (e.g., 'edit', 'delete', 'ban'). |
| `Reason` | `string` | The reason provided by the moderator for the action. |
| `ThreadID` | `*uint` | Optional foreign key to a related thread. |
| `PostID` | `*uint` | Optional foreign key to a related post. |
| `CategoryID` | `*uint` | Optional foreign key to a related category. |
| `Metadata` | `string` | JSONB field for storing additional data (e.g., old values). |
| `ActionDate` | `time.Time` | Timestamp of when the moderation action occurred. |

## Helper Files

### `dto.go`

Contains Data Transfer Objects (DTOs) used to shape the data sent in API responses. This helps to avoid exposing the raw database models and allows for tailored, efficient data structures for the frontend.

### `methods.go`

Includes helper methods attached to the model structs. These methods provide convenient, reusable logic for common operations (e.g., `user.CanModerate()`, `thread.ToSummary()`), which helps to keep the handler logic cleaner and more focused.

### `database_optimizations.go`

This file is for documentation purposes and outlines recommended SQL indexes and constraints to ensure optimal database performance. These are not applied automatically by GORM and should be managed through database migrations.
