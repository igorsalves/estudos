package campaign

import (
	"errors"

	"github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/contract"
	internalerrors "github.com/igorsalves/estudos/tree/main/udemy/go-do-zero-ao-avancado/email-notifier/internal/internal-errors"
)

type Service interface {
	Create(newCampaign contract.NewCampaign) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
	Cancel(id string) error
	Delete(id string) error
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaign) (string, error) {
	campaign, err := NewCampaign(
		newCampaign.Name,
		newCampaign.Content,
		newCampaign.Emails,
	)

	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaign)

	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return nil, internalerrors.ErrInternal
	}

	if campaign == nil {
		return nil, errors.New("Campaign not found")
	}

	return &contract.CampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil
}

func (s *ServiceImp) Cancel(id string) error {

	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return internalerrors.ErrInternal
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.Cancel()
	err = s.Repository.Update(campaign)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}

func (s *ServiceImp) Delete(id string) error {
	campaign, err := s.Repository.GetBy(id)

	if err != nil {
		return internalerrors.ErrInternal
	}

	if campaign.Status != Pending {
		return errors.New("Campaign status invalid")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)
	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
