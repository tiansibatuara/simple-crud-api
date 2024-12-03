package repository

import (
	"context"
	"simple-crud-api/model"
)

type SongRepository interface {
	Save(ctx context.Context, song model.Song)
	Update(ctx context.Context, song model.Song)
	Delete(ctx context.Context, songId int)
	FindById(ctx context.Context, songId int)(model.Song, error)
	FindAll(ctx context.Context) []model.Song
}