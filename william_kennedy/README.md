# William Kennedy project structure

https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html

## Structure

```
.
├── cmd
│   └── libraryd
│       └── main.go
├── internal
│   ├── platform
│   │   ├── database
│   │   │   └── postgres.go
│   │   └── http
│   │       └── server.go
│   └── users
│       ├── service.go
│       └── user.go
└── README.md

7 directories, 6 files

```
