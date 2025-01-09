package schema

import (
	"time"
)

// ============= user ===================
type User struct {
	UserID    string     `gorm:"type:uuid;primaryKey"`
	UserName  string     `gorm:"type:varchar;unique;not null"`
	Email     string     `gorm:"type:varchar;unique;not null"`
	Password  string     `gorm:"type:varchar;not null"`
	Status    string     `gorm:"type:status_enum;default:'inactive'"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	UserRoles []UserRole `gorm:"foreignKey:UserID"`
}

type Role struct {
	RoleID      string     `gorm:"type:uuid;primaryKey"`
	RoleName    string     `gorm:"type:varchar;not null"`
	Description string     `gorm:"type:varchar"`
	UserRoles   []UserRole `gorm:"foreignKey:RoleID"`
}

type UserRole struct {
	UserRoleID string `gorm:"type:uuid;primaryKey"`
	UserID     string `gorm:"type:uuid;not null"`
	RoleID     string `gorm:"type:uuid;not null"`
}

type Permission struct {
	PermissionID   string `gorm:"type:uuid;primaryKey"`
	PermissionName string `gorm:"type:varchar;not null"`
	Description    string `gorm:"type:varchar"`
}

type RolePermission struct {
	RolePermissionID string `gorm:"type:uuid;primaryKey"`
	RoleID           string `gorm:"type:uuid;not null"`
	PermissionID     string `gorm:"type:uuid;not null"`
}

// ========== product ===============
type Category struct {
	CategoryID      string        `gorm:"type:uuid;primaryKey"`
	CategoryNameVN  string        `gorm:"type:varchar;unique;not null"`
	CategoryNameENG string        `gorm:"type:varchar;unique;not null"`
	CategoryNameDE  string        `gorm:"type:varchar;unique;not null"`
	CategoryNameTH  string        `gorm:"type:varchar;unique;not null"`
	ImageURL        string        `gorm:"type:varchar"`
	SubCategories   []SubCategory `gorm:"foreignKey:CategoryID"`
}

type SubCategory struct {
	SubCategoryID      string    `gorm:"type:uuid;primaryKey"`
	SubCategoryNameVN  string    `gorm:"type:varchar;unique;not null"`
	SubCategoryNameENG string    `gorm:"type:varchar;unique;not null"`
	SubCategoryNameDE  string    `gorm:"type:varchar;unique;not null"`
	SubCategoryNameTH  string    `gorm:"type:varchar;unique;not null"`
	ImageURL           string    `gorm:"type:varchar"`
	Dph                int32     `gorm:"type:integer;not null"`
	CategoryID         string    `gorm:"type:uuid;not null"`
	Products           []Product `gorm:"foreignKey:CategoryID"`
}

type InStockProduct struct {
	InStockProductID   string    `gorm:"type:uuid;primaryKey"`
	Quantity           int       `gorm:"type:int;not null"`
	ProductID          string    `gorm:"type:uuid;not null"`
	SupplierID         string    `gorm:"type:uuid;not null"`
	EanCode            string    `gorm:"type:varchar;unique;not null"`
	EanCodeBox         string    `gorm:"type:varchar;unique;not null"`
	LeadTime           int       `gorm:"type:int;not null"`
	CountryOfOrigin    string    `gorm:"type:varchar"`
	CountrySubOfOrigin string    `gorm:"type:varchar"`
	ExpiryDate         time.Time `gorm:"type:date;not null"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	UpdatedAt          time.Time `gorm:"autoUpdateTime"`
}

// ================= supplier ================
type Supplier struct {
	SupplierID         string          `gorm:"type:uuid;primaryKey"`
	SupplierName       string          `gorm:"type:varchar;not null"`
	SupplierCode       string          `gorm:"type:varchar;unique;not null"`
	TaxID              string          `gorm:"type:varchar;unique;not null"`
	CurrencyID         string          `gorm:"type:uuid;not null"`
	Description        string          `gorm:"type:varchar"`
	ContactInfo        string          `gorm:"type:varchar"`
	EmailPurchase      string          `gorm:"type:varchar;unique;not null"`
	Note               string          `gorm:"type:varchar"`
	Rate               float64         `gorm:"type:decimal(10,2);default:1.0"`
	OutstandingBalance float64         `gorm:"type:decimal(10,2);default:0.0"`
	Status             string          `gorm:"type:varchar;default:'active'"`
	PurchasePrices     []PurchasePrice `gorm:"foreignKey:SupplierID"`
	Invoices           []Invoice       `gorm:"foreignKey:SupplierID"`
	Products           []Product       `gorm:"many2many:supplier_products;"` // many to many
	DurationPakage     int             `gorm:"type:int;default:30"`
	Suppliers          []Supplier      `gorm:"many2many:supplier_brands;"` // many to many
	CreatedAt          time.Time       `gorm:"autoCreateTime"`
	UpdatedAt          time.Time       `gorm:"autoUpdateTime"`
}

type Brand struct {
	BrandID     string     `gorm:"type:uuid;primaryKey"`
	BrandName   string     `gorm:"type:varchar;unique;not null"`
	Description string     `gorm:"type:varchar"`
	Products    []Product  `gorm:"foreignKey:BrandID"`
	Suppliers   []Supplier `gorm:"many2many:supplier_brands;"` // many to many
}

type SupplierBrand struct {
	SupplierID string `gorm:"type:uuid;not null"`
	BrandID    string `gorm:"type:uuid;not null"`
}

type SupplierProduct struct {
	SupplierID string `gorm:"type:uuid;primaryKey"`
	ProductID  string `gorm:"type:uuid;primaryKey"`
}

type Product struct {
	ProductID                string           `gorm:"type:uuid;primaryKey"`
	ProductNameVN            string           `gorm:"type:varchar;not null"`
	ProductNameENG           string           `gorm:"type:varchar;not null"`
	ProductNameDE            string           `gorm:"type:varchar;not null"`
	ProductNameTH            string           `gorm:"type:varchar;not null"`
	ProductCode              string           `gorm:"type:varchar;not null;unique"`
	Dph                      int32            `gorm:"type:integer"`
	Description              string           `gorm:"type:varchar"`
	ImageURL                 string           `gorm:"type:varchar"`
	Status                   string           `gorm:"type:status_product_enum;default:'available'"`
	CategoryID               string           `gorm:"type:uuid;not null"`
	MinimumOrderQuantity     int              `gorm:"type:int;default:100"`
	MaximumOrderQuantity     int              `gorm:"type:int;default:1000"`
	ReorderLevel             int              `gorm:"type:int;default:100"`
	IsStackability           bool             `gorm:"type:boolean;default:true"`
	TemperatureRequirement   float64          `gorm:"type:decimal(10,2);default:0.0"`
	IsFragility              bool             `gorm:"type:boolean;default:false"`
	ShelfLife                int              `gorm:"type:int;not null"`
	Note                     string           `gorm:"type:varchar"`
	Season                   string           `gorm:"type:season_enum"`
	IsPublished              bool             `gorm:"type:boolean;default:false"`
	PublishedAt              time.Time        `gorm:"type:date"`
	PreOrder                 time.Time        `gorm:"type:date;not null"`
	Length                   string           `gorm:"type:length_enum;default:'cm'"`
	Width                    float64          `gorm:"type:decimal(10,2)"`
	Height                   float64          `gorm:"type:decimal(10,2)"`
	NetWeight                float64          `gorm:"type:decimal(10,2);not null"`
	GrossWeight              float64          `gorm:"type:decimal(10,2);not null"`
	Cubic                    float64          `gorm:"type:decimal(10,2);not null"`
	AttitudeProductPackageID string           `gorm:"type:uuid"`
	PurchasePrices           []PurchasePrice  `gorm:"foreignKey:ProductID"`
	SalesPrices              []SalesPrice     `gorm:"foreignKey:ProductID"`
	InStockProducts          []InStockProduct `gorm:"foreignKey:ProductID"`
	Suppliers                []Supplier       `gorm:"many2many:supplier_products;"`  // many to many
	Products                 []Product        `gorm:"many2many:promotion_products;"` // many to many
	BrandID                  string           `gorm:"type:uuid"`
	CreatedAt                time.Time        `gorm:"autoCreateTime"`
	UpdatedAt                time.Time        `gorm:"autoUpdateTime"`
}

type AttitudeProductPackage struct {
	AttitudeProductPackageID   string    `gorm:"type:uuid;primaryKey"`
	AttitudeProductPackageCode string    `gorm:"type:varchar;unique"`
	PackageLength              string    `gorm:"type:length_enum;default:'cm'"`
	PackageWidth               float64   `gorm:"type:decimal(10,2);not null"`
	PackageHeight              float64   `gorm:"type:decimal(10,2);not null"`
	PackageNetWeight           float64   `gorm:"type:decimal(10,2);not null"`
	PackageGrossWeight         float64   `gorm:"type:decimal(10,2);not null"`
	PackageCubic               float64   `gorm:"type:decimal(10,2);not null"`
	UnitsPerBox                int       `gorm:"type:int;not null"`
	Products                   []Product `gorm:"foreignKey:AttitudeProductPackageID"`
}

type AttitudeContainer struct {
	AttitudeContainerID string  `gorm:"type:uuid;primaryKey"`
	ContainerType       string  `gorm:"type:container_type_enum;default:'dry container'"`
	ConatinerSize       string  `gorm:"type:container_size_enum;default:'20ft'"`
	ConatinerWidth      float64 `gorm:"type:decimal(10,2);not null"`
	ConatinerHeight     float64 `gorm:"type:decimal(10,2);not null"`
	ConatinerLength     float64 `gorm:"type:decimal(10,2);not null"`
	ConatinerCubic      float64 `gorm:"type:decimal(10,2);not null"`
}

// ======== currency ===============
type Currency struct {
	CurrencyID           string                `gorm:"type:uuid;primaryKey"`
	CurrencyName         string                `gorm:"type:varchar;not null"`
	CurrencyCode         string                `gorm:"type:varchar;not null"`
	ExchangeRate         float64               `gorm:"type:decimal(10,2)"`
	Suppliers            []Supplier            `gorm:"foreignKey:CurrencyID"`
	PurchasePrices       []PurchasePrice       `gorm:"foreignKey:CurrencyID"`
	SalesPrices          []SalesPrice          `gorm:"foreignKey:CurrencyID"`
	Invoices             []Invoice             `gorm:"foreignKey:CurrencyID"`
	PurchaseTransactions []PurchaseTransaction `gorm:"foreignKey:CurrencyID"`
	Customers            []Customer            `gorm:"foreignKey:CurrencyID"`
}

// ============= customer =============
type Customer struct {
	CustomerID          string    `gorm:"type:uuid;primaryKey"`
	CustomerName        string    `gorm:"type:varchar;not null"`
	TaxID               string    `gorm:"type:varchar;unique"`
	ContactInfo         string    `gorm:"type:varchar"`
	EmailSales          string    `gorm:"type:varchar;unique"`
	Note                string    `gorm:"type:varchar"`
	Rate                float64   `gorm:"type:decimal(10,2);default:1.0"`
	TotalPurchaseAmount float64   `gorm:"type:decimal(10,2);default:0.0"`
	ZoneID              string    `gorm:"type:uuid;not null"`
	CurrencyID          string    `gorm:"type:uuid;not null"`
	SalesPriceID        string    `gorm:"type:uuid"`
	ZonePriceID         string    `gorm:"type:uuid;not null"`
	CreatedAt           time.Time `gorm:"autoCreateTime"`
	UpdatedAt           time.Time `gorm:"autoUpdateTime"`
}

type ZonePrice struct {
	ZonePriceID  string     `gorm:"type:uuid;primaryKey"`
	ZoneName     string     `gorm:"type:varchar;not null"`
	Customers    []Customer `gorm:"foreignKey:ZoneID"`
	SalesPriceID string     `gorm:"type:uuid;not null"`
}

// ============== pricing ===============
type PurchasePrice struct {
	PurchasePriceID string    `gorm:"type:uuid;primaryKey"`
	ProductID       string    `gorm:"type:uuid;not null"`
	CurrencyID      string    `gorm:"type:uuid;not null"`
	SupplierID      string    `gorm:"type:uuid;not null"`
	Season          string    `gorm:"type:season_enum;not null"`
	RetailPrice     float64   `gorm:"type:decimal(10,2); not null"`
	BoxPrice        float64   `gorm:"type:decimal(10,2); not null"`
	PalletPrice     float64   `gorm:"type:decimal(10,2); not null"`
	ContainerPrice  float64   `gorm:"type:decimal(10,2); not null"`
	CreatedAt       time.Time `gorm:"autoCreateTime"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime"`
}

type SalesPrice struct {
	SalesPriceID   string      `gorm:"type:uuid;primaryKey"`
	ProductID      string      `gorm:"type:uuid;not null"`
	CurrencyID     string      `gorm:"type:uuid;not null"`
	RetailPrice    float64     `gorm:"type:decimal(10,2); not null"`
	BoxPrice       float64     `gorm:"type:decimal(10,2); not null"`
	PalletPrice    float64     `gorm:"type:decimal(10,2); not null"`
	ContainerPrice float64     `gorm:"type:decimal(10,2); not null"`
	Season         string      `gorm:"type:season_enum;not null"`
	Customers      []Customer  `gorm:"foreignKey:SalesPriceID"`
	ZonePrices     []ZonePrice `gorm:"foreignKey:SalesPriceID"`
	CreatedAt      time.Time   `gorm:"autoCreateTime"`
	UpdatedAt      time.Time   `gorm:"autoUpdateTime"`
}

// ================ transaction ===========
type Invoice struct {
	InvoiceID            string                `gorm:"type:uuid;primaryKey"`
	ExpiryDate           time.Time             `gorm:"type:date;not null"`
	CurrencyID           string                `gorm:"uuid;not null"`
	Description          string                `gorm:"type:varchar"`
	Note                 string                `gorm:"type:varchar"`
	Status               string                `gorm:"type:varchar;default:'pending'"`
	Amount               float64               `gorm:"type:decimal(10,2);not null"`
	ImageURL             string                `gorm:"type:varchar;not null"`
	SupplierID           string                `gorm:"type:uuid;not null"`
	PurchaseTransactions []PurchaseTransaction `gorm:"many2many:purchase_transaction_invoices;"` // many to many
	CreatedAt            time.Time             `gorm:"autoCreateTime"`
	UpdatedAt            time.Time             `gorm:"autoUpdateTime"`
}

type PurchaseTransaction struct {
	PurchaseTransactionID string    `gorm:"type:uuid;primaryKey"`
	CurrencyID            string    `gorm:"uuid;not null"`
	TransferNumber        string    `gorm:"type:varchar;not null"`
	Amount                float64   `gorm:"type:decimal(10,2); not null"`
	Description           string    `gorm:"type:varchar"`
	Invoices              []Invoice `gorm:"many2many:purchase_transaction_invoices;"` // many to many
	ImageURL              string    `gorm:"type:varchar;not null"`
	CreatedAt             time.Time `gorm:"autoCreateTime"`
	UpdatedAt             time.Time `gorm:"autoUpdateTime"`
}

type PurchaseTransactionInvoice struct {
	InvoiceID             string `gorm:"type:uuid;not null"`
	PurchaseTransactionID string `gorm:"type:uuid;not null"`
}

// ============== Marketing ===========
type Promotion struct {
	PromotionID string    `gorm:"type:uuid;primaryKey"`
	Description string    `gorm:"type:varchar"`
	StartDate   time.Time `gorm:"type:date;not null"`
	EndDate     time.Time `gorm:"type:date;not null"`
	Discount    float64   `gorm:"type:decimal(10,2);not null"`
	Products    []Product `gorm:"many2many:promotion_products;"` // many to many
}

type PromotionProduct struct {
	PromotionID string `gorm:"type:uuid;not null"`
	ProductID   string `gorm:"type:uuid;not null"`
}
