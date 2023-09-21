package db

import (
	"errors"
	"fmt"
	"net/http"
	"tyranno/backend/domain/model"
	"tyranno/backend/utils/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (conn *gorm.DB, err error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		config.DbHost,
		config.DbUser,
		config.DbPass,
		config.DbName,
		config.DbPort,
	)
	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("%d : database error", http.StatusInternalServerError))
	}

	if err = conn.AutoMigrate(&model.User{}, &model.Post{}); err != nil {
		return nil, errors.New(fmt.Sprintf("%d : faild migrate", http.StatusInternalServerError))
	}

	return conn, nil
}
