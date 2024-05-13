# Ch 01 Basics 


## Multiple main packages 

Can have multiple `main` packages. The example below can generate at least 2 diff programs - `server` and `demo`: 
```
some-app/
  cmd/
    server/
      main.go # package main
    demo/
      main.go # package main 
    blah.go
    foo.go
```