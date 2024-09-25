# WIP

## User Endpoints

> `GET /users`

Gets all registered users

**Returns**

```json
[
  {
    "username": "example-user",
    "displayname": "Example",
  }
]
```

> `POST /users`

Creates a new user

**Expects** 

```json
{ 
  "username": "example-user", 
  "displayname": "Example", 
  "password": "terrible-password" 
}
```

**Returns** 

```json
{ 
  "username": "example-user",
  "displayname": "Example",
}
```

## Auth Endpoints

> `POST /auth`

Creates a user session, returning a session key if
authentication is successful

**Expects**

```json
{
  "username": "example-user",
  "password": "terrible-password"
}
```

**Returns**

```json
{
 "key": "OdT7yQCl1a4xoCXc4OB1X7oSZH4q1bSpCuSEtxwLAu3YKaBd1MMwYfTVP/HbJKZJiNQKayi",
 "displayname": "example-user"
}
```

`Key` is always unique

> `DELETE /auth`

Deletes a user session, an action that should be taken when
logging out.

**Expects**

```json
{
 "key": "OdT7yQCl1a4xoCXc4OB1X7oSZH4q1bSpCuSEtxwLAu3YKaBd1MMwYfTVP/HbJKZJiNQKayi",
}
```

The key is the session that will be deleted

**Returns**

`200 OK` if all goes well
