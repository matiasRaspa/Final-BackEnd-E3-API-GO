package domain

// Dentist represents a dentist.
type Dentist struct {
	Id       int64  `json:"id"`
	LastName string `json:"last_name" binding:"required"`
	Name     string `json:"name" binding:"required"`
	License  string `json:"license" binding:"required"`
}
