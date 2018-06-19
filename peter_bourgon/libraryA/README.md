# Peter Bourgon complex project structure

https://peter.bourgon.org/go-best-practices-2016/#repository-structure

## Structure
```
.
├── cmd
│   └── server
│       └── main.go
├── pkg
│   ├── http
│   │   ├── handler.go
│   │   └── server.go
│   ├── library
│   │   ├── book.go
│   │   └── user.go
│   └── postgres
│       ├── connection.go
│       └── user_service.go
└── README.md

6 directories, 8 files
```
