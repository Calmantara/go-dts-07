package user

import (
	"context"
	"errors"
	"log"

	"github.com/Calmantara/go-dts-user/config"
	"github.com/Calmantara/go-dts-user/module/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserGormRepoImpl struct {
	db *gorm.DB
}

func NewUserGormRepo(db *gorm.DB) UserRepo {
	userRepo := &UserGormRepoImpl{
		db: db,
	}

	if config.Load.DataSource.Migrate {
		userRepo.doMigration()
	}

	return &UserGormRepoImpl{
		db: db,
	}
}

func (u *UserGormRepoImpl) doMigration() (err error) {
	// create user table
	if err = u.db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}
	log.Println("successfully create user table")
	// create user photo table
	if err = u.db.AutoMigrate(&model.UserPhoto{}); err != nil {
		panic(err)
	}
	log.Println("successfully create user photo table")
	return
}

func (u *UserGormRepoImpl) FindUserByIdEager(ctx context.Context, userId uint64) (user model.User, err error) {
	// Eager
	// akan memasukkan semua value
	// dari relasi ke dalam struct

	// ex: User ada property Photos []UserPhoto `json:"photos"
	// eager loading, dia akan mengquery photos
	// dan langsung dimasukkan ke dalam user struct

	// eager from user to photo
	tx := u.db.
		Model(&model.User{}).
		Preload("Photos").
		Where("id = ?", userId).
		Find(&user)
	if err = tx.Error; err != nil {
		return
	}

	// eager from photo to user
	var userPhoto model.UserPhoto
	tx = u.db.
		Model(&model.UserPhoto{}).
		Preload("UserDetail").
		Where("id = ?", 1).
		Find(&userPhoto)
	if err = tx.Error; err != nil {
		return
	}

	if user.Id <= 0 {
		err = errors.New("user is not found")
	}
	return
}

func (u *UserGormRepoImpl) FindUserByIdJoin(ctx context.Context, userId uint64) (err error) {

	type UserCustom struct {
		Id       uint64 `json:"id" gorm:"column:id"`
		Name     string `json:"name" gorm:"column:name"`
		Email    string `json:"email" gorm:"column:email"`
		PhotoId  uint64 `json:"photo_id" gorm:"column:photo_id"`
		PhotoUrl string `json:"photo_url" gorm:"column:photo_url"`
	}

	// join with new struct type
	var users []UserCustom
	tx := u.db.
		Table("users").
		Select(`
			users.id as id, 
			user_photos.id as photo_id,
			users.name as name,
			users.email as email,
			user_photos.url as photo_url
		`).
		Joins(`JOIN user_photos 
			   ON users.id = user_photos.user_id
			   AND user_photos.deleted_at IS NULL`).
		Where("users.id = ?", userId).
		Find(&users)
	if err = tx.Error; err != nil {
		return
	}

	// eager join
	var userModel []model.User
	tx = u.db.
		Joins(`Photo`).
		Where("users.id = ?", userId).
		Find(&userModel)
	if err = tx.Error; err != nil {
		return
	}

	return
}

func (u *UserGormRepoImpl) FindUserById(ctx context.Context, userId uint64) (user model.User, err error) {

	tx := u.db.
		Model(&model.User{}).
		Where("id = ?", userId).
		Find(&user)
	// ketika kita menggunakan auto migrate
	// di method, kita tidak perlu menambahkan
	// deleted_at is null
	// karena gorm udah auto nambahin itu
	if err = tx.Error; err != nil {
		return
	}

	if user.Id <= 0 {
		err = errors.New("user is not found")
	}

	return
}

func (u *UserGormRepoImpl) FindAllUsers(ctx context.Context) (users []model.User, err error) {

	tx := u.db.
		Model(&model.User{}).
		Find(&users).
		Order("created_at ASC")

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (u *UserGormRepoImpl) InsertUser(ctx context.Context, userIn model.User) (user model.User, err error) {
	// untuk insert user,
	// di gorm menyediakan beberapa method
	// Save
	// Create
	// FirstOrCreate

	tx := u.db.
		Model(&model.User{}).
		Create(&userIn)

	if err = tx.Error; err != nil {
		return
	}

	return userIn, err
}

func (u *UserGormRepoImpl) BulkInsertUser(ctx context.Context, userIn []model.User) (err error) {
	// untuk insert user,
	// di gorm menyediakan beberapa method
	// Save
	// Create
	// FirstOrCreate

	// kita mau validasi
	// sebelum create, kita harus check dulu
	// kalau name admin, dia tidak akan valid
	// dengan gorm, kita bisa menambahkan HOOK
	// HOOK -> suato function yang akan execute
	// sebelum melakukan create

	tx := u.db.
		Model(&model.User{}).
		Create(&userIn)

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (u *UserGormRepoImpl) UpdateUser(ctx context.Context, userIn model.User) (err error) {
	// kalau kita hanya ingin update 1 column
	// bisa menggunakan update
	// multiple column, kita bisa menggunakan Updates

	tx := u.db.
		Model(&model.User{}).
		Where("id = ?", userIn.Id).
		Updates(&userIn)

	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("user is not found")
		return
	}

	return
}

func (u *UserGormRepoImpl) DeleteUserById(ctx context.Context, userId uint64) (user model.User, err error) {
	tx := u.db.
		Model(&model.User{}).
		Clauses(clause.Returning{}). // clause to return data after delete
		Where("id = ?", userId).
		Delete(&user)
		// by default, func delete
		// di gorm akan mengupdate column deleted_at
	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("user is not found")
		return
	}
	return
}

func (u *UserGormRepoImpl) HardDeleteUserById(ctx context.Context, userId uint64) (user model.User, err error) {
	// untuk custom query, kita bisa menggunakan
	// Raw method di gorm
	// u.db.Raw()

	tx := u.db.
		Unscoped().
		Model(&model.User{}).
		Where("id = ?", userId).
		Delete(&model.User{})
		// unscope akan menandakan gorm
		// untuk hard delete row
	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("user is not found")
		return
	}
	return
}
