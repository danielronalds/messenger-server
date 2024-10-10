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
 "key": "OdT7yQCl1a4xoCXc4OB1X7oSZH4q1bSpCuSEtxwLAu3YKaBd1MMwYfTVP/HbJKZJiNQKayi"
}
```

The key is the session that will be deleted

**Returns**

`200 OK` if all goes well

## Message Endpoints

> `POST /message`

Sends a message from the user to another user. Uses
authentication.

**Expects**

```json
{
 "key": "OdT7yQCl1a4xoCXc4OB1X7oSZH4q1bSpCuSEtxwLAu3YKaBd1MMwYfTVP/HbJKZJiNQKayi",
 "to": "example-user",
 "content": "Hello world!"
}
```

**Returns**

```json
{
	"id": 1
	"sender": "jonsnow"
	"receiver": "example-user"
	"content": "Hello, world!"
	"delivered": time.Now()
	"isRead": false
}
```

## Inbox Endpoint

> `POST /inbox`

Gets a conversation between the logged in user and the
supplied contact. Does not return unread messages from the
contact.

**Expects**

```json
{
 "key": "OdT7yQCl1a4xoCXc4OB1X7oSZH4q1bSpCuSEtxwLAu3YKaBd1MMwYfTVP/HbJKZJiNQKayi",
 "contact": "jonsnow"
}
```

> `POST /inbox/unread`

Gets unread messages for all contacts.

**Expects**

```json
{
 "key": "OdT7yQCl1a4xoCXc4OB1X7oSZH4q1bSpCuSEtxwLAu3YKaBd1MMwYfTVP/HbJKZJiNQKayi",
}
```
