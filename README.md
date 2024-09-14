# WIP

## User Endpoints

> `GET /users`

Returns all registered users

Will return the following object

```json
[
  {
    "UserName": "example-user",
    "DisplayName": "Example",
  }
]
```

## Auth Endpoints

> `POST /auth`

Creates a user session, returning a session key if
authentication is successful

Expected Body

```json
{
  "UserName": "example-user",
  "Password": "terrible-password"
}
```

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

**Expected**

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
