// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package queries

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
)

type ProjectRatingType string

const (
	ProjectRatingTypeIndividual ProjectRatingType = "individual"
	ProjectRatingTypeInitial    ProjectRatingType = "initial"
	ProjectRatingTypeFinal      ProjectRatingType = "final"
)

func (e *ProjectRatingType) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProjectRatingType(s)
	case string:
		*e = ProjectRatingType(s)
	default:
		return fmt.Errorf("unsupported scan type for ProjectRatingType: %T", src)
	}
	return nil
}

type NullProjectRatingType struct {
	ProjectRatingType ProjectRatingType
	Valid             bool // Valid is true if ProjectRatingType is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProjectRatingType) Scan(value interface{}) error {
	if value == nil {
		ns.ProjectRatingType, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProjectRatingType.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProjectRatingType) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProjectRatingType), nil
}

type ProjectState string

const (
	ProjectStateDraft              ProjectState = "draft"
	ProjectStateSubmitted          ProjectState = "submitted"
	ProjectStateAccepted           ProjectState = "accepted"
	ProjectStateAcceptedIndividual ProjectState = "accepted_individual"
	ProjectStateAcceptedInitial    ProjectState = "accepted_initial"
	ProjectStateAcceptedFinal      ProjectState = "accepted_final"
	ProjectStateRejected           ProjectState = "rejected"
)

func (e *ProjectState) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ProjectState(s)
	case string:
		*e = ProjectState(s)
	default:
		return fmt.Errorf("unsupported scan type for ProjectState: %T", src)
	}
	return nil
}

type NullProjectState struct {
	ProjectState ProjectState
	Valid        bool // Valid is true if ProjectState is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullProjectState) Scan(value interface{}) error {
	if value == nil {
		ns.ProjectState, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.ProjectState.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullProjectState) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.ProjectState), nil
}

type EditionAwardsRepresentativeContest struct {
	AwardsRepresentativeContestID int32
	AwardsRepresentativeID        int32
	ContestID                     int32
}

type EditionContest struct {
	ContestID                   int32
	Year                        int32
	MasterJuryID                int32
	EstTimeIndividualAssessment sql.NullTime
	EstTimeInitialAssessment    sql.NullTime
	EstTimeFinalAssessment      sql.NullTime
	EstTimeJuryQuestions        sql.NullTime
	MinProjectDurationDays      int32
	MinParticipantTeamSize      int32
	MinSubcontractors           int32
	MaxProjectCompletionMonths  int32
	UrlCodeOfConduct            sql.NullString
	UrlSchedule                 sql.NullString
	UrlFlyer                    sql.NullString
	UrlFinalists                sql.NullString
	UrlResults                  sql.NullString
}

type EditionJuryMemberContest struct {
	JuryMemberContestID int32
	JuryMemberID        int32
	ContestID           int32
}

type EditionPemCriterion struct {
	PemCriterionID int32
	ContestID      int32
	Name           string
	Description    string
	Area           string
	Criteria       string
	Subcriteria    sql.NullString
}

type PersonApplicant struct {
	ApplicantID int32
	PersonID    int32
}

type PersonAssessor struct {
	AssessorID   int32
	IpmaExpertID int32
}

type PersonAwardsRepresentative struct {
	AwardsRepresentativeID int32
	PersonID               int32
}

type PersonBase struct {
	PersonID  int32
	FirstName string
	LastName  string
	Email     string
}

type PersonIpmaExpert struct {
	IpmaExpertID int32
	PersonID     int32
}

type PersonJuryMember struct {
	JuryMemberID int32
	PersonID     int32
}

type ProjectApplicantSubmission struct {
	ApplicantSubmissionID int32
	ApplicantID           int32
	SubmissionID          int32
}

type ProjectApplicationReport struct {
	ApplicationReportID   int32
	SubmissionID          int32
	IsDraft               bool
	ReportSubmissionDate  sql.NullTime
	ProjectGoals          sql.NullString
	OrganisationStructure sql.NullString
	DivisionOfWork        sql.NullString
	ProjectSchedule       sql.NullString
	Attatchments          sql.NullString
}

type ProjectAssessorSubmission struct {
	AssessorSubmissionID int32
	AssessorID           int32
	SubmissionID         int32
	IsLeading            bool
}

type ProjectAssessorsAnswer struct {
	AssessorsAnswerID int32
	JuryQuestionID    int32
	AssessorID        int32
	Answer            sql.NullString
	Files             sql.NullString
}

type ProjectIpmaExpertSubmission struct {
	IpmaExpertSubmissionID int32
	IpmaExpertID           int32
	SubmissionID           int32
}

type ProjectJuryQuestion struct {
	JuryQuestionID int32
	SubmissionID   int32
	CriterionID    int32
	Question       string
	IsDraft        bool
}

type ProjectPartialRating struct {
	PartialRatingID int32
	RatingID        int32
	CriterionID     int32
	Points          int32
	Justification   string
	Modified        time.Time
	ModifiedByID    int32
}

type ProjectRating struct {
	RatingID                int32
	SubmissionID            int32
	AssessorID              int32
	IsDraft                 bool
	Type                    ProjectRatingType
	CustomEstAssessmentTime sql.NullTime
}

type ProjectSubmission struct {
	SubmissionID         int32
	ContestID            int32
	Name                 string
	DurationDays         int32
	TeamSize             int32
	FinishDate           time.Time
	Status               ProjectState
	Budget               pgtype.Numeric
	Description          string
	ReferenceLetter      sql.NullString
	Photos               sql.NullString
	Video                sql.NullString
	Logo                 sql.NullString
	InitiatorDeclaration sql.NullString
	ApplicantDeclaration sql.NullString
}
