package service

import (
	"context"
	"simple-crud-api/data/request"
	"simple-crud-api/data/response"
)

type SongService interface {
	Create(ctx context.Context, request request.SongCreateRequest)
	Update(ctx context.Context, request request.SongUpdateRequest)
	Delete(ctx context.Context, songId int)
	FindById(ctx context.Context, songId int) response.SongResponse
	FindAll(ctx context.Context) []response.SongResponse
}