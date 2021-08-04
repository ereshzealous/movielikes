# Movie Likes
## Setup local development

- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    brew install golang-migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    brew install sqlc
    ```

### Setup local infra
I have used local postgres set up not docker. But I will try to update that as well in coming days.
You have to create database schema by name "movieDB". It's up to you can change the name as well. But do change all other references too.

- Run db migration down all versions:

    ```bash
    make migratedown
    ```

### How to generate code

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Create a new db migration:

    ```bash
    migrate create -ext sql -dir db/migration -seq <migration_name>
    ```

- Run test cases

    ```bash
    make test
    ```
