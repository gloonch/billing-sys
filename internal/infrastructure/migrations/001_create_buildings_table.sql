CREATE TABLE buildings (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    address VARCHAR(255) NOT NULL,
    total_units INT NOT NULL,
    total_area FLOAT NOT NULL
);

INSERT INTO buildings (name, address, total_units, total_area)
VALUES ('Building A', '123 Main street', 10, 2500.50);

SELECT * FROM buildings;