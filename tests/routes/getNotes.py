import utils
import config
from testRestApi import testRoute, Test, GET

def testToken():
    workspaceID = utils.createWorkspace()
    res = testRoute(GET, f"{config.server}/api/v1/notes/{workspaceID}")
    return res.equals({ "message": "invalid-token" })

def testGetNotes():
    workspaceID = utils.createWorkspace()
    res = testRoute(GET, f"{config.server}/api/v1/notes/{workspaceID}", headers={ "X-Token": config.token })
    print(res.body)
    return res.hasKeys("notes")

tests = [
    Test("Get Notes Invalid Token", testToken),
    Test("Get Notes", testGetNotes),
]
