package model

import (
	"database/sql"
	"fmt"
	mysqlx "github.com/jmoiron/sqlx"
	"strings"
	"uapply-micro/common/filed"

	"github.com/tal-tech/go-zero/core/stores/builder"
	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
)

var (
	registerFieldNames          = builder.RawFieldNames(&Register{})
	registerRows                = strings.Join(registerFieldNames, ",")
	registerRowsExpectAutoSet   = strings.Join(stringx.Remove(registerFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	registerRowsWithPlaceHolder = strings.Join(stringx.Remove(registerFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheRegisterIdPrefix       = "cache:register:id:"
	cacheRegisterUidOrgIdPrefix = "cache:register:uid:orgId:"
)

type (
	RegisterModel interface {
		Insert(data *Register) (sql.Result, error)
		FindOne(id int64) (*Register, error)
		FindOneByUidOrgId(uid int64, orgId int64) (*Register, error)
		Update(data *Register) error
		Delete(id int64) error
		DeleteByDepAndUids(depId int64, uids []int, mysql *mysqlx.DB) error
		FindAll(uid int64) ([]*Register, error)
		FindFinalStatus(t int, depId int64) ([]*Register, error)
		FindStatus(uid []int, num string, mysql *mysqlx.DB) ([]int, error)
		FindDepAll(depId int64) ([]*Register, error)
		FindUsersInterviewedFirstTurn(depId int64) ([]*Register, error)
		UpdateMany(uid []int, depid int, mysql *mysqlx.DB, num string, t int) error
		FindByUid(uid int64) (*Register, error)
		UpdateMark(uid, depid int64, mark int32, marktag, evatag, eva string) error
		CountByOrgAndUid(orgId int64, uid int64, db *mysqlx.DB) (int64, error)
	}

	defaultRegisterModel struct {
		sqlc.CachedConn
		table string
	}

	Register struct {
		Id           int64          `db:"id"`
		Uid          int64          `db:"uid"`
		OrgId        int64          `db:"org_id"`
		OrgName      string         `db:"org_name"`
		DepId        int64          `db:"dep_id"`
		DepName      string         `db:"dep_name"`
		FirstStatus  int64          `db:"first_status"`
		SecondStatus int64          `db:"second_status"`
		FinalStatus  int64          `db:"final_status"`
		FirstEva     sql.NullString `db:"first_eva"`
		FirstMark    int64          `db:"first_mark"`
		SecondEva    sql.NullString `db:"second_eva"`
		SecondMark   int64          `db:"second_mark"`
	}
)

func NewRegisterModel(conn sqlx.SqlConn, c cache.CacheConf) RegisterModel {
	return &defaultRegisterModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`register`",
	}
}

func (m *defaultRegisterModel) FindByUid(uid int64) (*Register, error) {
	var rsp Register
	query := fmt.Sprintf("SELECT %s FROM %s WHERE `uid`=?", registerRows, m.table)
	err := m.QueryRowNoCache(&rsp, query, uid)
	switch err {
	case nil:
		return &rsp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRegisterModel) FindUsersInterviewedFirstTurn(depId int64) ([]*Register, error) {
	var rsp []*Register
	query := fmt.Sprintf("select %s from %s where `dep_id` = ? and `first_status` = ? order by `uid`", registerRows, m.table)
	err := m.QueryRowsNoCache(&rsp, query, depId, filed.PASS)
	switch err {
	case nil:
		return rsp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRegisterModel) FindDepAll(depId int64) ([]*Register, error) {
	var rsp []*Register
	query := fmt.Sprintf("select %s from %s where `dep_id` = ? order by `uid`", registerRows, m.table)
	err := m.QueryRowsNoCache(&rsp, query, depId)
	switch err {
	case nil:
		return rsp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRegisterModel) FindStatus(uids []int, num string, mysql *mysqlx.DB) ([]int, error) {
	var rsp []int
	query := fmt.Sprintf("select %s from %s where `uid` in (?)", num, m.table)

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

func (m *defaultRegisterModel) FindFinalStatus(t int, depId int64) ([]*Register, error) {
	var rsp []*Register
	query := fmt.Sprintf("select %s from %s where `final_status` = ? and `dep_id` = ? order by `uid`", registerRows, m.table)
	err := m.QueryRowsNoCache(&rsp, query, t, depId)
	switch err {
	case nil:
		return rsp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRegisterModel) FindAll(uid int64) ([]*Register, error) {
	var rsp []*Register
	query := fmt.Sprintf("select %s from %s where `uid` = ?", registerRows, m.table)
	err := m.QueryRowsNoCache(&rsp, query, uid)
	switch err {
	case nil:
		return rsp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRegisterModel) Insert(data *Register) (sql.Result, error) {
	registerIdKey := fmt.Sprintf("%s%v", cacheRegisterIdPrefix, data.Id)
	registerUidOrgIdKey := fmt.Sprintf("%s%v:%v", cacheRegisterUidOrgIdPrefix, data.Uid, data.OrgId)
	ret, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, registerRowsExpectAutoSet)
		return conn.Exec(query, data.Uid, data.OrgId, data.OrgName, data.DepId, data.DepName, data.FirstStatus, data.SecondStatus, data.FinalStatus, data.FirstEva, data.FirstMark, data.SecondEva, data.SecondMark)
	}, registerUidOrgIdKey, registerIdKey)
	return ret, err
}

func (m *defaultRegisterModel) FindOne(id int64) (*Register, error) {
	registerIdKey := fmt.Sprintf("%s%v", cacheRegisterIdPrefix, id)
	var resp Register
	err := m.QueryRow(&resp, registerIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", registerRows, m.table)
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

func (m *defaultRegisterModel) FindOneByUidOrgId(uid int64, orgId int64) (*Register, error) {
	registerUidOrgIdKey := fmt.Sprintf("%s%v:%v", cacheRegisterUidOrgIdPrefix, uid, orgId)
	var resp Register
	err := m.QueryRowIndex(&resp, registerUidOrgIdKey, m.formatPrimary, func(conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where `uid` = ? and `org_id` = ? limit 1", registerRows, m.table)
		if err := conn.QueryRow(&resp, query, uid, orgId); err != nil {
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

func (m *defaultRegisterModel) Update(data *Register) error {
	registerIdKey := fmt.Sprintf("%s%v", cacheRegisterIdPrefix, data.Id)
	registerUidOrgIdKey := fmt.Sprintf("%s%v:%v", cacheRegisterUidOrgIdPrefix, data.Uid, data.OrgId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, registerRowsWithPlaceHolder)
		return conn.Exec(query, data.Uid, data.OrgId, data.OrgName, data.DepId, data.DepName, data.FirstStatus, data.SecondStatus, data.FinalStatus, data.FirstEva, data.FirstMark, data.SecondEva, data.SecondMark, data.Id)
	}, registerIdKey, registerUidOrgIdKey)
	return err
}

func (m *defaultRegisterModel) UpdateMany(uid []int, depid int, mysql *mysqlx.DB, num string, t int) error {
	query := fmt.Sprintf("UPDATE %s SET %s=%d WHERE uid in (?) and dep_id = %d", m.table, num, t, depid)

	query, args, err := mysqlx.In(query, uid)
	if err != nil {
		return err
	}
	query = mysql.Rebind(query)
	_, err = mysql.Exec(query, args...)
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return ErrNotFound
	default:
		return err
	}
}

func (m *defaultRegisterModel) UpdateMark(uid, depid int64, mark int32, marktag, evatag, eva string) error {
	query := fmt.Sprintf("UPDATE %s SET %s=%d,%s='%s' WHERE `uid` = ? AND `dep_id` = ?", m.table, marktag, mark, evatag, eva)

	_, err := m.ExecNoCache(query, uid, depid)
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return ErrNotFound
	default:
		return err
	}
}

func (m *defaultRegisterModel) Delete(id int64) error {
	data, err := m.FindOne(id)
	if err != nil {
		return err
	}

	registerIdKey := fmt.Sprintf("%s%v", cacheRegisterIdPrefix, id)
	registerUidOrgIdKey := fmt.Sprintf("%s%v:%v", cacheRegisterUidOrgIdPrefix, data.Uid, data.OrgId)
	_, err = m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, registerIdKey, registerUidOrgIdKey)
	return err
}

func (m *defaultRegisterModel) DeleteByDepAndUids(depId int64, uids []int, mysql *mysqlx.DB) error {
	query := fmt.Sprintf("delete from %s where dep_id = %d and uid in (?)", m.table, depId)

	query, args, err := mysqlx.In(query, uids)
	if err != nil {
		return err
	}
	query = mysql.Rebind(query)

	_, err = mysql.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (m *defaultRegisterModel) CountByOrgAndUid(orgId int64, uid int64, db *mysqlx.DB) (int64, error) {
	query := fmt.Sprintf("select count(*) from %s where `org_id` = ? and `uid` = ?", m.table)

	var cnt int64
	err := db.Get(&cnt, query, orgId, uid)

	switch err {
	case nil:
		return cnt, nil
	case sql.ErrNoRows:
		return 0, nil
	default:
		return 0, err
	}
}

func (m *defaultRegisterModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRegisterIdPrefix, primary)
}

func (m *defaultRegisterModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", registerRows, m.table)
	return conn.QueryRow(v, query, primary)
}
