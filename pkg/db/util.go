package db

import (
	"context"
	"database/sql"
	"fmt"
)

var (
	CommunicationTypes []CommunicationType
	EntityTypes        []EntityType
	RelationshipTypes  []RelationshipType
)

func Seed(ctx context.Context, db *sql.DB, q *Queries) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction for seeding: %s", err)
	}

	defer tx.Rollback()

	qtx := q.WithTx(tx)

	if err := qtx.SeedAttributeTypes(ctx); err != nil {
		return fmt.Errorf("failed to seed attribute types: %s", err)
	}

	if err := qtx.SeedCommunicationTypes(ctx); err != nil {
		return fmt.Errorf("failed to seed communication types: %s", err)
	}

	if err := qtx.SeedEntityTypes(ctx); err != nil {
		return fmt.Errorf("failed to seed entity types: %s", err)
	}

	if err := qtx.SeedRelationshipTypes(ctx); err != nil {
		return fmt.Errorf("failed to seed relationship types: %s", err)
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction for seeding: %s", err)
	}

	if CommunicationTypes, err = q.ListCommunicationTypes(ctx); err != nil {
		return fmt.Errorf("failed to list communication types: %s", err)
	}

	if EntityTypes, err = q.ListEntityTypes(ctx); err != nil {
		return fmt.Errorf("failed to list entity types: %s", err)
	}

	if RelationshipTypes, err = q.ListRelationshipTypes(ctx); err != nil {
		return fmt.Errorf("failed to list relationship types: %s", err)
	}

	return nil
}
