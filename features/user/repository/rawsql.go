package repository

// import (
// 	"be13/ca/features/user"
// 	"database/sql"
// )

// type userRawRepository struct {
// 	db *sql.DB
// }

// func NewRaw(db *sql.DB) user.RepositoryInterface {
// 	return &userRawRepository{
// 		db: db,
// 	}
// }

// // Create implements user.RepositoryInterface
// func (*userRawRepository) Create(input user.Core) (row int, err error) {
// 	panic("unimplemented")
// }

// // GetAll implements user.RepositoryInterface
// func (*userRawRepository) GetAll() (data []user.Core, err error) {
// 	panic("unimplemented")
// }
