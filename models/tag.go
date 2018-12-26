package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Tag struct {
	Model

	Name 		string 		`json:"name"`
	CreatedBy 	string		`json:"created_by"`
	ModifiedBy	string		`json:"modified_by"`
	State		int			`json:"state"`
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("created_on", time.Now())
	scope.SetColumn("modified_on", time.Now())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now())
	return nil
}

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)
	return
}

func GetTagsTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)
	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func AddTag(name string, state int, createBy string) bool {
	tag := &Tag {
		Name: name,
		State: state,
		CreatedBy: createBy,
	}
	db.Create(tag)

	return true
}