import utils
import config
from testRestApi import testRoute, Test, GET

def testToken():
    noteID = utils.createNote()
    res = testRoute(GET, f"{config.server}/api/v1/notes/single/{noteID}")
    return res.equals({ "message": "invalid-token" })

def testGetSingleNote():
    noteID = utils.createNote()
    res = testRoute(GET, f"{config.server}/api/v1/notes/single/{noteID}", headers={ "X-Token": config.token })
    return res.hasKeys("note")

tests = [
    Test("Get Single Note Invalid Token", testToken),
    Test("Get Single Note", testGetSingleNote),
]
