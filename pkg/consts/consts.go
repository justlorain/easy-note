package consts

const (
	NoteTableName   = "note"
	UserTableName   = "user"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Notes           = "notes"
	ApiServiceName  = "demoapi"
	NoteServiceName = "demonote"
	UserServiceName = "demouser"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:9910)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	LocalHost       = "127.0.0.1"
	NacosPort       = 8848
	UserServiceAddr = "127.0.0.1:9000"
	NoteServiceAddr = "127.0.0.1:10000"
	DefaultLimit    = 10
)
