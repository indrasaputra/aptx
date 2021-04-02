package entity

import "time"

// URL defines the logical data of an URL.
type URL struct {
	// Code is a unique code for short URL.
	// ShortURL is built using this code, prefixed by domain.
	Code string
	// ShortURL is a shortened version of original URL created by the system.
	ShortURL string
	// OriginalURL is a raw URL from user.
	OriginalURL string
	// ExpireAt defines the expire time of the URL.
	// URL can only be accessed if current time is still smaller than expired time.
	ExpiredAt time.Time
	// CreatedAt defines when the URL was created.
	CreatedAt time.Time
}
