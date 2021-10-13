module wp-server

go 1.16

// +heroku goVersion go1.16

require (
	github.com/jackc/pgx/v4 v4.13.0 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5 // indirect
	golang.org/x/text v0.3.7 // indirect
	gorm.io/driver/postgres v1.1.0
	gorm.io/driver/sqlite v1.1.4
	gorm.io/gorm v1.21.13
)
