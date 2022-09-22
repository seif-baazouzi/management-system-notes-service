# Create Note

Used to create new note.

**URL**: `/api/v1/notes/:workspace`

**Method**: `POST`

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
    "message": "success",
    "noteID": "noteID"
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
