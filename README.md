# boilerplate

# Installation

on this project, you need to install tool on your local :
- Git
- Docker
- Docker Compose
- Postgres (if you need to use local database)

### Architecture Structure
```
- boilerplate
 |- api
   |- versioning
    |- controllers
    |- middleware
    |- objects
    |- responses
    |- services
  |- database
  |- utils
 |- main.go
```

# How to Setup Local

1. Clone repository using command : 
```bash
git clone git@github.com:zhuangalbert/app-go.git
```

2. enter the directory
```bash
cd app-go
```

3. Create file .env, by copying from .env.example, don't forget to modify .env value 
```bash
cp .env.example .env
```
my .env :
```bash
API_SECRET=go-beta
API_PORT=8099

DB_HOST=docker.for.mac.localhost
DB_DRIVER=postgres 
DB_USER=ab
DB_PASSWORD=abganteng
DB_NAME=db
DB_PORT=5432
```

you can use the mysql database on the local machine with the following configuration : 
```bash
DB_HOST=127.0.0.1
DB_PORT=3306
DB_DATABASE=database_name
DB_USERNAME=database_username
DB_PASSWORD=database_password
```

4. run docker-compose with command : 
```bash
docker-compose up --build
```
this command take a long time.
If you get a screen like the one below, the setup is successful and you are ready to develop your project:

```bash
app_golang   | Connection is created
app_golang   | [GIN-debug] POST   /v1/users/login           --> github.com/zhuangalbert/boilerplate/src/api/v1/controllers.(*UserController).Login-fm (3 handlers)
app_golang   | [GIN-debug] POST   /v1/users/register        --> github.com/zhuangalbert/boilerplate/src/api/v1/controllers.(*UserController).Register-fm (3 handlers)
app_golang   | [GIN-debug] GET    /v1/users/:id             --> github.com/zhuangalbert/boilerplate/src/api/v1/controllers.(*UserController).GetUser-fm (4 handlers)
app_golang   | [GIN-debug] GET    /v1/users/:id/update      --> github.com/zhuangalbert/boilerplate/src/api/v1/controllers.(*UserController).Update-fm (4 handlers)
app_golang   | :8080
app_golang   | [GIN-debug] Listening and serving HTTP on :8080
```
