package unitofwork

import "gorm.io/gorm"

type UnitOfWork struct {
	DbMaster *gorm.DB
	Tx       *gorm.DB
}

type UnitOfWorkInterface interface {
	Start() *UnitOfWork
	Complete() error
	Dispose() error
	Finish(err error) error
}

func (uow *UnitOfWork) Start() *UnitOfWork {
	tx := uow.DbMaster.Begin()
	return &UnitOfWork{
		Tx: tx,
	}
}

func (uow *UnitOfWork) Finish(err error) error {
	if err != nil {
		errRollback := uow.Dispose()
		if errRollback != nil {
			return errRollback
		}
		return err
	}
	if err := uow.Complete(); err != nil {
		errRollback := uow.Dispose()
		if errRollback != nil {
			return errRollback
		}
		return err
	}
	return nil
}
func (uow *UnitOfWork) Complete() error {
	return uow.Tx.Commit().Error
}

func (uow *UnitOfWork) Dispose() error {
	return uow.Tx.Rollback().Error
}
