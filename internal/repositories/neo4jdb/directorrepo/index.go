package directorrepo

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

func NewNeo4jRepository(driver neo4j.DriverWithContext, ctx context.Context) ports.DirectorRepository {
	return &Repository{
		driver: driver,
		ctx:    ctx,
	}
}

func (r *Repository) ListAll() ([]entities.Director, error) {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	directorsResult, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (d:Director) RETURN d`, nil)
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
		return nil, fmt.Errorf("Erro ao listar diretores -> %w", err)
	}

	var directors []entities.Director
	for _, record := range directorsResult.([]*neo4j.Record) {
		director := DirectorFromRecord(record)
		directors = append(directors, *director)
	}

	return directors, nil
}

func (r *Repository) Find(id string) (*entities.Director, error) {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	directorRecord, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (d:Director {directorID: $id}) RETURN d`, map[string]any{
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
		return nil, fmt.Errorf("Erro ao encontrar diretor -> %w", err)
	}

	director := DirectorFromRecord(directorRecord.(*db.Record))

	return director, nil
}

func (r *Repository) Insert(director *entities.Director) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx,
				`MATCH (d:Director)
				WITH d ORDER BY d.directorID DESC LIMIT 1
				CREATE (gn:Director {directorID: toString(toInteger(d.directorID) + 1),name:$name})`,
				map[string]any{
					"name": director.Name,
				})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao criar diretor -> %w", err)
	}

	return nil
}

func (r *Repository) Save(director *entities.Director) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (d:Director {directorID: $id}) SET d.name = $name`, map[string]any{
				"id":   director.DirectorID,
				"name": director.Name,
			})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao atualizar diretor -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (d:Director {directorID: $id}) DELETE d`, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao excluir diretor -> %w", err)
	}

	return nil
}

func DirectorFromRecord(record *db.Record) *entities.Director {
	rawDirectorNode, _ := record.Get("d")
	directorNode := rawDirectorNode.(neo4j.Node)

	var director entities.Director
	director.DirectorID, _ = neo4j.GetProperty[string](directorNode, "directorID")
	director.Name, _ = neo4j.GetProperty[string](directorNode, "name")

	return &director
}
