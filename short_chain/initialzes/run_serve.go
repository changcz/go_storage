package initialzes

func RunServe() {
	InitConfig()
	InitMysql()
	InitTable()
	InitRouter()
}
