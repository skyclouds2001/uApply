package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	interInfoFieldNames          = builder.RawFieldNames(&InterInfo{})
	interInfoRows                = strings.Join(interInfoFieldNames, ",")
	interInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(interInfoFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	interInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(interInfoFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheInterInfoIdPrefix = "cache:interInfo:id:"
)

type (
	InterInfoModel interface {
		Insert(data *InterInfo) (sql.Result, error)
		FindOne(id int64) (*InterInfo, error)
		FindOneByUidAndDepId(uid int64, depId int64) (*InterInfo, error)
		Update(data *InterInfo) error
		Delete(id int64) error
		FindByUid(uid int64) ([]*InterInfo, error)
		FindByDepId(depId int64) ([]*InterInfo, error)
		DeleteByUidAndDepId(uid int32, depid int64) error
	}

	defaultInterInfoModel struct {
		sqlc.CachedConn
		table string
	}

	InterInfo struct {
		Id               int64  `db:"id"`
		Uid              int64  `db:"uid"`
		Name             string `db:"name"`
		OrganizationId   int64  `db:"organization_id"`
		OrganizationName string `db:"organization_name"`
		DepartmentId     int64  `db:"department_id"`
		DepartmentName   string `db:"department_name"`
	}
)

func NewInterInfoModel(conn sqlx.SqlConn, c cache.CacheConf) InterInfoModel {
	return &defaultInterInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`inter_info`",
	}
}

func (m *defaultInterInfoModel) FindByUid(uid int64) ([]*InterInfo, error) {
	var interInfo []*InterInfo
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `uid` = ? ", interInfoRows, m.table)
	err := m.QueryRowsNoCache(&interInfo, query, uid)

	switch err {
	case nil:
		return interInfo, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInterInfoModel) FindByDepId(depId int64) ([]*InterInfo, error) {
	var interInfo []*InterInfo
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `department_id` = ? ", interInfoRows, m.table)
	err := m.QueryRowsNoCache(&interInfo, query, depId)

	switch err {
	case nil:
		return interInfo, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInterInfoModel) Insert(data *InterInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, interInfoRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Uid, data.Name, data.OrganizationId, data.OrganizationName, data.DepartmentId, data.DepartmentName)

	return ret, err
}

func (m *defaultInterInfoModel) FindOne(id int64) (*InterInfo, error) {
	interInfoIdKey := fmt.Sprintf("%s%v", cacheInterInfoIdPrefix, id)
	var resp InterInfo
	err := m.QueryRow(&resp, interInfoIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", interInfoRows, m.table)
		return conn.QueryRow(v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInterInfoModel) FindOneByUidAndDepId(uid int64, depId int64) (*InterInfo, error) {
	var resp InterInfo
	query := fmt.Sprintf("select %s from %s where `uid` = ? and `department_id` = ? limit 1", interInfoRows, m.table)
	err := m.QueryRowNoCache(&resp, query, uid, depId)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultInterInfoModel) Update(data *InterInfo) error {
	interInfoIdKey := fmt.Sprintf("%s%v", cacheInterInfoIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, interInfoRowsWithPlaceHolder)
		return conn.Exec(query, data.Uid, data.Name, data.OrganizationId, data.OrganizationName, data.DepartmentId, data.DepartmentName, data.Id)
	}, interInfoIdKey)
	return err
}

func (m *defaultInterInfoModel) Delete(id int64) error {

	interInfoIdKey := fmt.Sprintf("%s%v", cacheInterInfoIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, interInfoIdKey)
	return err
}

func (m *defaultInterInfoModel) DeleteByUidAndDepId(uid int32, depid int64) error {
	query := fmt.Sprintf("delete from %s where `uid` = ? and `department_id` = ?", m.table)
	_, err := m.ExecNoCache(query, uid, depid)
	return err
}

func (m *defaultInterInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheInterInfoIdPrefix, primary)
}

func (m *defaultInterInfoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", interInfoRows, m.table)
	return conn.QueryRow(v, query, primary)
}
