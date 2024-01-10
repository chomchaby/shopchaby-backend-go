-------------------------------------------------------------------
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE OR REPLACE FUNCTION update_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ language 'plpgsql';

-------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS "user" (
	email VARCHAR(50) PRIMARY KEY,
	username VARCHAR(30) NOT NULL UNIQUE,
	pwd_hash VARCHAR(100) NOT NULL,
	phone VARCHAR(15) NOT NULL,
	address VARCHAR(150) NOT NULL,
	balance INTEGER DEFAULT 0 NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT check_phone_format CHECK (phone ~ '^[0-9]+$')
    
);


CREATE OR REPLACE TRIGGER update_user_updated_at
    BEFORE UPDATE
    ON "user"
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

-------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS store (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    user_email VARCHAR(50) NOT NULL UNIQUE,
	name VARCHAR(30) NOT NULL,
	description VARCHAR(300) NOT NULL,
    address VARCHAR(150) NOT NULL,
	phone VARCHAR(15) NOT NULL,
    pic_url VARCHAR(2048) NOT NULL,
    balance INTEGER DEFAULT 0 NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_user FOREIGN KEY (user_email) REFERENCES "user"(email) ON DELETE CASCADE,
    CONSTRAINT check_phone_format CHECK (phone ~ '^[0-9]+$')
);

CREATE OR REPLACE TRIGGER  update_store_updated_at
    BEFORE UPDATE
    ON store
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();


-------------------------------------------------------------------


CREATE TABLE IF NOT EXISTS product_category (
	name VARCHAR(30) PRIMARY KEY,
	pic_url VARCHAR(2048) NOT NULL UNIQUE,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL
);

CREATE OR REPLACE TRIGGER  update_product_category_updated_at
    BEFORE UPDATE
    ON product_category
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

-------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS product (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	category VARCHAR(30) NOT NULL,
	name VARCHAR(50) NOT NULL,
	description VARCHAR(300) NOT NULL,
	store_id UUID NOT NULL,
    suspend BOOLEAN DEFAULT FALSE NOT NULL,
    max_price INTEGER,
	min_price INTEGER,
    onsale BOOLEAN DEFAULT FALSE NOT NULL,
    vendible BOOLEAN DEFAULT FALSE NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_product_category FOREIGN KEY (category) REFERENCES product_category(name) ON DELETE SET NULL ON UPDATE CASCADE,
	CONSTRAINT fk_store FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE CASCADE
);

CREATE OR REPLACE TRIGGER update_product_updated_at
    BEFORE UPDATE
    ON product
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

-------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS product_image (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_id UUID NOT NULL,
    title VARCHAR(30) NOT NULL,
    pic_url VARCHAR(2048) NOT NULL,
    is_default BOOLEAN DEFAULT TRUE NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE
);

CREATE OR REPLACE TRIGGER update_product_image_updated_at
    BEFORE UPDATE
    ON product_image
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

-------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS subproduct (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	product_id UUID NOT NULL,
    variation VARCHAR(30) NOT NULL,
    stock_amount INTEGER DEFAULT 0 NOT NULL,
    price INTEGER NOT NULL,
    sale_price INTEGER,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE,
    CONSTRAINT check_sale_price_less_than_price CHECK (sale_price IS NULL OR sale_price < price)
    -- CONSTRAINT uc_product_variation UNIQUE (product_id, variation)
);

CREATE OR REPLACE FUNCTION update_product_prices()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'DELETE' THEN
        IF NOT EXISTS (SELECT 1 FROM subproduct WHERE product_id = OLD.product_id) THEN
            UPDATE product
            SET min_price = NULL,
                max_price = NULL
            WHERE id = OLD.product_id;
        ELSE
            UPDATE product
            SET min_price = (
                SELECT MIN(price) FROM subproduct WHERE product_id = OLD.product_id
            ),
            max_price = (
                SELECT MAX(price) FROM subproduct WHERE product_id = OLD.product_id
            )
            WHERE id = OLD.product_id;
        END IF;
    ELSE
        IF NOT EXISTS (SELECT 1 FROM subproduct WHERE product_id = NEW.product_id) THEN
            UPDATE product
            SET min_price = NULL,
                max_price = NULL
            WHERE id = NEW.product_id;
        ELSE
            UPDATE product
            SET min_price = (
                SELECT MIN(price) FROM subproduct WHERE product_id = NEW.product_id
            ),
            max_price = (
                SELECT MAX(price) FROM subproduct WHERE product_id = NEW.product_id
            )
            WHERE id = NEW.product_id;
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_product_prices_trigger
AFTER INSERT OR UPDATE OR DELETE
ON subproduct
FOR EACH ROW
EXECUTE FUNCTION update_product_prices();


CREATE OR REPLACE FUNCTION update_product_onsale()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'DELETE' THEN
        IF NOT EXISTS (SELECT 1 FROM subproduct WHERE product_id = OLD.product_id AND sale_price IS NOT NULL) THEN
            UPDATE product
            SET onsale = FALSE
            WHERE id = OLD.product_id;
        ELSE
            UPDATE product
            SET onsale = TRUE
            WHERE id = OLD.product_id;
        END IF;
    ELSE
        IF NOT EXISTS (SELECT 1 FROM subproduct WHERE product_id = NEW.product_id AND sale_price IS NOT NULL) THEN
            UPDATE product
            SET onsale = FALSE
            WHERE id = NEW.product_id;
        ELSE
            UPDATE product
            SET onsale = TRUE
            WHERE id = NEW.product_id;
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_product_onsale_trigger
AFTER INSERT OR UPDATE OR DELETE
ON subproduct
FOR EACH ROW
EXECUTE FUNCTION update_product_onsale();


CREATE OR REPLACE FUNCTION update_product_vendible()
RETURNS TRIGGER AS $$
BEGIN
    IF TG_OP = 'DELETE' THEN
        IF NOT EXISTS (SELECT 1 FROM subproduct WHERE product_id = OLD.product_id) THEN
            UPDATE product
            SET vendible = FALSE
            WHERE id = OLD.product_id;
        ELSE
            UPDATE product
            SET vendible = TRUE
            WHERE id = OLD.product_id;
        END IF;
    ELSE
        IF NOT EXISTS (SELECT 1 FROM subproduct WHERE product_id = NEW.product_id) THEN
            UPDATE product
            SET vendible = FALSE
            WHERE id = NEW.product_id;
        ELSE
            UPDATE product
            SET vendible = TRUE
            WHERE id = NEW.product_id;
        END IF;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE TRIGGER update_product_vendible_trigger
AFTER INSERT OR UPDATE OR DELETE
ON subproduct
FOR EACH ROW
EXECUTE FUNCTION update_product_vendible();


-------------------------------------------------------------------


CREATE TABLE IF NOT EXISTS favorite (
	user_email VARCHAR(50) NOT NULL,
    product_id UUID NOT NULL,
	timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (user_email, product_id),
    CONSTRAINT fk_user FOREIGN KEY (user_email) REFERENCES "user"(email) ON DELETE CASCADE,
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE CASCADE
);


-------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS cart (
	user_email VARCHAR(50) NOT NULL,
    subproduct_id UUID NOT NULL,
	quantity INTEGER NOT NULL,
	created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    PRIMARY KEY (user_email, subproduct_id),
    CONSTRAINT fk_user FOREIGN KEY (user_email) REFERENCES "user"(email) ON DELETE CASCADE,
    CONSTRAINT fk_subproduct FOREIGN KEY (subproduct_id) REFERENCES subproduct(id) on DELETE CASCADE
);

CREATE OR REPLACE TRIGGER update_cart_updated_at
    BEFORE UPDATE
    ON cart
    FOR EACH ROW
EXECUTE PROCEDURE update_updated_at();

-------------------------------------------------------------------

CREATE TABLE IF NOT EXISTS purchase (
	id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
	buyer_email VARCHAR(50) NOT NULL,
    buyer_username VARCHAR(30) NOT NULL,
    seller_email VARCHAR(50) NOT NULL,
    seller_username VARCHAR(30) NOT NULL,
    product_id UUID,
    store_id UUID,
    product_name VARCHAR(50) NOT NULL,
    variation VARCHAR(30) NOT NULL,
    price INTEGER NOT NULL,
    quantity INTEGER NOT NULL,
    store_name VARCHAR(30) NOT NULL,
    pic_url VARCHAR(2024) NOT NULL,
    shipping_address VARCHAR(100) NOT NULL,
    receiver VARCHAR(50) NOT NULL,
	timestamp TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,

    CONSTRAINT chk_different_buyer_seller CHECK (buyer_email <> seller_email),
    CONSTRAINT fk_buyer FOREIGN KEY (buyer_email) REFERENCES "user"(email),
    CONSTRAINT fk_seller FOREIGN KEY (seller_email) REFERENCES "user"(email), 
    CONSTRAINT fk_product FOREIGN KEY (product_id) REFERENCES product(id) ON DELETE SET NULL,
    CONSTRAINT fk_store FOREIGN KEY (store_id) REFERENCES store(id) ON DELETE SET NULL

);


-------------------------------------------------------------------
-------------------------------------------------------------------