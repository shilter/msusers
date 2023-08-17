
# ms users

1. create database
```bash
    CREATE DATABASE <database name>;
    CREATE USER <username db> WITH PASSWORD '<your password>';
    ALTER ROLE <username db> SET client_encoding TO 'utf8';
    ALTER ROLE <username db> SET default_transaction_isolation TO 'read committed';
    ALTER ROLE <username db> SET timezone TO 'UTC';
    GRANT ALL PRIVILEGES ON DATABASE <database name> TO <username db>;
```
2. change .Env to your credential Env


## Before Run

export go 

```bash
   export GOROOT=/usr/local/go
   export GOPATH=$HOME/go-projects
   export PATH=$GOPATH/bin:$GOROOT/bin:$PATH
```

after that please run 
```bash
    go mod tidy
    go build
    go run main.go
```
or using air 
```bash 
    go mod tidy
    go build
    air server
```

## Color Reference

| Color             | Hex                                                                |
| ----------------- | ------------------------------------------------------------------ |
| Example Color | ![#0a192f](https://via.placeholder.com/10/0a192f?text=+) #0a192f |
| Example Color | ![#f8f8f8](https://via.placeholder.com/10/f8f8f8?text=+) #f8f8f8 |
| Example Color | ![#00b48a](https://via.placeholder.com/10/00b48a?text=+) #00b48a |
| Example Color | ![#00d1a0](https://via.placeholder.com/10/00b48a?text=+) #00d1a0 |

