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
