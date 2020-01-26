-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

CREATE TABLE IF NOT EXISTS plans (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    cleanName varchar(255) NOT NULL,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS cycles (
    id int NOT NULL AUTO_INCREMENT,
    idPlan int NOT NULL,
    type varchar(255) NOT NULL,
    priceRenew varchar(255) NOT NULL,
    priceOrder varchar(255) NOT NULL,
    months int NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (idPlan) REFERENCES plans(id)
);

INSERT INTO plans VALUES 
(5, "Plano P", "planoP"), 
(6, "Plano M", "planoM"), 
(335, "Plano Turbo", "planoTurbo");


INSERT INTO cycles (idPlan, type, priceRenew, priceOrder, months) VALUES 
(5, "monthly", "24.19", "29.19", 1),
(6, "monthly", "29.69", "29.69", 1),
(335, "monthly", "44.99", "44.99", 1),

(5, "semiannually", "128.34", "128.34", 6), 
(6, "semiannually", "159.54", "159.54", 6), 
(335, "semiannually", "257.94", "257.94", 6), 

(5, "biennially", "393.36", "393.36", 24), 
(6, "biennially", "532.56", "532.56", 24), 
(335, "biennially", "983.76", "983.76", 24), 

(5, "triennially", "561.13", "561.13", 36), 
(6, "triennially", "764.22", "764.22", 36), 
(335, "triennially", "1439.64", "1439.64", 36), 

(5, "quarterly", "67.17", "67.17", 3), 
(6, "quarterly", "82.77", "82.77", 3), 
(335, "quarterly", "131.97", "131.97", 3), 

(5, "annually", "220.66", "220.66", 12), 
(6, "annually", "286.66", "286.66", 12), 
(335, "annually", "503.88", "503.88", 12);

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE prices;
DROP TABLE cycles;