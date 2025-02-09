CREATE TABLE IF NOT EXISTS projects (
  id   INTEGER PRIMARY KEY,
  name text    NOT NULL,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS requests (
  id         INTEGER PRIMARY KEY,
  project_id INTEGER NOT NULL,
  name       TEXT    NOT NULL,
  curl       TEXT    NOT NULL,
  method     TEXT    NOT NULL,
  url        TEXT    NOT NULL,
  headers    TEXT,
  body       TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
  FOREIGN KEY (project_id) REFERENCES projects(id)
);

CREATE TABLE IF NOT EXISTS selected_project (
  project_id INTEGER NOT NULL,
  FOREIGN KEY (project_id) REFERENCES projects(id)
);
