CREATE DATABASE notes;

\c notes

CREATE TABLE notes (
    noteID UUID PRIMARY KEY,
    title VARCHAR NOT NULL, 
    body VARCHAR NOT NULL, 
    userID UUID NOT NULL,
    workspaceID UUID NOT NULL,
    createdAt TIMESTAMP DEFAULT NOW() 
);

CREATE INDEX UserIdInex ON notes USING hash (userID);
CREATE INDEX WorkspaceInex ON notes USING hash (workspaceID);
