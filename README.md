# WIP

## User Endpoints

> `GET /users`

Gets all registered users

**Returns**

```json
[
  {
    "UserName": "example-user",
    "DisplayName": "Example",
  }
]
```

> `POST /users`

Creates a new user

**Expects** 

```json
{ 
  "UserName": "example-user", 
  "DisplayName": "Example", 
  "Password": "terrible-password" 
}
```

**Returns** 

```json
{ 
  "UserName": "example-user", 
  "DisplayName": "Example", 
}
```

## Auth Endpoints

> `POST /auth`

Creates a user session, returning a session key if
authentication is successful

**Expects**

```json
{
  "UserName": "example-user",
  "Password": "terrible-password"
}
```

**Returns**

```json
{
 "Key": "OdT7yQCl1a4xoCXc4OB1X7oSZH4q1bSpCuSEtxwLAu3YKaBd1MMwYfTVP/HbJKZJiNQKayi",
 "DisplayName": "example-user"
}
```

`Key` is always unique
