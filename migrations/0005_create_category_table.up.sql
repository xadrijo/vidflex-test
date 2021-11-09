CREATE TABLE IF NOT EXISTS category(
    id smallserial PRIMARY KEY,
    label varchar (50),
    parent_id integer,
    createdAt timestamp NOT NULL,
    updatedAt timestamp NOT NULL
);
