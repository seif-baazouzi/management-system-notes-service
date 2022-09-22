# Delete Workspace Notes

Used to delete workspace notes.

**URL**: `/api/v1/notes/workspace/:workspaceID`

**Method**: `DELETE`

**Auth required**: Send User `JWT token` in `X-Token` header

### Success Response

**CODE**: `200`

```json
{
    "message": "success"
}
```

### Server Error Response

**CODE**: `500`

```json
{
    "message": "server-error"
}
```
