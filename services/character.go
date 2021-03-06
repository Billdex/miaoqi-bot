package services

import (
	"fmt"
	"github.com/ProtobufBot/go-pbbot"
	"miaoqi-bot/models"
	"miaoqi-bot/scheduler"
	"miaoqi-bot/utils"
	"miaoqi-bot/utils/logger"
	"strconv"
)

func Character(c *scheduler.Context) {
	param := c.PretreatedMessage

	characters := make([]models.Character, 0)
	numId, err := strconv.Atoi(param)
	if err != nil {
		err = models.DB.Where("name like ?", "%"+param+"%").Find(&characters)
		if err != nil {
			logger.Error("查询数据库出错!", err)
			msg := pbbot.NewMsg().Text(utils.SystemErrorNote)
			c.Reply(msg, false)
			return
		}
	} else {
		galleryId := fmt.Sprintf("%03d", numId)
		err = models.DB.Where("no = ?", galleryId).Find(&characters)
		if err != nil {
			logger.Error("查询数据库出错!", err)
			msg := pbbot.NewMsg().Text(utils.SystemErrorNote)
			c.Reply(msg, false)
			return
		}
	}

	var s string
	if len(characters) == 0 {
		s = "没有查询到角色信息"
	} else if len(characters) == 1 {
		s = characters[0].Format()
	} else {
		s = "查询到以下角色:"
		for _, character := range characters {
			s += fmt.Sprintf("\n[%s] %s", character.No, character.Name)
		}
	}
	msg := pbbot.NewMsg().Text(s)
	_, err = c.Reply(msg, false)
	if err != nil {
		logger.Error("发送消息失败！", err)
	}

}
