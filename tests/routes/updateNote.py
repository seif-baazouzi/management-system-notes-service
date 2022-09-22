import utils
import config
from testRestApi import testRoute, Test, PUT

def testToken():
    noteID = utils.createNote()
    res = testRoute(PUT, f"{config.server}/api/v1/notes/{noteID}")
    return res.equals({ "message": "invalid-token" })

def testEmptyFields():
    noteID = utils.createNote()
    res = testRoute(PUT, f"{config.server}/api/v1/notes/{noteID}", headers={ "X-Token": config.token })
    return res.equals({ "title": "Must not be empty" })

def testUpdateNote():
    noteID = utils.createNote()
    body = { "title": utils.randomString(10), "body": utils.randomString(50) }

    res = testRoute(PUT, f"{config.server}/api/v1/notes/{noteID}", headers={ "X-Token": config.token }, body=body)
    return res.equals({ "message": "success"})


tests = [
    Test("Update Note Invalid Token", testToken),
    Test("Update Note Empty Fields", testEmptyFields),
    Test("Update Note", testUpdateNote),
]
