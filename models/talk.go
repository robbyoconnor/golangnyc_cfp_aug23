package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Talk struct {
	ID          uuid.UUID `json:"id" db:"id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
	Title       string    `json:"title" db:"title"`
	Abstract    string    `json:"abstract" db:"abstract"`
	Description string    `json:"description" db:"description"`
	UserID      uuid.UUID `json:"user_id" db:"user_id"`
}

// String is not required by pop and may be deleted
func (t Talk) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Talks is not required by pop and may be deleted
type Talks []Talk

// String is not required by pop and may be deleted
func (t Talks) String() string {
	jt, _ := json.Marshal(t)
	return string(jt)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (t *Talk) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: t.Title, Name: "Title"},
		&validators.StringIsPresent{Field: t.Abstract, Name: "Abstract"},
		&validators.StringIsPresent{Field: t.Description, Name: "Description"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (t *Talk) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (t *Talk) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
