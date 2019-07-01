package psql

import (
	"context"
	"fmt"

	"github.com/Mrcampbell/pgo2/pokemon-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

type PokemonService struct {
	DB              *pg.DB
	breedClient     pokemon.BreedServiceClient
	breedMoveClient pokemon.BreedMoveServiceClient
}

func NewPokemonService(breedClient pokemon.BreedServiceClient, breedMoveClient pokemon.BreedMoveServiceClient) (*PokemonService, error) {
	db := pg.Connect(&pg.Options{
		Addr:     config.DBHost + ":" + config.DBUser,
		Database: config.DBDatabase,
		User:     config.DBUser,
		Password: config.DBPassword,
		OnConnect: func(conn *pg.Conn) error {
			fmt.Println("Connected to DB")
			return nil
		},
	})

	err := createSchema(db)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = initialize(db)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &PokemonService{
		DB:              db,
		breedClient:     breedClient,
		breedMoveClient: breedMoveClient,
	}, nil
}

func (ps *PokemonService) GetPokemon(ctx context.Context, req *pokemon.GetPokemonRequest) (*pokemon.GetPokemonResponse, error) {
	p := pokemon.Pokemon{Id: req.Id}
	err := ps.DB.Select(&p)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	b, err := ps.breedClient.GetBreedSummary(ctx, &pokemon.GetBreedSummaryRequest{
		Id: p.BreedId,
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	p.BreedSummary = b.Summary

	return &pokemon.GetPokemonResponse{
		Pokemon: &p,
	}, nil
}

func (ps *PokemonService) ListPokemon(ctx context.Context, req *pokemon.ListPokemonRequest) (*pokemon.ListPokemonResponse, error) {
	var p []*pokemon.Pokemon
	err := ps.DB.Model(&p).Select()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pokemon.ListPokemonResponse{
		Pokemon: p,
	}, nil
}

func (ps *PokemonService) InternalCreatePokemon(ctx context.Context, req *pokemon.InternalCreatePokemonRequest) (*pokemon.InternalCreatePokemonResponse, error) {
	fmt.Println("PRE BREED REQUEST")
	breed, err := ps.breedClient.GetBreedDetail(ctx, &pokemon.GetBreedDetailRequest{
		Id: req.BreedId,
		// VersionGroup: req.,
	})

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result, err := ps.buildPokemon(ctx, *breed.Detail, *req, pokemon.VersionGroup_ULTRA_SUN_ULTRA_MOON) // todo not hard code
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = Upsert(ps.DB, &result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &pokemon.InternalCreatePokemonResponse{
		Pokemon: &result,
	}, nil
}

func createSchema(db *pg.DB) error {
	for _, model := range []interface{}{(*pokemon.Pokemon)(nil)} {

		// todo: remove after debug
		// db.DropTable(model, &orm.DropTableOptions{
		// 	Cascade:  true,
		// 	IfExists: true,
		// })
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        false,
			IfNotExists: true,
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func initialize(db *pg.DB) error {
	p1Iv := pokemon.StatGroup{
		Hp:             1,
		Defense:        4,
		SpecialAttack:  23,
		SpecialDefense: 99,
	}
	p1Ev := pokemon.StatGroup{
		Speed: 23,
		Hp:    12,
	}
	p1Stat := pokemon.StatGroup{}

	p1 := pokemon.Pokemon{
		Id:      "7",
		BreedId: "7",
		Iv:      &p1Iv,
		Ev:      &p1Ev,
		Stat:    &p1Stat,
	}

	err := Upsert(db, &p1)
	if err != nil {
		return err
	}

	return nil
}

func Upsert(db *pg.DB, model interface{}) error {
	err := db.Update(model)
	if err != nil {
		err = db.Insert(model)
		if err != nil {
			return err
		}
	}
	return nil
}
