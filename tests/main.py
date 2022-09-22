from testRestApi import runTests

import routes.getNotes as getNotes
import routes.getSingleNote as getSingleNote
import routes.createNote as createNote
import routes.updateNote as updateNote
import routes.deleteNote as deleteNote

if __name__ == "__main__":
    runTests([
        *getNotes.tests,
        *getSingleNote.tests,
        *createNote.tests,
        *updateNote.tests,
        *deleteNote.tests,
    ])
