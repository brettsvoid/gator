# Gator 🐊

An RSS feed aggregator CLI written in Go. Collect posts from across the internet, store them in Postgres, and view them right in your terminal.

---

## 🚀 Getting Started

### Prerequisites

- [Go](https://go.dev/doc/install) 1.21+
- [PostgreSQL](https://www.postgresql.org/download/)

### Install

```bash
go install github.com/brettsvoid/gator@latest
```

### First Run

1. Create a Postgres database (e.g. `gator`).
2. Set up your config file:

   ```bash
   echo '{"db_url":"postgres://user:password@localhost:5432/gator?sslmode=disable"}' >~/.gatorconfig.json
   ```

   This stores your DB connection string and user info.

3. Add a user:

   ```bash
   gator register john
   ```

4. Add a feed:

   ```bash
   gator addfeed https://techcrunch.com/feed/
   ```

5. View your posts:

   ```bash
   gator following
   ```

---

## 🔧 How-To Guides

- **Add multiple feeds and view them together**

  ```bash
  gator addfeed https://blog1.com/rss
  gator addfeed https://blog2.com/rss
  gator feeds
  ```

- **Follow a feed**

  ```bash
  gator feeds
  gator follow <feed_url>
  ```

- **Unfollow a feed**

  ```bash
  gator unfollow <feed_id>
  ```

  Stops tracking posts from that feed.

- **Browse posts**

  ```bash
  gator browse --limit 5 --sort asc --filter go
  ```

  Stops tracking posts from that feed.

- **Agg and refetch feeds on a 10 minute interval**

  ```bash
  gator agg 10m0s
  ```

---

## 📖 Reference

### CLI Commands

- `gator addfeed <url>` – add a feed
- `gator agg <time>` – view aggregated posts
- `gator browse` – view posts
- `gator login <user>` – login with a user
- `gator feeds` – view feeds
- `gator follow <feed_url>` – follow a new feed
- `gator following` – view followed feeds
- `gator register <user>` – add a new user
- `gator users` – view users
- `gator unfollow <feed_id>` – unfollow a feed

### Config File

Stored at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://user:password@localhost:5432/gator",
  "current_user_name": "your-username"
}
```

---

## 💡 Explanation

- **Why RSS?**
  RSS is an open standard for following updates from blogs, news, and podcasts.

- **Why Postgres?**
  It provides persistence and lets multiple users share feeds.

- **How it works**
  - CLI parses commands.
  - Feeds are fetched and stored in Postgres.
  - Posts can be listed, followed, and filtered.

---

## 📎 Links

- Project repo: [https://github.com/brettvoid/gator](https://github.com/brettsvoid/gator)
