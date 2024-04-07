## The simplest API architecture example 👨🏻‍🏫
This repository presents the simplest API architecture for your project.
## Clone the project

```
$ git clone https://github.com/kurerid/api-architecture
```

## Example 👨🏻‍💻
This repository describes default CRUD application with "TODO" entity.

### DB schema
```postgresql
create table "Todo"
(
    "id"   serial primary key,
    "note" varchar(250)
)
```

### Config 📁
The config is located in `/configs` directory.
Without correct config server will not start.
```json
{
  "host": "localhost",
  "port": "5432",
  "username": "postgres",
  "password": "1234",
  "DBName": "todo",
  "SSLMode": "disable"
}
```

### Server
HTTP server config is hardcoded. You can customize it.