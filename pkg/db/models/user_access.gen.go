// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package models

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/turbovladimir/csv-parser/pkg/db/model"
)

func newUserAccess(db *gorm.DB, opts ...gen.DOOption) userAccess {
	_userAccess := userAccess{}

	_userAccess.userAccessDo.UseDB(db, opts...)
	_userAccess.userAccessDo.UseModel(&model.UserAccess{})

	tableName := _userAccess.userAccessDo.TableName()
	_userAccess.ALL = field.NewAsterisk(tableName)
	_userAccess.ID = field.NewInt32(tableName, "id")
	_userAccess.IP = field.NewString(tableName, "ip")
	_userAccess.AddedAt = field.NewTime(tableName, "added_at")
	_userAccess.Limit_ = field.NewInt32(tableName, "_limit")

	_userAccess.fillFieldMap()

	return _userAccess
}

type userAccess struct {
	userAccessDo

	ALL     field.Asterisk
	ID      field.Int32
	IP      field.String
	AddedAt field.Time
	Limit_  field.Int32

	fieldMap map[string]field.Expr
}

func (u userAccess) Table(newTableName string) *userAccess {
	u.userAccessDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userAccess) As(alias string) *userAccess {
	u.userAccessDo.DO = *(u.userAccessDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userAccess) updateTableName(table string) *userAccess {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt32(table, "id")
	u.IP = field.NewString(table, "ip")
	u.AddedAt = field.NewTime(table, "added_at")
	u.Limit_ = field.NewInt32(table, "_limit")

	u.fillFieldMap()

	return u
}

func (u *userAccess) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userAccess) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 4)
	u.fieldMap["id"] = u.ID
	u.fieldMap["ip"] = u.IP
	u.fieldMap["added_at"] = u.AddedAt
	u.fieldMap["_limit"] = u.Limit_
}

func (u userAccess) clone(db *gorm.DB) userAccess {
	u.userAccessDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u userAccess) replaceDB(db *gorm.DB) userAccess {
	u.userAccessDo.ReplaceDB(db)
	return u
}

type userAccessDo struct{ gen.DO }

type IUserAccessDo interface {
	gen.SubQuery
	Debug() IUserAccessDo
	WithContext(ctx context.Context) IUserAccessDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IUserAccessDo
	WriteDB() IUserAccessDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IUserAccessDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IUserAccessDo
	Not(conds ...gen.Condition) IUserAccessDo
	Or(conds ...gen.Condition) IUserAccessDo
	Select(conds ...field.Expr) IUserAccessDo
	Where(conds ...gen.Condition) IUserAccessDo
	Order(conds ...field.Expr) IUserAccessDo
	Distinct(cols ...field.Expr) IUserAccessDo
	Omit(cols ...field.Expr) IUserAccessDo
	Join(table schema.Tabler, on ...field.Expr) IUserAccessDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IUserAccessDo
	RightJoin(table schema.Tabler, on ...field.Expr) IUserAccessDo
	Group(cols ...field.Expr) IUserAccessDo
	Having(conds ...gen.Condition) IUserAccessDo
	Limit(limit int) IUserAccessDo
	Offset(offset int) IUserAccessDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IUserAccessDo
	Unscoped() IUserAccessDo
	Create(values ...*model.UserAccess) error
	CreateInBatches(values []*model.UserAccess, batchSize int) error
	Save(values ...*model.UserAccess) error
	First() (*model.UserAccess, error)
	Take() (*model.UserAccess, error)
	Last() (*model.UserAccess, error)
	Find() ([]*model.UserAccess, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAccess, err error)
	FindInBatches(result *[]*model.UserAccess, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.UserAccess) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IUserAccessDo
	Assign(attrs ...field.AssignExpr) IUserAccessDo
	Joins(fields ...field.RelationField) IUserAccessDo
	Preload(fields ...field.RelationField) IUserAccessDo
	FirstOrInit() (*model.UserAccess, error)
	FirstOrCreate() (*model.UserAccess, error)
	FindByPage(offset int, limit int) (result []*model.UserAccess, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IUserAccessDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (u userAccessDo) Debug() IUserAccessDo {
	return u.withDO(u.DO.Debug())
}

func (u userAccessDo) WithContext(ctx context.Context) IUserAccessDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userAccessDo) ReadDB() IUserAccessDo {
	return u.Clauses(dbresolver.Read)
}

func (u userAccessDo) WriteDB() IUserAccessDo {
	return u.Clauses(dbresolver.Write)
}

func (u userAccessDo) Session(config *gorm.Session) IUserAccessDo {
	return u.withDO(u.DO.Session(config))
}

func (u userAccessDo) Clauses(conds ...clause.Expression) IUserAccessDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userAccessDo) Returning(value interface{}, columns ...string) IUserAccessDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userAccessDo) Not(conds ...gen.Condition) IUserAccessDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userAccessDo) Or(conds ...gen.Condition) IUserAccessDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userAccessDo) Select(conds ...field.Expr) IUserAccessDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userAccessDo) Where(conds ...gen.Condition) IUserAccessDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userAccessDo) Order(conds ...field.Expr) IUserAccessDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userAccessDo) Distinct(cols ...field.Expr) IUserAccessDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userAccessDo) Omit(cols ...field.Expr) IUserAccessDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userAccessDo) Join(table schema.Tabler, on ...field.Expr) IUserAccessDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userAccessDo) LeftJoin(table schema.Tabler, on ...field.Expr) IUserAccessDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userAccessDo) RightJoin(table schema.Tabler, on ...field.Expr) IUserAccessDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userAccessDo) Group(cols ...field.Expr) IUserAccessDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userAccessDo) Having(conds ...gen.Condition) IUserAccessDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userAccessDo) Limit(limit int) IUserAccessDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userAccessDo) Offset(offset int) IUserAccessDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userAccessDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IUserAccessDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userAccessDo) Unscoped() IUserAccessDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userAccessDo) Create(values ...*model.UserAccess) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userAccessDo) CreateInBatches(values []*model.UserAccess, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userAccessDo) Save(values ...*model.UserAccess) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userAccessDo) First() (*model.UserAccess, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccess), nil
	}
}

func (u userAccessDo) Take() (*model.UserAccess, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccess), nil
	}
}

func (u userAccessDo) Last() (*model.UserAccess, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccess), nil
	}
}

func (u userAccessDo) Find() ([]*model.UserAccess, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserAccess), err
}

func (u userAccessDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserAccess, err error) {
	buf := make([]*model.UserAccess, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userAccessDo) FindInBatches(result *[]*model.UserAccess, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userAccessDo) Attrs(attrs ...field.AssignExpr) IUserAccessDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userAccessDo) Assign(attrs ...field.AssignExpr) IUserAccessDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userAccessDo) Joins(fields ...field.RelationField) IUserAccessDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userAccessDo) Preload(fields ...field.RelationField) IUserAccessDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userAccessDo) FirstOrInit() (*model.UserAccess, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccess), nil
	}
}

func (u userAccessDo) FirstOrCreate() (*model.UserAccess, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserAccess), nil
	}
}

func (u userAccessDo) FindByPage(offset int, limit int) (result []*model.UserAccess, count int64, err error) {
	result, err = u.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = u.Offset(-1).Limit(-1).Count()
	return
}

func (u userAccessDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userAccessDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u userAccessDo) Delete(models ...*model.UserAccess) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *userAccessDo) withDO(do gen.Dao) *userAccessDo {
	u.DO = *do.(*gen.DO)
	return u
}