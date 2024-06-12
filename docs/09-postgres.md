# Databases and PostgreSQL

Install, then switch to postgres acct and connect:
```sh
sudo apt install postgresql postgresql-contrib
sudo -i -u postgres
psql
```
To quit: `\q`
To login directly: 
```
sudo -u postgres psql
```

## Basic setup 

```sh
$ createuser --interactive
## sudo -u postgres createuser --interactive
Enter name of role to add: llocked
Shall the new role be a superuser? (y/n) y
```
```sh
createdb llocked
su - adminacct
sudo adduser llocked
sudo -i -u llocked
```
```sql
psql (12.18 (Ubuntu 12.18-0ubuntu0.20.04.1))
llocked=# \conninfo
You are connected to database "llocked" as user "llocked" via socket in "/var/run/postgresql" at port "5433".
```
Show databases: 
Show tables: 
