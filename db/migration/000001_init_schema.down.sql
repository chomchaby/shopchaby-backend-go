-- Drop Unique and Check
-- ALTER TABLE subproduct DROP CONSTRAINT IF EXISTS uc_product_variation;
ALTER TABLE "user" DROP CONSTRAINT IF EXISTS check_phone_format;
ALTER TABLE store DROP CONSTRAINT IF EXISTS check_phone_format;
ALTER TABLE subproduct DROP CONSTRAINT IF EXISTS check_sale_price_less_than_price;

-- Drop Foreign Keys and Triggers

-- Drop Foreign Keys in 'purchase' table
ALTER TABLE purchase DROP CONSTRAINT IF EXISTS fk_buyer;
ALTER TABLE purchase DROP CONSTRAINT IF EXISTS fk_seller;
ALTER TABLE purchase DROP CONSTRAINT IF EXISTS fk_product;
ALTER TABLE purchase DROP CONSTRAINT IF EXISTS fk_store;

-- Drop Foreign Keys in 'cart' table
ALTER TABLE cart DROP CONSTRAINT IF EXISTS fk_user;
ALTER TABLE cart DROP CONSTRAINT IF EXISTS fk_subproduct;

-- Drop Foreign Keys in 'favorite' table
ALTER TABLE favorite DROP CONSTRAINT IF EXISTS fk_user;
ALTER TABLE favorite DROP CONSTRAINT IF EXISTS fk_product;

-- Drop Foreign Keys in 'subproduct' table
ALTER TABLE subproduct DROP CONSTRAINT IF EXISTS fk_product;

-- Drop Foreign Keys in 'product_image' table
ALTER TABLE product_image DROP CONSTRAINT IF EXISTS fk_product;

-- Drop Foreign Keys in 'product' table
ALTER TABLE product DROP CONSTRAINT IF EXISTS fk_product_category;

-- Drop Foreign Keys in 'store' table
ALTER TABLE store DROP CONSTRAINT IF EXISTS fk_user;

-- Drop Triggers

-- Drop Triggers on 'cart' table
DROP TRIGGER IF EXISTS update_cart_updated_at ON cart;

-- Drop Triggers on 'subproduct' table
DROP TRIGGER IF EXISTS update_subproduct_updated_at ON subproduct;
DROP TRIGGER IF EXISTS update_product_prices_trigger ON subproduct;
DROP TRIGGER IF EXISTS update_product_onsale_trigger ON subproduct;
DROP TRIGGER IF EXISTS update_product_vendible_trigger ON subproduct;

-- Drop Triggers on 'product_image' table
DROP TRIGGER IF EXISTS update_product_image_updated_at ON product_image;

-- Drop Triggers on 'product' table
DROP TRIGGER IF EXISTS update_product_updated_at ON product;

-- Drop Triggers on 'store' table
DROP TRIGGER IF EXISTS update_store_updated_at ON store;

-- Drop Triggers on 'user' table
DROP TRIGGER IF EXISTS update_user_updated_at ON "user";

-- Drop Triggers on 'product_category' table
DROP TRIGGER IF EXISTS update_product_category_updated_at ON product_category;


-- Drop Functions and Extensions

-- Drop Function 'update_updated_at'
DROP FUNCTION IF EXISTS update_updated_at();

-- Drop Function 'update_product_prices'
DROP FUNCTION IF EXISTS update_product_prices();
DROP FUNCTION IF EXISTS update_product_onsale();
DROP FUNCTION IF EXISTS update_product_vendible();


-- Drop Tables

DROP TABLE IF EXISTS purchase;
DROP TABLE IF EXISTS cart;
DROP TABLE IF EXISTS favorite;
DROP TABLE IF EXISTS subproduct;
DROP TABLE IF EXISTS product_image;
DROP TABLE IF EXISTS product;
DROP TABLE IF EXISTS store;
DROP TABLE IF EXISTS "user";
DROP TABLE IF EXISTS product_category;

-- Drop Extension 'uuid-ossp'
DROP EXTENSION IF EXISTS "uuid-ossp";