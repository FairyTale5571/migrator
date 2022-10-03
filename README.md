# Migrate Extension for Arma

This extension allow migrate your database updates automatically.

## Installation

Set environment variables for logs into discord webhook
`LOG_SNOWFLAKE` and `LOG_HOOK`

Set environment variables for path to migrates files `MIGRATIONS_PATH`

Setup grc_config.json

```json
{
  "ip": "ip_database",
  "port": "port_database",
  "database": "name_database",
  "user": "user_database",
  "password": "password_database"
}
```

Place grc_config.json to `@extensions` folder

## Usage
Execute in arma script
```sqf
"migrate" callExtension "migrate";
```

## Build

```bash
make build
```
