# Update Note

Used to update an existing note.

**URL**: `/api/v1/notes/:noteID`

**Method**: `PUT`

**Auth required**: Send User `JWT token` in `X-Token` header

### Request

```json
{
    "title": "title", 
    "body": "body", // optional
}
```

### Success Response

**CODE**: `200`

```json
{
    "message": "success"
}
```

### Invalid Input Response

**CODE**: `200`

```json
{
    "title": "title error message", 
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
