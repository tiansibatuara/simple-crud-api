package service

import (
	"context"
	"simple-crud-api/data/request"
	"simple-crud-api/data/response"
	"simple-crud-api/helper"
	"simple-crud-api/model"
	"simple-crud-api/repository"
)

type SongServiceImpl struct {
	SongRepository repository.SongRepository
}

func NewSongRepositoryImpl(songRepository repository.SongRepository) SongService {
	return &SongServiceImpl{SongRepository: songRepository}
}

// Create implements SongService.
func (s *SongServiceImpl) Create(ctx context.Context, request request.SongCreateRequest) {
	song := model.Song{
		Name: request.Name,
	}
	s.SongRepository.Save(ctx, song)
}

// Update implements SongService.
func (s *SongServiceImpl) Update(ctx context.Context, request request.SongUpdateRequest) {
	song, err := s.SongRepository.FindById(ctx, request.Id)
	helper.PanicIfError(err)

	song.Name = request.Name
	s.SongRepository.Update(ctx, song)
}

// Delete implements SongService.
func (s *SongServiceImpl) Delete(ctx context.Context, songId int) {
	song, err := s.SongRepository.FindById(ctx, songId)
	helper.PanicIfError(err)
	s.SongRepository.Delete(ctx, song.Id)
}

// FindAll implements SongService.
func (s *SongServiceImpl) FindAll(ctx context.Context) []response.SongResponse {
	songs := s.SongRepository.FindAll(ctx)

	var songResp []response.SongResponse

	for _, value := range songs {
		book := response.SongResponse{Id: value.Id, Name: value.Name}
		songResp = append(songResp, book)
	}
	return songResp
}

// FindById implements SongService.
func (s *SongServiceImpl) FindById(ctx context.Context, songId int) response.SongResponse {
	song, err := s.SongRepository.FindById(ctx, songId)
	helper.PanicIfError(err)
	return response.SongResponse(song)	
}
