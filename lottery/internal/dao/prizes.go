package dao

import (
	"context"
	"database/sql"

	"app/lottery/internal/model"
)

const (
	_findPrizeSql = `SELECT id, prize, total, stock, version FROM lottery_prizes `
	_updatePrizeSql = `UPDATE lottery_prizes SET stock = stock - 1, version = version + 1 WHERE stock > 0 AND prize = ?`
)


func (d *Dao) FetchPrizes(c context.Context) (list []*model.Prize, err error) {
	rows, err := d.db.QueryContext(c, _findPrizeSql)
	if err != nil {
		d.Logger.Printf("d.FetchPrizes d.db.QueryContext error(%v)", err)
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	for rows.Next() {
		t := new(model.Prize)
		err = rows.Scan(
			&t.Id, &t.PrizeId, &t.Total, &t.Stock, &t.Version,
		)
		if err != nil {
			d.Logger.Printf("d.FetchPrizes rows.Scan error(%v)", err)
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

func (d *Dao) FindOnePrize(c context.Context, conditions string) (prize *model.Prize, err error) {
	prize = new(model.Prize)
	err = d.db.QueryRowContext(c, _findPrizeSql + conditions).Scan(&prize.Id, &prize.PrizeId, &prize.Total, &prize.Stock, &prize.Version)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return
	}
	return
}

func (d *Dao) UpdatePrize(c context.Context, prize model.PrizeKey) (rows int64, err error) {
	res, err := d.db.Exec(_updatePrizeSql, prize)
	if err != nil {
		d.Logger.Printf("d.UpdatePrize d.db.Exec error(%v)", err)
		return
	}
	rows, err = res.RowsAffected()
	return
}

