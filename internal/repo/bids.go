package repo

import (
	"context"
	"database/sql"
	"github.com/fanfaronDo/test_avito/internal/domain"
	log "github.com/sirupsen/logrus"
	"time"
)

type BidsRepo struct {
	db *sql.DB
}

func NewBidsRepo(db *sql.DB) *BidsRepo {
	return &BidsRepo{db}
}

func (b *BidsRepo) CreateBids(tenderUUID, descr string, bid domain.Bid) (domain.Bid, error) {
	var uuid string
	query := `INSERT INTO bids (name, description, status, tender_id, author_type, author_id) 
			  VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	ctx, cencelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cencelFn()
	err := b.db.QueryRowContext(ctx, query,
		bid.Name,
		descr,
		bid.Status,
		tenderUUID,
		bid.AuthorType,
		bid.AuthorID).Scan(&uuid)

	if err != nil {
		log.Debugf("%s: %v", ErrCreationBid, err)
		return domain.Bid{}, ErrCreationBid
	}
	bid.ID = uuid

	return bid, nil
}
