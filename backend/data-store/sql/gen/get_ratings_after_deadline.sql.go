// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: get_ratings_after_deadline.sql

package queries

import (
	"context"
)

const getDelayedRatings = `-- name: GetDelayedRatings :many
select
    rating.rating_id
from project.rating as rating
inner join project.submission as submission using (submission_id)
inner join edition.contest    as contest    using (contest_id)
where
    rating.is_draft = true
    and
    coalesce(
        rating.custom_est_assessment_time,
        case rating.type
            when 'individual' then contest.est_time_individual_assessment
            when 'initial'    then contest.est_time_initial_assessment
            when 'final'      then contest.est_time_final_assessment
            else '9999-12-31'::date -- should never happen
        end
    ) < now()
`

func (q *Queries) GetDelayedRatings(ctx context.Context) ([]int32, error) {
	rows, err := q.db.Query(ctx, getDelayedRatings)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int32
	for rows.Next() {
		var rating_id int32
		if err := rows.Scan(&rating_id); err != nil {
			return nil, err
		}
		items = append(items, rating_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
