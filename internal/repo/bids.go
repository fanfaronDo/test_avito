package repo

import (
	"context"
	"database/sql"
	"fmt"
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

func (b *BidsRepo) GetBids(limit, offset int, userUUID string) ([]domain.Bid, error) {
	var bids []domain.Bid
	query := `SELECT id, name, status, author_type, author_id, version, created_at 
				FROM bids WHERE author_id = $1 LIMIT $2 OFFSET $3;`

	ctx, cencelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cencelFn()

	rows, err := b.db.QueryContext(ctx, query, userUUID, limit, offset)
	if err != nil {
		return []domain.Bid{}, err
	}

	for rows.Next() {
		var bid domain.Bid
		err = rows.Scan(
			&bid.ID,
			&bid.Name,
			&bid.Status,
			&bid.AuthorType,
			&bid.AuthorID,
			&bid.Version,
			&bid.CreatedAt)

		if err != nil {
			return []domain.Bid{}, ErrBidsNotFound
		}
		bids = append(bids, bid)
	}

	return bids, nil
}

func (r *BidsRepo) GetBidsByTenderId(limit, offset int, tenderUUID string) ([]domain.Bid, error) {
	var bids []domain.Bid
	query := `SELECT id, name, status, author_type, author_id, version, created_at
	FROM bids WHERE tender_id = $1 LIMIT $2 OFFSET $3;`

	ctx, cencelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cencelFn()

	rows, err := r.db.QueryContext(ctx, query, tenderUUID, limit, offset)
	if err != nil {
		return []domain.Bid{}, err
	}

	for rows.Next() {
		var bid domain.Bid
		err = rows.Scan(
			&bid.ID,
			&bid.Name,
			&bid.Status,
			&bid.AuthorType,
			&bid.AuthorID,
			&bid.Version,
			&bid.CreatedAt)

		if err != nil {
			return []domain.Bid{}, ErrBidsNotFound
		}
		fmt.Println(bid)
		bids = append(bids, bid)
	}
	fmt.Println(bids)
	return bids, nil
}
