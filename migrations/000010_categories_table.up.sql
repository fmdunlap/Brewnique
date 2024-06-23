CREATE TABLE IF NOT EXISTS categories (
id BIGSERIAL PRIMARY KEY,
name TEXT NOT NULL,
parent_id BIGINT REFERENCES categories (id)
);

-- Insert main categories with specific IDs
INSERT INTO categories (id, name, parent_id) OVERRIDING SYSTEM VALUE VALUES
(1, 'Ales', NULL),
(2, 'Lagers', NULL),
(3, 'Hybrid Beers', NULL),
(4, 'Mead', NULL),
(5, 'Cider', NULL),
(6, 'Specialty/Experimental', NULL);

-- Reset the sequence
SELECT setval('categories_id_seq', (SELECT MAX(id) FROM categories));

-- Subcategories
-- Ales
INSERT INTO categories (name, parent_id) VALUES
('American Ale', 1),
('Belgian Ale', 1),
('British Ale', 1),
('German Ale', 1),
('IPA', 1),
('Wheat Beer', 1),
('Sour Ale', 1),
('Saison', 1),
('Porter', 1),
('Stout', 1);

-- Lagers
INSERT INTO categories (name, parent_id) VALUES
('Pilsner', 2),
('Bock', 2),
('Dunkel', 2),
('Märzen/Oktoberfest', 2),
('Vienna Lager', 2),
('Helles', 2),
('Schwarzbier', 2);

-- Hybrid Beers
INSERT INTO categories (name, parent_id) VALUES
('Kölsch', 3),
('Altbier', 3),
('California Common/Steam Beer', 3),
('Cream Ale', 3);

-- Mead
INSERT INTO categories (name, parent_id) VALUES
('Traditional Mead', 4),
('Melomel', 4),
('Metheglin', 4),
('Braggot', 4);

-- Cider
INSERT INTO categories (name, parent_id) VALUES
('Standard Cider', 5),
('New England Cider', 5),
('French Cider', 5),
('Spanish Cider', 5);

-- Specialty/Experimental
INSERT INTO categories (name, parent_id) VALUES
('Fruit Beer', 6),
('Spice/Herb/Vegetable Beer', 6),
('Smoke-Flavored Beer', 6),
('Wood-Aged Beer', 6),
('Gluten-Free Beer', 6);