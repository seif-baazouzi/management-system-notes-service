import utils
import config
from testRestApi import testRoute, Test, DELETE

def testToken():
    noteID = utils.createNote()
    res = testRoute(DELETE, f"{config.server}/api/v1/notes/{noteID}")
    return res.equals({ "message": "invalid-token" })

def testDeleteNote():
    noteID = utils.createNote()
    res = testRoute(DELETE, f"{config.server}/api/v1/notes/{noteID}", headers={ "X-Token": config.token })
    return res.equals({ "message": "success"})


tests = [
    Test("Delete Note Invalid Token", testToken),
    Test("Delete Note", testDeleteNote),
]
