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

func newReservation(db *gorm.DB, opts ...gen.DOOption) reservation {
	_reservation := reservation{}

	_reservation.reservationDo.UseDB(db, opts...)
	_reservation.reservationDo.UseModel(&models.Reservation{})

	tableName := _reservation.reservationDo.TableName()
	_reservation.ALL = field.NewAsterisk(tableName)
	_reservation.ID = field.NewUint(tableName, "id")
	_reservation.CreatedAt = field.NewTime(tableName, "created_at")
	_reservation.UpdatedAt = field.NewTime(tableName, "updated_at")
	_reservation.DeletedAt = field.NewField(tableName, "deleted_at")
	_reservation.GuildID = field.NewString(tableName, "guild_id")
	_reservation.ChannelID = field.NewString(tableName, "channel_id")
	_reservation.GameSubscriptions = reservationHasManyGameSubscriptions{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("GameSubscriptions", "models.GameSubscription"),
	}

	_reservation.YoutubeSubscriptions = reservationHasManyYoutubeSubscriptions{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("YoutubeSubscriptions", "models.YoutubeSubscription"),
	}

	_reservation.BlacklistedUsers = reservationHasManyBlacklistedUsers{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("BlacklistedUsers", "models.BlacklistedUser"),
	}

	_reservation.fillFieldMap()

	return _reservation
}

type reservation struct {
	reservationDo

	ALL               field.Asterisk
	ID                field.Uint
	CreatedAt         field.Time
	UpdatedAt         field.Time
	DeletedAt         field.Field
	GuildID           field.String
	ChannelID         field.String
	GameSubscriptions reservationHasManyGameSubscriptions

	YoutubeSubscriptions reservationHasManyYoutubeSubscriptions

	BlacklistedUsers reservationHasManyBlacklistedUsers

	fieldMap map[string]field.Expr
}

func (r reservation) Table(newTableName string) *reservation {
	r.reservationDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r reservation) As(alias string) *reservation {
	r.reservationDo.DO = *(r.reservationDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *reservation) updateTableName(table string) *reservation {
	r.ALL = field.NewAsterisk(table)
	r.ID = field.NewUint(table, "id")
	r.CreatedAt = field.NewTime(table, "created_at")
	r.UpdatedAt = field.NewTime(table, "updated_at")
	r.DeletedAt = field.NewField(table, "deleted_at")
	r.GuildID = field.NewString(table, "guild_id")
	r.ChannelID = field.NewString(table, "channel_id")

	r.fillFieldMap()

	return r
}

func (r *reservation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *reservation) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 9)
	r.fieldMap["id"] = r.ID
	r.fieldMap["created_at"] = r.CreatedAt
	r.fieldMap["updated_at"] = r.UpdatedAt
	r.fieldMap["deleted_at"] = r.DeletedAt
	r.fieldMap["guild_id"] = r.GuildID
	r.fieldMap["channel_id"] = r.ChannelID

}

func (r reservation) clone(db *gorm.DB) reservation {
	r.reservationDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r reservation) replaceDB(db *gorm.DB) reservation {
	r.reservationDo.ReplaceDB(db)
	return r
}

type reservationHasManyGameSubscriptions struct {
	db *gorm.DB

	field.RelationField
}

func (a reservationHasManyGameSubscriptions) Where(conds ...field.Expr) *reservationHasManyGameSubscriptions {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a reservationHasManyGameSubscriptions) WithContext(ctx context.Context) *reservationHasManyGameSubscriptions {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a reservationHasManyGameSubscriptions) Session(session *gorm.Session) *reservationHasManyGameSubscriptions {
	a.db = a.db.Session(session)
	return &a
}

func (a reservationHasManyGameSubscriptions) Model(m *models.Reservation) *reservationHasManyGameSubscriptionsTx {
	return &reservationHasManyGameSubscriptionsTx{a.db.Model(m).Association(a.Name())}
}

type reservationHasManyGameSubscriptionsTx struct{ tx *gorm.Association }

func (a reservationHasManyGameSubscriptionsTx) Find() (result []*models.GameSubscription, err error) {
	return result, a.tx.Find(&result)
}

func (a reservationHasManyGameSubscriptionsTx) Append(values ...*models.GameSubscription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a reservationHasManyGameSubscriptionsTx) Replace(values ...*models.GameSubscription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a reservationHasManyGameSubscriptionsTx) Delete(values ...*models.GameSubscription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a reservationHasManyGameSubscriptionsTx) Clear() error {
	return a.tx.Clear()
}

func (a reservationHasManyGameSubscriptionsTx) Count() int64 {
	return a.tx.Count()
}

type reservationHasManyYoutubeSubscriptions struct {
	db *gorm.DB

	field.RelationField
}

func (a reservationHasManyYoutubeSubscriptions) Where(conds ...field.Expr) *reservationHasManyYoutubeSubscriptions {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a reservationHasManyYoutubeSubscriptions) WithContext(ctx context.Context) *reservationHasManyYoutubeSubscriptions {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a reservationHasManyYoutubeSubscriptions) Session(session *gorm.Session) *reservationHasManyYoutubeSubscriptions {
	a.db = a.db.Session(session)
	return &a
}

func (a reservationHasManyYoutubeSubscriptions) Model(m *models.Reservation) *reservationHasManyYoutubeSubscriptionsTx {
	return &reservationHasManyYoutubeSubscriptionsTx{a.db.Model(m).Association(a.Name())}
}

type reservationHasManyYoutubeSubscriptionsTx struct{ tx *gorm.Association }

func (a reservationHasManyYoutubeSubscriptionsTx) Find() (result []*models.YoutubeSubscription, err error) {
	return result, a.tx.Find(&result)
}

func (a reservationHasManyYoutubeSubscriptionsTx) Append(values ...*models.YoutubeSubscription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a reservationHasManyYoutubeSubscriptionsTx) Replace(values ...*models.YoutubeSubscription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a reservationHasManyYoutubeSubscriptionsTx) Delete(values ...*models.YoutubeSubscription) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a reservationHasManyYoutubeSubscriptionsTx) Clear() error {
	return a.tx.Clear()
}

func (a reservationHasManyYoutubeSubscriptionsTx) Count() int64 {
	return a.tx.Count()
}

type reservationHasManyBlacklistedUsers struct {
	db *gorm.DB

	field.RelationField
}

func (a reservationHasManyBlacklistedUsers) Where(conds ...field.Expr) *reservationHasManyBlacklistedUsers {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a reservationHasManyBlacklistedUsers) WithContext(ctx context.Context) *reservationHasManyBlacklistedUsers {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a reservationHasManyBlacklistedUsers) Session(session *gorm.Session) *reservationHasManyBlacklistedUsers {
	a.db = a.db.Session(session)
	return &a
}

func (a reservationHasManyBlacklistedUsers) Model(m *models.Reservation) *reservationHasManyBlacklistedUsersTx {
	return &reservationHasManyBlacklistedUsersTx{a.db.Model(m).Association(a.Name())}
}

type reservationHasManyBlacklistedUsersTx struct{ tx *gorm.Association }

func (a reservationHasManyBlacklistedUsersTx) Find() (result []*models.BlacklistedUser, err error) {
	return result, a.tx.Find(&result)
}

func (a reservationHasManyBlacklistedUsersTx) Append(values ...*models.BlacklistedUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a reservationHasManyBlacklistedUsersTx) Replace(values ...*models.BlacklistedUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a reservationHasManyBlacklistedUsersTx) Delete(values ...*models.BlacklistedUser) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a reservationHasManyBlacklistedUsersTx) Clear() error {
	return a.tx.Clear()
}

func (a reservationHasManyBlacklistedUsersTx) Count() int64 {
	return a.tx.Count()
}

type reservationDo struct{ gen.DO }

type IReservationDo interface {
	gen.SubQuery
	Debug() IReservationDo
	WithContext(ctx context.Context) IReservationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IReservationDo
	WriteDB() IReservationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IReservationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IReservationDo
	Not(conds ...gen.Condition) IReservationDo
	Or(conds ...gen.Condition) IReservationDo
	Select(conds ...field.Expr) IReservationDo
	Where(conds ...gen.Condition) IReservationDo
	Order(conds ...field.Expr) IReservationDo
	Distinct(cols ...field.Expr) IReservationDo
	Omit(cols ...field.Expr) IReservationDo
	Join(table schema.Tabler, on ...field.Expr) IReservationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IReservationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IReservationDo
	Group(cols ...field.Expr) IReservationDo
	Having(conds ...gen.Condition) IReservationDo
	Limit(limit int) IReservationDo
	Offset(offset int) IReservationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IReservationDo
	Unscoped() IReservationDo
	Create(values ...*models.Reservation) error
	CreateInBatches(values []*models.Reservation, batchSize int) error
	Save(values ...*models.Reservation) error
	First() (*models.Reservation, error)
	Take() (*models.Reservation, error)
	Last() (*models.Reservation, error)
	Find() ([]*models.Reservation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Reservation, err error)
	FindInBatches(result *[]*models.Reservation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Reservation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IReservationDo
	Assign(attrs ...field.AssignExpr) IReservationDo
	Joins(fields ...field.RelationField) IReservationDo
	Preload(fields ...field.RelationField) IReservationDo
	FirstOrInit() (*models.Reservation, error)
	FirstOrCreate() (*models.Reservation, error)
	FindByPage(offset int, limit int) (result []*models.Reservation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IReservationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r reservationDo) Debug() IReservationDo {
	return r.withDO(r.DO.Debug())
}

func (r reservationDo) WithContext(ctx context.Context) IReservationDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r reservationDo) ReadDB() IReservationDo {
	return r.Clauses(dbresolver.Read)
}

func (r reservationDo) WriteDB() IReservationDo {
	return r.Clauses(dbresolver.Write)
}

func (r reservationDo) Session(config *gorm.Session) IReservationDo {
	return r.withDO(r.DO.Session(config))
}

func (r reservationDo) Clauses(conds ...clause.Expression) IReservationDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r reservationDo) Returning(value interface{}, columns ...string) IReservationDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r reservationDo) Not(conds ...gen.Condition) IReservationDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r reservationDo) Or(conds ...gen.Condition) IReservationDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r reservationDo) Select(conds ...field.Expr) IReservationDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r reservationDo) Where(conds ...gen.Condition) IReservationDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r reservationDo) Order(conds ...field.Expr) IReservationDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r reservationDo) Distinct(cols ...field.Expr) IReservationDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r reservationDo) Omit(cols ...field.Expr) IReservationDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r reservationDo) Join(table schema.Tabler, on ...field.Expr) IReservationDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r reservationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IReservationDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r reservationDo) RightJoin(table schema.Tabler, on ...field.Expr) IReservationDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r reservationDo) Group(cols ...field.Expr) IReservationDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r reservationDo) Having(conds ...gen.Condition) IReservationDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r reservationDo) Limit(limit int) IReservationDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r reservationDo) Offset(offset int) IReservationDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r reservationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IReservationDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r reservationDo) Unscoped() IReservationDo {
	return r.withDO(r.DO.Unscoped())
}

func (r reservationDo) Create(values ...*models.Reservation) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r reservationDo) CreateInBatches(values []*models.Reservation, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r reservationDo) Save(values ...*models.Reservation) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r reservationDo) First() (*models.Reservation, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Reservation), nil
	}
}

func (r reservationDo) Take() (*models.Reservation, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Reservation), nil
	}
}

func (r reservationDo) Last() (*models.Reservation, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Reservation), nil
	}
}

func (r reservationDo) Find() ([]*models.Reservation, error) {
	result, err := r.DO.Find()
	return result.([]*models.Reservation), err
}

func (r reservationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Reservation, err error) {
	buf := make([]*models.Reservation, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r reservationDo) FindInBatches(result *[]*models.Reservation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r reservationDo) Attrs(attrs ...field.AssignExpr) IReservationDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r reservationDo) Assign(attrs ...field.AssignExpr) IReservationDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r reservationDo) Joins(fields ...field.RelationField) IReservationDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r reservationDo) Preload(fields ...field.RelationField) IReservationDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r reservationDo) FirstOrInit() (*models.Reservation, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Reservation), nil
	}
}

func (r reservationDo) FirstOrCreate() (*models.Reservation, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Reservation), nil
	}
}

func (r reservationDo) FindByPage(offset int, limit int) (result []*models.Reservation, count int64, err error) {
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

func (r reservationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r reservationDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r reservationDo) Delete(models ...*models.Reservation) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *reservationDo) withDO(do gen.Dao) *reservationDo {
	r.DO = *do.(*gen.DO)
	return r
}
