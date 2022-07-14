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

	"github.com/mc_nft/dal/model"
)

func newUserProp(db *gorm.DB) userProp {
	_userProp := userProp{}

	_userProp.userPropDo.UseDB(db)
	_userProp.userPropDo.UseModel(&model.UserProp{})

	tableName := _userProp.userPropDo.TableName()
	_userProp.ALL = field.NewField(tableName, "*")
	_userProp.ID = field.NewInt64(tableName, "id")
	_userProp.UserID = field.NewInt64(tableName, "user_id")
	_userProp.PropsID = field.NewInt64(tableName, "props_id")
	_userProp.Num = field.NewInt32(tableName, "num")
	_userProp.CreatedAt = field.NewTime(tableName, "created_at")
	_userProp.UpdatedAt = field.NewTime(tableName, "updated_at")

	_userProp.fillFieldMap()

	return _userProp
}

type userProp struct {
	userPropDo userPropDo

	ALL       field.Field
	ID        field.Int64
	UserID    field.Int64
	PropsID   field.Int64
	Num       field.Int32
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (u userProp) Table(newTableName string) *userProp {
	u.userPropDo.UseTable(newTableName)
	return u.updateTableName(newTableName)
}

func (u userProp) As(alias string) *userProp {
	u.userPropDo.DO = *(u.userPropDo.As(alias).(*gen.DO))
	return u.updateTableName(alias)
}

func (u *userProp) updateTableName(table string) *userProp {
	u.ALL = field.NewField(table, "*")
	u.ID = field.NewInt64(table, "id")
	u.UserID = field.NewInt64(table, "user_id")
	u.PropsID = field.NewInt64(table, "props_id")
	u.Num = field.NewInt32(table, "num")
	u.CreatedAt = field.NewTime(table, "created_at")
	u.UpdatedAt = field.NewTime(table, "updated_at")

	u.fillFieldMap()

	return u
}

func (u *userProp) WithContext(ctx context.Context) *userPropDo { return u.userPropDo.WithContext(ctx) }

func (u userProp) TableName() string { return u.userPropDo.TableName() }

func (u userProp) Alias() string { return u.userPropDo.Alias() }

func (u *userProp) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (u *userProp) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 6)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["props_id"] = u.PropsID
	u.fieldMap["num"] = u.Num
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
}

func (u userProp) clone(db *gorm.DB) userProp {
	u.userPropDo.ReplaceDB(db)
	return u
}

type userPropDo struct{ gen.DO }

func (u userPropDo) Debug() *userPropDo {
	return u.withDO(u.DO.Debug())
}

func (u userPropDo) WithContext(ctx context.Context) *userPropDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userPropDo) ReadDB() *userPropDo {
	return u.Clauses(dbresolver.Read)
}

func (u userPropDo) WriteDB() *userPropDo {
	return u.Clauses(dbresolver.Write)
}

func (u userPropDo) Clauses(conds ...clause.Expression) *userPropDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userPropDo) Returning(value interface{}, columns ...string) *userPropDo {
	return u.withDO(u.DO.Returning(value, columns...))
}

func (u userPropDo) Not(conds ...gen.Condition) *userPropDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userPropDo) Or(conds ...gen.Condition) *userPropDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userPropDo) Select(conds ...field.Expr) *userPropDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userPropDo) Where(conds ...gen.Condition) *userPropDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userPropDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) *userPropDo {
	return u.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (u userPropDo) Order(conds ...field.Expr) *userPropDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userPropDo) Distinct(cols ...field.Expr) *userPropDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userPropDo) Omit(cols ...field.Expr) *userPropDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userPropDo) Join(table schema.Tabler, on ...field.Expr) *userPropDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userPropDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userPropDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userPropDo) RightJoin(table schema.Tabler, on ...field.Expr) *userPropDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userPropDo) Group(cols ...field.Expr) *userPropDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userPropDo) Having(conds ...gen.Condition) *userPropDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userPropDo) Limit(limit int) *userPropDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userPropDo) Offset(offset int) *userPropDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userPropDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userPropDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userPropDo) Unscoped() *userPropDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userPropDo) Create(values ...*model.UserProp) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userPropDo) CreateInBatches(values []*model.UserProp, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userPropDo) Save(values ...*model.UserProp) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userPropDo) First() (*model.UserProp, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProp), nil
	}
}

func (u userPropDo) Take() (*model.UserProp, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProp), nil
	}
}

func (u userPropDo) Last() (*model.UserProp, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProp), nil
	}
}

func (u userPropDo) Find() ([]*model.UserProp, error) {
	result, err := u.DO.Find()
	return result.([]*model.UserProp), err
}

func (u userPropDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.UserProp, err error) {
	buf := make([]*model.UserProp, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userPropDo) FindInBatches(result *[]*model.UserProp, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userPropDo) Attrs(attrs ...field.AssignExpr) *userPropDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userPropDo) Assign(attrs ...field.AssignExpr) *userPropDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userPropDo) Joins(fields ...field.RelationField) *userPropDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Joins(_f))
	}
	return &u
}

func (u userPropDo) Preload(fields ...field.RelationField) *userPropDo {
	for _, _f := range fields {
		u = *u.withDO(u.DO.Preload(_f))
	}
	return &u
}

func (u userPropDo) FirstOrInit() (*model.UserProp, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProp), nil
	}
}

func (u userPropDo) FirstOrCreate() (*model.UserProp, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.UserProp), nil
	}
}

func (u userPropDo) FindByPage(offset int, limit int) (result []*model.UserProp, count int64, err error) {
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

func (u userPropDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u userPropDo) Scan(result interface{}) (err error) {
	return u.DO.Scan(result)
}

func (u *userPropDo) withDO(do gen.Dao) *userPropDo {
	u.DO = *do.(*gen.DO)
	return u
}
