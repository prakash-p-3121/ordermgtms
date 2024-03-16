CREATE TABLE orders (
    id VARBINARY(2000) NOT NULL PRIMARY KEY,
    id_bit_count BIGINT UNSIGNED NOT NULL,
    product_id VARBINARY(2000) NOT NULL,
    seller_id VARBINARY(2000) NOT NULL,
    buyer_id VARBINARY(2000) NOT NULL,
    listing_id VARBINARY(2000) NOT NULL,
    amount DOUBLE NOT NULL,
    currency VARCHAR(10) NOT NULL,
    state_id INT UNSIGNED NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    INDEX orders_product_id (product_id),
    INDEX orders_seller_id (seller_id),
    INDEX orders_buyer_id (buyer_id),
    INDEX orders_listing_id (listing_id),
    INDEX orders_amount (amount)
);
