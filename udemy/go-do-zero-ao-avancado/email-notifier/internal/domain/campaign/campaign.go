package campaign

import (
	"time"

	internalerrors "github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/internal-errors"
	"github.com/rs/xid"
)

const (
	Pending string = "Pending"
	Started string = "Started"
	Done    string = "Done"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {
	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index] = Contact{Email: email}
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
	}

	err := internalerrors.ValidateStruct(campaign)

	if err != nil {
		return nil, err
	}

	return campaign, nil
}
