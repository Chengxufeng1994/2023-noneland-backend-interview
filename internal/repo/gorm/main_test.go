package gorm

import (
	"noneland/backend/interview/configs"
	"noneland/backend/interview/internal/db"
	"noneland/backend/interview/internal/entity"
	"os"
	"path"
	"runtime"
	"testing"
)

var repo entity.Repository

func TestMain(m *testing.M) {
	_, filename, _, _ := runtime.Caller(0)
	dir := path.Join(path.Dir(filename), "..", "..", "..")
	err := os.Chdir(dir)
	if err != nil {
		panic(err)
	}

	cfg := configs.NewConfig()
	db := db.NewDb()
	repo = NewRepository(db, cfg)

	os.Exit(m.Run())
}
