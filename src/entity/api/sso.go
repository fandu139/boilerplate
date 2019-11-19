package api

import "google.golang.org/api/googleapi"

// SsoGoogleRequest mapping
type SsoGoogleRequest struct {
	IDToken string `json:"id_token" form:"id_token"`
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Image   string `json:"image" form:"image"`
}

// GoogleReturnResponse mapping
type GoogleReturnResponse struct {
	AccessType               string `json:"access_type,omitempty"`
	Audience                 string `json:"audience,omitempty"`
	Email                    string `json:"email,omitempty"`
	ExpiresIn                int64  `json:"expires_in,omitempty"`
	IssuedTo                 string `json:"issued_to,omitempty"`
	Scope                    string `json:"scope,omitempty"`
	TokenHandle              string `json:"token_handle,omitempty"`
	UserId                   string `json:"user_id,omitempty"`
	VerifiedEmail            bool   `json:"verified_email,omitempty"`
	googleapi.ServerResponse `json:"-"`
	ForceSendFields          []string `json:"-"`
	NullFields               []string `json:"-"`
}

// SsoRequest mapping
type SsoRequest struct {
	GrantType    string `json:"grant_type" form:"grant_type"`
	ClientSecret string `json:"client_secret" form:"client_secret"`
	ClientKey    string `json:"client_key" form:"client_key"`
}

// ClientCreateRalaliSso mapping
type ClientCreateRalaliSso struct {
	ClientName   string   `json:"client_name"`
	IsFirtsParty bool     `json:"is_first_party"`
	RedirectUris []string `json:"redirect_uris"`
}

// SsoUserRequest mapping
type SsoUserRequest struct {
	GrantType string `json:"grant_type" form:"grant_type"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"password" form:"password"`
}

// SsoAuthByPassword mapping
type SsoAuthByPassword struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

// SsoResponseBeforeUsers mapping for unixtime
type SsoResponseBeforeUsers struct {
	TokenType   string `json:"token_type"`
	ExpiresAt   int64  `json:"expires_at"`
	AccessToken string `json:"access_token"`
}

// SsoRalaliClientCreate mapping
type SsoRalaliClientCreate struct {
	ID           uint   `json:"id"`
	ClientName   uint   `json:"client_name"`
	ClientSecret string `json:"client_secret"`
	ClientKey    string `json:"client_key"`
	ClientPubKey string `json:"client_public_key"`
	IsActive     bool   `json:"is_active"`
	IsFirstParty bool   `json:"is_first_party"`
	RedirectUris string `json:"redirect_uris"`
}

// SsoResponseAfterUsers mapping for unixtime
type SsoResponseAfterUsers struct {
	TokenType    string `json:"token_type"`
	ExpiresAt    int64  `json:"expires_at"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// ProfileOauthResponse mapping
type ProfileOauthResponse struct {
	ApprovedAt       string `json:"approved_at"`
	ConfirmationCode string `json:"confirmation_code"`
	Confirmed        uint   `json:"confirmed"`
	ConfirmedAt      string `json:"confirmed_at"`
	CreatedAt        string `json:"created_at"`
	DefaultShipAddr  uint   `json:"default_shipping_address_id"`
	Email            string `json:"email"`
	GmailAlias       string `json:"gmail_alias"`
	RalaliID         uint   `json:"id"`
	ImageProfile     string `json:"image_profile"`
	Language         string `json:"language"`
	Name             string `json:"name"`
	Password         string `json:"password"`
	Phone            string `json:"phone"`
	PhoneVerified    uint   `json:"phone_verified"`
	ReferallCode     string `json:"referral_code"`
	RememberToken    string `json:"remember_token"`
	UpdateAt         string `json:"updated_at"`
	Status           uint   `json:"status"`
}

// ResultBirthOfDate detail
type ResultBirthOfDate struct {
	Date  string `json:"date"`
	Month string `json:"month"`
	Year  string `json:"year"`
}

// ResultAddresses detail
type ResultAddresses struct {
	SetDefault string         `json:"set_default"`
	Address    string         `json:"address"`
	Lat        string         `json:"lat"`
	Long       string         `json:"long"`
	District   ResultDistrict `json:"district"`
}

// ResultDistrict detail
type ResultDistrict struct {
	Name string     `json:"name"`
	City ResultCity `json:"city"`
}

// ResultCity detail
type ResultCity struct {
	Name     string         `json:"name"`
	Province ResultProvince `json:"province"`
}

// ResultProvince detail
type ResultProvince struct {
	Name string `json:"name"`
}
