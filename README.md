Gator
Gator is a CLI RSS feed aggregator built with Go and PostgreSQL. It lets you register users, subscribe to RSS feeds, run a background aggregator that continuously fetches new posts, and browse the latest content — all from your terminal.
Prerequisites
You'll need both of the following installed before getting started:
Go (1.22 or later) — https://go.dev/doc/install
PostgreSQL (16 or later) — https://www.postgresql.org/download/
Make sure `psql` is accessible from your terminal and that you have a running Postgres instance with a database created for Gator:
```sql
CREATE DATABASE gator;
```
Installation
Install the `gator` CLI directly with `go install`:
```bash
go install github.com/bootdotdev/gator@latest
```
Make sure your Go binary path is in your `PATH`. If the command isn't found after installing, add this to your shell config (`.bashrc`, `.zshrc`, etc.):
```bash
export PATH=$PATH:$(go env GOPATH)/bin
```
Configuration
Gator reads its config from a file called `.gatorconfig.json` in your home directory. Create it manually:
```bash
touch ~/.gatorconfig.json
```
Then add the following content, replacing the connection string as needed:
```json
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": ""
}
```
The `db_url` follows the standard PostgreSQL connection string format: `postgres://<user>:<password>@<host>:<port>/<dbname>?sslmode=disable`. Update the username, password, and database name to match your local Postgres setup.
Usage
Managing Users
Register a new user (this also sets them as the current user):
```bash
gator register <username>
```
Switch to an existing user:
```bash
gator login <username>
```
List all registered users (the current user is marked):
```bash
gator users
```
Managing Feeds
Add a new RSS feed and automatically follow it as the current user:
```bash
gator addfeed <feed-name> <url>
```
Example:
```bash
gator addfeed "The Verge" https://www.theverge.com/rss/index.xml
```
List all feeds in the database:
```bash
gator feeds
```
Follow an existing feed (by URL):
```bash
gator follow <url>
```
Unfollow a feed:
```bash
gator unfollow <url>
```
List all feeds the current user is following:
```bash
gator following
```
Aggregating Posts
Start the background aggregator, which fetches new posts from all feeds on a set interval:
```bash
gator agg <interval>
```
The interval accepts Go duration strings like `30s`, `5m`, `1h`, `2.5h`, etc. Keep this running in a terminal while you use other commands in a separate session.
Browsing Posts
Browse the most recently aggregated posts for the feeds you follow:
```bash
gator browse
```
You can optionally pass a limit to control how many posts are shown:
```bash
gator browse 10
```
Resetting the Database
To wipe all users and data (useful during development):
```bash
gator reset
```
