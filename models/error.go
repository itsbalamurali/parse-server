package models

// An Error from the Parse API. When a valid Parse JSON error is found, the
// returned error will be of this type.
type Error struct {
	Message string `json:"error"`
	Code    int    `json:"code"`
}
