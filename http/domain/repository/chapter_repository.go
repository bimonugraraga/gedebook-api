package repository

import (
	"context"
	"fmt"

	"gedebook.com/api/db"
	"gedebook.com/api/domain"
	"github.com/uptrace/bun"
)

type ChapterRepository interface {
	CreateOneChapter(ctx context.Context, src *domain.Chapter) (err error)
	GetOneChapter(ctx context.Context, book_id int, chapter_id int, chapter_status []domain.ChapterPublishedStatus, user_id int) (domain.Chapter, error)
	UpdateOneChapter(ctx context.Context, book_id int, chapter_id int, src *domain.Chapter) (err error)
}

type chapterRepository struct {
	db *bun.DB
}

func NewChapterRepository() ChapterRepository {
	return &chapterRepository{
		db: db.GetConn(),
	}
}

func (r *chapterRepository) CreateOneChapter(ctx context.Context, src *domain.Chapter) (err error) {
	_, err = r.db.NewInsert().
		Model(src).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}
	return
}

func (r *chapterRepository) GetOneChapter(ctx context.Context, book_id int, chapter_id int, chapter_status []domain.ChapterPublishedStatus, user_id int) (domain.Chapter, error) {
	res := domain.Chapter{}
	if err := r.db.NewSelect().
		Model(&res).
		Where(fmt.Sprintf("Chapter.id = %d AND book_id = %d", chapter_id, book_id)).
		//!User Can Get Their Own Draft, Cancelled, And Published
		Where("Chapter.published_status IN (?)", bun.In(chapter_status)).
		Relation("Book", func(query *bun.SelectQuery) *bun.SelectQuery {
			if user_id != 0 {
				query = query.Where("user_id = ?", user_id)
			} else {
				query = query.Where("Book.published_status = ?", domain.BookPublishedStatusPublished)
			}
			return query
		}).
		Scan(ctx); err != nil {
		return res, err
	}
	return res, nil
}

func (r *chapterRepository) UpdateOneChapter(ctx context.Context, book_id int, chapter_id int, src *domain.Chapter) (err error) {
	_, err = r.db.NewUpdate().
		Model(src).
		Where(fmt.Sprintf("Chapter.id = %d AND book_id = %d", chapter_id, book_id)).
		Returning("*").
		Exec(ctx)
	if err != nil {
		return err
	}
	return
}
