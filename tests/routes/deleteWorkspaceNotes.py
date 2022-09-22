import utils
import config
from testRestApi import testRoute, Test, DELETE

def testToken():
    workspaceID = utils.createWorkspace()
    res = testRoute(DELETE, f"{config.server}/api/v1/notes/workspace/{workspaceID}")
    return res.equals({ "message": "invalid-token" })


def testNotExistingWorkspace():
    res = testRoute(DELETE, f"{config.server}/api/v1/notes/workspace/{utils.genUUID()}", headers={ "X-Token": config.token })
    return res.status == 400

def testDeleteNote():
    workspaceID = utils.createWorkspace()
    res = testRoute(DELETE, f"{config.server}/api/v1/notes/workspace/{workspaceID}", headers={ "X-Token": config.token })
    return res.equals({ "message": "success"})


tests = [
    Test("Delete Note Invalid Token", testToken),
    Test("Delete Note Not Existing Workspace", testNotExistingWorkspace),
    Test("Delete Note", testDeleteNote),
]
