package model

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	mysqlx "github.com/jmoiron/sqlx"
	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"strings"
)

var (
	userInfoFieldNames          = builder.RawFieldNames(&UserInfo{})
	userInfoRows                = strings.Join(userInfoFieldNames, ",")
	userInfoRowsExpectAutoSet   = strings.Join(stringx.Remove(userInfoFieldNames, "`create_time`", "`update_time`"), ",")
	userInfoRowsWithPlaceHolder = strings.Join(stringx.Remove(userInfoFieldNames, "`uid`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserInfoUidPrefix = "cache:userInfo:uid:"
)

type (
	UserInfoModel interface {
		Insert(data *UserInfo) (sql.Result, error)
		FindOne(uid int64) (*UserInfo, error)
		Update(data *UserInfo) error
		Delete(uid int64) error
		FindSexInfo(uids []int, mysql *mysqlx.DB) ([]string, error)
		FindAll(uid []int, mysql *mysqlx.DB) ([]*UserInfo, error)
	}

	defaultUserInfoModel struct {
		sqlc.CachedConn
		table string
	}

	UserInfo struct {
		Uid        int64          `db:"uid"`
		Name       string         `db:"name"`
		Sex        string         `db:"sex"`
		StudentNum string         `db:"student_num"`
		AddressNum sql.NullString `db:"address_num"`
		Major      sql.NullString `db:"major"`
		PhoneNum   string         `db:"phone_num"`
		Email      string         `db:"email"`
		Intro      string         `db:"intro"`
	}
)

func NewUserInfoModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserInfoModel {
	return &defaultUserInfoModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user_info`",
	}
}

func (m *defaultUserInfoModel) Insert(data *UserInfo) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userInfoRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Uid, data.Name, data.Sex, data.StudentNum, data.AddressNum, data.Major, data.PhoneNum, data.Email, data.Intro)
	return ret, err
}

func (m *defaultUserInfoModel) FindSexInfo(uids []int, mysql *mysqlx.DB) ([]string, error) {
	var rsp []string
	query := fmt.Sprintf("select %s from %s where `uid` in (?)", `sex`, m.table)

	query, args, err := mysqlx.In(query, uids)
	if err != nil {
		return nil, err
	}
	query = mysql.Rebind(query)
	err = mysql.Select(&rsp, query, args...)
	switch err {
	case nil:
		return rsp, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) FindAll(uid []int, mysql *mysqlx.DB) ([]*UserInfo, error) {
	var rsp []*UserInfo
	query := fmt.Sprintf("select %s from %s WHERE uid in (?) order by `uid`", userInfoRows, m.table)
	query, args, err := mysqlx.In(query, uid)
	if err != nil {
		return nil, err
	}
	query = mysql.Rebind(query)
	err = mysql.Select(&rsp, query, args...)
	switch err {
	case nil:
		return rsp, nil
	case sql.ErrNoRows:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserInfoModel) FindOne(uid int64) (*UserInfo, error) {
	userInfoUidKey := fmt.Sprintf("%s%v", cacheUserInfoUidPrefix, uid)
	var resp UserInfo
	err := m.QueryRow(&resp, userInfoUidKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userInfoRows, m.table)
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

func (m *defaultUserInfoModel) Update(data *UserInfo) error {
	userInfoUidKey := fmt.Sprintf("%s%v", cacheUserInfoUidPrefix, data.Uid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `uid` = ?", m.table, userInfoRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Sex, data.StudentNum, data.AddressNum, data.Major, data.PhoneNum, data.Email, data.Intro, data.Uid)
	}, userInfoUidKey)
	return err
}

func (m *defaultUserInfoModel) Delete(uid int64) error {

	userInfoUidKey := fmt.Sprintf("%s%v", cacheUserInfoUidPrefix, uid)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `uid` = ?", m.table)
		return conn.Exec(query, uid)
	}, userInfoUidKey)
	return err
}

func (m *defaultUserInfoModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserInfoUidPrefix, primary)
}

func (m *defaultUserInfoModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `uid` = ? limit 1", userInfoRows, m.table)
	return conn.QueryRow(v, query, primary)
}
