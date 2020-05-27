package model

import (
	"strconv"
)

type DrawReply struct {
	IsWin bool `json:"is_win"`
	Prize PrizeKey `json:"prize"`
	Msg string `json:"msg"`
}

type PrizeListReq struct {

}

type PrizeListReply struct {
	List []*Prize `json:"list"`
}



type Prize struct {
	Id int `json:"id"`
	PrizeId PrizeKey `json:"prize_id"`
	Name string `json:"name"`
	Total int `json:"-"`
	Stock int `json:"-"`
	Version int `json:"-"`
}


// 奖品中奖概率
type PrizeRate struct {
	Start int // 中奖概率起始编码（包含）
	End   int // 中奖概率终止编码（包含）
}

// 奖品列表
var PrizeList = []PrizeKey{
	Paster,
	PhoneCard,
	Phone,
}

var PrizeNames = map[PrizeKey]string{ // 可以并发读
	Paster: "贴纸一张",
	PhoneCard: "电话卡一张",
	Phone: "手机一部",
	NoPrize: "谢谢参与", // 有限奖品耗尽时
}

const BaseNum = 10000
// 奖品的中奖概率设置，与上面的 prizeList 对应的设置
// Rate = (End - Start) / BaseNum
var RateList = []PrizeRate{
	{0, 9399}, // 94%
	{9400, 9899}, // 5%
	{9900, 9999}, // 1%
}


type PrizeKey int8
const (
	Paster PrizeKey = 0
	PhoneCard PrizeKey = 1
	Phone PrizeKey = 2
	NoPrize PrizeKey = 3
)

func(pk PrizeKey) IsUnlimited() bool {
	return pk == Paster
}

func(pk PrizeKey) Where() string {
	return " WHERE prize = " + strconv.Itoa(int(pk))
}

func(pk PrizeKey) GetLockName() string {
	return PrizeNames[pk]
}

const (
	// 锁未正常释放时 100毫秒
	HandleTimeOut = 100
	StockDeductSuccess = 1
)

