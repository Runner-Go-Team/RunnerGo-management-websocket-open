// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"RunnerGo-management/cmd/generate/internal/pkg/dal/model"
)

func newReportMachine(db *gorm.DB, opts ...gen.DOOption) reportMachine {
	_reportMachine := reportMachine{}

	_reportMachine.reportMachineDo.UseDB(db, opts...)
	_reportMachine.reportMachineDo.UseModel(&model.ReportMachine{})

	tableName := _reportMachine.reportMachineDo.TableName()
	_reportMachine.ALL = field.NewAsterisk(tableName)
	_reportMachine.ID = field.NewInt64(tableName, "id")
	_reportMachine.ReportID = field.NewInt64(tableName, "report_id")
	_reportMachine.IP = field.NewString(tableName, "ip")
	_reportMachine.CreatedAt = field.NewTime(tableName, "created_at")
	_reportMachine.UpdatedAt = field.NewTime(tableName, "updated_at")
	_reportMachine.DeletedAt = field.NewField(tableName, "deleted_at")

	_reportMachine.fillFieldMap()

	return _reportMachine
}

type reportMachine struct {
	reportMachineDo reportMachineDo

	ALL       field.Asterisk
	ID        field.Int64
	ReportID  field.Int64  // 报告id
	IP        field.String // 机器ip
	CreatedAt field.Time
	UpdatedAt field.Time
	DeletedAt field.Field

	fieldMap map[string]field.Expr
}

func (r reportMachine) Table(newTableName string) *reportMachine {
	r.reportMachineDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reportMachine) As(alias string) *reportMachine {
	r.reportMachineDo.DO = *(r.reportMachineDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reportMachine) updateTableName(table string) *reportMachine {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewInt64(table, "id")
	r.ReportID = field.NewInt64(table, "report_id")
	r.IP = field.NewString(table, "ip")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")

	r.fillFieldMap()

	return r
}

func (r *reportMachine) WithContext(ctx context.Context) *reportMachineDo {
	return r.reportMachineDo.WithContext(ctx)
}

func (r reportMachine) TableName() string { return r.reportMachineDo.TableName() }

func (r reportMachine) Alias() string { return r.reportMachineDo.Alias() }

func (r *reportMachine) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reportMachine) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 6)
	r.fieldMap["id"] = r.ID
	r.fieldMap["report_id"] = r.ReportID
	r.fieldMap["ip"] = r.IP
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
}

func (r reportMachine) clone(db *gorm.DB) reportMachine {
	r.reportMachineDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reportMachine) replaceDB(db *gorm.DB) reportMachine {
	r.reportMachineDo.ReplaceDB(db)
	return r
}

type reportMachineDo struct{ gen.DO }

func (r reportMachineDo) Debug() *reportMachineDo {
	return r.withDO(r.DO.Debug())
}

func (r reportMachineDo) WithContext(ctx context.Context) *reportMachineDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reportMachineDo) ReadDB() *reportMachineDo {
	return r.Clauses(dbresolver.Read)
}

func (r reportMachineDo) WriteDB() *reportMachineDo {
	return r.Clauses(dbresolver.Write)
}

func (r reportMachineDo) Session(config *gorm.Session) *reportMachineDo {
	return r.withDO(r.DO.Session(config))
}

func (r reportMachineDo) Clauses(conds ...clause.Expression) *reportMachineDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reportMachineDo) Returning(value interface{}, columns ...string) *reportMachineDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reportMachineDo) Not(conds ...gen.Condition) *reportMachineDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reportMachineDo) Or(conds ...gen.Condition) *reportMachineDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reportMachineDo) Select(conds ...field.Expr) *reportMachineDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reportMachineDo) Where(conds ...gen.Condition) *reportMachineDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reportMachineDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *reportMachineDo {
	return r.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (r reportMachineDo) Order(conds ...field.Expr) *reportMachineDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reportMachineDo) Distinct(cols ...field.Expr) *reportMachineDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reportMachineDo) Omit(cols ...field.Expr) *reportMachineDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reportMachineDo) Join(table schema.Tabler, on ...field.Expr) *reportMachineDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reportMachineDo) LeftJoin(table schema.Tabler, on ...field.Expr) *reportMachineDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reportMachineDo) RightJoin(table schema.Tabler, on ...field.Expr) *reportMachineDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reportMachineDo) Group(cols ...field.Expr) *reportMachineDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reportMachineDo) Having(conds ...gen.Condition) *reportMachineDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reportMachineDo) Limit(limit int) *reportMachineDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reportMachineDo) Offset(offset int) *reportMachineDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reportMachineDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *reportMachineDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reportMachineDo) Unscoped() *reportMachineDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reportMachineDo) Create(values ...*model.ReportMachine) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reportMachineDo) CreateInBatches(values []*model.ReportMachine, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reportMachineDo) Save(values ...*model.ReportMachine) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reportMachineDo) First() (*model.ReportMachine, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportMachine), nil
	}
}

func (r reportMachineDo) Take() (*model.ReportMachine, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportMachine), nil
	}
}

func (r reportMachineDo) Last() (*model.ReportMachine, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportMachine), nil
	}
}

func (r reportMachineDo) Find() ([]*model.ReportMachine, error) {
	result, err := r.DO.Find()
	return result.([]*model.ReportMachine), err
}

func (r reportMachineDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ReportMachine, err error) {
	buf := make([]*model.ReportMachine, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reportMachineDo) FindInBatches(result *[]*model.ReportMachine, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reportMachineDo) Attrs(attrs ...field.AssignExpr) *reportMachineDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reportMachineDo) Assign(attrs ...field.AssignExpr) *reportMachineDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reportMachineDo) Joins(fields ...field.RelationField) *reportMachineDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reportMachineDo) Preload(fields ...field.RelationField) *reportMachineDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reportMachineDo) FirstOrInit() (*model.ReportMachine, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportMachine), nil
	}
}

func (r reportMachineDo) FirstOrCreate() (*model.ReportMachine, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ReportMachine), nil
	}
}

func (r reportMachineDo) FindByPage(offset int, limit int) (result []*model.ReportMachine, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r reportMachineDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reportMachineDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reportMachineDo) Delete(models ...*model.ReportMachine) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reportMachineDo) withDO(do gen.Dao) *reportMachineDo {
	r.DO = *do.(*gen.DO)
	return r
}