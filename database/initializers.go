package database

import (
	"fmt"
	"log"
	"server-veggie/config"
	"server-veggie/database/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Initializers() *gorm.DB {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalf("error loading config: %v", err)
	}
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s TimeZone=UTC",
		config.RDSHostName, config.RDSPort, config.RDSUser, config.RDSPassword, config.RDSDBName, config.SSMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	// Tạo kiểu ENUM nếu chưa tồn tại
	err = db.Exec(`
        DO $$ BEGIN
            IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'season_enum') THEN
                CREATE TYPE season_enum AS ENUM ('spring', 'summer', 'autumn','winter');
            END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN
                CREATE TYPE status_enum AS ENUM ('active','inactive');
            END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_product_enum') THEN
                CREATE TYPE status_product_enum AS ENUM ('available','unavailable','expired','promotion','pending');
            END IF;
        END $$;
    `).Error
	if err != nil {
		log.Fatal("Failed to create ENUM type:", err)
	}

	if err := db.AutoMigrate(
		&schema.User{},
		&schema.Role{},
		&schema.UserRole{},
		&schema.Permission{},
		&schema.RolePermission{},
		&schema.Category{},
		&schema.InStockProduct{},
		&schema.Product{},
		&schema.Supplier{},
		&schema.SupplierProduct{},
		&schema.Currency{},
		&schema.Customer{},
		&schema.ZonePrice{},
		&schema.PurchasePrice{},
		&schema.SalesPrice{},
		&schema.Invoice{},
		&schema.PurchaseTransaction{},
		&schema.PurchaseTransactionInvoice{},
	); err != nil {
		log.Fatalln(err)
	}
	return db
}
