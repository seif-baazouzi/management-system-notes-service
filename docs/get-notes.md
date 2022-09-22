# Get Notes

Used to get user workspace notes.

**URL**: `/api/v1/notes/:workspaceID`

**Method**: `GET`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "notes": [
        {
            "noteID": "noteID", 
            "title": "title", 
            "body": "body", 
            "userID": "userID", 
            "workspaceID": "workspaceID", 
            "createdAt": "created date" 
        },
        ...
    ]
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
