# Gossip With Go

This is a repository for a Go backend for a web forum application, and it has endpoints which facilitate CRUD 
functionality for the application. 

* This project uses [go-chi](https://github.com/go-chi/chi) as a web framework.
* [golangci](https://github.com/golangci/golangci-lint) has been used to lint the code. 


## Getting Started

### Installing Go

Download and install Go by following the instructions [here](https://go.dev/doc/install).

### Running the app
1. [Fork](https://docs.github.com/en/get-started/quickstart/fork-a-repo#forking-a-repository) this repo.
2. [Clone](https://docs.github.com/en/get-started/quickstart/fork-a-repo#cloning-your-forked-repository) **your** forked repo.
3. Open your terminal and navigate to the directory containing your cloned project.
4. Run `go run cmd/server/main.go` and head over to http://localhost:8000/users to view the response.


### Navigating the code
This is the main file structure, which takes the Model-View-Controller design paradigm into account.
```
.
├── README.md
├── cmd
│   └── server
│       └── main.go
├── erdiagram.png
├── go.mod
├── go.sum
└── internal
    ├── api                             # Encapsulates types and utilities related to the API
    │   └── api.go
    ├── dataaccess                      # Data Access layer accesses data from the database
    │   ├── posts
    │   │   └── posts.go
    │   ├── tags
    │   │   └── tags.go
    │   ├── threads
    │   │   └── threads.go
    │   └── users
    │       └── users.go
    ├── database                        # Encapsulates the types and utilities related to the database
    │   ├── database.go
    │   └── seed.sql
    ├── handlers                        # Handler functions to respond to requests
    │   ├── posts.go
    │   ├── tags.go
    │   ├── threads.go
    │   └── users.go
    ├── models                          # Definitions of objects used in the application
    │   ├── post.go
    │   ├── tag.go
    │   ├── thread.go
    │   └── user.go
    ├── router                          # Encapsulates types and utilities related to the router
    │   └── router.go
    └── routes                          # Defines routes that are used in the application
        └── routes.go

```
### Schema
<br/>
<img src="./erdiagram.png" width="400" alt="Entity relationship diagram">
<br/>
The entity relationship diagram above depicts the schema of the database and the relationships each entity has.

Main directories/files to note:
* `cmd` contains the main entry point for the application
* `internal` holds most of the functional code for your project that is specific to the core logic of your application
* `README.md` is a form of documentation about the project. It is what you are reading right now.
* `go.mod` contains important metadata, for example, the dependencies in the project. See [here](https://go.dev/ref/mod) for more information
* `go.sum` See [here](https://go.dev/ref/mod) for more information
