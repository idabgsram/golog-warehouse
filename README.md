# Go-Log Warehouse
## Go simple store log to warehouse

This library is made for golog migration purposes. The log will be stored to Redis, where they can be accessed by another sender worker.

## Installation

Install go depedency on your projects

```sh
go get github.com/idabgsram/golog-warehouse
```

## Set Evironment Variable

| Env Key | Note |
| ------- | ----- |
| GOLOG_CHANNEL | Your Log Channel Name | 
| GOLOG_USERNAME | Your Log Username | 
| GOLOG_WAREHOUSE_KEY | Your Log Warehouse Key (optional) | 
| GOLOG_REDIS_URL | Your Redis URL | 
| GOLOG_REDIS_HOST | Your Redis Host (if Redis URL Not Defined) |
| GOLOG_REDIS_PASSWORD | Your Redis Password (if Redis URL Not Defined) |
| GOLOG_REDIS_PORT | Your Redis Port (if Redis URL Not Defined) |

