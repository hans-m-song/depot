-- name: ListAttributesByEntityID :many
SELECT attribute.*,
  attribute_type.attribute_type_name
FROM attribute
  INNER JOIN attribute_type ON attribute.attribute_type_id = attribute_type.attribute_type_id
WHERE entity_id = $entity_id;
