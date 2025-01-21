# Go-Log Warehouse
## Go simple store log to warehouse

This library is made for golog migration purposes. The log will be stored to Redis, where they can be accessed by another sender worker.

## Installation

Install go depedency on your projects

```sh
go get github.com/idabgsram/golog-warehouse
```

## Set Evironment Variable

| Env Key | Note | Example |
| ------- | ----- | ----- |
| GOLOG_CHANNEL | Your Log Channel Name | #log-channel |
| GOLOG_USERNAME | Your Log Username | LogService |
| GOLOG_WAREHOUSE_KEY | Your Log Warehouse Key (optional) | golog_warehouse |
| GOLOG_DISTRIBUTION_MEDIA | Your Log Warehouse Distribution Media (optional) | discord,telegram |
| GOLOG_REDIS_URL | Your Redis URL | redis://user:password@localhost:6379/2?protocol=3 |
| GOLOG_REDIS_HOST | Your Redis Host (if Redis URL Not Defined) | localhost |
| GOLOG_REDIS_PASSWORD | Your Redis Password (if Redis URL Not Defined) | password |
| GOLOG_REDIS_PORT | Your Redis Port (if Redis URL Not Defined) | 6379 |

