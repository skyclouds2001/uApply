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
	departmentFieldNames          = builder.RawFieldNames(&Department{})
	departmentRows                = strings.Join(departmentFieldNames, ",")
	departmentRowsExpectAutoSet   = strings.Join(stringx.Remove(departmentFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	departmentRowsWithPlaceHolder = strings.Join(stringx.Remove(departmentFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheDepartmentIdPrefix      = "cache:department:id:"
	cacheDepartmentAccountPrefix = "cache:department:account:"
	cacheDepartmentNamePrefix    = "cache:department:name:"
)

type (
	DepartmentModel interface {
		Insert(data *Department) (sql.Result, error)
		FindOne(id int64) (*Department, error)
		FindOneByAccount(account string) (*Department, error)
		FindOneByName(name string) (*Department, error)
		Update(data *Department) error
		Delete(data *Department, id int64) error
		FindAll(orgid int64) ([]*Department, error)
	}

	defaultDepartmentModel struct {
		sqlc.CachedConn
		table string
	}

	Department struct {
		Id       int64  `db:"id"`
		Name     string `db:"name"`
		Account  string `db:"account"`
		Password string `db:"password"`
		OrgId    int64  `db:"org_id"` // 对应的组织id
		OrgName  string `db:"org_name"`
	}
)

func NewDepartmentModel(conn sqlx.SqlConn, c cache.CacheConf) DepartmentModel {
	return &defaultDepartmentModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`department`",
	}
}

func (m *defaultDepartmentModel) FindAll(orgid int64) ([]*Department, error) {
	var resp = make([]*Department, 0)
	query := fmt.Sprintf("select %s from %s where `org_id` = ?", departmentRows, m.table)
	err := m.QueryRowsNoCache(&resp, query, orgid)
	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultDepartmentModel) Insert(data *Department) (sql.Result, error) {
	departmentIdKey := fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, data.Id)
	departmentAccountKey := fmt.Sprintf("%s%v", cacheDepartmentAccountPrefix, data.Account)
	departmentNameKey := fmt.Sprintf("%s%v", cacheDepartmentNamePrefix, data.Name)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, departmentRowsExpectAutoSet)
		return conn.Exec(query, data.Name, data.Account, data.Password, data.OrgId, data.OrgName)
	}, departmentIdKey, departmentAccountKey, departmentNameKey)
	return ret, err
}

func (m *defaultDepartmentModel) FindOne(id int64) (*Department, error) {
	departmentIdKey := fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, id)
	var resp Department
	err := m.QueryRow(&resp, departmentIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", departmentRows, m.table)
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

func (m *defaultDepartmentModel) FindOneByAccount(account string) (*Department, error) {
	departmentAccountKey := fmt.Sprintf("%s%v", cacheDepartmentAccountPrefix, account)
	var resp Department
	err := m.QueryRowIndex(&resp, departmentAccountKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", departmentRows, m.table)
		if err := conn.QueryRow(&resp, query, account); err != nil {
			return nil, err
		}
		return resp.Id, nil
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

func (m *defaultDepartmentModel) FindOneByName(name string) (*Department, error) {
	departmentNameKey := fmt.Sprintf("%s%v", cacheDepartmentNamePrefix, name)
	var resp Department
	err := m.QueryRowIndex(&resp, departmentNameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", departmentRows, m.table)
		if err := conn.QueryRow(&resp, query, name); err != nil {
			return nil, err
		}
		return resp.Id, nil
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

func (m *defaultDepartmentModel) Update(data *Department) error {
	departmentIdKey := fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, data.Id)
	departmentAccountKey := fmt.Sprintf("%s%v", cacheDepartmentAccountPrefix, data.Account)
	departmentNameKey := fmt.Sprintf("%s%v", cacheDepartmentNamePrefix, data.Name)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, departmentRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Account, data.Password, data.OrgId, data.OrgName, data.Id)
	}, departmentIdKey, departmentAccountKey, departmentNameKey)
	return err
}

func (m *defaultDepartmentModel) Delete(data *Department, id int64) error {
	departmentAccountKey := fmt.Sprintf("%s%v", cacheDepartmentAccountPrefix, data.Account)
	departmentNameKey := fmt.Sprintf("%s%v", cacheDepartmentNamePrefix, data.Name)
	departmentIdKey := fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, departmentIdKey, departmentAccountKey, departmentNameKey)
	return err
}

func (m *defaultDepartmentModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheDepartmentIdPrefix, primary)
}

func (m *defaultDepartmentModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", departmentRows, m.table)
	return conn.QueryRow(v, query, primary)
}
