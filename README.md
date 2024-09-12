# WIP

## User Endpoints

> `GET /users`

Returns all registered users

> `POST /users`

Creates a new user

Expected Body 

```json
{ 
  "UserName": "example-user", 
  "DisplayName": "Example", 
  "Password": "terrible-password" 
}
```

## Auth Endpoints

> `POST /auth`

Creates a user session, returning a session key if
authentication is successful

Expected Body

```json
{
  "Id": 0,
  "Password": "terrible-password"
}
```
