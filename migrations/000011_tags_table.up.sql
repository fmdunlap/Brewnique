CREATE TABLE IF NOT EXISTS tags (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

-- Flavor Profile
INSERT INTO tags (name) VALUES
('Hoppy'),
('Malty'),
('Bitter'),
('Sweet'),
('Sour'),
('Fruity'),
('Spicy'),
('Roasty'),
('Smoky'),
('Crisp'),
('Dry'),
('Floral'),
('Citrusy'),
('Piney'),
('Caramel'),
('Chocolate'),
('Coffee'),
('Vanilla'),
('Tart'),
('Funky');

-- Strength
INSERT INTO tags (name) VALUES
('Session'),
('Standard Strength'),
('High ABV'),
('Imperial'),
('Double');

-- Color
INSERT INTO tags (name) VALUES
('Light'),
('Amber'),
('Brown'),
('Dark'),
('Black');

-- Body
INSERT INTO tags (name) VALUES
('Light-bodied'),
('Medium-bodied'),
('Full-bodied');

-- Ingredients
INSERT INTO tags (name) VALUES
('Cascade Hops'),
('Citra Hops'),
('Mosaic Hops'),
('Saaz Hops'),
('Pilsner Malt'),
('Munich Malt'),
('Wheat Malt'),
('Oats'),
('Rye'),
('Belgian Yeast'),
('American Yeast'),
('British Yeast'),
('Lager Yeast'),
('Brett'),
('Lactobacillus');

-- Process
INSERT INTO tags (name) VALUES
('Dry-hopped'),
('Barrel-aged'),
('Oak-aged'),
('Bottle-conditioned'),
('Unfiltered'),
('Kettle-soured'),
('Cold-crashed'),
('Decoction mashed');

-- Special Ingredients
INSERT INTO tags (name) VALUES
('Fruit-added'),
('Spice-added'),
('Honey'),
('Coffee-infused'),
('Chocolate-added'),
('Wood-chips'),
('Vanilla beans'),
('Lactose');

-- Seasonal
INSERT INTO tags (name) VALUES
('Summer'),
('Winter'),
('Spring'),
('Fall'),
('Christmas'),
('Oktoberfest');

-- Regional Style
INSERT INTO tags (name) VALUES
('American-style'),
('Belgian-style'),
('German-style'),
('British-style'),
('Czech-style'),
('Irish-style');

-- Clarity
INSERT INTO tags (name) VALUES
('Clear'),
('Hazy'),
('Cloudy');

-- Carbonation
INSERT INTO tags (name) VALUES
('Still'),
('Lightly Carbonated'),
('Highly Carbonated');