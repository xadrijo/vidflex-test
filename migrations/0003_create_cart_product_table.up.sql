CREATE TABLE IF NOT EXISTS cart_product(
    id  serial PRIMARY KEY,
    cart_id integer REFERENCES cart (id),
    product_id integer REFERENCES product (id),
    quantity smallint,
    createdAt timestamp NOT NULL,
    updatedAt timestamp NOT NULL
);
