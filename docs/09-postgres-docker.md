# Databases and PostgreSQL 

## Docker 

```
docker pull postgres
docker images 
docker ps-a 
docker start | stop
```
Committing image changes: 
```go
docker commit -m "What you did to the image" -a "Author Name" container_id repository/new_image_name
```
```
docker run ubuntu
docker start -it # interactively
rm | stop
```

```
$ sudo docker commit -m "add curl" -a "dell" 72ee1c2013 images/ubuntu
sha256:280623a2d726c5c89aa5a3e2e1b4fda661125492024b64f51745a80fc29de31a
$ sudo docker images | grep ubuntu
images/ubuntu   latest    280623a2d726   51 seconds ago   124MB
ubuntu          latest    17c0145030df   13 days ago      76.2MB
```

```sh
$ sudo docker commit -m "add curl" -a "dell" 72ee1c2013 images/ubuntu2
sha256:280623a2d726c5c89aa5a3e2e1b4fda661125492024b64f51745a80fc29de31a
$ docker login -u vl
Password: 
WARNING! Your password will be stored unencrypted in /home/dell/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store
Login Succeeded
```

## More Docker

Rebuild services from scratch: `docker compose down`
Run out docker containers in detouched mode: `docker compose up -d`
To stop cont's running detouched: `docker compose stop`
Check running apps: `docker compose ls`
Reset database: 
```
docker compose down
docker compose up
```

Run a binary inside a docker cont: `docker compose exec -it db psql -U vl -d llocked`
* exec: execute a binary inside a container run by docker compose
* -it: enable us to interact w/ the terminal session after running the command
* db: name of the service we wanna use to exec the command
* psql: run psql binary that's inside the docker container 
* -U xxx -d yyy: username and database 

Connect db: 
```
\c llocked
```

A couple queries: 
```sql 
CREATE TABLE users ( id SERIAL PRIMARY KEY, email TEXT );
INSERT INTO users ( email ) VALUES ( 'vl@chammy.info' );
SELECT * FROM users
-- RES
 id |     email      
----+----------------
  1 | vl@chammy.info
```
Now: `docker ps` and check what's in there:  
CONTAINER ID   | IMAGE      | COMMAND                  | CREATED       | STATUS         | PORTS                                       | NAMES
---|---|---|---|---|---|---
5f5ac5236981   | adminer    | "entrypoint.sh php -…"   | 5 hours ago   | Up 9 minutes   | 0.0.0.0:3333->8080/tcp, :::3333->8080/tcp   | llocked-adminer-1
86968536ab7a   | postgres   | "docker-entrypoint.s…"   | 5 hours ago   | Up 9 minutes   | 0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   | llocked-db-1

Important: NAMES - `llocked-adminer-1`, `llocked-db-1`
Example: `docker exec -it llocked-db-1 /usr/bin/psql -U vl -d llocked`

## Types and constraints 

Type | Description
---|---
int | 2,147,483,648...2,147,483,647
serial | 1...2,147,483,647
varchar | similar to Go's `string` 
text | another text 

Constraint | Description
---|---
UNIQUE | Unique val
NOT NULL  | 
PRIMARY KEY | 



## Changing tables + adding data

```
drop table if exists users;
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    age INT,
    first_name TEXT,
    last_name TEXT,
    email TEXT UNIQUE NOT NULL
);
```
Inserting data: if no val, serial starts with 1, so first error then ok:
```
llocked=# INSERT INTO users (age, email, first_name, last_name)
VALUES (40, 'V', 'L', 'vl@chammy.info');
ERROR:  duplicate key value violates unique constraint "users_pkey"
DETAIL:  Key (id)=(1) already exists.
llocked=# INSERT INTO users (age, email, first_name, last_name)
VALUES (40, 'V', 'L', 'vl@chammy.info');
INSERT 0 1
```

## Querying, updating, and deleting records 

```sql
SELECT * FROM users;
SELECT id, email FROM users;
SELECT * FROM users
WHERE age < 30 OR last_name = 'Smith';
```
Updating: 
```sql
UPDATE users
SET first_name = 'John', last_name = 'Appleseed'
WHERE id = 2;
-- Update all users
UPDATE users
SET first_name = 'John';
UPDATE users
SET first_name = 'Jon'
WHERE first_name = 'John';
```
Deleting:  
```sql
INSERT INTO users (age, email, first_name, last_name)
VALUES (40, 'gztrk@engineer.com', 'V', 'L');
DELETE FROM users WHERE id = 3;
```


