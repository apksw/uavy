package model

import (
	"time"

	"gitlab.com/adrianpk/uavy/auth/pkg/base"
)

type (
	// User struct
	User struct {
		base.Identification
		Username          string       `bson:"username"`
		Password          string       `bson:"-"`
		PasswordDigest    string       `bson:"password_digest"`
		Email             string       `bson:"email"`
		EmailConfirmation string       `bson:"-"`
		LastIP            string       `bson:"last_ip"`
		ConfirmationToken string       `bson:"confirmation_token"`
		IsConfirmed       bool         `bson:"is_confirmed"`
		GeoLocation       base.GeoJson `bson:"geolocations"`
		StartsAt          time.Time    `bson:"starts_at"`
		EndsAt            time.Time    `bson:"ends_at"`
		IsActive          bool         `bson:"is_active"`
		IsDeleted         bool         `bson:"is_deleted"`
		base.Audit
	}

	// Auth struct
	Auth struct {
		User
		PermissionTags string `bson:"permission_tags"`
	}
)
