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

func newPush(db *gorm.DB, opts ...gen.DOOption) push {
	_push := push{}

	_push.pushDo.UseDB(db, opts...)
	_push.pushDo.UseModel(&model.Push{})

	tableName := _push.pushDo.TableName()
	_push.ALL = field.NewAsterisk(tableName)
	_push.ID = field.NewInt32(tableName, "id")
	_push.AddedAt = field.NewTime(tableName, "added_at")
	_push.URI = field.NewString(tableName, "uri")
	_push.Target = field.NewString(tableName, "target")
	_push.Text = field.NewString(tableName, "text")
	_push.ShowDelaySecs = field.NewInt32(tableName, "show_delay_secs")
	_push.IsActive = field.NewBool(tableName, "is_active")

	_push.fillFieldMap()

	return _push
}

type push struct {
	pushDo

	ALL           field.Asterisk
	ID            field.Int32
	AddedAt       field.Time
	URI           field.String
	Target        field.String
	Text          field.String
	ShowDelaySecs field.Int32
	IsActive      field.Bool

	fieldMap map[string]field.Expr
}

func (p push) Table(newTableName string) *push {
	p.pushDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p push) As(alias string) *push {
	p.pushDo.DO = *(p.pushDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *push) updateTableName(table string) *push {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.AddedAt = field.NewTime(table, "added_at")
	p.URI = field.NewString(table, "uri")
	p.Target = field.NewString(table, "target")
	p.Text = field.NewString(table, "text")
	p.ShowDelaySecs = field.NewInt32(table, "show_delay_secs")
	p.IsActive = field.NewBool(table, "is_active")

	p.fillFieldMap()

	return p
}

func (p *push) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *push) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["added_at"] = p.AddedAt
	p.fieldMap["uri"] = p.URI
	p.fieldMap["target"] = p.Target
	p.fieldMap["text"] = p.Text
	p.fieldMap["show_delay_secs"] = p.ShowDelaySecs
	p.fieldMap["is_active"] = p.IsActive
}

func (p push) clone(db *gorm.DB) push {
	p.pushDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p push) replaceDB(db *gorm.DB) push {
	p.pushDo.ReplaceDB(db)
	return p
}

type pushDo struct{ gen.DO }

type IPushDo interface {
	gen.SubQuery
	Debug() IPushDo
	WithContext(ctx context.Context) IPushDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPushDo
	WriteDB() IPushDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPushDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPushDo
	Not(conds ...gen.Condition) IPushDo
	Or(conds ...gen.Condition) IPushDo
	Select(conds ...field.Expr) IPushDo
	Where(conds ...gen.Condition) IPushDo
	Order(conds ...field.Expr) IPushDo
	Distinct(cols ...field.Expr) IPushDo
	Omit(cols ...field.Expr) IPushDo
	Join(table schema.Tabler, on ...field.Expr) IPushDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPushDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPushDo
	Group(cols ...field.Expr) IPushDo
	Having(conds ...gen.Condition) IPushDo
	Limit(limit int) IPushDo
	Offset(offset int) IPushDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPushDo
	Unscoped() IPushDo
	Create(values ...*model.Push) error
	CreateInBatches(values []*model.Push, batchSize int) error
	Save(values ...*model.Push) error
	First() (*model.Push, error)
	Take() (*model.Push, error)
	Last() (*model.Push, error)
	Find() ([]*model.Push, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Push, err error)
	FindInBatches(result *[]*model.Push, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Push) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPushDo
	Assign(attrs ...field.AssignExpr) IPushDo
	Joins(fields ...field.RelationField) IPushDo
	Preload(fields ...field.RelationField) IPushDo
	FirstOrInit() (*model.Push, error)
	FirstOrCreate() (*model.Push, error)
	FindByPage(offset int, limit int) (result []*model.Push, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPushDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p pushDo) Debug() IPushDo {
	return p.withDO(p.DO.Debug())
}

func (p pushDo) WithContext(ctx context.Context) IPushDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p pushDo) ReadDB() IPushDo {
	return p.Clauses(dbresolver.Read)
}

func (p pushDo) WriteDB() IPushDo {
	return p.Clauses(dbresolver.Write)
}

func (p pushDo) Session(config *gorm.Session) IPushDo {
	return p.withDO(p.DO.Session(config))
}

func (p pushDo) Clauses(conds ...clause.Expression) IPushDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p pushDo) Returning(value interface{}, columns ...string) IPushDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p pushDo) Not(conds ...gen.Condition) IPushDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p pushDo) Or(conds ...gen.Condition) IPushDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p pushDo) Select(conds ...field.Expr) IPushDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p pushDo) Where(conds ...gen.Condition) IPushDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p pushDo) Order(conds ...field.Expr) IPushDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p pushDo) Distinct(cols ...field.Expr) IPushDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p pushDo) Omit(cols ...field.Expr) IPushDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p pushDo) Join(table schema.Tabler, on ...field.Expr) IPushDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p pushDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPushDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p pushDo) RightJoin(table schema.Tabler, on ...field.Expr) IPushDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p pushDo) Group(cols ...field.Expr) IPushDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p pushDo) Having(conds ...gen.Condition) IPushDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p pushDo) Limit(limit int) IPushDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p pushDo) Offset(offset int) IPushDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p pushDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPushDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p pushDo) Unscoped() IPushDo {
	return p.withDO(p.DO.Unscoped())
}

func (p pushDo) Create(values ...*model.Push) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p pushDo) CreateInBatches(values []*model.Push, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p pushDo) Save(values ...*model.Push) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p pushDo) First() (*model.Push, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Push), nil
	}
}

func (p pushDo) Take() (*model.Push, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Push), nil
	}
}

func (p pushDo) Last() (*model.Push, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Push), nil
	}
}

func (p pushDo) Find() ([]*model.Push, error) {
	result, err := p.DO.Find()
	return result.([]*model.Push), err
}

func (p pushDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Push, err error) {
	buf := make([]*model.Push, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p pushDo) FindInBatches(result *[]*model.Push, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p pushDo) Attrs(attrs ...field.AssignExpr) IPushDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p pushDo) Assign(attrs ...field.AssignExpr) IPushDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p pushDo) Joins(fields ...field.RelationField) IPushDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p pushDo) Preload(fields ...field.RelationField) IPushDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p pushDo) FirstOrInit() (*model.Push, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Push), nil
	}
}

func (p pushDo) FirstOrCreate() (*model.Push, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Push), nil
	}
}

func (p pushDo) FindByPage(offset int, limit int) (result []*model.Push, count int64, err error) {
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

func (p pushDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p pushDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p pushDo) Delete(models ...*model.Push) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *pushDo) withDO(do gen.Dao) *pushDo {
	p.DO = *do.(*gen.DO)
	return p
}
