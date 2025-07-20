# go-gator

A simple RSS aggregator Go program with PostgreSQL database.

## Requirement

- Go 1.24.5
- PostgreSQL 17

## Build

```bash
clone https://github.com/clementine-tw/go-gator.git
cd go-gator
go build
```

## Install

```bash
clone https://github.com/clementine-tw/go-gator.git
cd go-gator
go install
```

## Setup

1. Create a config file named `.gatorconfig.json` in your home directory
 with the following content
    ```json
    {
        "db_url": ""
    }
    ```

2. Change the `db_url` to the link of your PostgreSQL database

## Usage

You need to register a user at first time.

- Register a user
    ```bash
    go-gator register <user_name>
    ```

- Login a user
    ```bash
    go-gator login <user_name>
    ```

- List users
    ```bash
    go-gator users
    ```

- Add feed
    ```bash
    go-gator addfeed <feed_name> <feed_url>
    ```

- List feeds
    ```bash
    go-gator feeds
    ```

- Follow a feed
    ```bash
    go-gator follow <feed_url>
    ```

- Unfollow a feed
    ```bash
    go-gator unfollow <feed_url>
    ```

- Browse posts
    ```bash
    go-gator browse [limit]
    # limit is the number of posts you want to list
    ```

- Start aggregator feeds
    ```bash
    go-gator agg <time_interval>
    # time_interval is the time between requesting feed info
    # eg., 10s, 1m and 1h
    ```
