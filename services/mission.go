package services

import (
	"github.com/ProtobufBot/go-pbbot"
	"miaoqi-bot/models"
	"miaoqi-bot/scheduler"
	"miaoqi-bot/utils"
	"miaoqi-bot/utils/logger"
	"strconv"
	"strings"
)

func Mission(c *scheduler.Context) {
	params := strings.Split(c.PretreatedMessage, " ")
	start, err := strconv.Atoi(params[0])
	if err != nil {
		msg := pbbot.NewMsg().Text("查询格式有误")
		c.Reply(msg, false)
		return
	} else {
		start = utils.LimitNumberSection(start, 1, 435)
	}
	num := 1
	if len(params) > 1 {
		n, err := strconv.Atoi(params[1])
		if err == nil {
			num = utils.LimitNumberSection(n, 1, 5)
		}
	}

	missionList := make([]int, num)
	for i := 0; i < num; i++ {
		missionList[i] = start + i
	}

	missions := make([]models.Mission, 0)
	err = models.DB.Where("type = ?", "main").In("no", missionList).Find(&missions)
	if err != nil {
		msg := pbbot.NewMsg().Text(utils.SystemErrorNote)
		c.Reply(msg, false)
		return
	}

	s := "[主线任务查询]"
	for _, mission := range missions {
		s += "\n"
		s += mission.Format()
	}

	msg := pbbot.NewMsg().Text(s)
	_, err = c.Reply(msg, false)
	if err != nil {
		logger.Error("发送消息失败！", err)
	}
}
