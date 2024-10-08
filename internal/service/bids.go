package service

import (
	"github.com/fanfaronDo/test_avito/internal/domain"
	"github.com/fanfaronDo/test_avito/internal/repo"
)

const (
	created   = "Created"
	published = "Published"
	canceled  = "Canceled"
)

type BidsService struct {
	repo *repo.Repository
}

func NewBidService(repo *repo.Repository) *BidsService {
	return &BidsService{repo: repo}
}

func (b *BidsService) CreateBid(bidCreator domain.BidCreator) (domain.Bid, error) {
	bid := initBids(&bidCreator)
	return b.repo.Bid.CreateBids(bidCreator.TenderID, bidCreator.Description, bid)
}

func (b *BidsService) GetBids(limit, offset int, userUUID string) ([]domain.Bid, error) {
	return b.repo.Bid.GetBids(limit, offset, userUUID)
}

func (b *BidsService) GetBidsByTenderId(limit, offset int, tenderUUID string) ([]domain.Bid, error) {
	return b.repo.Bid.GetBidsByTenderId(limit, offset, tenderUUID)
}

func initBids(creator *domain.BidCreator) domain.Bid {
	return domain.Bid{
		Name:       creator.Name,
		Status:     created,
		AuthorType: creator.AuthorType,
		AuthorID:   creator.AuthorID,
		Version:    1,
	}
}
