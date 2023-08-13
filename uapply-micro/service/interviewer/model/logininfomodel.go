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
	loginInfoFieldNames          = builder.RawFieldNames(&LoginInfo{})
	loginInfoRows                = strings.Join(loginInfoFieldNames, ",")
	loginInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(loginInfoFieldNames, "`uid`", "`create_time`", "`update_time`"), ",")
	loginInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(loginInfoFieldNames, "`uid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheLoginInfoUidPrefix = "cache:loginInfo:uid:"
)

type (
	LoginInfoModel interface {
		Insert(data *LoginInfo) (sql.Result, error)
		FindOne(uid int64) (*LoginInfo, error)
		Update(data *LoginInfo) error
		Delete(uid int64) error
		FindOneByOpenId(openId string) (*LoginInfo, error)
	}

	defaultLoginInfoModel struct {
		sqlc.CachedConn
		table string
	}

	LoginInfo struct {
		Uid    int64  `db:"uid"`
		OpenId string `db:"open_id"`
	}
)

func NewLoginInfoModel(conn sqlx.SqlConn, c cache.CacheConf) LoginInfoModel {
	return &defaultLoginInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`login_info`",
	}
}

func (m *defaultLoginInfoModel) FindOneByOpenId(openId string) (*LoginInfo, error) {
	loginInfoOpenIdKey := fmt.Sprintf("%s%v", cacheLoginInfoUidPrefix, openId)
	var resp LoginInfo
	err := m.QueryRowIndex(&resp, loginInfoOpenIdKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `open_id` = ? limit 1", loginInfoRows, m.table)
		if err := conn.QueryRow(&resp, query, openId); err != nil {
			return nil, err
		}
		return resp.Uid, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
func (m *defaultLoginInfoModel) Insert(data *LoginInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?)", m.table, loginInfoRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.OpenId)

	return ret, err
}

func (m *defaultLoginInfoModel) FindOne(uid int64) (*LoginInfo, error) {
	loginInfoUidKey := fmt.Sprintf("%s%v", cacheLoginInfoUidPrefix, uid)
	var resp LoginInfo
	err := m.QueryRow(&resp, loginInfoUidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", loginInfoRows, m.table)
		return conn.QueryRow(v, query, uid)
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

func (m *defaultLoginInfoModel) Update(data *LoginInfo) error {
	loginInfoUidKey := fmt.Sprintf("%s%v", cacheLoginInfoUidPrefix, data.Uid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `uid` = ?", m.table, loginInfoRowsWithPlaceHolder)
		return conn.Exec(query, data.OpenId, data.Uid)
	}, loginInfoUidKey)
	return err
}

func (m *defaultLoginInfoModel) Delete(uid int64) error {

	loginInfoUidKey := fmt.Sprintf("%s%v", cacheLoginInfoUidPrefix, uid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `uid` = ?", m.table)
		return conn.Exec(query, uid)
	}, loginInfoUidKey)
	return err
}

func (m *defaultLoginInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheLoginInfoUidPrefix, primary)
}

func (m *defaultLoginInfoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", loginInfoRows, m.table)
	return conn.QueryRow(v, query, primary)
}
