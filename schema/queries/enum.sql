-- name: ListAttributeTypes :many
SELECT *
FROM attribute_type;

-- name: SeedAttributeTypes :exec
INSERT
  OR IGNORE INTO attribute_type (attribute_type_name)
VALUES ('USER_EXTERNAL_ID');

-- name: ListCommunicationTypes :many
SELECT *
FROM communication_type;

-- name: SeedCommunicationTypes :exec
INSERT
  OR IGNORE INTO communication_type (communication_type_name)
VALUES ('EMAIL'),
  ('SLACK');

-- name: ListEntityTypes :many
SELECT *
FROM entity_type;

-- name: SeedEntityTypes :exec
INSERT
  OR IGNORE INTO entity_type (entity_type_name)
VALUES ('APPLICATION'),
  ('TEAM'),
  ('USER');

-- name: ListRelationshipTypes :many
SELECT *
FROM relationship_type;

-- name: SeedRelationshipTypes :exec
INSERT
  OR IGNORE INTO relationship_type (relationship_type_name)
VALUES ('DEPENDS_ON'),
  ('OWNS'),
  ('SUPERVISES');
