package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contacts struct {
	Email string
}

type Campaign struct {
	ID        string
	Name      string
	CreatedOn time.Time
	Content   string
	Contacts  []Contacts
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	if name == "" {
		return nil, errors.New("name is required")
	}

	if content == "" {
		return nil, errors.New("content is required")
	}

	if len(emails) == 0 {
		return nil, errors.New("at least one contact is required")
	}

	contacts := make([]Contacts, len(emails))
	for index, email := range emails {
		contacts[index] = Contacts{Email: email}
	}

	return &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}, nil
}
