CREATE TABLE animals (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

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
