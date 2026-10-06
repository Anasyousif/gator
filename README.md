# Gator 🐊

Gator is a command-line RSS feed aggregator built in Go and backed by PostgreSQL. It allows you to manage feeds, automatically follow feeds you create, follow/unfollow other feeds, continuously scrape posts in the background, and browse stored posts directly from your terminal.

---

## Prerequisites

Before installing Gator, make sure you have the following installed on your system:

- **Go** (version 1.20 or later): [Install Go](https://go.dev/doc/install)
- **PostgreSQL**: [Install PostgreSQL](https://www.postgresql.org/download/)

Ensure PostgreSQL is running locally and that you have a database created for Gator (e.g., `gator`).

---

## Installation

You can install the `gator` CLI tool directly using `go install`:

```bash
go install [github.com/](https://github.com/)<your-github-username>/gator@latest

Note: Ensure $GOPATH/bin (typically ~/go/bin) is in your system's PATH environment variable so you can invoke gator directly.

Configuration Setup

Gator requires a JSON configuration file named .gatorconfig.json in your home directory (~/.gatorconfig.json).

Create ~/.gatorconfig.json with the following contents, updating db_url with your local PostgreSQL connection details:
JSON

{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}

Database Migrations

Run your database migrations before executing CLI commands. If you are using goose:
Bash

goose postgres "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable" up

Available Commands
User Management

    Register a new user:
    Bash

    gator register <username>

    Log in as an existing user:
    Bash

    gator login <username>

    List registered users:
    Bash

    gator users

Feed Management

    Add and automatically follow a feed:
    Bash

    gator addfeed "TechCrunch" "[https://techcrunch.com/feed/](https://techcrunch.com/feed/)"

    List all feeds in the database:
    Bash

    gator feeds

    Follow a feed by URL:
    Bash

    gator follow "[https://news.ycombinator.com/rss](https://news.ycombinator.com/rss)"

    List feeds followed by current user:
    Bash

    gator following

    Unfollow a feed:
    Bash

    gator unfollow "[https://news.ycombinator.com/rss](https://news.ycombinator.com/rss)"

Aggregation & Browsing

    Start the continuous background scraper:
    Bash

    gator agg 1m

    (Runs continuously until stopped with Ctrl+C. Keep this running in a separate terminal window.)

    Browse saved posts from followed feeds:
    Bash

    # View default (2) posts
    gator browse

    # View a custom number of posts
    gator browse 10