-- name: CreateEntity :one
INSERT INTO entity (entity_type_id, entity_name)
VALUES ($entity_type_id, $entity_name)
RETURNING *;

-- name: GetEntityByID :one
SELECT entity.*,
  entity_type.entity_type_name
FROM entity
  INNER JOIN entity_type ON entity.entity_type_id = entity_type.entity_type_id
WHERE entity.entity_id = $entity_id
LIMIT 1;

-- name: GetEntityByName :one
SELECT entity.*,
  entity_type.entity_type_name
FROM entity
  INNER JOIN entity_type ON entity.entity_type_id = entity_type.entity_type_id
WHERE entity_name = $entity_name
LIMIT 1;

-- name: ListEntities :many
SELECT entity.*,
  entity_type.entity_type_name
FROM entity
  INNER JOIN entity_type ON entity.entity_type_id = entity_type.entity_type_id
LIMIT $limit OFFSET $offset;

-- name: ListEntitiesByType :many
SELECT entity.*,
  entity_type.entity_type_name
FROM entity
  INNER JOIN entity_type ON entity.entity_type_id = entity_type.entity_type_id
WHERE entity.entity_type_id = $entity_type_id
LIMIT $limit OFFSET $offset;

-- name: UpdateEntityName :one
UPDATE entity
SET entity_name = $entity_name
WHERE entity_id = $entity_id
RETURNING *;

-- name: DeleteEntity :exec
DELETE FROM entity
WHERE entity_id = $entity_id;
