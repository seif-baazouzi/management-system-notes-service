from testRestApi import runTests

import routes.getNotes as getNotes

if __name__ == "__main__":
    runTests([
        *getNotes.tests,
    ])
