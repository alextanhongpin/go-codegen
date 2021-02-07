// Code generated by generator, DO NOT EDIT.
package model

import (
	uuid "github.com/google/uuid"
	"net"
	"time"
)

type User struct {
	allowPasswordChange    bool
	birthdate              *time.Time
	confirmationSentAt     *time.Time
	confirmationToken      *string
	confirmedAt            *time.Time
	country                string
	createdAt              time.Time
	currentSignInAt        *time.Time
	currentSignInIp        net.IP
	currentSignInUserAgent string
	deletedAt              *time.Time
	email                  *string
	emailVerified          bool
	familyName             string
	gender                 string
	givenName              string
	id                     uuid.UUID
	lastSignInAt           *time.Time
	lastSignInIp           net.IP
	lastSignInUserAgent    string
	lastSignOutAt          *time.Time
	lastSignOutIp          net.IP
	lastSignOutUserAgent   string
	locale                 string
	locality               string
	middleName             string
	name                   string
	nickname               string
	phoneNumber            string
	phoneNumberVerified    bool
	picture                string
	postalCode             string
	preferredUsername      string
	profile                string
	region                 string
	resetPasswordSentAt    *time.Time
	resetPasswordToken     *string
	signInCount            int32
	streetAddress          string
	unconfirmedEmail       string
	updatedAt              time.Time
	website                string
	zoneinfo               string
}

func NewUser(allowPasswordChange bool, birthdate *time.Time, confirmationSentAt *time.Time, confirmationToken *string, confirmedAt *time.Time, country string, createdAt time.Time, currentSignInAt *time.Time, currentSignInIp net.IP, currentSignInUserAgent string, deletedAt *time.Time, email *string, emailVerified bool, familyName string, gender string, givenName string, id uuid.UUID, lastSignInAt *time.Time, lastSignInIp net.IP, lastSignInUserAgent string, lastSignOutAt *time.Time, lastSignOutIp net.IP, lastSignOutUserAgent string, locale string, locality string, middleName string, name string, nickname string, phoneNumber string, phoneNumberVerified bool, picture string, postalCode string, preferredUsername string, profile string, region string, resetPasswordSentAt *time.Time, resetPasswordToken *string, signInCount int32, streetAddress string, unconfirmedEmail string, updatedAt time.Time, website string, zoneinfo string) {
	return &User{
		allowPasswordChange:    allowPasswordChange,
		birthdate:              birthdate,
		confirmationSentAt:     confirmationSentAt,
		confirmationToken:      confirmationToken,
		confirmedAt:            confirmedAt,
		country:                country,
		createdAt:              createdAt,
		currentSignInAt:        currentSignInAt,
		currentSignInIp:        currentSignInIp,
		currentSignInUserAgent: currentSignInUserAgent,
		deletedAt:              deletedAt,
		email:                  email,
		emailVerified:          emailVerified,
		familyName:             familyName,
		gender:                 gender,
		givenName:              givenName,
		id:                     id,
		lastSignInAt:           lastSignInAt,
		lastSignInIp:           lastSignInIp,
		lastSignInUserAgent:    lastSignInUserAgent,
		lastSignOutAt:          lastSignOutAt,
		lastSignOutIp:          lastSignOutIp,
		lastSignOutUserAgent:   lastSignOutUserAgent,
		locale:                 locale,
		locality:               locality,
		middleName:             middleName,
		name:                   name,
		nickname:               nickname,
		phoneNumber:            phoneNumber,
		phoneNumberVerified:    phoneNumberVerified,
		picture:                picture,
		postalCode:             postalCode,
		preferredUsername:      preferredUsername,
		profile:                profile,
		region:                 region,
		resetPasswordSentAt:    resetPasswordSentAt,
		resetPasswordToken:     resetPasswordToken,
		signInCount:            signInCount,
		streetAddress:          streetAddress,
		unconfirmedEmail:       unconfirmedEmail,
		updatedAt:              updatedAt,
		website:                website,
		zoneinfo:               zoneinfo,
	}
}

func (u User) AllowPasswordChange() bool {
	return u.allowPasswordChange
}

func (u User) Birthdate() *time.Time {
	return u.birthdate
}

func (u User) ConfirmationSentAt() *time.Time {
	return u.confirmationSentAt
}

func (u User) ConfirmationToken() *string {
	return u.confirmationToken
}

func (u User) ConfirmedAt() *time.Time {
	return u.confirmedAt
}

func (u User) Country() string {
	return u.country
}

func (u User) CreatedAt() time.Time {
	return u.createdAt
}

func (u User) CurrentSignInAt() *time.Time {
	return u.currentSignInAt
}

func (u User) CurrentSignInIp() net.IP {
	return u.currentSignInIp
}

func (u User) CurrentSignInUserAgent() string {
	return u.currentSignInUserAgent
}

func (u User) DeletedAt() *time.Time {
	return u.deletedAt
}

func (u User) Email() *string {
	return u.email
}

func (u User) EmailVerified() bool {
	return u.emailVerified
}

func (u User) FamilyName() string {
	return u.familyName
}

func (u User) Gender() string {
	return u.gender
}

func (u User) GivenName() string {
	return u.givenName
}

func (u User) ID() uuid.UUID {
	return u.id
}

func (u User) LastSignInAt() *time.Time {
	return u.lastSignInAt
}

func (u User) LastSignInIp() net.IP {
	return u.lastSignInIp
}

func (u User) LastSignInUserAgent() string {
	return u.lastSignInUserAgent
}

func (u User) LastSignOutAt() *time.Time {
	return u.lastSignOutAt
}

func (u User) LastSignOutIp() net.IP {
	return u.lastSignOutIp
}

func (u User) LastSignOutUserAgent() string {
	return u.lastSignOutUserAgent
}

func (u User) Locale() string {
	return u.locale
}

func (u User) Locality() string {
	return u.locality
}

func (u User) MiddleName() string {
	return u.middleName
}

func (u User) Name() string {
	return u.name
}

func (u User) Nickname() string {
	return u.nickname
}

func (u User) PhoneNumber() string {
	return u.phoneNumber
}

func (u User) PhoneNumberVerified() bool {
	return u.phoneNumberVerified
}

func (u User) Picture() string {
	return u.picture
}

func (u User) PostalCode() string {
	return u.postalCode
}

func (u User) PreferredUsername() string {
	return u.preferredUsername
}

func (u User) Profile() string {
	return u.profile
}

func (u User) Region() string {
	return u.region
}

func (u User) ResetPasswordSentAt() *time.Time {
	return u.resetPasswordSentAt
}

func (u User) ResetPasswordToken() *string {
	return u.resetPasswordToken
}

func (u User) SignInCount() int32 {
	return u.signInCount
}

func (u User) StreetAddress() string {
	return u.streetAddress
}

func (u User) UnconfirmedEmail() string {
	return u.unconfirmedEmail
}

func (u User) UpdatedAt() time.Time {
	return u.updatedAt
}

func (u User) Website() string {
	return u.website
}

func (u User) Zoneinfo() string {
	return u.zoneinfo
}