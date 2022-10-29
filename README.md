# mongodb-migration-tool

Uses the [Golang Migrate](https://github.com/golang-migrate/migrate/) package within a binary for easy migrations to MongoDB.

## How to use

1. You can use the Makefile to first `build` the binary into a _dist_ folder. 
2. run the binary using three required arguments; __command__, __mongodb connection string__, and __file path to a folder containing migrations__

### Examples

Example, This runs all __up__ migrations in the folder _./example/migration_ with the extensions, `<NAME>.up.json`,
>dist/mongodb_migration up "mongodb://localhost:27017/my_db" file://./example/migration

Example, This runs all __down__ migrations in the folder _./example/migration_ with the extensions, `<NAME>.down.json`,
>dist/mongodb_migration down "mongodb://localhost:27017/my_db" file://./example/migration

Example, Connection with authentication,
>dist/mongodb_migration up "mongodb://user:pass@localhost:27017/my_db?authSource=admin" file://./example/migration

Connection strings will vary so it is recommended you use a connection string that has been used before.

## How does it work

The first run will create two new collections in your database, _schema_migrations_ & _migrate_advisory_lock_. These are used to signal what state the migrations are in.

If you create something in your __up__ always destroy it in your __down__ scripts this will always control the state of your database.

__Don't drop anything created from your migrations__, if you do then the state will think it still exists and leave your migration in an inconsistent state.

## How to create migrations

You can look at the example folder in this repo but looking at Golang-Migrate's is the best source.

- Writing migrations,
https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md

- More examples,
  https://github.com/golang-migrate/migrate/tree/master/database/mongodb/examples/migrations
