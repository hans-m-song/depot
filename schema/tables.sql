CREATE TABLE entity_type (
  entity_type_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  entity_type_name TEXT UNIQUE NOT NULL
);

CREATE VIEW source_entity_type AS
SELECT *
FROM entity_type;

CREATE VIEW target_entity_type AS
SELECT *
FROM entity_type;

CREATE TABLE entity (
  entity_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  entity_name TEXT NOT NULL,
  entity_type_id INTEGER NOT NULL REFERENCES entity_type (entity_type_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE VIEW source_entity AS
SELECT *
FROM entity;

CREATE VIEW target_entity AS
SELECT *
FROM entity;

CREATE TABLE relationship_type (
  relationship_type_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  relationship_type_name TEXT UNIQUE NOT NULL
);

CREATE TABLE relationship (
  relationship_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  source_entity_id INTEGER NOT NULL REFERENCES entity (entity_id) ON UPDATE CASCADE ON DELETE CASCADE,
  target_entity_id INTEGER NOT NULL REFERENCES entity (entity_id) ON UPDATE CASCADE ON DELETE CASCADE,
  relationship_type_id INTEGER NOT NULL REFERENCES relationship_type (relationship_type_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE UNIQUE INDEX index_relationship_source_entity_id ON relationship (source_entity_id);

CREATE UNIQUE INDEX index_relationship_target_entity_id ON relationship (target_entity_id);

CREATE TABLE communication_type (
  communication_type_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  communication_type_name TEXT UNIQUE NOT NULL
);

CREATE TABLE communication (
  communication_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  entity_id INTEGER NOT NULL REFERENCES entity (entity_id) ON UPDATE CASCADE ON DELETE CASCADE,
  communication_name TEXT NOT NULL,
  communication_destination TEXT NOT NULL,
  communication_type_id INTEGER NOT NULL REFERENCES communication_type (communication_type_id) ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE UNIQUE INDEX index_communication_entity_id ON communication (entity_id);

CREATE TABLE attribute_type (
  attribute_type_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  attribute_type_name TEXT UNIQUE NOT NULL
);

CREATE TABLE attribute (
  attribute_id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  entity_id INTEGER NOT NULL REFERENCES entity (entity_id) ON UPDATE CASCADE ON DELETE CASCADE,
  attribute_type_id INTEGER NOT NULL REFERENCES attribute_type (attribute_type_id) ON UPDATE CASCADE ON DELETE RESTRICT,
  attribute_value TEXT NOT NULL
);

CREATE UNIQUE INDEX index_attribute_entity_id ON attribute (entity_id);
