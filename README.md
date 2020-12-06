# Stateful rest sample
![GitHub](https://img.shields.io/github/license/jakobhostnik/jakobhostnik.github.io.svg?i)![Maintenance](https://img.shields.io/maintenance/yes/2023.svg)

Sample of rest application written in go using Postgres (or CrateDB) for persistence.
Mainly used for testing kubernetes setup.
This application is really simple and is meant to be used only for local testing!
Application was written with no securiy in mind.

# How does it work?
### There are two api routes:
```GET /users```
which returns JSON list of users from database table "users"

```POST /users```
as a body user should be sent
which adds user in JSON to database table "users"

### Example:
```bash
curl -X POST localhost:8080/users --data '{"name": "John", "lastname": "Doe"}'
curl localhost:8080/users
```

## Environment variables to set
##### PORT
Set on what port rest api will be listening.
##### DATABASE_URL
URL to your database.

### Setup database
You can setup your database in quite a few ways. So here is a simple example how to setup Postgres in Docker. 
```bash
docker run 
  -e POSTGRES_PASSWORD=password 
  -e POSTGRES_USER=user 
  -e POSTGRES_DB=testdb 
  -p 5432:5432 
  -d postgres
```
Then you have to manually create table ```users```
```sql
CREATE TABLE users ( name TEXT, lastname TEXT);
```
If you are using this example you should check your ip of ```docker0``` interface.
Mine was ```172.17.0.1```. So mine database url was:
```postgres://user:password@localhost:5432/testdb```

### Run using docker
```
docker run 
  -e PORT=8080 
  -e DATABASE_URL=postgres://{DATABASE_URL}:{DATABASE_PASSWORD}@{DATABASE_HOSTNAME}:5432/{DATABASE_NAME}
  -p 8080:8080 
  hostops/stateful-rest-sample
```

### Build using docker
```
git clone https://github.com/hostops/stateful-rest-sample.git
cd stateful-rest-sample
docker build . -t stateful-rest-sample
docker run 
  -e PORT=8080 
  -e DATABASE_URL=postgres://{DATABASE_URL}:{DATABASE_PASSWORD}@{DATABASE_HOSTNAME}:5432/{DATABASE_NAME}
  -p 8080:8080 
  stateful-rest-sample
```

### Building yourself
```bash
git clone https://github.com/hostops/stateful-rest-sample.git
cd stateful-rest-sample
go run main.go 
```

# Contributing
Feel free to open pull request or issue on github.
It is simpler to extend this project than to create a new one with same features or requirements.


