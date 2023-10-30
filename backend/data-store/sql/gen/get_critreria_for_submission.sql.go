// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: get_critreria_for_submission.sql

package queries

import (
	"context"
	"database/sql"
)

const getCriteriaForSubmission = `-- name: GetCriteriaForSubmission :many
select
    criteria.pem_criterion_id,
    criteria.name,
    criteria.description,
    criteria.area,
    criteria.criteria,
    criteria.subcriteria
from project.submission as proj_submission
inner join edition.contest as contest on proj_submission.contest_id = contest.contest_id
inner join edition.pem_criterion as criteria on contest.contest_id = criteria.contest_id
where proj_submission.submission_id = $1
`

type GetCriteriaForSubmissionRow struct {
	PemCriterionID int32
	Name           string
	Description    string
	Area           string
	Criteria       string
	Subcriteria    sql.NullString
}

func (q *Queries) GetCriteriaForSubmission(ctx context.Context, submissionID int32) ([]GetCriteriaForSubmissionRow, error) {
	rows, err := q.db.Query(ctx, getCriteriaForSubmission, submissionID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []GetCriteriaForSubmissionRow
	for rows.Next() {
		var i GetCriteriaForSubmissionRow
		if err := rows.Scan(
			&i.PemCriterionID,
			&i.Name,
			&i.Description,
			&i.Area,
			&i.Criteria,
			&i.Subcriteria,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
