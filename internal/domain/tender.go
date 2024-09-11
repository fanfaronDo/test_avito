package domain

import "time"

type Tender struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	ServiceType string    `json:"service_type"`
	Version     int       `json:"version"`
	CreatedAt   time.Time `json:"created_at"`
}

type TenderCreator struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	Status          string `json:"status"`
	ServiceType     string `json:"serviceType"`
	OrganizationID  string `json:"organizationId"`
	CreatorUsername string `json:"creatorUsername"`
}

type TenderEditor struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ServiceType string `json:"serviceType"`
}
