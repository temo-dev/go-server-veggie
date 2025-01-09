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
	            CREATE TYPE season_enum AS ENUM ('spring', 'summer', 'autumn','winter','all');
	        END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_enum') THEN
	            CREATE TYPE status_enum AS ENUM ('active','inactive');
	        END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'status_product_enum') THEN
	            CREATE TYPE status_product_enum AS ENUM ('available','unavailable','expired','promotion','pending');
	        END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'container_type_enum') THEN
	            CREATE TYPE container_type_enum AS ENUM ('dry container','reefer container','special container');
	        END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'length_enum') THEN
	            CREATE TYPE length_enum AS ENUM ('cm','inch');
	        END IF;
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'container_size_enum') THEN
	            CREATE TYPE container_size_enum AS ENUM ('20ft','40ft','40ft hc','20ft hc');
	        END IF;
	    END $$;
	`).Error
	if err != nil {
		log.Fatal("Failed to create ENUM type:", err)
	}

	if err := db.AutoMigrate(
		&schema.User{},
		&schema.Role{},
		&schema.Currency{},
		&schema.UserRole{},
		&schema.Permission{},
		&schema.RolePermission{},
		&schema.Category{},
		&schema.SubCategory{},
		&schema.AttitudeProductPackage{},
		&schema.AttitudeContainer{},
		&schema.InStockProduct{},
		&schema.Supplier{},
		&schema.Brand{},
		&schema.SupplierBrand{},
		&schema.SupplierProduct{},
		&schema.Product{},
		&schema.Customer{},
		&schema.ZonePrice{},
		&schema.PurchasePrice{},
		&schema.SalesPrice{},
		&schema.Invoice{},
		&schema.PurchaseTransaction{},
		&schema.PurchaseTransactionInvoice{},
		&schema.Promotion{},
		&schema.PromotionProduct{},
		&schema.Tag{},
		&schema.TagProduct{},
	); err != nil {
		log.Fatalln(err)
	}
	return db
}
