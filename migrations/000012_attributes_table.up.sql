CREATE TABLE IF NOT EXISTS attributes (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    type TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS attribute_values (
    id BIGSERIAL PRIMARY KEY,
    attribute_id BIGINT NOT NULL REFERENCES attributes (id),
    value TEXT NOT NULL
);

-- Attributes
INSERT INTO attributes (name, type) VALUES
('Difficulty', 'difficulty'),
('Brew Time', 'brew_time'),
('Equipment Needed', 'equipment'),
('Fermentation Type', 'fermentation'),
('Batch Size', 'batch_size'),
('Clarity', 'clarity'),
('Original Gravity (OG) Range', 'og_range'),
('Final Gravity (FG) Range', 'fg_range'),
('IBU Range', 'ibu_range'),
('ABV Range', 'abv_range'),
('SRM Range', 'srm_range');

-- Attribute Values
-- Difficulty
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'Difficulty'), 'Beginner'),
((SELECT id FROM attributes WHERE name = 'Difficulty'), 'Intermediate'),
((SELECT id FROM attributes WHERE name = 'Difficulty'), 'Advanced');

-- Brew Time
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'Brew Time'), 'Quick (< 2 weeks)'),
((SELECT id FROM attributes WHERE name = 'Brew Time'), 'Standard (2-4 weeks)'),
((SELECT id FROM attributes WHERE name = 'Brew Time'), 'Extended (> 4 weeks)');

-- Equipment Needed
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'Equipment Needed'), 'Basic'),
((SELECT id FROM attributes WHERE name = 'Equipment Needed'), 'Intermediate'),
((SELECT id FROM attributes WHERE name = 'Equipment Needed'), 'Advanced');

-- Fermentation Type
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'Fermentation Type'), 'Ale'),
((SELECT id FROM attributes WHERE name = 'Fermentation Type'), 'Lager'),
((SELECT id FROM attributes WHERE name = 'Fermentation Type'), 'Mixed');

-- Batch Size
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'Batch Size'), 'Small (1-2 gallons)'),
((SELECT id FROM attributes WHERE name = 'Batch Size'), 'Standard (5 gallons)'),
((SELECT id FROM attributes WHERE name = 'Batch Size'), 'Large (10+ gallons)');

-- Clarity
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'Clarity'), 'Clear'),
((SELECT id FROM attributes WHERE name = 'Clarity'), 'Slightly Hazy'),
((SELECT id FROM attributes WHERE name = 'Clarity'), 'Hazy');

-- Original Gravity (OG) Range
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'Original Gravity (OG) Range'), 'Low (1.030-1.039)'),
((SELECT id FROM attributes WHERE name = 'Original Gravity (OG) Range'), 'Medium (1.040-1.059)'),
((SELECT id FROM attributes WHERE name = 'Original Gravity (OG) Range'), 'High (1.060-1.075)'),
((SELECT id FROM attributes WHERE name = 'Original Gravity (OG) Range'), 'Very High (1.076+)');

-- Final Gravity (FG) Range
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'Final Gravity (FG) Range'), 'Very Dry (0.998-1.004)'),
((SELECT id FROM attributes WHERE name = 'Final Gravity (FG) Range'), 'Dry (1.005-1.010)'),
((SELECT id FROM attributes WHERE name = 'Final Gravity (FG) Range'), 'Medium (1.011-1.014)'),
((SELECT id FROM attributes WHERE name = 'Final Gravity (FG) Range'), 'Sweet (1.015-1.020)'),
((SELECT id FROM attributes WHERE name = 'Final Gravity (FG) Range'), 'Very Sweet (1.021+)');

-- IBU Range
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'IBU Range'), 'Low (0-20 IBU)'),
((SELECT id FROM attributes WHERE name = 'IBU Range'), 'Medium (21-40 IBU)'),
((SELECT id FROM attributes WHERE name = 'IBU Range'), 'High (41-60 IBU)'),
((SELECT id FROM attributes WHERE name = 'IBU Range'), 'Very High (61+ IBU)');

-- ABV Range
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'ABV Range'), 'Session (< 4% ABV)'),
((SELECT id FROM attributes WHERE name = 'ABV Range'), 'Standard (4-6% ABV)'),
((SELECT id FROM attributes WHERE name = 'ABV Range'), 'Strong (6-9% ABV)'),
((SELECT id FROM attributes WHERE name = 'ABV Range'), 'Very Strong (> 9% ABV)');

-- SRM Range (Color)
INSERT INTO attribute_values (attribute_id, value) VALUES
((SELECT id FROM attributes WHERE name = 'SRM Range'), 'Very Light (1-3 SRM)'),
((SELECT id FROM attributes WHERE name = 'SRM Range'), 'Light (4-6 SRM)'),
((SELECT id FROM attributes WHERE name = 'SRM Range'), 'Gold (7-10 SRM)'),
((SELECT id FROM attributes WHERE name = 'SRM Range'), 'Amber (11-18 SRM)'),
((SELECT id FROM attributes WHERE name = 'SRM Range'), 'Copper/Brown (19-35 SRM)'),
((SELECT id FROM attributes WHERE name = 'SRM Range'), 'Dark (36+ SRM)');