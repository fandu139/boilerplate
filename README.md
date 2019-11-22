# BOILERPLATE

## Installing
This support For go version 1.13 

### Local Development

Fork this repo for your repo then clone in your local
```
git clone https://github.com/sofyan48/BOILERGOLANG.git
```

Get Project Moduls

```
go get github.com/sofyan48/BOILERGOLANG
```

#### Environment Setup
For Development Mode Setup dotenv
```
cp .env.example .env
```
Setting up your local configuration see example
```
SERVER_ADDRESS=0.0.0.0
SERVER_PORT=3000
SERVER_TIMEZONE=Asia/Jakarta

DB_MYSQL_USERNAME=root
DB_MYSQL_PASSWORD=password
DB_MYSQL_HOST=localhost
DB_MYSQL_PORT=3306
DB_MYSQL_DATABASE=db
```

After environment setting then run your server

```
go run src/main.go
```

for building
```
go build src/main.go
```
#### Live Reload
To activate Live Reload install air 
##### on macOS

```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/darwin/air
chmod +x /usr/local/bin/air
```

##### on Linux

```
curl -fLo /usr/local/bin/air \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/linux/air
chmod +x /usr/local/bin/air
```

##### on Windows

```
curl -fLo ~/air.exe \
    https://raw.githubusercontent.com/cosmtrek/air/master/bin/windows/air.exe
```

see watcher.conf setting watcher file for air config now

##### Starting Live Reload
go to your project path
```
air -c watcher.conf
```

### Production Mode

#### Dockerizing
Building Image
```
docker build -t bigevent_api
```
Edit Environment In docker-compose.yml then Run Compose
```
docker-compose up
```
Stop Container
```
docker-compose stop
```
Remove your container
```
docker-compose rm -f
```

## Database Migration
### Golang Migrate
Documentation Mode 
[Release Downloads](https://github.com/golang-migrate/migrate/blob/master/cmd/migrate/README.md)

#### Installing
##### MAC
```
brew install golang-migrate
```

##### Linux And Windows
```
curl -L https://github.com/golang-migrate/migrate/releases/download/$version/migrate.$platform-amd64.tar.gz | tar xvz
```
### Migrating Database

```
migrate -path path_migration/ -database 'mysql://root:root@tcp(localhost:3306)/bigevent' up
```
in this boilerplate migration path : src/migration/mysql


### Setup Swagger Docs
See Documentation 
[Swag Docs](https://github.com/swaggo/swag)
