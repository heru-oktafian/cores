package configs

import (
	context "context"
	log "log"
	os "os"
	strconv "strconv"
	time "time"

	// models "github.com/heru-oktafian/cores/models"
	redis "github.com/redis/go-redis/v9"
	postgres "gorm.io/driver/postgres"
	gorm "gorm.io/gorm"
	logger "gorm.io/gorm/logger"
)

var (
	DB  *gorm.DB
	RDB *redis.Client
)

func SetupDB() (err error) {

	ctx := context.Background()

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASS")
	db_name := os.Getenv("DB_NAME")
	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_pass := os.Getenv("REDIS_PASS")

	redis_short := os.Getenv("REDIS_SHORT")

	redis_db, err := strconv.Atoi(redis_short)

	dsn := "user=" + db_user + " password=" + db_pass + " host=" + db_host + " port=" + db_port + " dbname=" + db_name + "  sslmode=disable TimeZone=Asia/Jakarta"
	// DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // Output log
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Silent, // ⬅️ Ubah sesuai kebutuhan: Silent, Error, Warn, Info
				Colorful:      true,
			},
		),
	})

	DB.AutoMigrate(
	// &models.AnotherIncomes{},
	// &models.BalanceReport{},
	// &models.Branch{},
	// &models.BuyReturnItems{},
	// &models.BuyReturns{},
	// &models.DailyProfitReport{},
	// &models.DailyAsset{},
	// &models.Defectas{},
	// &models.DefectaItems{},
	// &models.DuplicateReceiptItems{},
	// &models.DuplicateReceipts{},
	// &models.Expenses{},
	// &models.FirstStockItems{},
	// &models.FirstStocks{},
	// &models.MemberCategory{},
	// &models.Member{},
	// &models.OpnameItems{},
	// &models.Opnames{},
	// &models.ProductCategory{},
	// &models.Product{},
	// &models.PurchaseItems{},
	// &models.Purchases{},
	// &models.Defectas{},
	// &models.DefectaItems{},
	// &models.SaleItems{},
	// &models.Sales{},
	// &models.SupplierCategory{},
	// &models.Supplier{},
	// &models.TransactionReports{},
	// &models.UnitConversion{},
	// &models.Unit{},
	// &models.UserBranch{},
	// &models.User{},
	)

	// Check connection Postgres
	if err != nil {
		log.Fatalf("failed to connect to PostgreSQL database: %v", err)
	}

	// Panggil fungsi migrasi
	// err = runMigrations(DB)
	// if err != nil {
	// 	panic("failed to migrate: " + err.Error())
	// }

	// Connect to database Redis
	RDB = redis.NewClient(&redis.Options{
		Addr:     redis_host + ":" + redis_port,
		Password: redis_pass,
		DB:       redis_db,
	})

	// Check connection Redis
	_, err = RDB.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("failed to connect to Redis database: %v", err)
	}

	return nil
}

// runMigrations bertanggung jawab untuk menambahkan kolom baru
func runMigrations(db *gorm.DB) error {
	//m := db.Migrator()

	// Tambahkan kolom 'subscription_type' jika belum ada
	// if !m.HasColumn(&models.Branch{}, "SubscriptionType") {
	// 	err := m.AddColumn(&models.Branch{}, "SubscriptionType")
	// 	if err != nil {
	// 		return fmt.Errorf("failed to add column 'subscription_type': %w", err)
	// 	}
	// 	fmt.Println("Column 'subscription_type' successfully added.")
	// } else {
	// 	fmt.Println("Column 'subscription_type' already exists. Skipping.")
	// }

	// // Tambahkan kolom 'quota' jika belum ada
	// if !m.HasColumn(&models.Branch{}, "Quota") {
	// 	err := m.AddColumn(&models.Branch{}, "Quota")
	// 	if err != nil {
	// 		return fmt.Errorf("failed to add column 'quota': %w", err)
	// 	}
	// 	fmt.Println("Column 'quota' successfully added.")
	// } else {
	// 	fmt.Println("Column 'quota' already exists. Skipping.")
	// }

	// //Tambahkan kolom 'expired_date' Opname Item jika belum ada
	// if !m.HasColumn(&models.FirstStockItems{}, "ExpiredDate") {
	// 	err := m.AddColumn(&models.FirstStockItems{}, "ExpiredDate")
	// 	if err != nil {
	// 		return fmt.Errorf("failed to add column 'expired_date': %w", err)
	// 	}
	// 	fmt.Println("Column 'expired_date' Opname Items successfully added.")
	// } else {
	// 	fmt.Println("Column 'expired_date' Opname Items already exists. Skipping.")
	// }

	return nil
}
