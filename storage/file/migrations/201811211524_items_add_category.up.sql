CREATE TABLE items_temp
(
  id        INTEGER NOT NULL PRIMARY KEY,
  name      TEXT    DEFAULT '',
  filename  TEXT    DEFAULT '',
  path      TEXT    DEFAULT '',
  size      INTEGER DEFAULT 0,
  category  INTEGER DEFAULT 2,
  available INTEGER DEFAULT 0
);
INSERT INTO items_temp(id, name, filename, path, size, available)
SELECT id, name, filename, path, size, available
FROM items;
DROP TABLE items;
ALTER TABLE items_temp RENAME TO items;
