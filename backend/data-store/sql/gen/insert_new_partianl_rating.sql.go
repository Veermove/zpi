// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: insert_new_partianl_rating.sql

package queries

import (
	"context"
)

const createPartialRating = `-- name: CreatePartialRating :one
insert into project.partial_rating (
    "rating_id",
    "criterion_id",

    "points",
    "justification",

    "modified",
    "modified_by_id"
) values (
    $1,
    $2,
    $3,
    $4,
    current_timestamp,
    $5
) returning partial_rating_id, rating_id, criterion_id, points, justification, modified, modified_by_id
`

type CreatePartialRatingParams struct {
	RatingID      int32
	CriterionID   int32
	Points        int32
	Justification string
	ModifiedByID  int32
}

func (q *Queries) CreatePartialRating(ctx context.Context, arg CreatePartialRatingParams) (ProjectPartialRating, error) {
	row := q.db.QueryRow(ctx, createPartialRating,
		arg.RatingID,
		arg.CriterionID,
		arg.Points,
		arg.Justification,
		arg.ModifiedByID,
	)
	var i ProjectPartialRating
	err := row.Scan(
		&i.PartialRatingID,
		&i.RatingID,
		&i.CriterionID,
		&i.Points,
		&i.Justification,
		&i.Modified,
		&i.ModifiedByID,
	)
	return i, err
}
