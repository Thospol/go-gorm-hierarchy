# go-gorm-hierarchy

<img align="center" width="180px" src="https://images-wixmp-ed30a86b8c4ca887773594c2.wixmp.com/f/c7d894cb-8d37-4495-a454-89c868b12375/dcycwca-813a3b2d-1eae-4f6a-beab-27f1264b364b.png?token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJzdWIiOiJ1cm46YXBwOjdlMGQxODg5ODIyNjQzNzNhNWYwZDQxNWVhMGQyNmUwIiwiaXNzIjoidXJuOmFwcDo3ZTBkMTg4OTgyMjY0MzczYTVmMGQ0MTVlYTBkMjZlMCIsIm9iaiI6W1t7InBhdGgiOiJcL2ZcL2M3ZDg5NGNiLThkMzctNDQ5NS1hNDU0LTg5Yzg2OGIxMjM3NVwvZGN5Y3djYS04MTNhM2IyZC0xZWFlLTRmNmEtYmVhYi0yN2YxMjY0YjM2NGIucG5nIn1dXSwiYXVkIjpbInVybjpzZXJ2aWNlOmZpbGUuZG93bmxvYWQiXX0.KijY-p4GWjczqKcWqY3xgRmvPgK8SUgbHDdHsDIQvYc">

## Installation PostgreSQL using docker-compose
```bash
docker-compose up
```

## Installation for Migration

To install the library and command line program, use the following:

```bash
go get -v github.com/rubenv/sql-migrate/...
```

## Usage

```
$ sql-migrate --help
usage: sql-migrate [--version] [--help] <command> [<args>]

Available commands are:
    down      Undo a database migration
    new       Create a new migration
    redo      Reapply the last migration
    skip      Sets the database level to the most recent version available, without running the migrations
    status    Show migration status
    up        Migrates the database to the most recent version available
```

The `table` save logs version is `migrations`.

The environment that will be used can be specified with the `-env` flag (defaults to `development`).

Use the `--help` flag in combination with any of the commands to get an overview of its usage

## Example Migration

### create migrate

run this command `sql-migrate new {name}`

### up migrate

- run migration (development) run this command `sql-migrate up`
- run migration (production) run this command `sql-migrate up -env=production`

## Example DB Docs

Prerequisite: Make sure NodeJS and NPM have been installed on your computer before the installation.

1. Install dbdocs via terminal

```bash
npm install -g dbdocs
```

3. Login to dbdocs

```bash
dbdocs login
```

3. Generate dbdocs

```bash
dbdocs build doc/db.dbml
```

4. Convert a DBML file to SQL

```bash
npm install -g @dbml/cli
dbml2sql -o ./db/docs/db.sql ./db/docs/db.dbml
```
