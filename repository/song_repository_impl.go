package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-crud-api/helper"
	"simple-crud-api/model"
)

type SongRepositoryImpl struct {
	Db *sql.DB
}

func NewSongRepository(Db *sql.DB) SongRepository {
	return &SongRepositoryImpl{Db: Db}
}

// Save implements SongRepository.
func (s *SongRepositoryImpl) Save(ctx context.Context, song model.Song) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "insert into song (name) values($1)"
	_, err = tx.ExecContext(ctx, SQL, song.Name)
	helper.PanicIfError(err)
}

// Update implements SongRepository.
func (s *SongRepositoryImpl) Update(ctx context.Context, song model.Song) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "UPDATE SONG set name = $1 where id=$2"
	_, err = tx.ExecContext(ctx, SQL, song.Name, song.Id)
	helper.PanicIfError(err)
}

// Delete implements SongRepository.
func (s *SongRepositoryImpl) Delete(ctx context.Context, songId int) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "DELETE FROM SONG where id = $1"
	_, errExec := tx.ExecContext(ctx, SQL, songId)
	helper.PanicIfError(errExec)
}

// FindAll implements SongRepository.
func (s *SongRepositoryImpl) FindAll(ctx context.Context) []model.Song {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name FROM SONG"
	result, errQuery := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var songs []model.Song

	for result.Next() {
		song := model.Song{}
		err := result.Scan(&song.Id, &song.Name)
		helper.PanicIfError(err)

		songs = append(songs, song)
	}

	return songs
}

// FindById implements SongRepository.
func (s *SongRepositoryImpl) FindById(ctx context.Context, songId int) (model.Song, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	SQL := "SELECT id, name FROM BOOK where id = $1"
	result, errQuery := tx.QueryContext(ctx, SQL, songId)
	helper.PanicIfError(errQuery)
	defer result.Close()

	song := model.Song{}

	if result.Next() {
		err := result.Scan(&song.Id, &song.Name)
		helper.PanicIfError(err)
		return song, nil
	} else {
		return song, errors.New("song id not found")
	}
}