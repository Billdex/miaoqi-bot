package services

import "miaoqi-bot/scheduler"

func SetUp() *scheduler.Scheduler {
	s := scheduler.New()

	filter := s.Group("*")
	{
		filter.Bind("任务", Mission).Alias("主线")
		filter.Bind("角色", Character).Alias("人物")
	}

	return s
}
