package models

import (
	"fmt"
)

type CharacterSkin struct {
	Id          int    `xorm:"id"`
	CharacterNo string `xorm:"character_no"`
	Pos         int    `xorm:"pos"`
	Name        string `xorm:"name"`
	SkillName   string `xorm:"skill_name"`
	SkillDesc   string `xorm:"skill_desc"`
}

func (CharacterSkin) TableName() string {
	return "character_skins"
}

func (cs CharacterSkin) Format() string {
	return fmt.Sprintf("[%d]%s", cs.Pos, cs.SkillDesc)
}
