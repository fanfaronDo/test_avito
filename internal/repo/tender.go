package repo

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/fanfaronDo/test_avito/internal/domain"
	"time"
)

const (
	timeuotCtx = 5 * time.Second
)

type TenderRepo struct {
	db *sql.DB
}

func NewTenderRepo(db *sql.DB) *TenderRepo {
	return &TenderRepo{db: db}
}

func (t *TenderRepo) CreateTender(tender domain.Tender) (domain.Tender, error) {
	query := `INSERT INTO tenders (name, description, service_type, status, organization_id, creator_id, version) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	var uuid string
	err := t.db.QueryRowContext(ctx, query,
		tender.Name,
		tender.Description,
		tender.ServiceType,
		tender.Status,
		tender.OrganizationID,
		tender.CreatorID,
		tender.Version).Scan(&uuid)

	if err != nil {
		return domain.Tender{}, fmt.Errorf("Error %s when inserting row into tenders table", err)
	}
	tender.ID = uuid
	return tender, nil
}

func (t *TenderRepo) GetUserUUIDCharge(username, organisation_id string) (string, error) {
	query := `SELECT e.id FROM organization_responsible o 
				LEFT JOIN employee e ON o.user_id = e.id 
				WHERE e.username = $1 AND o.organization_id = $2;`

	var uuid string
	ctx, cancelFn := context.WithTimeout(context.Background(), timeuotCtx)
	defer cancelFn()
	err := t.db.QueryRowContext(ctx, query, username, organisation_id).Scan(&uuid)
	if err != nil {
		return "", fmt.Errorf("Error %s when checking user charge", err)
	}

	return uuid, nil
}

// Get
//exaple
//package main
//
//import (
//"database/sql"
//"fmt"
//"log"
//"net/http"
//"strconv"
//
//_ "github.com/go-sql-driver/mysql"
//"github.com/gorilla/mux"
//)
//
//type Tender struct {
//	ID           int    `json:"id"`
//	Title        string `json:"title"`
//	Description  string `json:"description"`
//	ServiceType  string `json:"service_type"`
//	CreatedAt    string `json:"created_at"`
//}
//
//type PaginatedResponse struct {
//	Total       int       `json:"total"`
//	TotalPages  int       `json:"total_pages"`
//	CurrentPage int       `json:"current_page"`
//	PerPage     int       `json:"per_page"`
//	Data        []Tender  `json:"data"`
//}
//
//func main() {
//	db, err := sql.Open("mysql", "username:password@/database")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer db.Close()
//
//	router := mux.NewRouter()
//	router.HandleFunc("/api/tenders", getTenders(db)).Methods("GET")
//
//	fmt.Println("Server started on port 8000")
//	log.Fatal(http.ListenAndServe(":8000", router))
//}
//
//func getTenders(db *sql.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		serviceType := r.URL.Query().Get("service_type")
//		page, _ := strconv.Atoi(r.URL.Query().Get("page"))
//		if page == 0 {
//			page = 1
//		}
//		limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
//		if limit == 0 {
//			limit = 10
//		}
//		offset := (page - 1) * limit
//
//		var tenders []Tender
//		var total int
//
//		query := "SELECT * FROM tenders"
//		countQuery := "SELECT COUNT(*) AS total FROM tenders"
//		params := []interface{}{}
//
//		if serviceType != "" {
//			query += " WHERE service_type = ?"
//			countQuery += " WHERE service_type = ?"
//			params = append(params, serviceType)
//		}
//
//		query += " LIMIT ? OFFSET ?"
//		params = append(params, limit, offset)
//
//		err := db.QueryRow(countQuery, params...).Scan(&total)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//
//		rows, err := db.Query(query, params...)
//		if err != nil {
//			http.Error(w, err.Error(), http.StatusInternalServerError)
//			return
//		}
//		defer rows.Close()
//
//		for rows.Next() {
//			var tender Tender
//			err := rows.Scan(&tender.ID, &tender.Title, &tender.Description, &tender.ServiceType, &tender.CreatedAt)
//			if err != nil {
//				http.Error(w, err.Error(), http.StatusInternalServerError)
//				return
//			}
//			tenders = append(tenders, tender)
//		}
//
//		totalPages := (total + limit - 1) / limit
//
//		response := PaginatedResponse{
//			Total:       total,
//			TotalPages:  totalPages,
//			CurrentPage: page,
//			PerPage:     limit,
//			Data:        tenders,
//		}
//
//		writeJSON(w, http.StatusOK, response)
//	}
//}
//
//func writeJSON(w http.ResponseWriter, status int, v interface{}) {
//	w.Header().Set("Content-Type", "application/json")
//	w.WriteHeader(status)
//	encoder := json.NewEncoder(w)
//	if err := encoder.Encode(v); err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
