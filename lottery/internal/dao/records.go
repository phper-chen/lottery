package dao

import (
	"context"

	"app/lottery/internal/model"
)

const (
	_findRecordsSql = `SELECT id, phone, prize_id, draw_date, ctime FROM lottery_records `
	_saveRecordsSql = `INSERT INTO lottery_records (prize_id, phone, draw_date) VALUES (?,?,?)`
)



func (d *Dao) FetchRecords(c context.Context) (list []*model.PrizeRecord, err error) {
	rows, err := d.db.QueryContext(c, _findRecordsSql)
	if err != nil {
		d.Logger.Printf("d.FetchRecords d.db.QueryContext error(%v)", err)
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		t := new(model.PrizeRecord)
		err = rows.Scan(
			&t.Id, &t.Phone, &t.PrizeId, &t.DrawDate, &t.Ctime,
		)
		if err != nil {
			d.Logger.Printf("d.FetchRecords rows.Scan error(%v)", err)
			return
		}
		if _, ok := model.PrizeNames[t.PrizeId]; ok { // 不分配变量 只寻址
			t.Name = model.PrizeNames[t.PrizeId]
		}
		list = append(list, t)
	}
	err = rows.Err()
	return
}

func (d *Dao) SaveRecords(c context.Context, prizeId model.PrizeKey, phone int, drawDate string) (insertId int64, err error) {
	res, err := d.db.Exec(_saveRecordsSql, prizeId, phone, drawDate)
	if err != nil {
		d.Logger.Printf("d.SaveRecords d.db.Exec error(%v)", err)
		return
	}
	insertId, err = res.LastInsertId()
	return
}