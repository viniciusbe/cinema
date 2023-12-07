package buyerrepo

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

func NewNeo4jRepository(driver neo4j.DriverWithContext, ctx context.Context) ports.BuyerRepository {
	return &Repository{
		driver: driver,
		ctx:    ctx,
	}
}

func (r *Repository) ListAll() ([]entities.Buyer, error) {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	buyersResult, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (g:Buyer) RETURN g`, nil)
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

	var buyers []entities.Buyer
	for _, record := range buyersResult.([]*neo4j.Record) {
		buyer := BuyerFromRecord(record)
		buyers = append(buyers, *buyer)
	}

	return buyers, nil
}

func (r *Repository) Find(id string) (*entities.Buyer, error) {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	buyerRecord, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (g:Buyer {buyerID: $id}) RETURN g`, map[string]any{
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
		return nil, fmt.Errorf("Erro ao encontrar pagante -> %w", err)
	}

	buyer := BuyerFromRecord(buyerRecord.(*db.Record))

	return buyer, nil
}

func (r *Repository) Insert(buyer *entities.Buyer) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx,
				`MATCH (b:Buyer)
				WITH b ORDER BY b.buyerID DESC LIMIT 1
				CREATE (bu:Buyer {buyerID: toString(toInteger(b.buyerID) + 1),document:$document,name:$name})`,
				map[string]any{
					"document": buyer.Document,
					"name":     buyer.Name,
				})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao criar pagante -> %w", err)
	}

	return nil
}

func (r *Repository) Save(buyer *entities.Buyer) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (b:Buyer {buyerID: $id}) SET b.name = $name,b.document = $document`,
				map[string]any{
					"id":       buyer.BuyerID,
					"name":     buyer.Name,
					"document": buyer.Document,
				})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao atualizar pagante -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (g:Buyer {buyerID: $id}) DELETE g`, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao excluir pagante -> %w", err)
	}

	return nil
}

func BuyerFromRecord(record *db.Record) *entities.Buyer {
	rawBuyerNode, _ := record.Get("g")
	buyerNode := rawBuyerNode.(neo4j.Node)

	var buyer entities.Buyer
	buyer.BuyerID, _ = neo4j.GetProperty[string](buyerNode, "buyerID")
	buyer.Name, _ = neo4j.GetProperty[string](buyerNode, "name")
	buyer.Document, _ = neo4j.GetProperty[string](buyerNode, "document")

	return &buyer
}
