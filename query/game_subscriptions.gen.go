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

	"streambot/models"
)

func newGameSubscription(db *gorm.DB, opts ...gen.DOOption) gameSubscription {
	_gameSubscription := gameSubscription{}

	_gameSubscription.gameSubscriptionDo.UseDB(db, opts...)
	_gameSubscription.gameSubscriptionDo.UseModel(&models.GameSubscription{})

	tableName := _gameSubscription.gameSubscriptionDo.TableName()
	_gameSubscription.ALL = field.NewAsterisk(tableName)
	_gameSubscription.ID = field.NewUint(tableName, "id")
	_gameSubscription.CreatedAt = field.NewTime(tableName, "created_at")
	_gameSubscription.UpdatedAt = field.NewTime(tableName, "updated_at")
	_gameSubscription.DeletedAt = field.NewField(tableName, "deleted_at")
	_gameSubscription.ReservationID = field.NewInt(tableName, "reservation_id")
	_gameSubscription.GameID = field.NewString(tableName, "game_id")
	_gameSubscription.Name = field.NewString(tableName, "name")

	_gameSubscription.fillFieldMap()

	return _gameSubscription
}

type gameSubscription struct {
	gameSubscriptionDo

	ALL           field.Asterisk
	ID            field.Uint
	CreatedAt     field.Time
	UpdatedAt     field.Time
	DeletedAt     field.Field
	ReservationID field.Int
	GameID        field.String
	Name          field.String

	fieldMap map[string]field.Expr
}

func (g gameSubscription) Table(newTableName string) *gameSubscription {
	g.gameSubscriptionDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g gameSubscription) As(alias string) *gameSubscription {
	g.gameSubscriptionDo.DO = *(g.gameSubscriptionDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *gameSubscription) updateTableName(table string) *gameSubscription {
	g.ALL = field.NewAsterisk(table)
	g.ID = field.NewUint(table, "id")
	g.CreatedAt = field.NewTime(table, "created_at")
	g.UpdatedAt = field.NewTime(table, "updated_at")
	g.DeletedAt = field.NewField(table, "deleted_at")
	g.ReservationID = field.NewInt(table, "reservation_id")
	g.GameID = field.NewString(table, "game_id")
	g.Name = field.NewString(table, "name")

	g.fillFieldMap()

	return g
}

func (g *gameSubscription) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *gameSubscription) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 7)
	g.fieldMap["id"] = g.ID
	g.fieldMap["created_at"] = g.CreatedAt
	g.fieldMap["updated_at"] = g.UpdatedAt
	g.fieldMap["deleted_at"] = g.DeletedAt
	g.fieldMap["reservation_id"] = g.ReservationID
	g.fieldMap["game_id"] = g.GameID
	g.fieldMap["name"] = g.Name
}

func (g gameSubscription) clone(db *gorm.DB) gameSubscription {
	g.gameSubscriptionDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g gameSubscription) replaceDB(db *gorm.DB) gameSubscription {
	g.gameSubscriptionDo.ReplaceDB(db)
	return g
}

type gameSubscriptionDo struct{ gen.DO }

type IGameSubscriptionDo interface {
	gen.SubQuery
	Debug() IGameSubscriptionDo
	WithContext(ctx context.Context) IGameSubscriptionDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGameSubscriptionDo
	WriteDB() IGameSubscriptionDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGameSubscriptionDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGameSubscriptionDo
	Not(conds ...gen.Condition) IGameSubscriptionDo
	Or(conds ...gen.Condition) IGameSubscriptionDo
	Select(conds ...field.Expr) IGameSubscriptionDo
	Where(conds ...gen.Condition) IGameSubscriptionDo
	Order(conds ...field.Expr) IGameSubscriptionDo
	Distinct(cols ...field.Expr) IGameSubscriptionDo
	Omit(cols ...field.Expr) IGameSubscriptionDo
	Join(table schema.Tabler, on ...field.Expr) IGameSubscriptionDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGameSubscriptionDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGameSubscriptionDo
	Group(cols ...field.Expr) IGameSubscriptionDo
	Having(conds ...gen.Condition) IGameSubscriptionDo
	Limit(limit int) IGameSubscriptionDo
	Offset(offset int) IGameSubscriptionDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGameSubscriptionDo
	Unscoped() IGameSubscriptionDo
	Create(values ...*models.GameSubscription) error
	CreateInBatches(values []*models.GameSubscription, batchSize int) error
	Save(values ...*models.GameSubscription) error
	First() (*models.GameSubscription, error)
	Take() (*models.GameSubscription, error)
	Last() (*models.GameSubscription, error)
	Find() ([]*models.GameSubscription, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.GameSubscription, err error)
	FindInBatches(result *[]*models.GameSubscription, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.GameSubscription) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGameSubscriptionDo
	Assign(attrs ...field.AssignExpr) IGameSubscriptionDo
	Joins(fields ...field.RelationField) IGameSubscriptionDo
	Preload(fields ...field.RelationField) IGameSubscriptionDo
	FirstOrInit() (*models.GameSubscription, error)
	FirstOrCreate() (*models.GameSubscription, error)
	FindByPage(offset int, limit int) (result []*models.GameSubscription, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGameSubscriptionDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g gameSubscriptionDo) Debug() IGameSubscriptionDo {
	return g.withDO(g.DO.Debug())
}

func (g gameSubscriptionDo) WithContext(ctx context.Context) IGameSubscriptionDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g gameSubscriptionDo) ReadDB() IGameSubscriptionDo {
	return g.Clauses(dbresolver.Read)
}

func (g gameSubscriptionDo) WriteDB() IGameSubscriptionDo {
	return g.Clauses(dbresolver.Write)
}

func (g gameSubscriptionDo) Session(config *gorm.Session) IGameSubscriptionDo {
	return g.withDO(g.DO.Session(config))
}

func (g gameSubscriptionDo) Clauses(conds ...clause.Expression) IGameSubscriptionDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g gameSubscriptionDo) Returning(value interface{}, columns ...string) IGameSubscriptionDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g gameSubscriptionDo) Not(conds ...gen.Condition) IGameSubscriptionDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g gameSubscriptionDo) Or(conds ...gen.Condition) IGameSubscriptionDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g gameSubscriptionDo) Select(conds ...field.Expr) IGameSubscriptionDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g gameSubscriptionDo) Where(conds ...gen.Condition) IGameSubscriptionDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g gameSubscriptionDo) Order(conds ...field.Expr) IGameSubscriptionDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g gameSubscriptionDo) Distinct(cols ...field.Expr) IGameSubscriptionDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g gameSubscriptionDo) Omit(cols ...field.Expr) IGameSubscriptionDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g gameSubscriptionDo) Join(table schema.Tabler, on ...field.Expr) IGameSubscriptionDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g gameSubscriptionDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGameSubscriptionDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g gameSubscriptionDo) RightJoin(table schema.Tabler, on ...field.Expr) IGameSubscriptionDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g gameSubscriptionDo) Group(cols ...field.Expr) IGameSubscriptionDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g gameSubscriptionDo) Having(conds ...gen.Condition) IGameSubscriptionDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g gameSubscriptionDo) Limit(limit int) IGameSubscriptionDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g gameSubscriptionDo) Offset(offset int) IGameSubscriptionDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g gameSubscriptionDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGameSubscriptionDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g gameSubscriptionDo) Unscoped() IGameSubscriptionDo {
	return g.withDO(g.DO.Unscoped())
}

func (g gameSubscriptionDo) Create(values ...*models.GameSubscription) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g gameSubscriptionDo) CreateInBatches(values []*models.GameSubscription, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g gameSubscriptionDo) Save(values ...*models.GameSubscription) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g gameSubscriptionDo) First() (*models.GameSubscription, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.GameSubscription), nil
	}
}

func (g gameSubscriptionDo) Take() (*models.GameSubscription, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.GameSubscription), nil
	}
}

func (g gameSubscriptionDo) Last() (*models.GameSubscription, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.GameSubscription), nil
	}
}

func (g gameSubscriptionDo) Find() ([]*models.GameSubscription, error) {
	result, err := g.DO.Find()
	return result.([]*models.GameSubscription), err
}

func (g gameSubscriptionDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.GameSubscription, err error) {
	buf := make([]*models.GameSubscription, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g gameSubscriptionDo) FindInBatches(result *[]*models.GameSubscription, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g gameSubscriptionDo) Attrs(attrs ...field.AssignExpr) IGameSubscriptionDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g gameSubscriptionDo) Assign(attrs ...field.AssignExpr) IGameSubscriptionDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g gameSubscriptionDo) Joins(fields ...field.RelationField) IGameSubscriptionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g gameSubscriptionDo) Preload(fields ...field.RelationField) IGameSubscriptionDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g gameSubscriptionDo) FirstOrInit() (*models.GameSubscription, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.GameSubscription), nil
	}
}

func (g gameSubscriptionDo) FirstOrCreate() (*models.GameSubscription, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.GameSubscription), nil
	}
}

func (g gameSubscriptionDo) FindByPage(offset int, limit int) (result []*models.GameSubscription, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g gameSubscriptionDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g gameSubscriptionDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g gameSubscriptionDo) Delete(models ...*models.GameSubscription) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *gameSubscriptionDo) withDO(do gen.Dao) *gameSubscriptionDo {
	g.DO = *do.(*gen.DO)
	return g
}
