import utils
import config
from testRestApi import testRoute, Test, POST

def testToken():
    workspaceID = utils.createWorkspace()
    res = testRoute(POST, f"{config.server}/api/v1/notes/{workspaceID}")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    workspaceID = utils.createWorkspace()
    res = testRoute(POST, f"{config.server}/api/v1/notes/{workspaceID}", headers={ "X-Token": config.token })
    return res.equals({ "title": "Must not be empty" })

def testNotExistingWorkspace():
    res = testRoute(POST, f"{config.server}/api/v1/notes/{utils.genUUID()}", headers={ "X-Token": config.token })
    return res.status == 400

def testCreateNote():
    workspaceID = utils.createWorkspace()
    body = { "title": utils.randomString(10), "body": utils.randomString(50) }

    res = testRoute(POST, f"{config.server}/api/v1/notes/{workspaceID}", headers={ "X-Token": config.token }, body=body)
    return res.equals({ "message": "success"}) and res.hasKeys("noteID")


tests = [
    Test("Create Note Invalid Token", testToken),
    Test("Create Note Empty Fields", testEmptyFields),
    Test("Create Note Not Existing Workspace", testNotExistingWorkspace),
    Test("Create Note", testCreateNote),
]
