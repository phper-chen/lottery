package dao

import (
	"app/lottery/internal/model"
	"context"
	"database/sql"
	"fmt"
)

const (
	_findUserSql = `SELECT id, phone, draw_right, article, ctime FROM lottery_users %s`
	_saveUserSql  = `INSERT INTO lottery_users (phone, draw_right, article) VALUES (?,?,?)`
)


func (d *Dao) FetchUsers(c context.Context) (list []*model.UserInvolvesInfo, err error) {
	rows, err := d.db.QueryContext(c, fmt.Sprintf(_findUserSql, ""))
	if err != nil {
		d.Logger.Printf("d.FetchUsers d.db.QueryContext error(%v)", err)
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		t := new(model.UserInvolvesInfo)
		err = rows.Scan(
			&t.Id, &t.Phone, &t.DrawRight, &t.Article, &t.Ctime,
		)
		if err != nil {
			d.Logger.Printf("d.FetchUsers rows.Scan error(%v)", err)
			return
		}
		list = append(list, t)
	}
	err = rows.Err()
	return
}

func (d *Dao) FindOneUser(c context.Context, conditions string) (info *model.UserInvolvesInfo, err error) {
	info = new(model.UserInvolvesInfo)
	err = d.db.QueryRowContext(c, _findPrizeSql + conditions).Scan(&info.Id, &info.Phone, &info.DrawRight, &info.Article, &info.Ctime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return
	}
	return
}


func (d *Dao) SaveUsers(c context.Context, phone int,  article string) (insertId int64, err error) {
	res, err := d.db.Exec(_saveUserSql, phone, model.HaveDrawRight, article)
	if err != nil {
		d.Logger.Printf("d.SaveUsers d.db.Exec error(%v)", err)
		return
	}
	insertId, err = res.LastInsertId()
	return
}