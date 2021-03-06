package models

import (
	"fmt"
	"miaoqi-bot/utils/logger"
	"strings"
)

type Character struct {
	Id           int    `xorm:"id"`
	No           string `xorm:"no"`
	Name         string `xorm:"name"`
	Rarity       int    `xorm:"rarity"`
	Planet       int    `xorm:"from_planet_no"`
	Gender       int    `xorm:"gender"`
	OriginDesc   string `xorm:"origin_desc"`
	SkillName    string `xorm:"skill_name"`
	SkillDesc    string `xorm:"skill_desc"`
	Life         int    `xorm:"LIF"`
	Attack       int    `xorm:"ATK"`
	Defense      int    `xorm:"DEF"`
	Critical     int    `xorm:"CRI"`
	Learn        int    `xorm:"LER"`
	Charm        int    `xorm:"CHR"`
	Strength     int    `xorm:"'STR'"`
	Intelligence int    `xorm:"'INT'"`
}

func (Character) TableName() string {
	return "characters"
}

func (c Character) Format() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("[%s] %s", c.No, c.Name))
	s.WriteString(fmt.Sprintf("\n性别:%s    星球:%s", c.GetGender(), c.GetPlanet()))
	s.WriteString(fmt.Sprintf("\n来源:%s", c.OriginDesc))
	s.WriteString(fmt.Sprintf("\n技能:%s", c.SkillDesc))
	s.WriteString(fmt.Sprintf("\n生命:%d     攻击:%d", c.Life, c.Attack))
	s.WriteString(fmt.Sprintf("\n忍耐:%d     暴击:%d", c.Defense, c.Critical))
	s.WriteString(fmt.Sprintf("\n技巧:%d     亲和:%d", c.Learn, c.Charm))
	s.WriteString(fmt.Sprintf("\n力气:%d     智力:%d", c.Strength, c.Intelligence))
	s.WriteString("\n" + c.GetSkins())
	s.WriteString("\n" + c.GetEquipments())
	return s.String()
}

func (c Character) GetGender() string {
	switch c.Gender {
	case 1:
		return "男"
	case 2:
		return "女"
	case 3:
		return "双性"
	default:
		return "未知"
	}
}

func (c Character) GetPlanet() string {
	switch c.Planet {
	case 0:
		return "星际流民"
	case 1:
		return "冒险星"
	case 2:
		return "美食星"
	case 3:
		return "博物馆星"
	case 4:
		return "悬疑星"
	default:
		return "未知"
	}
}

func (c Character) GetSkins() string {
	s := strings.Builder{}
	s.WriteString("[皮肤]")
	skins := make([]CharacterSkin, 0)
	err := DB.Where("character_no = ?", c.No).Asc("pos").Find(&skins)
	if err != nil {
		logger.Error("查询数据库出错!", err)
		return s.String()
	}
	for _, skin := range skins {
		s.WriteString("\n")
		s.WriteString(skin.Format())
	}
	return s.String()
}

func (c Character) GetEquipments() string {
	s := strings.Builder{}
	s.WriteString("[装备]")
	equipments := make([]CharacterEquipment, 0)
	err := DB.Where("character_no = ?", c.No).Asc("pos").Find(&equipments)
	if err != nil {
		logger.Error("查询数据库出错!", err)
		return s.String()
	}
	for _, equipment := range equipments {
		s.WriteString("\n")
		s.WriteString(equipment.Format())
	}
	return s.String()
}
