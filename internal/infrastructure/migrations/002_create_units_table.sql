CREATE TABLE units (
    id SERIAL PRIMARY KEY,
    unit_number INT NOT NULL,
    occupants_count INT NOT NULL,
    area FLOAT NOT NULL,
    building_id INT NOT NULL,
    FOREIGN KEY (building_id) REFERENCES buildings (id)
);


INSERT INTO buildings (name, address, total_units, total_area)
VALUES ('Building A', '123 Main street', 10, 2500.50);

SELECT * FROM buildings;