package service

import (
	"context"
	"fmt"
	"time"

	"app/lottery/internal/model"
)

func (s *Svc) asyncSaveRecords(draw *model.DrawReply, phone int) {
	if !draw.IsWin {
		return
	}
	now := time.Now()
	drawDate := now.Format(model.LayoutByDay)
	// 增加用户获奖记录
	go func() { // TODO 通过全局goroutine pool限制 goroutine暴涨
		//s.dao.Logger.Printf("s.asyncSaveRecords draw(%v) phone(%d) excute time(%s) ", draw, phone, now.Format(model.Layout))

		// TODO 最好丢给worker批量录入数据库
		if _, err := s.dao.SaveRecords(context.Background(), draw.Prize, phone, drawDate); err != nil {
			s.dao.Logger.Printf("s.asyncSaveRecords draw(%v) phone(%d) excute time(%s) err(%v)", draw, phone, now.Format(model.Layout), err)
		}
		return
	}()
}


func (s *Svc) ExportDrawRecords(c context.Context) (data *model.PrizeRecordsReply){
	list, err := s.dao.FetchRecords(c)
	if err != nil {
		return
	}

	data = &model.PrizeRecordsReply{
		Title:   fmt.Sprintf("中奖记录-%s-导出", time.Now().Format(model.LayoutByDay)),
		Content: make([][]string, 0, len(list)+1),
	}
	data.Content = append(data.Content, (&model.PrizeRecord{}).ToCsvHeader())
	for _, v := range list {
		data.Content = append(data.Content, v.ToCsvRow())
	}
	return
}