-- name: ListChildrenEntities :many
SELECT relationship.*,
  relationship_type.relationship_type_name,
  entity.entity_name,
  entity_type.entity_type_name
FROM relationship
  INNER JOIN relationship_type ON relationship.relationship_type_id = relationship_type.relationship_type_id
  INNER JOIN entity ON relationship.target_entity_id = entity.entity_id
  INNER JOIN entity_type ON entity.entity_type_id = entity_type.entity_type_id
WHERE source_entity_id = $entity_id
LIMIT $limit OFFSET $offset;

-- name: ListParentEntities :many
SELECT relationship.*,
  relationship_type.relationship_type_name,
  entity.entity_name,
  entity_type.entity_type_name
FROM relationship
  INNER JOIN relationship_type ON relationship.relationship_type_id = relationship_type.relationship_type_id
  INNER JOIN entity ON relationship.source_entity_id = entity.entity_id
  INNER JOIN entity_type ON relationship.source_entity_type_id = entity_type.entity_type_id
WHERE target_entity_id = $entity_id
LIMIT $limit OFFSET $offset;

-- name: ListRelationships :many
SELECT relationship.*,
  relationship_type.relationship_type_name,
  source_entity.entity_name AS source_entity_name,
  source_entity_type.entity_type_name AS source_entity_type_name,
  target_entity.entity_name AS target_entity_name,
  target_entity_type.entity_type_name AS target_entity_type_name
FROM relationship
  INNER JOIN source_entity ON relationship.source_entity_id = source_entity.entity_id
  INNER JOIN target_entity ON relationship.target_entity_id = target_entity.entity_id
  INNER JOIN relationship_type ON relationship.relationship_type_id = relationship_type.relationship_type_id
  INNER JOIN source_entity_type ON source_entity.entity_type_id = source_entity_type.entity_type_id
  INNER JOIN target_entity_type ON target_entity.entity_type_id = target_entity_type.entity_type_id
LIMIT $limit OFFSET $offset;
