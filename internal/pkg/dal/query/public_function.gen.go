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

	"RunnerGo-management/internal/pkg/dal/model"
)

func newPublicFunction(db *gorm.DB, opts ...gen.DOOption) publicFunction {
	_publicFunction := publicFunction{}

	_publicFunction.publicFunctionDo.UseDB(db, opts...)
	_publicFunction.publicFunctionDo.UseModel(&model.PublicFunction{})

	tableName := _publicFunction.publicFunctionDo.TableName()
	_publicFunction.ALL = field.NewAsterisk(tableName)
	_publicFunction.ID = field.NewInt32(tableName, "id")
	_publicFunction.Function = field.NewString(tableName, "function")
	_publicFunction.FunctionName = field.NewString(tableName, "function_name")
	_publicFunction.Remark = field.NewString(tableName, "remark")
	_publicFunction.CreatedAt = field.NewTime(tableName, "created_at")
	_publicFunction.UpdatedAt = field.NewTime(tableName, "updated_at")
	_publicFunction.DeletedAt = field.NewField(tableName, "deleted_at")

	_publicFunction.fillFieldMap()

	return _publicFunction
}

type publicFunction struct {
	publicFunctionDo publicFunctionDo

	ALL          field.Asterisk
	ID           field.Int32  // 主键id
	Function     field.String // 函数
	FunctionName field.String // 函数名称
	Remark       field.String // 备注
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 修改时间
	DeletedAt    field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (p publicFunction) Table(newTableName string) *publicFunction {
	p.publicFunctionDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p publicFunction) As(alias string) *publicFunction {
	p.publicFunctionDo.DO = *(p.publicFunctionDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *publicFunction) updateTableName(table string) *publicFunction {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt32(table, "id")
	p.Function = field.NewString(table, "function")
	p.FunctionName = field.NewString(table, "function_name")
	p.Remark = field.NewString(table, "remark")
	p.CreatedAt = field.NewTime(table, "created_at")
	p.UpdatedAt = field.NewTime(table, "updated_at")
	p.DeletedAt = field.NewField(table, "deleted_at")

	p.fillFieldMap()

	return p
}

func (p *publicFunction) WithContext(ctx context.Context) *publicFunctionDo {
	return p.publicFunctionDo.WithContext(ctx)
}

func (p publicFunction) TableName() string { return p.publicFunctionDo.TableName() }

func (p publicFunction) Alias() string { return p.publicFunctionDo.Alias() }

func (p *publicFunction) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *publicFunction) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 7)
	p.fieldMap["id"] = p.ID
	p.fieldMap["function"] = p.Function
	p.fieldMap["function_name"] = p.FunctionName
	p.fieldMap["remark"] = p.Remark
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
	p.fieldMap["deleted_at"] = p.DeletedAt
}

func (p publicFunction) clone(db *gorm.DB) publicFunction {
	p.publicFunctionDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p publicFunction) replaceDB(db *gorm.DB) publicFunction {
	p.publicFunctionDo.ReplaceDB(db)
	return p
}

type publicFunctionDo struct{ gen.DO }

func (p publicFunctionDo) Debug() *publicFunctionDo {
	return p.withDO(p.DO.Debug())
}

func (p publicFunctionDo) WithContext(ctx context.Context) *publicFunctionDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p publicFunctionDo) ReadDB() *publicFunctionDo {
	return p.Clauses(dbresolver.Read)
}

func (p publicFunctionDo) WriteDB() *publicFunctionDo {
	return p.Clauses(dbresolver.Write)
}

func (p publicFunctionDo) Session(config *gorm.Session) *publicFunctionDo {
	return p.withDO(p.DO.Session(config))
}

func (p publicFunctionDo) Clauses(conds ...clause.Expression) *publicFunctionDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p publicFunctionDo) Returning(value interface{}, columns ...string) *publicFunctionDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p publicFunctionDo) Not(conds ...gen.Condition) *publicFunctionDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p publicFunctionDo) Or(conds ...gen.Condition) *publicFunctionDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p publicFunctionDo) Select(conds ...field.Expr) *publicFunctionDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p publicFunctionDo) Where(conds ...gen.Condition) *publicFunctionDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p publicFunctionDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *publicFunctionDo {
	return p.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (p publicFunctionDo) Order(conds ...field.Expr) *publicFunctionDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p publicFunctionDo) Distinct(cols ...field.Expr) *publicFunctionDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p publicFunctionDo) Omit(cols ...field.Expr) *publicFunctionDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p publicFunctionDo) Join(table schema.Tabler, on ...field.Expr) *publicFunctionDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p publicFunctionDo) LeftJoin(table schema.Tabler, on ...field.Expr) *publicFunctionDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p publicFunctionDo) RightJoin(table schema.Tabler, on ...field.Expr) *publicFunctionDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p publicFunctionDo) Group(cols ...field.Expr) *publicFunctionDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p publicFunctionDo) Having(conds ...gen.Condition) *publicFunctionDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p publicFunctionDo) Limit(limit int) *publicFunctionDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p publicFunctionDo) Offset(offset int) *publicFunctionDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p publicFunctionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *publicFunctionDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p publicFunctionDo) Unscoped() *publicFunctionDo {
	return p.withDO(p.DO.Unscoped())
}

func (p publicFunctionDo) Create(values ...*model.PublicFunction) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p publicFunctionDo) CreateInBatches(values []*model.PublicFunction, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p publicFunctionDo) Save(values ...*model.PublicFunction) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p publicFunctionDo) First() (*model.PublicFunction, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublicFunction), nil
	}
}

func (p publicFunctionDo) Take() (*model.PublicFunction, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublicFunction), nil
	}
}

func (p publicFunctionDo) Last() (*model.PublicFunction, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublicFunction), nil
	}
}

func (p publicFunctionDo) Find() ([]*model.PublicFunction, error) {
	result, err := p.DO.Find()
	return result.([]*model.PublicFunction), err
}

func (p publicFunctionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.PublicFunction, err error) {
	buf := make([]*model.PublicFunction, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p publicFunctionDo) FindInBatches(result *[]*model.PublicFunction, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p publicFunctionDo) Attrs(attrs ...field.AssignExpr) *publicFunctionDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p publicFunctionDo) Assign(attrs ...field.AssignExpr) *publicFunctionDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p publicFunctionDo) Joins(fields ...field.RelationField) *publicFunctionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p publicFunctionDo) Preload(fields ...field.RelationField) *publicFunctionDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p publicFunctionDo) FirstOrInit() (*model.PublicFunction, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublicFunction), nil
	}
}

func (p publicFunctionDo) FirstOrCreate() (*model.PublicFunction, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.PublicFunction), nil
	}
}

func (p publicFunctionDo) FindByPage(offset int, limit int) (result []*model.PublicFunction, count int64, err error) {
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

func (p publicFunctionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p publicFunctionDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p publicFunctionDo) Delete(models ...*model.PublicFunction) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *publicFunctionDo) withDO(do gen.Dao) *publicFunctionDo {
	p.DO = *do.(*gen.DO)
	return p
}
