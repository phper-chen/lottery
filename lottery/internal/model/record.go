package model

import (
	"strconv"
	"time"
)

const (
	LayoutByDay = "2006-01-02"
	Layout = "2006-01-02 15:04:05"
)

type PrizeRecordsReply struct{
	Title   string
	Content [][]string
}

type PrizeRecord struct {
	Id int `json:"id"`
	Phone int `json:"phone"`
	PrizeId PrizeKey `json:"-"`
	Name string `json:"name"`
	DrawDate string `json:"-"`
	Ctime time.Time `json:"ctime"`
}

func (p *PrizeRecord) ToCsvHeader() []string {
	return []string{
		"ID",  // 0
		"用户(手机号)",  // 1
		"奖品", // 2
		"时间",    // 4
	}
}

func (p *PrizeRecord) ToCsvRow() []string {
	return []string{
		strconv.Itoa(p.Id), // 0
		strconv.Itoa(p.Phone),                   // 1
		p.Name,                     // 2
		p.Ctime.Format(Layout),                            // 4
	}
}