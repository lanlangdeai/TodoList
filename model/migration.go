package model

//执行数据迁移
func migration() {
	//自动迁移模式
	DB.Set("gorm:table_options", "charset=utf8").
		AutoMigrate(&User{}).
		AutoMigrate(&Task{})
	DB.Model(&Task{}).AddForeignKey("uid","User(id)","CASCADE","CASCADE")
}