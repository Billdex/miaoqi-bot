package models

import (
	"fmt"
)

type CharacterEquipment struct {
	Id          int    `xorm:"id"`
	CharacterNo string `xorm:"character_no"`
	Pos         int    `xorm:"pos"`
	Name        string `xorm:"name"`
	SkillDesc   string `xorm:"skill_desc"`
}

func (CharacterEquipment) TableName() string {
	return "character_equipment"
}

func (ce CharacterEquipment) Format() string {
	return fmt.Sprintf("[%d]%s", ce.Pos, ce.SkillDesc)
}
