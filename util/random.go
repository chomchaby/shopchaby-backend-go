package util

import (
	"database/sql"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

// RandomInt gererates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

// RandomBoolean generates true or false randomly
func RandomBoolean() bool {
	// Generate a random number (0 or 1)
	randomNum := rand.Intn(2)

	// Convert the random number to a boolean
	return randomNum == 1
}

// RandomUUID generates a random UUID (Version 4)
func RandomUUID() uuid.UUID {
	uuid.SetRand(rand.New(rand.NewSource(time.Now().UnixNano())))
	return uuid.New()
}

/////////////////////////////////////////////////////
// Random User

type RandomUserDetailParams struct {
	Email    string
	Username string
	Phone    string
	Address  string
}

// RandomUser generates a random user email
func RandomUserDetail() RandomUserDetailParams {
	name := RandomString(int(RandomInt(4, 6)))
	user := RandomUserDetailParams{}
	user.Email = name + "@test.com"
	user.Username = name
	user.Phone = strconv.Itoa(int(RandomInt(100000000, 999999999)))
	user.Address = name + " address"
	return user
}

/////////////////////////////////////////////////////
// Random Store

type RandomStoreDetailParams struct {
	ID          uuid.UUID
	Name        string
	Description string
	Address     string
	Phone       string
	PicUrl      string
}

func RandomStoreDetail() RandomStoreDetailParams {
	name := RandomString(int(RandomInt(4, 6)))
	store := RandomStoreDetailParams{}
	store.ID = RandomUUID()
	store.Name = name
	store.Description = "Description of " + name + " store"
	store.Address = name + " address"
	store.Phone = strconv.Itoa(int(RandomInt(100000000, 999999999)))
	store.PicUrl = "www.pic-store/" + name
	return store
}

// ///////////////////////////////////////////////////
// Random Product
type RandomProductDetailParams struct {
	ID          uuid.UUID
	Category    string
	Name        string
	Description string
	Suspend     bool
}

func RandomProductDetail() RandomProductDetailParams {
	categories := []string{"Clothes", "Shoes", "Bags", "Beauty & Personal Care", "Fashion Accessories", "Computers & Laptops", "Mobile & Gadgets", "Food & Beverages", "Home Appliances"}

	category := categories[RandomInt(0, int64(len(categories)-1))] // Pick a random category
	name := RandomString(int(RandomInt(4, 6)))

	product := RandomProductDetailParams{}
	product.ID = RandomUUID()
	product.Category = category
	product.Name = name
	product.Description = "Description of " + name + " product"
	product.Suspend = RandomBoolean()

	return product
}

type RandomSubproductDetailParams struct {
	ID          uuid.UUID
	Variation   string
	StockAmount int32
	Price       int32
	SalePrice   sql.NullInt32
}

func RandomSubproductDetail(onsale bool) RandomSubproductDetailParams {

	subproduct := RandomSubproductDetailParams{}
	subproduct.ID = RandomUUID()
	subproduct.Variation = RandomString(int(RandomInt(4, 6)))
	subproduct.StockAmount = int32(RandomInt(0, 999))
	subproduct.Price = int32(RandomInt(5, 999999))

	// Create a sql.NullInt32 for the SalePrice field
	saleprice := int32(RandomInt(1, int64(subproduct.Price)-1))
	var nullableSalePrice sql.NullInt32
	if onsale {
		nullableSalePrice = sql.NullInt32{Int32: saleprice, Valid: true} // Assign the integer value with validity
	} else {
		nullableSalePrice = sql.NullInt32{Int32: 0, Valid: false} // If not valid, assign zero value with false validity
	}

	subproduct.SalePrice = nullableSalePrice

	return subproduct
}

type RandomProductImageDetailParams struct {
	ID        uuid.UUID
	Title     string
	PicUrl    string
	IsDefault bool
}

func RandomProductImageDetail() RandomProductImageDetailParams {

	title := RandomString(int(RandomInt(4, 6)))

	image := RandomProductImageDetailParams{}
	image.ID = RandomUUID()
	image.Title = title
	image.PicUrl = "www.pic-product/" + title
	image.IsDefault = RandomBoolean()

	return image
}
