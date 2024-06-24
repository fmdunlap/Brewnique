ALTER TABLE recipes ADD COLUMN category_id BIGINT REFERENCES categories (id);
ALTER TABLE recipes ADD COLUMN subcategory_id BIGINT REFERENCES categories (id);

UPDATE recipes SET category_id = (SELECT id FROM categories WHERE name = 'Ales');
UPDATE recipes SET subcategory_id = (SELECT id FROM categories WHERE name = 'American Ale');