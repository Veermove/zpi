// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: get_submission_details.sql

package queries

import (
	"context"
	"database/sql"
	"time"

	"github.com/jackc/pgtype"
)

const getSubmissionDetails = `-- name: GetSubmissionDetails :one
select
    submission.team_size,
    submission.finish_date,
    submission.status,
    submission.budget,
    submission.description,
    report.is_draft,
    report.report_submission_date,
    report.project_goals,
    report.organisation_structure,
    report.division_of_work,
    report.project_schedule,
    report.attatchments
from project.submission as submission
inner join project.application_report as report on report.submission_id = submission.submission_id
where submission.submission_id = $1
`

type GetSubmissionDetailsRow struct {
	TeamSize              int32
	FinishDate            time.Time
	Status                ProjectState
	Budget                pgtype.Numeric
	Description           string
	IsDraft               bool
	ReportSubmissionDate  sql.NullTime
	ProjectGoals          sql.NullString
	OrganisationStructure sql.NullString
	DivisionOfWork        sql.NullString
	ProjectSchedule       sql.NullString
	Attatchments          sql.NullString
}

func (q *Queries) GetSubmissionDetails(ctx context.Context, submissionID int32) (GetSubmissionDetailsRow, error) {
	row := q.db.QueryRow(ctx, getSubmissionDetails, submissionID)
	var i GetSubmissionDetailsRow
	err := row.Scan(
		&i.TeamSize,
		&i.FinishDate,
		&i.Status,
		&i.Budget,
		&i.Description,
		&i.IsDraft,
		&i.ReportSubmissionDate,
		&i.ProjectGoals,
		&i.OrganisationStructure,
		&i.DivisionOfWork,
		&i.ProjectSchedule,
		&i.Attatchments,
	)
	return i, err
}
