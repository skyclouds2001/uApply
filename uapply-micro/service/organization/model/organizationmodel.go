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
	organizationFieldNames          = builder.RawFieldNames(&Organization{})
	organizationRows                = strings.Join(organizationFieldNames, ",")
	organizationRowsExpectAutoSet   = strings.Join(stringx.Remove(organizationFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	organizationRowsWithPlaceHolder = strings.Join(stringx.Remove(organizationFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheOrganizationIdPrefix      = "cache:organization:id:"
	cacheOrganizationAccountPrefix = "cache:organization:account:"
	cacheOrganizationNamePrefix    = "cache:organization:name:"
)

type (
	OrganizationModel interface {
		Insert(data *Organization) (sql.Result, error)
		FindOne(id int64) (*Organization, error)
		FindOneByAccount(account string) (*Organization, error)
		FindOneByName(name string) (*Organization, error)
		Update(data *Organization) error
		Delete(id int64) error
	}

	defaultOrganizationModel struct {
		sqlc.CachedConn
		table string
	}

	Organization struct {
		Id       int64  `db:"id"`
		Name     string `db:"name"`
		Account  string `db:"account"`
		Password string `db:"password"`
	}
)

func NewOrganizationModel(conn sqlx.SqlConn, c cache.CacheConf) OrganizationModel {
	return &defaultOrganizationModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`organization`",
	}
}

func (m *defaultOrganizationModel) Insert(data *Organization) (sql.Result, error) {
	organizationIdKey := fmt.Sprintf("%s%v", cacheOrganizationIdPrefix, data.Id)
	organizationAccountKey := fmt.Sprintf("%s%v", cacheOrganizationAccountPrefix, data.Account)
	organizationNameKey := fmt.Sprintf("%s%v", cacheOrganizationNamePrefix, data.Name)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, organizationRowsExpectAutoSet)
		return conn.Exec(query, data.Name, data.Account, data.Password)
	}, organizationIdKey, organizationAccountKey, organizationNameKey)
	return ret, err
}

func (m *defaultOrganizationModel) FindOne(id int64) (*Organization, error) {
	organizationIdKey := fmt.Sprintf("%s%v", cacheOrganizationIdPrefix, id)
	var resp Organization
	err := m.QueryRow(&resp, organizationIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", organizationRows, m.table)
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

func (m *defaultOrganizationModel) FindOneByAccount(account string) (*Organization, error) {
	organizationAccountKey := fmt.Sprintf("%s%v", cacheOrganizationAccountPrefix, account)
	var resp Organization
	err := m.QueryRowIndex(&resp, organizationAccountKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `account` = ? limit 1", organizationRows, m.table)
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

func (m *defaultOrganizationModel) FindOneByName(name string) (*Organization, error) {
	organizationNameKey := fmt.Sprintf("%s%v", cacheOrganizationNamePrefix, name)
	var resp Organization
	err := m.QueryRowIndex(&resp, organizationNameKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `name` = ? limit 1", organizationRows, m.table)
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

func (m *defaultOrganizationModel) Update(data *Organization) error {
	organizationIdKey := fmt.Sprintf("%s%v", cacheOrganizationIdPrefix, data.Id)
	organizationAccountKey := fmt.Sprintf("%s%v", cacheOrganizationAccountPrefix, data.Account)
	organizationNameKey := fmt.Sprintf("%s%v", cacheOrganizationNamePrefix, data.Name)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, organizationRowsWithPlaceHolder)
		return conn.Exec(query, data.Name, data.Account, data.Password, data.Id)
	}, organizationIdKey, organizationAccountKey, organizationNameKey)
	return err
}

func (m *defaultOrganizationModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	organizationAccountKey := fmt.Sprintf("%s%v", cacheOrganizationAccountPrefix, data.Account)
	organizationNameKey := fmt.Sprintf("%s%v", cacheOrganizationNamePrefix, data.Name)
	organizationIdKey := fmt.Sprintf("%s%v", cacheOrganizationIdPrefix, id)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, organizationAccountKey, organizationNameKey, organizationIdKey)
	return err
}

func (m *defaultOrganizationModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheOrganizationIdPrefix, primary)
}

func (m *defaultOrganizationModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", organizationRows, m.table)
	return conn.QueryRow(v, query, primary)
}
