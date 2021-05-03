package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// Entityを定義
type Sauna struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	City        string    `gorm:"size:255;not null;" json:"city"`
	User        User      `json:"user"`
	UserID      uint32    `gorm:"not null" json:"user_id"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (s *Sauna) Prepare() {
	s.ID = 0
	s.Name = html.EscapeString(strings.TrimSpace(s.Name))
	s.Description = html.EscapeString(strings.TrimSpace(s.Description))
	s.City = html.EscapeString(strings.TrimSpace(s.City))
	s.User = User{}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
}

func (s *Sauna) Validate() error {

	if s.Name == "" {
		return errors.New("Required Name")
	}
	if s.Description == "" {
		return errors.New("Required Description")
	}
	if s.City == "" {
		return errors.New("Required City")
	}
	if s.UserID < 1 {
		return errors.New("Required User")
	}
	return nil
}

func (s *Sauna) SaveSauna(db *gorm.DB) (*Sauna, error) {
	var err error
	err = db.Debug().Model(&Sauna{}).Create(&s).Error
	if err != nil {
		return &Sauna{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.UserID).Take(&s.User).Error
		if err != nil {
			return &Sauna{}, err
		}
	}
	return s, nil
}

func (s *Sauna) FindAllSaunas(db *gorm.DB) (*[]Sauna, error) {
	var err error
	saunas := []Sauna{}
	err = db.Debug().Model(&Sauna{}).Limit(100).Find(&saunas).Error
	if err != nil {
		return &[]Sauna{}, err
	}
	if len(saunas) > 0 {
		for i, _ := range saunas {
			err := db.Debug().Model(&User{}).Where("id = ?", saunas[i].UserID).Take(&saunas[i].User).Error
			if err != nil {
				return &[]Sauna{}, err
			}
		}
	}
	return &saunas, nil
}

func (s *Sauna) FindSaunaByID(db *gorm.DB, pid uint64) (*Sauna, error) {
	var err error
	err = db.Debug().Model(&Sauna{}).Where("id = ?", pid).Take(&s).Error
	if err != nil {
		return &Sauna{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.UserID).Take(&s.User).Error
		if err != nil {
			return &Sauna{}, err
		}
	}
	return s, nil
}

func (s *Sauna) UpdateSauna(db *gorm.DB) (*Sauna, error) {

	var err error

	err = db.Debug().Model(&Sauna{}).Where("id = ?", s.ID).Updates(Sauna{Name: s.Name, Description: s.Description, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Sauna{}, err
	}
	if s.ID != 0 {
		err = db.Debug().Model(&User{}).Where("id = ?", s.UserID).Take(&s.User).Error
		if err != nil {
			return &Sauna{}, err
		}
	}
	return s, nil
}

func (s *Sauna) DeleteSauna(db *gorm.DB, pid uint64, uid uint32) (int64, error) {

	db = db.Debug().Model(&Sauna{}).Where("id = ? and user_id = ?", pid, uid).Take(&Sauna{}).Delete(&Sauna{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Sauna not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
