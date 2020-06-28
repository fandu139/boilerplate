package entity

import "time"

// ExampleRequest ...
type ExampleRequest struct {
	ExampleNumber   string `json:"order_number"`
	UserUUID        string `json:"uuid_user"`
	IDExampleType   string `json:"id_order_type"`
	IDExampleStatus string `json:"id_order_status"`
	IDPaymentStatus string `json:"id_payment_status"`
	IDPaymentModel  string `json:"id_payment_model"`
	InquiryNumber   string `json:"inquiry_number"`
	PaymentExample  string `json:"payment_order"`
	NMBank          string `json:"nm_bank"`
}

// ExampleResponses ...
type ExampleResponses struct {
	ID              int        `json:"id_order"`
	UUID            string     `json:"uuid"`
	ExampleNumber   string     `json:"order_number"`
	UserUUID        string     `json:"uuid_user"`
	IDExampleType   int64      `json:"id_order_type"`
	IDStatusExample int64      `json:"id_order_status"`
	CreatedAt       *time.Time `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

// Pagination ...
type Pagination struct {
	Limit int `form:"limit" json:"limit"`
	Page  int `form:"page" json:"page"`
}
