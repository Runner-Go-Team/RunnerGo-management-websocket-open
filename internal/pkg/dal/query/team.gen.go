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

func newTeam(db *gorm.DB, opts ...gen.DOOption) team {
	_team := team{}

	_team.teamDo.UseDB(db, opts...)
	_team.teamDo.UseModel(&model.Team{})

	tableName := _team.teamDo.TableName()
	_team.ALL = field.NewAsterisk(tableName)
	_team.ID = field.NewInt64(tableName, "id")
	_team.TeamID = field.NewString(tableName, "team_id")
	_team.Name = field.NewString(tableName, "name")
	_team.Type = field.NewInt32(tableName, "type")
	_team.TrialExpirationDate = field.NewTime(tableName, "trial_expiration_date")
	_team.IsVip = field.NewInt32(tableName, "is_vip")
	_team.VipExpirationDate = field.NewTime(tableName, "vip_expiration_date")
	_team.VumNum = field.NewInt64(tableName, "vum_num")
	_team.MaxUserNum = field.NewInt64(tableName, "max_user_num")
	_team.CreatedUserID = field.NewString(tableName, "created_user_id")
	_team.TeamBuyVersionType = field.NewInt32(tableName, "team_buy_version_type")
	_team.CreatedAt = field.NewTime(tableName, "created_at")
	_team.UpdatedAt = field.NewTime(tableName, "updated_at")
	_team.DeletedAt = field.NewField(tableName, "deleted_at")

	_team.fillFieldMap()

	return _team
}

type team struct {
	teamDo teamDo

	ALL                 field.Asterisk
	ID                  field.Int64  // 主键ID
	TeamID              field.String // 团队ID
	Name                field.String // 团队名称
	Type                field.Int32  // 团队类型 1: 私有团队；2: 普通团队
	TrialExpirationDate field.Time   // 试用有效期
	IsVip               field.Int32  // 是否为付费团队 1-否 2-是
	VipExpirationDate   field.Time   // 付费有效期
	VumNum              field.Int64  // 当前可用VUM总数
	MaxUserNum          field.Int64  // 当前团队最大成员数量
	CreatedUserID       field.String // 创建者id
	TeamBuyVersionType  field.Int32  // 团队套餐类型：1-个人版，2-团队版，3-企业版，4-私有化部署
	CreatedAt           field.Time
	UpdatedAt           field.Time
	DeletedAt           field.Field

	fieldMap map[string]field.Expr
}

func (t team) Table(newTableName string) *team {
	t.teamDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t team) As(alias string) *team {
	t.teamDo.DO = *(t.teamDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *team) updateTableName(table string) *team {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.TeamID = field.NewString(table, "team_id")
	t.Name = field.NewString(table, "name")
	t.Type = field.NewInt32(table, "type")
	t.TrialExpirationDate = field.NewTime(table, "trial_expiration_date")
	t.IsVip = field.NewInt32(table, "is_vip")
	t.VipExpirationDate = field.NewTime(table, "vip_expiration_date")
	t.VumNum = field.NewInt64(table, "vum_num")
	t.MaxUserNum = field.NewInt64(table, "max_user_num")
	t.CreatedUserID = field.NewString(table, "created_user_id")
	t.TeamBuyVersionType = field.NewInt32(table, "team_buy_version_type")
	t.CreatedAt = field.NewTime(table, "created_at")
	t.UpdatedAt = field.NewTime(table, "updated_at")
	t.DeletedAt = field.NewField(table, "deleted_at")

	t.fillFieldMap()

	return t
}

func (t *team) WithContext(ctx context.Context) *teamDo { return t.teamDo.WithContext(ctx) }

func (t team) TableName() string { return t.teamDo.TableName() }

func (t team) Alias() string { return t.teamDo.Alias() }

func (t *team) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *team) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 14)
	t.fieldMap["id"] = t.ID
	t.fieldMap["team_id"] = t.TeamID
	t.fieldMap["name"] = t.Name
	t.fieldMap["type"] = t.Type
	t.fieldMap["trial_expiration_date"] = t.TrialExpirationDate
	t.fieldMap["is_vip"] = t.IsVip
	t.fieldMap["vip_expiration_date"] = t.VipExpirationDate
	t.fieldMap["vum_num"] = t.VumNum
	t.fieldMap["max_user_num"] = t.MaxUserNum
	t.fieldMap["created_user_id"] = t.CreatedUserID
	t.fieldMap["team_buy_version_type"] = t.TeamBuyVersionType
	t.fieldMap["created_at"] = t.CreatedAt
	t.fieldMap["updated_at"] = t.UpdatedAt
	t.fieldMap["deleted_at"] = t.DeletedAt
}

func (t team) clone(db *gorm.DB) team {
	t.teamDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t team) replaceDB(db *gorm.DB) team {
	t.teamDo.ReplaceDB(db)
	return t
}

type teamDo struct{ gen.DO }

func (t teamDo) Debug() *teamDo {
	return t.withDO(t.DO.Debug())
}

func (t teamDo) WithContext(ctx context.Context) *teamDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t teamDo) ReadDB() *teamDo {
	return t.Clauses(dbresolver.Read)
}

func (t teamDo) WriteDB() *teamDo {
	return t.Clauses(dbresolver.Write)
}

func (t teamDo) Session(config *gorm.Session) *teamDo {
	return t.withDO(t.DO.Session(config))
}

func (t teamDo) Clauses(conds ...clause.Expression) *teamDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t teamDo) Returning(value interface{}, columns ...string) *teamDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t teamDo) Not(conds ...gen.Condition) *teamDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t teamDo) Or(conds ...gen.Condition) *teamDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t teamDo) Select(conds ...field.Expr) *teamDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t teamDo) Where(conds ...gen.Condition) *teamDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t teamDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *teamDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t teamDo) Order(conds ...field.Expr) *teamDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t teamDo) Distinct(cols ...field.Expr) *teamDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t teamDo) Omit(cols ...field.Expr) *teamDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t teamDo) Join(table schema.Tabler, on ...field.Expr) *teamDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t teamDo) LeftJoin(table schema.Tabler, on ...field.Expr) *teamDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t teamDo) RightJoin(table schema.Tabler, on ...field.Expr) *teamDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t teamDo) Group(cols ...field.Expr) *teamDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t teamDo) Having(conds ...gen.Condition) *teamDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t teamDo) Limit(limit int) *teamDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t teamDo) Offset(offset int) *teamDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t teamDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *teamDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t teamDo) Unscoped() *teamDo {
	return t.withDO(t.DO.Unscoped())
}

func (t teamDo) Create(values ...*model.Team) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t teamDo) CreateInBatches(values []*model.Team, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t teamDo) Save(values ...*model.Team) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t teamDo) First() (*model.Team, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Team), nil
	}
}

func (t teamDo) Take() (*model.Team, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Team), nil
	}
}

func (t teamDo) Last() (*model.Team, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Team), nil
	}
}

func (t teamDo) Find() ([]*model.Team, error) {
	result, err := t.DO.Find()
	return result.([]*model.Team), err
}

func (t teamDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Team, err error) {
	buf := make([]*model.Team, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t teamDo) FindInBatches(result *[]*model.Team, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t teamDo) Attrs(attrs ...field.AssignExpr) *teamDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t teamDo) Assign(attrs ...field.AssignExpr) *teamDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t teamDo) Joins(fields ...field.RelationField) *teamDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t teamDo) Preload(fields ...field.RelationField) *teamDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t teamDo) FirstOrInit() (*model.Team, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Team), nil
	}
}

func (t teamDo) FirstOrCreate() (*model.Team, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Team), nil
	}
}

func (t teamDo) FindByPage(offset int, limit int) (result []*model.Team, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t teamDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t teamDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t teamDo) Delete(models ...*model.Team) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *teamDo) withDO(do gen.Dao) *teamDo {
	t.DO = *do.(*gen.DO)
	return t
}
