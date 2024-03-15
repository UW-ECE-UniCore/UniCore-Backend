package constants

const (
	MySQLDSN      = "unicore:unicorepwd@tcp(host.docker.internal:18001)/unicore?charset=utf8&parseTime=True"
	RedisAddr     = "localhost:18002"
	RedisPassword = "unicorepwd"

	UserTableName = "user"

	Salt = "UniCoreHashSalt!"
)

const (
	TimeZone = "Canada/Eastern"
)

const (
	PostTableName = "post"

	TypePublic = 0
	TypeSchool = 1

	StatusCreate = 0
)

const (
	LikeTableName = "Like"
)

const (
	JWTSecretKey = "UniCoreJWTKey"
)

// email stuff
const (
	SMTPHost     = "smtp-mail.outlook.com"
	SMTPUsername = "unicore.ca@outlook.com"
	SMTPPassword = "kexqa7-pozpuh-zeszYn"
	WebName      = "UniCore"
	Port         = 587
	CodeLength   = 6
	CharSet      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)
