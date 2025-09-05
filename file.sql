--данные для создания и заполнения ДБ-1(нужно создать программу автозаполнения)
DROP DATABASE IF EXISTS cars;
CREATE DATABASE cars;
DROP TABLE IF EXISTS options;
DROP TABLE IF EXISTS marks;


CREATE TABLE marks
(
	car_id serial PRIMARY KEY,
	name varchar (32) NOT NULL UNIQUE
);

CREATE TABLE options
(
	car_id serial PRIMARY KEY,
	mark integer REFERENCES Marks(car_id) ON DELETE CASCADE,
	model varchar(32) NOT NULL,
	engine varchar(32) NOT NULL,
	generation varchar(32),
	price float NOT NULL
);

INSERT INTO marks(name)
VALUES
('Toyota'),
('Honda'),
('BMW'),
('Mercedes-Benz'),
('Audi'),
('Ford'),
('Volkswagen'),
('Nissan'),
('Hyundai'),
('Kia'),
('Volvo'),
('Lexus'),
('Mazda'),
('Subaru'),
('Tesla'),
('Porsche'),
('Jeep'),
('Chevrolet'),
('Renault'),
('Peugeot');

INSERT INTO Options (mark, model, engine, generation, price) VALUES
(1, 'Camry', '2.5L 4-cylinder Hybrid', 'XV70', 28500.00),
(1, 'RAV4', '2.5L AWD Hybrid', 'XA50', 32500.00),
(2, 'Civic', '1.5L Turbo VTEC', '11th Gen', 24500.00),
(2, 'CR-V', '1.5L Turbo Hybrid', '6th Gen', 31500.00),
(3, '3 Series', '2.0L Turbo I4', 'G20', 43500.00),
(3, 'X5', '3.0L Turbo I6', 'G05', 62500.00),
(4, 'C-Class', '2.0L Turbo EQ Boost', 'W206', 46500.00),
(4, 'GLC', '2.0L Turbo 4MATIC', 'X254', 49500.00),
(5, 'A4', '2.0L TFSI Quattro', 'B9', 42500.00),
(5, 'Q5', '2.0L TFSI Hybrid', 'FY', 47500.00),
(6, 'Mustang', '5.0L Coyote V8', 'S650', 38500.00),
(6, 'Explorer', '2.3L EcoBoost', 'U625', 38500.00),
(7, 'Golf', '1.5L TSI eHybrid', 'Mk8', 29500.00),
(7, 'Tiguan', '2.0L TSI 4MOTION', '2nd Gen', 34500.00),
(8, 'Qashqai', '1.3L DIG-T Mild Hybrid', 'J12', 28500.00),
(8, 'X-Trail', '1.5L Turbo e-POWER', 'T33', 33500.00),
(9, 'Tucson', '1.6L T-GDI Hybrid', 'NX4', 32500.00),
(9, 'Santa Fe', '2.5L T-GDI', 'TM', 37500.00),
(10, 'Sportage', '1.6L T-GDI Hybrid', 'QL', 31500.00),
(10, 'Sorento', '2.2L CRDi', 'MQ4', 36500.00),
(11, 'XC60', '2.0L B5 Mild Hybrid', '2nd Gen', 45500.00),
(11, 'XC90', '2.0L B6 Mild Hybrid', '2nd Gen', 58500.00),
(12, 'RX', '2.4L Turbo Hybrid', 'AL10', 54500.00),
(12, 'NX', '2.5L Hybrid AWD', 'AL20', 44500.00),
(13, 'CX-5', '2.5L Turbo SkyActiv-G', 'KF', 32500.00),
(13, 'MX-5 Miata', '2.0L SkyActiv-G', 'ND', 28500.00),
(14, 'Outback', '2.5L Boxer', '7th Gen', 33500.00),
(14, 'Forester', '2.5L Boxer e-Boxer', 'SK', 31500.00),
(15, 'Model 3', 'Electric RWD', 'Highland', 42500.00),
(15, 'Model Y', 'Electric AWD', 'Juniper', 49500.00),
(16, 'Cayenne', '3.0L Turbo V6', 'E3', 78500.00),
(16, 'Macan', '2.0L Turbo I4', 'Facelift', 62500.00),
(17, 'Grand Cherokee', '3.6L Pentastar V6', 'WL', 42500.00),
(17, 'Wrangler', '2.0L Turbo eTorque', 'JL', 38500.00),
(18, 'Equinox', '1.5L Turbo', '3rd Gen', 28500.00),
(18, 'Tahoe', '5.3L EcoTec3 V8', '5th Gen', 53500.00),
(19, 'Duster', '1.3L Turbo TCe', '2nd Gen', 22500.00),
(19, 'Arkana', '1.3L Hybrid', '1st Gen', 27500.00),
(20, '3008', '1.2L PureTech', '2nd Gen', 29500.00),
(20, '5008', '1.6L Hybrid', '2nd Gen', 34500.00);

CREATE VIEW view_marks_options AS
SELECT name, model, engine, generation, price
FROM options
JOIN marks ON options.mark = marks.car_id

DROP VIEW view_marks_options

SELECT *
FROM view_marks_options


SELECT options.car_id, name, model, engine, generation, price FROM options JOIN marks ON options.mark = marks.car_id WHERE name='Renault' AND model='Duster'

CREATE DATABASE orders;
CREATE TABLE orders
(
	order_id serial PRIMARY KEY,
	title text NOT NULL,
	price float NOT NULL,
	created_at date NOT NULL
)

SELECT *
FROM view_marks_options
ORDER BY price ASC
LIMIT 3

--данные для создания ДБ-2
CREATE TABLE orders
(
	order_id serial PRIMARY KEY,
	title text NOT NULL,
	price float NOT NULL,
	created_at timestamp NOT NULL
);

SELECT *
FROM orders;

INSERT INTO orders(title, price, created_at)
VALUES ('kzkzkzk', 2313.32, '02.01.2006 15:04:32')

TRUNCATE TABLE orders RESTART IDENTITY

SELECT * 
FROM orders 

DROP TABLE orders

SELECT * 
FROM orders 
WHERE price BETWEEN 44500 AND 50000 
ORDER BY price ASC

SELECT * 
FROM orders 
WHERE created_at BETWEEN '1999-01-01 00:00:00' AND '2025-09-03 20:56:28'
ORDER BY 

SELECT * 
FROM orders 
ORDER BY created_at DESC

SELECT * 
FROM orders 
WHERE price BETWEEN '' AND '' 
ORDER BY price ASC