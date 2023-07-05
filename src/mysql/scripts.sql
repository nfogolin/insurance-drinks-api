CREATE TABLE Drink (
    drink_id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    type varchar(255) NOT NULL,
    price decimal(6,2) NOT NULL,
    aging int NULL,
    PRIMARY KEY (drink_id)
);