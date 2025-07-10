-- Revert equipment_items table changes
ALTER TABLE equipment_items DROP COLUMN damage_type;
ALTER TABLE equipment_items ALTER COLUMN damage TYPE VARCHAR(255); 