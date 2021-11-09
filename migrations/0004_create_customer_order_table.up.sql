CREATE TABLE IF NOT EXISTS customer_order(
    id serial PRIMARY KEY,
    cart_id integer REFERENCES cart (id),
    createdAt timestamp NOT NULL,
    updatedAt timestamp NOT NULL
);
