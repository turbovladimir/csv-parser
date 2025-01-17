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

func newSendingSmsJob(db *gorm.DB, opts ...gen.DOOption) sendingSmsJob {
	_sendingSmsJob := sendingSmsJob{}

	_sendingSmsJob.sendingSmsJobDo.UseDB(db, opts...)
	_sendingSmsJob.sendingSmsJobDo.UseModel(&model.SendingSmsJob{})

	tableName := _sendingSmsJob.sendingSmsJobDo.TableName()
	_sendingSmsJob.ALL = field.NewAsterisk(tableName)
	_sendingSmsJob.ID = field.NewInt32(tableName, "id")
	_sendingSmsJob.ContactID = field.NewInt32(tableName, "contact_id")
	_sendingSmsJob.SmsQueueID = field.NewInt32(tableName, "sms_queue_id")
	_sendingSmsJob.AddedAt = field.NewTime(tableName, "added_at")
	_sendingSmsJob.SendingTime = field.NewTime(tableName, "sending_time")
	_sendingSmsJob.Status = field.NewString(tableName, "status")
	_sendingSmsJob.ErrorText = field.NewString(tableName, "error_text")

	_sendingSmsJob.fillFieldMap()

	return _sendingSmsJob
}

type sendingSmsJob struct {
	sendingSmsJobDo

	ALL         field.Asterisk
	ID          field.Int32
	ContactID   field.Int32
	SmsQueueID  field.Int32
	AddedAt     field.Time
	SendingTime field.Time
	Status      field.String
	ErrorText   field.String

	fieldMap map[string]field.Expr
}

func (s sendingSmsJob) Table(newTableName string) *sendingSmsJob {
	s.sendingSmsJobDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s sendingSmsJob) As(alias string) *sendingSmsJob {
	s.sendingSmsJobDo.DO = *(s.sendingSmsJobDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *sendingSmsJob) updateTableName(table string) *sendingSmsJob {
	s.ALL = field.NewAsterisk(table)
	s.ID = field.NewInt32(table, "id")
	s.ContactID = field.NewInt32(table, "contact_id")
	s.SmsQueueID = field.NewInt32(table, "sms_queue_id")
	s.AddedAt = field.NewTime(table, "added_at")
	s.SendingTime = field.NewTime(table, "sending_time")
	s.Status = field.NewString(table, "status")
	s.ErrorText = field.NewString(table, "error_text")

	s.fillFieldMap()

	return s
}

func (s *sendingSmsJob) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *sendingSmsJob) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["id"] = s.ID
	s.fieldMap["contact_id"] = s.ContactID
	s.fieldMap["sms_queue_id"] = s.SmsQueueID
	s.fieldMap["added_at"] = s.AddedAt
	s.fieldMap["sending_time"] = s.SendingTime
	s.fieldMap["status"] = s.Status
	s.fieldMap["error_text"] = s.ErrorText
}

func (s sendingSmsJob) clone(db *gorm.DB) sendingSmsJob {
	s.sendingSmsJobDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s sendingSmsJob) replaceDB(db *gorm.DB) sendingSmsJob {
	s.sendingSmsJobDo.ReplaceDB(db)
	return s
}

type sendingSmsJobDo struct{ gen.DO }

type ISendingSmsJobDo interface {
	gen.SubQuery
	Debug() ISendingSmsJobDo
	WithContext(ctx context.Context) ISendingSmsJobDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISendingSmsJobDo
	WriteDB() ISendingSmsJobDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISendingSmsJobDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISendingSmsJobDo
	Not(conds ...gen.Condition) ISendingSmsJobDo
	Or(conds ...gen.Condition) ISendingSmsJobDo
	Select(conds ...field.Expr) ISendingSmsJobDo
	Where(conds ...gen.Condition) ISendingSmsJobDo
	Order(conds ...field.Expr) ISendingSmsJobDo
	Distinct(cols ...field.Expr) ISendingSmsJobDo
	Omit(cols ...field.Expr) ISendingSmsJobDo
	Join(table schema.Tabler, on ...field.Expr) ISendingSmsJobDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISendingSmsJobDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISendingSmsJobDo
	Group(cols ...field.Expr) ISendingSmsJobDo
	Having(conds ...gen.Condition) ISendingSmsJobDo
	Limit(limit int) ISendingSmsJobDo
	Offset(offset int) ISendingSmsJobDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISendingSmsJobDo
	Unscoped() ISendingSmsJobDo
	Create(values ...*model.SendingSmsJob) error
	CreateInBatches(values []*model.SendingSmsJob, batchSize int) error
	Save(values ...*model.SendingSmsJob) error
	First() (*model.SendingSmsJob, error)
	Take() (*model.SendingSmsJob, error)
	Last() (*model.SendingSmsJob, error)
	Find() ([]*model.SendingSmsJob, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SendingSmsJob, err error)
	FindInBatches(result *[]*model.SendingSmsJob, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.SendingSmsJob) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISendingSmsJobDo
	Assign(attrs ...field.AssignExpr) ISendingSmsJobDo
	Joins(fields ...field.RelationField) ISendingSmsJobDo
	Preload(fields ...field.RelationField) ISendingSmsJobDo
	FirstOrInit() (*model.SendingSmsJob, error)
	FirstOrCreate() (*model.SendingSmsJob, error)
	FindByPage(offset int, limit int) (result []*model.SendingSmsJob, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISendingSmsJobDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s sendingSmsJobDo) Debug() ISendingSmsJobDo {
	return s.withDO(s.DO.Debug())
}

func (s sendingSmsJobDo) WithContext(ctx context.Context) ISendingSmsJobDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s sendingSmsJobDo) ReadDB() ISendingSmsJobDo {
	return s.Clauses(dbresolver.Read)
}

func (s sendingSmsJobDo) WriteDB() ISendingSmsJobDo {
	return s.Clauses(dbresolver.Write)
}

func (s sendingSmsJobDo) Session(config *gorm.Session) ISendingSmsJobDo {
	return s.withDO(s.DO.Session(config))
}

func (s sendingSmsJobDo) Clauses(conds ...clause.Expression) ISendingSmsJobDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s sendingSmsJobDo) Returning(value interface{}, columns ...string) ISendingSmsJobDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s sendingSmsJobDo) Not(conds ...gen.Condition) ISendingSmsJobDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s sendingSmsJobDo) Or(conds ...gen.Condition) ISendingSmsJobDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s sendingSmsJobDo) Select(conds ...field.Expr) ISendingSmsJobDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s sendingSmsJobDo) Where(conds ...gen.Condition) ISendingSmsJobDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s sendingSmsJobDo) Order(conds ...field.Expr) ISendingSmsJobDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s sendingSmsJobDo) Distinct(cols ...field.Expr) ISendingSmsJobDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s sendingSmsJobDo) Omit(cols ...field.Expr) ISendingSmsJobDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s sendingSmsJobDo) Join(table schema.Tabler, on ...field.Expr) ISendingSmsJobDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s sendingSmsJobDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISendingSmsJobDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s sendingSmsJobDo) RightJoin(table schema.Tabler, on ...field.Expr) ISendingSmsJobDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s sendingSmsJobDo) Group(cols ...field.Expr) ISendingSmsJobDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s sendingSmsJobDo) Having(conds ...gen.Condition) ISendingSmsJobDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s sendingSmsJobDo) Limit(limit int) ISendingSmsJobDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s sendingSmsJobDo) Offset(offset int) ISendingSmsJobDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s sendingSmsJobDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISendingSmsJobDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s sendingSmsJobDo) Unscoped() ISendingSmsJobDo {
	return s.withDO(s.DO.Unscoped())
}

func (s sendingSmsJobDo) Create(values ...*model.SendingSmsJob) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s sendingSmsJobDo) CreateInBatches(values []*model.SendingSmsJob, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s sendingSmsJobDo) Save(values ...*model.SendingSmsJob) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s sendingSmsJobDo) First() (*model.SendingSmsJob, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendingSmsJob), nil
	}
}

func (s sendingSmsJobDo) Take() (*model.SendingSmsJob, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendingSmsJob), nil
	}
}

func (s sendingSmsJobDo) Last() (*model.SendingSmsJob, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendingSmsJob), nil
	}
}

func (s sendingSmsJobDo) Find() ([]*model.SendingSmsJob, error) {
	result, err := s.DO.Find()
	return result.([]*model.SendingSmsJob), err
}

func (s sendingSmsJobDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.SendingSmsJob, err error) {
	buf := make([]*model.SendingSmsJob, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s sendingSmsJobDo) FindInBatches(result *[]*model.SendingSmsJob, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s sendingSmsJobDo) Attrs(attrs ...field.AssignExpr) ISendingSmsJobDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s sendingSmsJobDo) Assign(attrs ...field.AssignExpr) ISendingSmsJobDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s sendingSmsJobDo) Joins(fields ...field.RelationField) ISendingSmsJobDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s sendingSmsJobDo) Preload(fields ...field.RelationField) ISendingSmsJobDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s sendingSmsJobDo) FirstOrInit() (*model.SendingSmsJob, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendingSmsJob), nil
	}
}

func (s sendingSmsJobDo) FirstOrCreate() (*model.SendingSmsJob, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.SendingSmsJob), nil
	}
}

func (s sendingSmsJobDo) FindByPage(offset int, limit int) (result []*model.SendingSmsJob, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s sendingSmsJobDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s sendingSmsJobDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s sendingSmsJobDo) Delete(models ...*model.SendingSmsJob) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *sendingSmsJobDo) withDO(do gen.Dao) *sendingSmsJobDo {
	s.DO = *do.(*gen.DO)
	return s
}
