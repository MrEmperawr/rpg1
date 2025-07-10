-- Update equipment_items table to use nullable integer damage and damage type
ALTER TABLE equipment_items 
ADD COLUMN damage_type VARCHAR(10) DEFAULT NULL CHECK (damage_type IS NULL OR damage_type IN ('S', 'L', 'A'));

-- Update existing damage values to use the new structure
-- This will be handled by the Go code when reseeding
-- For now, set all existing damage to NULL and damage_type to NULL
UPDATE equipment_items SET damage = NULL, damage_type = NULL WHERE damage IS NULL OR damage = '';

-- Change damage column from VARCHAR to INTEGER (nullable)
ALTER TABLE equipment_items ALTER COLUMN damage TYPE INTEGER USING 
  CASE 
    WHEN damage = '' OR damage IS NULL THEN NULL
    WHEN damage ~ '^[+-]?\d+[SLA]?$' THEN 
      CASE 
        WHEN damage ~ '^[+-]?\d+$' THEN CAST(REPLACE(REPLACE(damage, '+', ''), '-', '-') AS INTEGER)
        ELSE NULL
      END
    ELSE NULL
  END; 