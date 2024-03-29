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

	"github.com/Runner-Go-Team/RunnerGo-management-websocket-open/internal/pkg/dal/model"
)

func newUIPlan(db *gorm.DB, opts ...gen.DOOption) uIPlan {
	_uIPlan := uIPlan{}

	_uIPlan.uIPlanDo.UseDB(db, opts...)
	_uIPlan.uIPlanDo.UseModel(&model.UIPlan{})

	tableName := _uIPlan.uIPlanDo.TableName()
	_uIPlan.ALL = field.NewAsterisk(tableName)
	_uIPlan.ID = field.NewInt64(tableName, "id")
	_uIPlan.PlanID = field.NewString(tableName, "plan_id")
	_uIPlan.TeamID = field.NewString(tableName, "team_id")
	_uIPlan.RankID = field.NewInt64(tableName, "rank_id")
	_uIPlan.Name = field.NewString(tableName, "name")
	_uIPlan.TaskType = field.NewInt32(tableName, "task_type")
	_uIPlan.CreateUserID = field.NewString(tableName, "create_user_id")
	_uIPlan.HeadUserID = field.NewString(tableName, "head_user_id")
	_uIPlan.RunCount = field.NewInt64(tableName, "run_count")
	_uIPlan.InitStrategy = field.NewInt32(tableName, "init_strategy")
	_uIPlan.Description = field.NewString(tableName, "description")
	_uIPlan.Browsers = field.NewString(tableName, "browsers")
	_uIPlan.UIMachineKey = field.NewString(tableName, "ui_machine_key")
	_uIPlan.CreatedAt = field.NewTime(tableName, "created_at")
	_uIPlan.UpdatedAt = field.NewTime(tableName, "updated_at")
	_uIPlan.DeletedAt = field.NewField(tableName, "deleted_at")

	_uIPlan.fillFieldMap()

	return _uIPlan
}

type uIPlan struct {
	uIPlanDo uIPlanDo

	ALL          field.Asterisk
	ID           field.Int64  // 主键ID
	PlanID       field.String // 计划ID
	TeamID       field.String // 团队ID
	RankID       field.Int64  // 序号ID
	Name         field.String // 计划名称
	TaskType     field.Int32  // 计划类型：1-普通任务，2-定时任务
	CreateUserID field.String // 创建人id
	HeadUserID   field.String // 负责人id ,用分割
	RunCount     field.Int64  // 运行次数
	InitStrategy field.Int32  // 初始化策略：1-计划执行前重启浏览器，2-场景执行前重启浏览器，3-无初始化
	Description  field.String // 备注
	Browsers     field.String // 浏览器信息
	UIMachineKey field.String // 指定机器key
	CreatedAt    field.Time   // 创建时间
	UpdatedAt    field.Time   // 修改时间
	DeletedAt    field.Field  // 删除时间

	fieldMap map[string]field.Expr
}

func (u uIPlan) Table(newTableName string) *uIPlan {
	u.uIPlanDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u uIPlan) As(alias string) *uIPlan {
	u.uIPlanDo.DO = *(u.uIPlanDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *uIPlan) updateTableName(table string) *uIPlan {
	u.ALL = field.NewAsterisk(table)
	u.ID = field.NewInt64(table, "id")
	u.PlanID = field.NewString(table, "plan_id")
	u.TeamID = field.NewString(table, "team_id")
	u.RankID = field.NewInt64(table, "rank_id")
	u.Name = field.NewString(table, "name")
	u.TaskType = field.NewInt32(table, "task_type")
	u.CreateUserID = field.NewString(table, "create_user_id")
	u.HeadUserID = field.NewString(table, "head_user_id")
	u.RunCount = field.NewInt64(table, "run_count")
	u.InitStrategy = field.NewInt32(table, "init_strategy")
	u.Description = field.NewString(table, "description")
	u.Browsers = field.NewString(table, "browsers")
	u.UIMachineKey = field.NewString(table, "ui_machine_key")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")
	u.DeletedAt = field.NewField(table, "deleted_at")

	u.fillFieldMap()

	return u
}

func (u *uIPlan) WithContext(ctx context.Context) *uIPlanDo { return u.uIPlanDo.WithContext(ctx) }

func (u uIPlan) TableName() string { return u.uIPlanDo.TableName() }

func (u uIPlan) Alias() string { return u.uIPlanDo.Alias() }

func (u *uIPlan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *uIPlan) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 16)
	u.fieldMap["id"] = u.ID
	u.fieldMap["plan_id"] = u.PlanID
	u.fieldMap["team_id"] = u.TeamID
	u.fieldMap["rank_id"] = u.RankID
	u.fieldMap["name"] = u.Name
	u.fieldMap["task_type"] = u.TaskType
	u.fieldMap["create_user_id"] = u.CreateUserID
	u.fieldMap["head_user_id"] = u.HeadUserID
	u.fieldMap["run_count"] = u.RunCount
	u.fieldMap["init_strategy"] = u.InitStrategy
	u.fieldMap["description"] = u.Description
	u.fieldMap["browsers"] = u.Browsers
	u.fieldMap["ui_machine_key"] = u.UIMachineKey
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
	u.fieldMap["deleted_at"] = u.DeletedAt
}

func (u uIPlan) clone(db *gorm.DB) uIPlan {
	u.uIPlanDo.ReplaceConnPool(db.Statement.ConnPool)
	return u
}

func (u uIPlan) replaceDB(db *gorm.DB) uIPlan {
	u.uIPlanDo.ReplaceDB(db)
	return u
}

type uIPlanDo struct{ gen.DO }

func (u uIPlanDo) Debug() *uIPlanDo {
	return u.withDO(u.DO.Debug())
}

func (u uIPlanDo) WithContext(ctx context.Context) *uIPlanDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u uIPlanDo) ReadDB() *uIPlanDo {
	return u.Clauses(dbresolver.Read)
}

func (u uIPlanDo) WriteDB() *uIPlanDo {
	return u.Clauses(dbresolver.Write)
}

func (u uIPlanDo) Session(config *gorm.Session) *uIPlanDo {
	return u.withDO(u.DO.Session(config))
}

func (u uIPlanDo) Clauses(conds ...clause.Expression) *uIPlanDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u uIPlanDo) Returning(value interface{}, columns ...string) *uIPlanDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u uIPlanDo) Not(conds ...gen.Condition) *uIPlanDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u uIPlanDo) Or(conds ...gen.Condition) *uIPlanDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u uIPlanDo) Select(conds ...field.Expr) *uIPlanDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u uIPlanDo) Where(conds ...gen.Condition) *uIPlanDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u uIPlanDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *uIPlanDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u uIPlanDo) Order(conds ...field.Expr) *uIPlanDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u uIPlanDo) Distinct(cols ...field.Expr) *uIPlanDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u uIPlanDo) Omit(cols ...field.Expr) *uIPlanDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u uIPlanDo) Join(table schema.Tabler, on ...field.Expr) *uIPlanDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u uIPlanDo) LeftJoin(table schema.Tabler, on ...field.Expr) *uIPlanDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u uIPlanDo) RightJoin(table schema.Tabler, on ...field.Expr) *uIPlanDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u uIPlanDo) Group(cols ...field.Expr) *uIPlanDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u uIPlanDo) Having(conds ...gen.Condition) *uIPlanDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u uIPlanDo) Limit(limit int) *uIPlanDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u uIPlanDo) Offset(offset int) *uIPlanDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u uIPlanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *uIPlanDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u uIPlanDo) Unscoped() *uIPlanDo {
	return u.withDO(u.DO.Unscoped())
}

func (u uIPlanDo) Create(values ...*model.UIPlan) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u uIPlanDo) CreateInBatches(values []*model.UIPlan, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u uIPlanDo) Save(values ...*model.UIPlan) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u uIPlanDo) First() (*model.UIPlan, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIPlan), nil
	}
}

func (u uIPlanDo) Take() (*model.UIPlan, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIPlan), nil
	}
}

func (u uIPlanDo) Last() (*model.UIPlan, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIPlan), nil
	}
}

func (u uIPlanDo) Find() ([]*model.UIPlan, error) {
	result, err := u.DO.Find()
	return result.([]*model.UIPlan), err
}

func (u uIPlanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UIPlan, err error) {
	buf := make([]*model.UIPlan, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u uIPlanDo) FindInBatches(result *[]*model.UIPlan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u uIPlanDo) Attrs(attrs ...field.AssignExpr) *uIPlanDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u uIPlanDo) Assign(attrs ...field.AssignExpr) *uIPlanDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u uIPlanDo) Joins(fields ...field.RelationField) *uIPlanDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u uIPlanDo) Preload(fields ...field.RelationField) *uIPlanDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u uIPlanDo) FirstOrInit() (*model.UIPlan, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIPlan), nil
	}
}

func (u uIPlanDo) FirstOrCreate() (*model.UIPlan, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UIPlan), nil
	}
}

func (u uIPlanDo) FindByPage(offset int, limit int) (result []*model.UIPlan, count int64, err error) {
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

func (u uIPlanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u uIPlanDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u uIPlanDo) Delete(models ...*model.UIPlan) (result gen.ResultInfo, err error) {
	return u.DO.Delete(models)
}

func (u *uIPlanDo) withDO(do gen.Dao) *uIPlanDo {
	u.DO = *do.(*gen.DO)
	return u
}
