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