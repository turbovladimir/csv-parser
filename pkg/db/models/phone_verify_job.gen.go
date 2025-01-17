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

func newPhoneVerifyJob(db *gorm.DB, opts ...gen.DOOption) phoneVerifyJob {
	_phoneVerifyJob := phoneVerifyJob{}

	_phoneVerifyJob.phoneVerifyJobDo.UseDB(db, opts...)
	_phoneVerifyJob.phoneVerifyJobDo.UseModel(&model.PhoneVerifyJob{})

	tableName := _phoneVerifyJob.phoneVerifyJobDo.TableName()
	_phoneVerifyJob.ALL = field.NewAsterisk(tableName)
	_phoneVerifyJob.ID = field.NewInt32(tableName, "id")
	_phoneVerifyJob.SessionID = field.NewString(tableName, "session_id")
	_phoneVerifyJob.Phone = field.NewString(tableName, "phone")
	_phoneVerifyJob.AddedAt = field.NewTime(tableName, "added_at")
	_phoneVerifyJob.IsActive = field.NewBool(tableName, "is_active")
	_phoneVerifyJob.Code = field.NewInt32(tableName, "code")
	_phoneVerifyJob.IsVerified = field.NewBool(tableName, "is_verified")

	_phoneVerifyJob.fillFieldMap()

	return _phoneVerifyJob
}

type phoneVerifyJob struct {
	phoneVerifyJobDo

	ALL        field.Asterisk
	ID         field.Int32
	SessionID  field.String
	Phone      field.String
	AddedAt    field.Time
	IsActive   field.Bool
	Code       field.Int32
	IsVerified field.Bool

	fieldMap map[string]field.Expr
}

func (p phoneVerifyJob) Table(newTableName string) *phoneVerifyJob {
	p.phoneVerifyJobDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p phoneVerifyJob) As(alias string) *phoneVerifyJob {
	p.phoneVerifyJobDo.DO = *(p.phoneVerifyJobDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *phoneVerifyJob) updateTableName(table string) *phoneVerifyJob {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.SessionID = field.NewString(table, "session_id")
	p.Phone = field.NewString(table, "phone")
	p.AddedAt = field.NewTime(table, "added_at")
	p.IsActive = field.NewBool(table, "is_active")
	p.Code = field.NewInt32(table, "code")
	p.IsVerified = field.NewBool(table, "is_verified")

	p.fillFieldMap()

	return p
}

func (p *phoneVerifyJob) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *phoneVerifyJob) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["session_id"] = p.SessionID
	p.fieldMap["phone"] = p.Phone
	p.fieldMap["added_at"] = p.AddedAt
	p.fieldMap["is_active"] = p.IsActive
	p.fieldMap["code"] = p.Code
	p.fieldMap["is_verified"] = p.IsVerified
}

func (p phoneVerifyJob) clone(db *gorm.DB) phoneVerifyJob {
	p.phoneVerifyJobDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p phoneVerifyJob) replaceDB(db *gorm.DB) phoneVerifyJob {
	p.phoneVerifyJobDo.ReplaceDB(db)
	return p
}

type phoneVerifyJobDo struct{ gen.DO }

type IPhoneVerifyJobDo interface {
	gen.SubQuery
	Debug() IPhoneVerifyJobDo
	WithContext(ctx context.Context) IPhoneVerifyJobDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPhoneVerifyJobDo
	WriteDB() IPhoneVerifyJobDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPhoneVerifyJobDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPhoneVerifyJobDo
	Not(conds ...gen.Condition) IPhoneVerifyJobDo
	Or(conds ...gen.Condition) IPhoneVerifyJobDo
	Select(conds ...field.Expr) IPhoneVerifyJobDo
	Where(conds ...gen.Condition) IPhoneVerifyJobDo
	Order(conds ...field.Expr) IPhoneVerifyJobDo
	Distinct(cols ...field.Expr) IPhoneVerifyJobDo
	Omit(cols ...field.Expr) IPhoneVerifyJobDo
	Join(table schema.Tabler, on ...field.Expr) IPhoneVerifyJobDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPhoneVerifyJobDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPhoneVerifyJobDo
	Group(cols ...field.Expr) IPhoneVerifyJobDo
	Having(conds ...gen.Condition) IPhoneVerifyJobDo
	Limit(limit int) IPhoneVerifyJobDo
	Offset(offset int) IPhoneVerifyJobDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPhoneVerifyJobDo
	Unscoped() IPhoneVerifyJobDo
	Create(values ...*model.PhoneVerifyJob) error
	CreateInBatches(values []*model.PhoneVerifyJob, batchSize int) error
	Save(values ...*model.PhoneVerifyJob) error
	First() (*model.PhoneVerifyJob, error)
	Take() (*model.PhoneVerifyJob, error)
	Last() (*model.PhoneVerifyJob, error)
	Find() ([]*model.PhoneVerifyJob, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PhoneVerifyJob, err error)
	FindInBatches(result *[]*model.PhoneVerifyJob, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.PhoneVerifyJob) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPhoneVerifyJobDo
	Assign(attrs ...field.AssignExpr) IPhoneVerifyJobDo
	Joins(fields ...field.RelationField) IPhoneVerifyJobDo
	Preload(fields ...field.RelationField) IPhoneVerifyJobDo
	FirstOrInit() (*model.PhoneVerifyJob, error)
	FirstOrCreate() (*model.PhoneVerifyJob, error)
	FindByPage(offset int, limit int) (result []*model.PhoneVerifyJob, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPhoneVerifyJobDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p phoneVerifyJobDo) Debug() IPhoneVerifyJobDo {
	return p.withDO(p.DO.Debug())
}

func (p phoneVerifyJobDo) WithContext(ctx context.Context) IPhoneVerifyJobDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p phoneVerifyJobDo) ReadDB() IPhoneVerifyJobDo {
	return p.Clauses(dbresolver.Read)
}

func (p phoneVerifyJobDo) WriteDB() IPhoneVerifyJobDo {
	return p.Clauses(dbresolver.Write)
}

func (p phoneVerifyJobDo) Session(config *gorm.Session) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Session(config))
}

func (p phoneVerifyJobDo) Clauses(conds ...clause.Expression) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p phoneVerifyJobDo) Returning(value interface{}, columns ...string) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p phoneVerifyJobDo) Not(conds ...gen.Condition) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p phoneVerifyJobDo) Or(conds ...gen.Condition) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p phoneVerifyJobDo) Select(conds ...field.Expr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p phoneVerifyJobDo) Where(conds ...gen.Condition) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p phoneVerifyJobDo) Order(conds ...field.Expr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p phoneVerifyJobDo) Distinct(cols ...field.Expr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p phoneVerifyJobDo) Omit(cols ...field.Expr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p phoneVerifyJobDo) Join(table schema.Tabler, on ...field.Expr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p phoneVerifyJobDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p phoneVerifyJobDo) RightJoin(table schema.Tabler, on ...field.Expr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p phoneVerifyJobDo) Group(cols ...field.Expr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p phoneVerifyJobDo) Having(conds ...gen.Condition) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p phoneVerifyJobDo) Limit(limit int) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p phoneVerifyJobDo) Offset(offset int) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p phoneVerifyJobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p phoneVerifyJobDo) Unscoped() IPhoneVerifyJobDo {
	return p.withDO(p.DO.Unscoped())
}

func (p phoneVerifyJobDo) Create(values ...*model.PhoneVerifyJob) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p phoneVerifyJobDo) CreateInBatches(values []*model.PhoneVerifyJob, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p phoneVerifyJobDo) Save(values ...*model.PhoneVerifyJob) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p phoneVerifyJobDo) First() (*model.PhoneVerifyJob, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PhoneVerifyJob), nil
	}
}

func (p phoneVerifyJobDo) Take() (*model.PhoneVerifyJob, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PhoneVerifyJob), nil
	}
}

func (p phoneVerifyJobDo) Last() (*model.PhoneVerifyJob, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PhoneVerifyJob), nil
	}
}

func (p phoneVerifyJobDo) Find() ([]*model.PhoneVerifyJob, error) {
	result, err := p.DO.Find()
	return result.([]*model.PhoneVerifyJob), err
}

func (p phoneVerifyJobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PhoneVerifyJob, err error) {
	buf := make([]*model.PhoneVerifyJob, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p phoneVerifyJobDo) FindInBatches(result *[]*model.PhoneVerifyJob, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p phoneVerifyJobDo) Attrs(attrs ...field.AssignExpr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p phoneVerifyJobDo) Assign(attrs ...field.AssignExpr) IPhoneVerifyJobDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p phoneVerifyJobDo) Joins(fields ...field.RelationField) IPhoneVerifyJobDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p phoneVerifyJobDo) Preload(fields ...field.RelationField) IPhoneVerifyJobDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p phoneVerifyJobDo) FirstOrInit() (*model.PhoneVerifyJob, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PhoneVerifyJob), nil
	}
}

func (p phoneVerifyJobDo) FirstOrCreate() (*model.PhoneVerifyJob, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PhoneVerifyJob), nil
	}
}

func (p phoneVerifyJobDo) FindByPage(offset int, limit int) (result []*model.PhoneVerifyJob, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p phoneVerifyJobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p phoneVerifyJobDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p phoneVerifyJobDo) Delete(models ...*model.PhoneVerifyJob) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *phoneVerifyJobDo) withDO(do gen.Dao) *phoneVerifyJobDo {
	p.DO = *do.(*gen.DO)
	return p
}
