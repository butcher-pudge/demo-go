# Go Learning Demo
Example that shows core principles of the Clean Architecture in Golang projects.

### Project Description&Structure:
REST API with custom JWT-based authentication system. Core functionality is about creating and managing bookmarks (Simple clone of <a href="https://app.getpocket.com/">Pocket</a>).

#### Structure:
4 Domain layers:

- Models layer
- Repository layer
- Service layer
- Controller layer

## API:

### POST /auth/sign-up

Creates new user

##### Example Input:
```
{
	"username": "UncleBob",
	"password": "123456"
} 
```


### POST /auth/sign-in

Request to get JWT Token based on user credentials

##### Example Input:
```
{
	"username": "UncleBob",
	"password": "123456"
} 
```

##### Example Response:
```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTg0ODExOTAuMzI5NjY0LCJ1c2VyIjp7IklEIjoiNjUzYTE3MTEzYjcyMTBhZjk4NTlmNzVlIiwiVXNlck5hbWUiOiJ5eSIsIlBhc3N3b3JkIjoiZWI4ZDU2ZjFhMzExMzQ0NmM4OTI2OGY5OTRkNWYwYmY2YzQxMThlYSJ9fQ._Sn5Z16G66-c8WTYLmjpWOQlnouzncERzUqhB8nEcDk"
} 
```

### POST /api/bookmarks

Creates new bookmark

##### Example Input:
```
{
	"url": "https://github.com/xxxx",
	"title": "Go Learning Demo"
} 
```

### GET /api/bookmarks

Returns all user bookmarks

##### Example Response:
```
{
	"bookmarks": [
            {
                "id": "5da2d8aae9g637v5ddfae750",
                "url": "https://github.com/xxxx",
                "title": "Go Learning Demo"
            }
    ]
} 
```

## Requirements
- go 1.21
- docker & docker-compose

## Run Project

Use ```make run``` to build and run docker containers with application itself and mongodb instance

