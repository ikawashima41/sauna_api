package seed

import (
	"log"

	"saunaApi/api/models"

	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Name:     "admin",
		Email:    "admin@gmail.com",
		Password: "password",
	},
	models.User{
		Name:     "test",
		Email:    "test@gmail.com",
		Password: "password",
	},
}

var posts = []models.Sauna{
	models.Sauna{
		Name:        "アダム・イブ",
		Description: "芸能人御用達サウナ",
		City:        "南麻布",
	},
	models.Sauna{
		Name:        "アスティル",
		Description: "新橋サラリーマンの憩いの場",
		City:        "新橋",
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.Sauna{}, &models.User{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Sauna{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Sauna{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
		posts[i].UserID = users[i].ID

		err = db.Debug().Model(&models.Sauna{}).Create(&posts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed posts table: %v", err)
		}
	}
}
