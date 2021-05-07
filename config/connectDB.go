package config

func ConnectDB() *.DB {
	dsn := "root:root@tcp(localhost)/silih_a3?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}
}