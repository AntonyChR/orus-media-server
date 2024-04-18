package repositories

// Define a common interface that contains methods shared by the repositories.
type Repository interface {
	DropDatabase() error
	Migrate() error
}
