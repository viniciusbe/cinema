package genderrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
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
	var genders []entities.Gender
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
	for _, person := range gendersResult.([]*neo4j.Record) {

		var gender entities.Gender
		// gender.Description,_ = gendersResult
		rawPersonNode, _ := person.Get("g")

		// gender.Description, _, _ = neo4j.GetRecordValue[string](person, "gender")
		gender.Description, _ = neo4j.GetProperty[string](rawPersonNode.(neo4j.Node), "description")
		gender.GenderID, _ = person.AsMap()["name"].(string)

		fmt.Println(gender.Description)
	}

	return genders, nil
}

func (r *Repository) Find(id string) (*entities.Gender, error) {
	var gender *entities.Gender

	// if err != nil {
	// 	return nil, fmt.Errorf("Erro ao encontrar gênero -> %w", err)
	// }

	return gender, nil
}

func (r *Repository) Insert(gender *entities.Gender) error {

	// if err != nil {
	// 	return fmt.Errorf("Erro ao inserir gênero -> %w", err)
	// }

	return nil
}

func (r *Repository) Save(gender *entities.Gender) error {

	// if err != nil {
	// 	return fmt.Errorf("Erro ao atualizar gênero -> %w", err)
	// }

	return nil
}

func (r *Repository) Delete(id string) error {

	// if err != nil {
	// 	return fmt.Errorf("Erro ao excluir gênero -> %w", err)
	// }

	return nil
}
