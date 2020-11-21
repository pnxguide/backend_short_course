CREATE USER 'dbuser'@localhost
IDENTIFIED BY '12345678';

GRANT ALL PRIVILEGES 
ON *.* 
TO 'dbuser'@localhost 
IDENTIFIED BY '12345678';

FLUSH PRIVILEGES;

system mysql -u dbuser -p

CREATE DATABASE example;
use example;

CREATE TABLE product_types (
    product_type_id     INT PRIMARY KEY,
    product_type_name   VARCHAR(50)
);

INSERT INTO product_types (product_type_id, product_type_name)
VALUES (1, 'Food'), (2, 'Drink');

CREATE TABLE products (
    product_id          INT PRIMARY KEY,
    product_name        VARCHAR(50),
    product_quantity    INT,
    product_price       DECIMAL(10,4),
    product_product_type_id INT REFERENCES product_types
);

/*
ALTER TABLE products 
ADD CONSTRAINT fk_products_product_types 
FOREIGN KEY (product_product_type_id) 
    REFERENCES product_types (product_type_id)
    ON DELETE RESTRICT
    ON UPDATE RESTRICT;
*/

INSERT INTO products (product_id, product_name,
    product_quantity, product_price,
    product_product_type_id)
VALUES
(1, 'Apple', 5, 10.3, 1),
(2, 'Bread', 23, 3.6, 1),
(3, 'Lobster Roll', 3, 50.7, 1),
(4, 'Drinking Water', 30, 1.2, 2),
(5, 'Cola', 26, 2.5, 2);

SELECT product_name, product_type_name 
FROM products 
JOIN product_types ON product_product_type_id = product_type_id;

UPDATE products
SET product_price = 50.8
WHERE product_id = 3;

INSERT INTO products (product_id) VALUES (99);

DELETE FROM products
WHERE product_id = 99;

