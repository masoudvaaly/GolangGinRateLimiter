package config

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

//func ConnectDatabase() {
//	var err error
//	dsn := "postgresql://neondb_owner:npg_ex1LnEt0JfDW@ep-still-hat-adkwyjap-pooler.c-2.us-east-1.aws.neon.tech/crud_app?sslmode=disable"
//	DB, err = sql.Open("pgx", dsn)
//	if err != nil {
//		log.Fatalf("Could not connect to the database: %v", err)
//	}
//
//	// Always test the connection
//	if err := DB.Ping(); err != nil {
//		log.Fatalf("Database ping failed: %v", err)
//	}
//	fmt.Println("Database connected!")
//}

func ConnectDatabase() error {
	//db, err := sql.Open("sqlite3", "D:\\java\\go\\database.db")
	db, err := sql.Open("sqlite", "file:database.db?cache=shared&mode=rwc")
	if err != nil {
		return err
	}

	DB = db
	return nil
}

//func ConnectDatabase(ctx context.Context) {
//	//ctx := c.Request.Context() // standard context
//
//	socksDialer, _ := proxy.SOCKS5("tcp", "127.0.0.1:1080", nil, proxy.Direct)
//
//	dialer := func(ctx context.Context, network, addr string) (net.Conn, error) {
//		return socksDialer.Dial(network, addr)
//	}
//
//	config, _ := pgx.ParseConfig("postgresql://neondb_owner:npg_ex1LnEt0JfDW@ep-still-hat-adkwyjap-pooler.c-2.us-east-1.aws.neon.tech/crud_app?sslmode=require")
//	config.DialFunc = dialer
//
//	conn, err := pgx.ConnectConfig(ctx, config)
//	if err != nil {
//		fmt.Println("error:", err)
//		return
//	}
//	fmt.Println("Connected via SOCKS5 proxy ✅ Time:", conn)
//}
