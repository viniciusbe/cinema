package genderrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
)

type Repository struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

func NewNeo4jRepository(driver neo4j.DriverWithContext, ctx context.Context) ports.GenderRepository {
	return &Repository{
		driver: driver,
		ctx:    ctx,
	}
}

func (r *Repository) ListAll() ([]entities.Gender, error) {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	gendersResult, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (g:Gender) RETURN g`, nil)
			if err != nil {
				return nil, err
			}
			records, err := result.Collect(ctx)
			if err != nil {
				return nil, err
			}

			return records, nil
		})

	if err != nil {
		return nil, fmt.Errorf("Erro ao listar gêneros -> %w", err)
	}

	var genders []entities.Gender
	for _, record := range gendersResult.([]*neo4j.Record) {
		gender := GenderFromRecord(record)
		genders = append(genders, *gender)
	}

	return genders, nil
}

func (r *Repository) Find(id string) (*entities.Gender, error) {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	genderRecord, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (g:Gender {genderID: $id}) RETURN g`, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			record, err := result.Single(ctx)
			if err != nil {
				return nil, err
			}

			return record, nil
		})

	if err != nil {
		return nil, fmt.Errorf("Erro ao encontrar gênero -> %w", err)
	}

	gender := GenderFromRecord(genderRecord.(*db.Record))

	return gender, nil
}

func (r *Repository) Insert(gender *entities.Gender) error {

	// if err != nil {
	// 	return fmt.Errorf("Erro ao inserir gênero -> %w", err)
	// }

	return nil
}

func (r *Repository) Save(gender *entities.Gender) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (g:Gender {genderID: $id}) SET g.description = $description`, map[string]any{
				"id":          gender.GenderID,
				"description": gender.Description,
			})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao atualizar gênero -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (g:Gender {genderID: $id}) DELETE g`, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao excluir gênero -> %w", err)
	}

	return nil
}

func GenderFromRecord(record *db.Record) *entities.Gender {
	rawGenderNode, _ := record.Get("g")
	genderNode := rawGenderNode.(neo4j.Node)

	var gender entities.Gender
	gender.Description, _ = neo4j.GetProperty[string](genderNode, "description")
	gender.GenderID, _ = neo4j.GetProperty[string](genderNode, "genderID")

	return &gender
}
