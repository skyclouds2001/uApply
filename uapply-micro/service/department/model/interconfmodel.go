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
	interConfFieldNames          = builder.RawFieldNames(&InterConf{})
	interConfRows                = strings.Join(interConfFieldNames, ",")
	interConfRowsExpectAutoSet   = strings.Join(stringx.Remove(interConfFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	interConfRowsWithPlaceHolder = strings.Join(stringx.Remove(interConfFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheInterConfIdPrefix        = "cache:interConf:id:"
	cacheInterConfDepIdTurnPrefix = "cache:interConf:depId:turn:"
)

type (
	InterConfModel interface {
		Insert(data *InterConf) (sql.Result, error)
		FindOne(id int64) (*InterConf, error)
		FindOneByDepIdTurn(depId int64, turn int64) (*InterConf, error)
		Update(data *InterConf) error
		Delete(id int64) error
	}

	defaultInterConfModel struct {
		sqlc.CachedConn
		table string
	}

	InterConf struct {
		Id           int64  `db:"id"`
		DepId        int64  `db:"dep_id"`        // 部门 id
		Turn         int64  `db:"turn"`          // 轮次
		Start        int64  `db:"start"`         // 开始面试时间戳
		End          int64  `db:"end"`           // 面试结束时间戳
		InterAddress string `db:"inter_address"` // 面试地点
		ContactPhone string `db:"contact_phone"` // 负责人电话
	}
)

func NewInterConfModel(conn sqlx.SqlConn, c cache.CacheConf) InterConfModel {
	return &defaultInterConfModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`inter_conf`",
	}
}

func (m *defaultInterConfModel) Insert(data *InterConf) (sql.Result, error) {
	interConfIdKey := fmt.Sprintf("%s%v", cacheInterConfIdPrefix, data.Id)
	interConfDepIdTurnKey := fmt.Sprintf("%s%v:%v", cacheInterConfDepIdTurnPrefix, data.DepId, data.Turn)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, interConfRowsExpectAutoSet)
		return conn.Exec(query, data.DepId, data.Turn, data.Start, data.End, data.InterAddress, data.ContactPhone)
	}, interConfIdKey, interConfDepIdTurnKey)
	return ret, err
}

func (m *defaultInterConfModel) FindOne(id int64) (*InterConf, error) {
	interConfIdKey := fmt.Sprintf("%s%v", cacheInterConfIdPrefix, id)
	var resp InterConf
	err := m.QueryRow(&resp, interConfIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", interConfRows, m.table)
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

func (m *defaultInterConfModel) FindOneByDepIdTurn(depId int64, turn int64) (*InterConf, error) {
	interConfDepIdTurnKey := fmt.Sprintf("%s%v:%v", cacheInterConfDepIdTurnPrefix, depId, turn)
	var resp InterConf
	err := m.QueryRowIndex(&resp, interConfDepIdTurnKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `dep_id` = ? and `turn` = ? limit 1", interConfRows, m.table)
		if err := conn.QueryRow(&resp, query, depId, turn); err != nil {
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

func (m *defaultInterConfModel) Update(data *InterConf) error {
	interConfIdKey := fmt.Sprintf("%s%v", cacheInterConfIdPrefix, data.Id)
	interConfDepIdTurnKey := fmt.Sprintf("%s%v:%v", cacheInterConfDepIdTurnPrefix, data.DepId, data.Turn)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, interConfRowsWithPlaceHolder)
		return conn.Exec(query, data.DepId, data.Turn, data.Start, data.End, data.InterAddress, data.ContactPhone, data.Id)
	}, interConfIdKey, interConfDepIdTurnKey)
	return err
}

func (m *defaultInterConfModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	interConfIdKey := fmt.Sprintf("%s%v", cacheInterConfIdPrefix, id)
	interConfDepIdTurnKey := fmt.Sprintf("%s%v:%v", cacheInterConfDepIdTurnPrefix, data.DepId, data.Turn)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, interConfIdKey, interConfDepIdTurnKey)
	return err
}

func (m *defaultInterConfModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheInterConfIdPrefix, primary)
}

func (m *defaultInterConfModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", interConfRows, m.table)
	return conn.QueryRow(v, query, primary)
}
