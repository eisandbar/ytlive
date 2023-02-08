package store

import (
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

var (
	HOST        = os.Getenv("POSTGRES_HOST")
	USER        = os.Getenv("POSTGRES_USER")
	DB_NAME     = os.Getenv("POSTGRES_DB")
	DB_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
)

var dsn = fmt.Sprintf("host=%s port=5432 user=%s dbname=%s password=%s sslmode=disable",
	HOST,
	USER,
	DB_NAME,
	DB_PASSWORD,
)

func NewPGStore() PGStore {
	ps := PGStore{}
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to postgres: %v", err)
	}

	if err = db.AutoMigrate(&Stream{}); err != nil {
		log.Fatalf("Error migrating DB")
	}

	if err = db.AutoMigrate(&Category{}); err != nil {
		log.Fatalf("Error migrating DB")
	}

	fmt.Println("Postgres DB initialized")

	ps.db = db
	return ps

}

type PGStore struct {
	db *gorm.DB
}

func (ps *PGStore) FindOne(Id string) Stream {
	stream := Stream{Id: Id}
	ps.db.First(&stream)
	return stream
}

func (ps *PGStore) Add(video Stream) {
	ps.db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&video)
}

func (ps *PGStore) Delete(Id string) {
	log.Println("Deleting stream ID:", Id)
	ps.db.Where("id = ?", Id).Delete(&Stream{})
}

func (ps *PGStore) Update(video Stream) {
	ps.db.Save(&video)
}

func (ps *PGStore) List(opts ...ListOption) []Stream {
	var ls listSettings
	for _, opt := range opts {
		opt.Apply(&ls)
	}

	var streams []Stream

	tx := ps.db.Model(&Stream{})
	if len(ls.GetFilters()) > 0 {
		tx.Where("category IN ?", ls.GetFilters())
	}
	if ls.live {
		tx.Where("live_broadcast_content = ?", "live")
	}
	if ls.gaming {
		tx.Where("category_id = ?", "20")
	}
	tx.Order("Concurrent_Viewers desc").Limit(ls.GetMaxResults()).Offset(ls.GetOffset()).Find(&streams)
	return streams
}

func (ps *PGStore) Len() int {
	var count int64
	ps.db.Model(&Stream{}).Count(&count)
	return int(count)
}

func (ps *PGStore) Categories(gaming bool) []Category {
	var categories []Category

	subquery := ps.db.Model(&Stream{})
	subquery.Select("category, sum(concurrent_viewers) as concurrent_viewers")
	subquery.Not(map[string]interface{}{"category": []string{NoCategory, Pending, ""}})
	subquery.Where("live_broadcast_content = ?", "live")
	if gaming {
		subquery.Where("category_id = ?", "20")
	}
	subquery.Group("category")

	tx := ps.db.Select("foo.category, foo.concurrent_viewers, categories.url").Table("(?) as foo", subquery)
	tx.Joins("left join categories on foo.category = categories.category")
	tx.Order("foo.concurrent_viewers desc").Find(&categories)

	return categories
}

func (ps *PGStore) SaveCategory(category Category) {
	ps.db.Save(&category)
}

const NoCategory = "No category"
const Pending = "Pending"
