package campaign

import (
	"time"

	internalerrors "github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/internal-errors"
	"github.com/rs/xid"
)

type Contacts struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string     `validate:"required"`
	Name      string     `validate:"min=5,max=24"`
	CreatedOn time.Time  `validate:"required"`
	Content   string     `validate:"min=5,max=1024"`
	Contacts  []Contacts `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contacts, len(emails))

	for index, email := range emails {
		contacts[index] = Contacts{Email: email}
	}

	campagin := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

	err := internalerrors.ValidateStruct(campagin)

	if err != nil {
		return nil, err
	}

	return campagin, nil
}
