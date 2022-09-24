# Get Notes

Used to get user single note.

**URL**: `/api/v1/notes/single/:noteID`

**Method**: `GET`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "note": {
        "noteID": "noteID", 
        "title": "title", 
        "body": "body", 
        "userID": "userID", 
        "workspaceID": "workspaceID", 
        "createdAt": "created date" 
    }
}
```

### Note Not Found

**CODE**: `404`

```json
{
    "note": null
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
