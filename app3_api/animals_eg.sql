CREATE TABLE animals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

ALTER TABLE animals ADD CONSTRAINT animals_id_pk PRIMARY KEY (id);

INSERT INTO animals ("name") VALUES
	 ('Dog'),
	 ('Cat'),
	 ('Horse'),
	 ('Cow'),
	 ('Rabbit'),
	 ('Sheep'),
	 ('Duck'),
	 ('Parrot'),
	 ('Canary');

