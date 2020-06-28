package entity

// HealthResponse Types
type HealthResponse struct {
	Messages string `json:"messages"`
}

// HealthDetailResponse detail
type HealthDetailResponse struct {
	Database ServiceDetailResponse `json:"database"`
	Redis    ServiceDetailResponse `json:"redis"`
}

// ServiceDetailResponse service detail scheme
type ServiceDetailResponse struct {
	Name   string          `json:"name"`
	Health HealthyResponse `json:"health"`
	Notify bool            `json:"notify"`
	Strict bool            `json:"strict"`
}

// HealthyResponse scheme
type HealthyResponse struct {
	Healthy bool   `json:"healthy"`
	Message string `json:"message"`
}
