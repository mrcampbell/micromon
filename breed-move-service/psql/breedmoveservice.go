package psql

import (
	"context"
	"fmt"

	"github.com/Mrcampbell/pgo2/breed-move-service/config"
	"github.com/Mrcampbell/pgo2/protorepo/pokemon"
	"github.com/go-pg/pg"
)

type BreedMoveService struct {
	DB *pg.DB
}

func NewBreedMoveService() (*BreedMoveService, error) {
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

	// no need to initialize or create schema because it should exist.
	// It's a huge process right now

	return &BreedMoveService{
		DB: db,
	}, nil
}

// GetAllPreviousLevelUpLearnableMoveIDsForBreed gets ALL the possible moves for a pokemon to have learned in the wild.
//  Useful for when generating wild pokemon and you need a random moveset.
func (ps *BreedMoveService) GetAllPreviousLevelUpLearnableMoveIDsForBreed(breedID string, generation, level int) ([]int, error) {
	return nil, nil
}

// GetRandomMoveIDSetForBreed Calls GetAllPreviousLevelUpLearnableMoveIDsForBreed but only returns (up to) 4 id's
func (ps *BreedMoveService) GetRandomMoveIDSetForBreed(breedID string, generation, level int) ([]int, error) {
	return nil, nil
}

func (ps *BreedMoveService) GetMovesForBreed(ctx context.Context, breedID string, version_groupd_id string) ([]*pokemon.BreedMove, error) {
	var bms []*pokemon.BreedMove

	err := ps.DB.ModelContext(ctx, &bms).
		Where("breed_id = ?", breedID).
		Where("version_group_id = ?", "16").
		Select()

	if err != nil {
		return nil, err
	}

	return bms, nil
}
