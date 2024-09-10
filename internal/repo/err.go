package repo

import "errors"

var (
	ErrDatabaseConnectionFailed = errors.New("database connection failed")
	ErrUserChargeNotFound       = errors.New("user charge not found")
	ErrUserNotFound             = errors.New("user not found")
	ErrCreationTender           = errors.New("creation tender error ")
	ErrScanDataTender           = errors.New("scan data tender error ")
	ErrTenderNotFound           = errors.New("tender not found")
	ErrTenderStatusNotFound     = errors.New("tenders status not found")
	ErrInFailedTransaction      = errors.New("failed transaction")
	ErrUpdatedTender            = errors.New("updated tender error")
	ErrFetchingTender           = errors.New("fetching tender error ")
	ErrCommittingTransaction    = errors.New("committing tender error ")
)
