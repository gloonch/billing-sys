CREATE TABLE payments (
    id SERIAL PRIMARY KEY,
    unit_id INT NOT NULL,
    amount FLOAT NOT NULL,
    payment_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    description TEXT,
    FOREIGN KEY (unit_id) REFERENCES units(id) ON DELETE CASCADE
);


INSERT INTO payments (unit_id, amount, payment_date, description)
VALUES (2, 100.0, "2024/12/12", "this payment has been purchased on this particular date");

SELECT * FROM payments;