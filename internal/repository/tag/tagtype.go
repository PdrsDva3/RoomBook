package tag

import (
	"RoomBook/internal/models"
	"RoomBook/internal/repository"
	"RoomBook/pkg/cerr"
	"context"
	"github.com/jmoiron/sqlx"
)

type RepoTagType struct {
	db *sqlx.DB
}

func InitTagTypeRepository(db *sqlx.DB) repository.TagTypeRepo {
	return RepoTagType{db: db}
}

func (r RepoTagType) CreateType(ctx context.Context, types models.TypeCreate) (int, error) {
	var id int
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return 0, cerr.Err(cerr.Transaction, err).Error()
	}
	row := transaction.QueryRowContext(ctx, `INSERT INTO tag_type (type) VALUES ($1) returning id;`,
		types.Type)

	err = row.Scan(&id)
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Scan, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return 0, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return 0, cerr.Err(cerr.Commit, err).Error()
	}
	return id, nil
}

func (r RepoTagType) CreateTag(ctx context.Context, tag models.TagCreate) (*models.Tag, error) {
	var NewTag models.Tag
	transaction, err := r.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, cerr.Err(cerr.Transaction, err).Error()
	}
	row := r.db.QueryRowContext(ctx, `SELECT type from tag_type WHERE id = $1;`, tag.IDType)
	err = row.Scan(&NewTag.Type)
	if err != nil {
		return nil, cerr.Err(cerr.Scan, err).Error()
	}
	row = transaction.QueryRowContext(ctx, `INSERT INTO tags (id_type, type, tag) VALUES ($1, $2, $3) returning id;`,
		tag.IDType, NewTag.Type, tag.Tag)

	err = row.Scan(&NewTag.IDTag)
	NewTag.Tag = tag.Tag
	NewTag.IDType = tag.IDType
	if err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return nil, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return nil, cerr.Err(cerr.Scan, err).Error()
	}
	if err = transaction.Commit(); err != nil {
		if rbErr := transaction.Rollback(); rbErr != nil {
			return nil, cerr.Err(cerr.Rollback, rbErr).Error()
		}
		return nil, cerr.Err(cerr.Commit, err).Error()
	}
	return &NewTag, nil
}

func (r RepoTagType) Tags(ctx context.Context) ([]models.Tag, error) {
	var tags []models.Tag
	rows, err := r.db.QueryContext(ctx, `SELECT ID, ID_TYPE, TYPE, TAG from tags;`)
	if err != nil {
		return nil, cerr.Err(cerr.Rows, err).Error()
	}
	for rows.Next() {
		var tag models.Tag
		err = rows.Scan(&tag.IDTag, &tag.IDType, &tag.Type, &tag.Tag)
		if err != nil {
			return nil, cerr.Err(cerr.Scan, err).Error()
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func (r RepoTagType) Types(ctx context.Context) ([]models.TypeBase, error) {
	var types []models.TypeBase
	rows, err := r.db.QueryContext(ctx, `SELECT ID, TYPE from tag_type;`)
	if err != nil {
		return nil, cerr.Err(cerr.Rows, err).Error()
	}
	for rows.Next() {
		var type1 models.TypeBase
		err = rows.Scan(&type1.IDType, &type1.Type)
		if err != nil {
			return nil, cerr.Err(cerr.Scan, err).Error()
		}
		types = append(types, type1)
	}
	return types, nil
}

func (r RepoTagType) TagsType(ctx context.Context, idType int) (*models.TagsType, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT id, id_type, type, tag from tags where id_type=$1;`, idType)
	if err != nil {
		return nil, cerr.Err(cerr.Rows, err).Error()
	}
	var type1 models.TagsType
	var tags []models.TagBase
	for rows.Next() {
		var tag models.TagBase
		err = rows.Scan(&tag.IDTag, &type1.IDType, &type1.Type, &tag.Tag)
		if err != nil {
			return nil, cerr.Err(cerr.Scan, err).Error()
		}
		tags = append(tags, tag)
	}
	type1.Tags = tags

	return &type1, nil
}
