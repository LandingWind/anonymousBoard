# anonymousBoard
匿名板 / 分享板


### Quick Start

```bash
# config your redis addr and auth at ./constant/redis.go

# build the app in linux
CGO_ENABLED=0 GOOS=linux go build -o anonymousBoard_linux
# or build the app in unix
go build -o anonymousBoard

# run in prod env
GIN_MODE=release ./anonymousBoard_linux
# or
GIN_MODE=release ./anonymousBoard
```
