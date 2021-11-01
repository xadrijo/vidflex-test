CREATE TABLE IF NOT EXISTS product(
    id  smallserial PRIMARY KEY,
    label varchar (50),
    type varchar (50),
    url varchar (50),
    weight varchar (50),
    createdAt timestamp NOT NULL,
    updatedAt timestamp NOT NULL
);