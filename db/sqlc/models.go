// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	UserEmail    string    `json:"user_email"`
	SubproductID uuid.UUID `json:"subproduct_id"`
	Quantity     int32     `json:"quantity"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Favorite struct {
	UserEmail string    `json:"user_email"`
	ProductID uuid.UUID `json:"product_id"`
	Timestamp time.Time `json:"timestamp"`
}

type Product struct {
	ID          uuid.UUID     `json:"id"`
	Category    string        `json:"category"`
	Name        string        `json:"name"`
	Description string        `json:"description"`
	StoreID     uuid.UUID     `json:"store_id"`
	Suspend     bool          `json:"suspend"`
	MaxPrice    sql.NullInt32 `json:"max_price"`
	MinPrice    sql.NullInt32 `json:"min_price"`
	Onsale      bool          `json:"onsale"`
	Vendible    bool          `json:"vendible"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type ProductCategory struct {
	Name      string    `json:"name"`
	PicUrl    string    `json:"pic_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ProductImage struct {
	ID        uuid.UUID `json:"id"`
	ProductID uuid.UUID `json:"product_id"`
	Title     string    `json:"title"`
	PicUrl    string    `json:"pic_url"`
	IsDefault bool      `json:"is_default"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Purchase struct {
	ID              uuid.UUID     `json:"id"`
	BuyerEmail      string        `json:"buyer_email"`
	BuyerUsername   string        `json:"buyer_username"`
	SellerEmail     string        `json:"seller_email"`
	SellerUsername  string        `json:"seller_username"`
	ProductID       uuid.NullUUID `json:"product_id"`
	StoreID         uuid.NullUUID `json:"store_id"`
	ProductName     string        `json:"product_name"`
	Variation       string        `json:"variation"`
	Price           int32         `json:"price"`
	Quantity        int32         `json:"quantity"`
	StoreName       string        `json:"store_name"`
	PicUrl          string        `json:"pic_url"`
	ShippingAddress string        `json:"shipping_address"`
	Receiver        string        `json:"receiver"`
	Timestamp       time.Time     `json:"timestamp"`
}

type Store struct {
	ID          uuid.UUID `json:"id"`
	UserEmail   string    `json:"user_email"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Address     string    `json:"address"`
	Phone       string    `json:"phone"`
	PicUrl      string    `json:"pic_url"`
	Balance     int32     `json:"balance"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Subproduct struct {
	ID          uuid.UUID     `json:"id"`
	ProductID   uuid.UUID     `json:"product_id"`
	Variation   string        `json:"variation"`
	StockAmount int32         `json:"stock_amount"`
	Price       int32         `json:"price"`
	SalePrice   sql.NullInt32 `json:"sale_price"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
}

type User struct {
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	PwdHash   string    `json:"pwd_hash"`
	Phone     string    `json:"phone"`
	Address   string    `json:"address"`
	Balance   int32     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
