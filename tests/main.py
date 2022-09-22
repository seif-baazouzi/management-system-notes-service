from testRestApi import runTests

import routes.getNotes as getNotes
import routes.getSingleNote as getSingleNote

if __name__ == "__main__":
    runTests([
        *getNotes.tests,
        *getSingleNote.tests,
    ])
