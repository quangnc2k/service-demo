package persistance

// DefaultRepository is a global variable, used when connection to a repo is needed
var DefaultRepository *Repository

// Database is a struct containing all repo
type Repository struct {
	//List of all repo
	StaffRepository StaffRepository
}
