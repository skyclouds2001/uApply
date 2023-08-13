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
	enrollTimeFieldNames          = builder.RawFieldNames(&EnrollTime{})
	enrollTimeRows                = strings.Join(enrollTimeFieldNames, ",")
	enrollTimeRowsExpectAutoSet   = strings.Join(stringx.Remove(enrollTimeFieldNames, "`create_time`", "`update_time`"), ",")
	enrollTimeRowsWithPlaceHolder = strings.Join(stringx.Remove(enrollTimeFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheEnrollTimeIdPrefix = "cache:enrollTime:id:"
	cacheAllTimeKey         = "cache:enrollTime:all"
)

type (
	EnrollTimeModel interface {
		Insert(data *EnrollTime) (sql.Result, error)
		FindOne(id int64) (*EnrollTime, error)
		Update(data *EnrollTime) error
		Delete(id int64) error
		FindAll() ([]*EnrollTime, error)
	}

	defaultEnrollTimeModel struct {
		sqlc.CachedConn
		table string
	}

	EnrollTime struct {
		Id        int64  `db:"id"`         // 组织id
		StartTime int64  `db:"start_time"` // 开始的时间戳
		EndTime   int64  `db:"end_time"`   // 结束的时间戳
		Name      string `db:"name"`
	}
)

func NewEnrollTimeModel(conn sqlx.SqlConn, c cache.CacheConf) EnrollTimeModel {
	return &defaultEnrollTimeModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`enroll_time`",
	}
}

func (m *defaultEnrollTimeModel) Insert(data *EnrollTime) (sql.Result, error) {
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, enrollTimeRowsExpectAutoSet)
		return conn.Exec(query, data.Id, data.StartTime, data.EndTime, data.Name)
	}, cacheAllTimeKey)
	return ret, err
}

func (m *defaultEnrollTimeModel) FindAll() ([]*EnrollTime, error) {
	var resp []*EnrollTime
	err := m.QueryRow(&resp, cacheAllTimeKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s", enrollTimeRows, m.table)
		return conn.QueryRows(v, query)
	})
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultEnrollTimeModel) FindOne(id int64) (*EnrollTime, error) {
	enrollTimeIdKey := fmt.Sprintf("%s%v", cacheEnrollTimeIdPrefix, id)
	var resp EnrollTime
	err := m.QueryRow(&resp, enrollTimeIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", enrollTimeRows, m.table)
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

func (m *defaultEnrollTimeModel) Update(data *EnrollTime) error {
	enrollTimeIdKey := fmt.Sprintf("%s%v", cacheEnrollTimeIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, enrollTimeRowsWithPlaceHolder)
		return conn.Exec(query, data.StartTime, data.EndTime, data.Name, data.Id)
	}, enrollTimeIdKey, cacheAllTimeKey)
	return err
}

func (m *defaultEnrollTimeModel) Delete(id int64) error {

	enrollTimeIdKey := fmt.Sprintf("%s%v", cacheEnrollTimeIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, enrollTimeIdKey, cacheAllTimeKey)
	return err
}

func (m *defaultEnrollTimeModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheEnrollTimeIdPrefix, primary)
}

func (m *defaultEnrollTimeModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", enrollTimeRows, m.table)
	return conn.QueryRow(v, query, primary)
}
