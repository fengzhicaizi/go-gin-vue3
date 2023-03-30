package initialize

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	. "github.com/fengzhicaizi/gin-vue3/global"
)

func Grom() (db *gorm.DB) {
	var err error

	db, err = gorm.Open(GVA_VP.GetString("database.Type"), fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		GVA_VP.GetString("database.User"),
		GVA_VP.GetString("database.Password"),
		GVA_VP.GetString("database.Host"),
		GVA_VP.GetString("database.Name")))

	if err != nil {
		log.Fatalf("model.Setup err : %v", err)
	}

	log.Printf("--%v准备就绪,地址:%v,数据库:%v -- \n", GVA_VP.GetString("database.Type"), GVA_VP.GetString("database.Host"), GVA_VP.GetString("database.Name"))

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return viper.GetString("database.TablePrefix") + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetConnMaxIdleTime(10)
	db.DB().SetMaxOpenConns(100)
	db.LogMode(true)

	return
}
