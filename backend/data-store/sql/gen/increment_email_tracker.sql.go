// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: increment_email_tracker.sql

package queries

import (
	"context"
)

const incrementEmailTracker = `-- name: IncrementEmailTracker :exec
update emails.sent_for_one_rating
set
    emails_sent = emails_sent + 1
where
    assessor_id = $1
    and
    submission_id = $2
    and
    rating_type = $3
returning assessor_id, submission_id, rating_type, emails_sent
`

type IncrementEmailTrackerParams struct {
	AssessorID   int32
	SubmissionID int32
	RatingType   ProjectRatingType
}

func (q *Queries) IncrementEmailTracker(ctx context.Context, arg IncrementEmailTrackerParams) error {
	_, err := q.db.Exec(ctx, incrementEmailTracker, arg.AssessorID, arg.SubmissionID, arg.RatingType)
	return err
}
