package request

import "net/http"

// GetUserID returns the user ID from the request header.
func GetUserID(r *http.Request) string {
	return r.Header.Get("X-User-ID")
}

// SetUserID sets the user ID in the request header.
func SetUserID(r *http.Request, userID string) {
	r.Header.Set("X-User-ID", userID)
}

// GetUserFullName returns the user full name from the request header.
func GetUserFullName(r *http.Request) string {
	return r.Header.Get("X-User-FullName")
}

// SetUserFullName sets the user full name in the request header.
func SetUserFullName(r *http.Request, fullName string) {
	r.Header.Set("X-User-FullName", fullName)
}

// GetUserPhone returns the user phone from the request header.
func GetUserPhone(r *http.Request) string {
	return r.Header.Get("X-User-Phone")
}

// SetUserPhone sets the user phone in the request header.
func SetUserPhone(r *http.Request, phone string) {
	r.Header.Set("X-User-Phone", phone)
}
