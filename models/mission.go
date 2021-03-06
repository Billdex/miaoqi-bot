package models

import (
	"fmt"
	"strings"
)

type Mission struct {
	Id              int    `xorm:"id"`
	No              string `xorm:"no"`
	Type            string `xorm:"type"`
	Title           string `xorm:"title"`
	RequirementDesc string `xorm:"requirement_desc"`
	RewardsDesc     string `xorm:"rewards_desc"`
	Unlock          string `xorm:"unlock"`
	Hint            string `xorm:"hint"`
}

func (Mission) TableName() string {
	return "missions"
}

func (m Mission) Format() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("[%s] %s\n", m.No, m.Title))
	s.WriteString(fmt.Sprintf("描述:%s", m.RequirementDesc))

	if m.Unlock != "" {
		s.WriteString(fmt.Sprintf("\n可解锁:%s", m.Unlock))
	}
	if m.Hint != "" {
		s.WriteString(fmt.Sprintf("\n提示:%s", m.Hint))
	}
	return s.String()
}
