package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Project struct {
	ID            uuid.UUID     `json:"id" db:"id"`
	CreatedAt     time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at" db:"updated_at"`
	DisplayName   string        `json:"display_name" db:"display_name"`
	FirstColor    string        `json:"first_color" db:"first_color"`
	SecondColor   nulls.String  `json:"second_color" db:"second_color"`
	Description   string        `json:"description" db:"description"`
	Logo          string        `json:"logo" db:"logo"`
	Link          string        `json:"link" db:"link"`
	SetupDuration nulls.String  `json:"setup_duration" db:"setup_duration"`
	IssuesCount   int           `json:"issues_count" db:"issues_count"`
	HFOSS         bool			`json:"HFOSS" db:"HFOSS"`	
	Tags          slices.String `json:"tags" db:"tags"`
}

// String is not required by pop and may be deleted
func (p Project) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Projects is not required by pop and may be deleted
type Projects []Project

// String is not required by pop and may be deleted
func (p Projects) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Project) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: p.DisplayName, Name: "DisplayName"},
		&validators.StringIsPresent{Field: p.Description, Name: "Description"},
		&validators.StringIsPresent{Field: p.Logo, Name: "Logo"},
		&validators.StringIsPresent{Field: p.Link, Name: "Link"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Project) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Project) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
