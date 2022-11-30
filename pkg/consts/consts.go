package consts

const (
	NoteTableName   = "note"
	UserTableName   = "user"
	SecretKey       = "secret key"
	IdentityKey     = "id"
	Total           = "total"
	Notes           = "notes"
	ApiClientName   = "demoapi-client"
	NoteClientName  = "demonote-client"
	NoteServiceName = "demonote"
	UserServiceName = "demouser"
	MySQLDefaultDSN = "gorm:gorm@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True&loc=Local"
	TCP             = "tcp"
	UserServiceAddr = "127.0.0.1:9000"
	NoteServiceAddr = "127.0.0.1:10000"
	ExportEndpoint  = "127.0.0.1:4317"
	DefaultLimit    = 10
)
