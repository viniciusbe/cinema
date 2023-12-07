package filmrepo

import (
	"cinema/internal/core/domain/entities"
	"cinema/internal/core/ports"
	"cinema/internal/repositories/neo4jdb/directorrepo"
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/db"
)

type Repository struct {
	driver neo4j.DriverWithContext
	ctx    context.Context
}

func NewNeo4jRepository(driver neo4j.DriverWithContext, ctx context.Context) ports.FilmRepository {
	return &Repository{
		driver: driver,
		ctx:    ctx,
	}
}

func (r *Repository) ListAll() ([]entities.Film, error) {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	filmsResult, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (g:Gender)<-[HAS_GENDER]-(f:Film)-[DIRECTED_BY]->(d:Director) RETURN f,d,g`, nil)
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

	var films []entities.Film
	for _, record := range filmsResult.([]*neo4j.Record) {

		film := FilmFromRecord(record)
		film.Director = *directorrepo.DirectorFromRecord(record)
		film.Genders = GenderFromRecord(record)

		films = append(films, *film)
	}

	return films, nil
}

func (r *Repository) Find(id string) (*entities.Film, error) {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	filmRecord, err := session.ExecuteRead(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (f:Film {filmID: $id}) RETURN f`, map[string]any{
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
		return nil, fmt.Errorf("Erro ao encontrar film -> %w", err)
	}

	film := FilmFromRecord(filmRecord.(*db.Record))

	return film, nil
}

func (r *Repository) Insert(film *entities.Film) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx,
				`MATCH (f:Film)
				WITH f ORDER BY f.filmID DESC LIMIT 1
				CREATE (gn:Film {filmID: toString(toInteger(f.filmID) + 1),description:$description})`,
				map[string]any{
					"description": film.Name,
				})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao criar film -> %w", err)
	}

	return nil
}

func (r *Repository) Save(film *entities.Film) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (f:Film {filmID: $id}) SET f.description = $description`, map[string]any{
				"id":          film.FilmID,
				"description": film.Name,
			})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao atualizar film -> %w", err)
	}

	return nil
}

func (r *Repository) Delete(id string) error {
	ctx := r.ctx
	session := r.driver.NewSession(ctx, neo4j.SessionConfig{DatabaseName: "neo4j"})
	defer session.Close(ctx)

	_, err := session.ExecuteWrite(ctx,
		func(tx neo4j.ManagedTransaction) (any, error) {
			result, err := tx.Run(ctx, `MATCH (f:Film {filmID: $id}) DELETE f`, map[string]any{
				"id": id,
			})
			if err != nil {
				return nil, err
			}
			return result, nil
		})

	if err != nil {
		return fmt.Errorf("Erro ao excluir film -> %w", err)
	}

	return nil
}

func (r *Repository) FindGendersById(ids []string) ([]entities.Gender, error) {
	// var genders []entities.Gender
	// genderErr := r.DB.Find(&genders, ids).Error
	// if genderErr != nil {
	// 	return nil, fmt.Errorf("Erro ao buscar genero(s) -> %w", genderErr)
	// }

	// return genders, nil
	return nil, nil
}

func (r *Repository) FindDirectorById(id string) (*entities.Director, error) {
	// var director *entities.Director
	// err := r.DB.First(&director, id).Error
	// if err != nil {
	// 	return nil, fmt.Errorf("Erro ao buscar diretor -> %w", err)
	// }

	// return director, nil
	return nil, nil

}

func FilmFromRecord(record *db.Record) *entities.Film {
	rawFilmNode, _ := record.Get("f")
	filmNode := rawFilmNode.(neo4j.Node)

	var film entities.Film
	film.FilmID, _ = neo4j.GetProperty[string](filmNode, "filmID")
	film.Name, _ = neo4j.GetProperty[string](filmNode, "name")
	film.Duration, _ = neo4j.GetProperty[string](filmNode, "duration")
	film.Synopsis, _ = neo4j.GetProperty[string](filmNode, "synopsis")
	film.Age, _ = neo4j.GetProperty[string](filmNode, "age")

	return &film
}

func GenderFromRecord(record *db.Record) []entities.Gender {
	rawGenderNode, _ := record.Get("g")
	fmt.Println(rawGenderNode)
	genderNode := rawGenderNode.(neo4j.Node)

	var gender entities.Gender
	var genders []entities.Gender
	gender.Description, _ = neo4j.GetProperty[string](genderNode, "description")
	gender.GenderID, _ = neo4j.GetProperty[string](genderNode, "genderID")

	genders = append(genders, gender)

	return genders
}
