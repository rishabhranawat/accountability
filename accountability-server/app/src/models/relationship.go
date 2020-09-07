package models

import "github.com/jinzhu/gorm"

type Relationship struct {
  gorm.Model

  RelationshipFrom  User `gorm:"foreignkey:RelationshipFromID"`
  RelationshipFromID uint
  RelationshipTo    User `gorm:"foreignkey:RelationshipToID"`
  RelationshipToID  uint
  Approved          bool
}

