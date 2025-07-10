CREATE TABLE equipment_items (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    category VARCHAR(100) NOT NULL,
    subcategory VARCHAR(100) NOT NULL,
    era VARCHAR(50) NOT NULL,
    cost INTEGER NOT NULL,
    rarity INTEGER NOT NULL,
    size INTEGER NOT NULL,
    weight DECIMAL(5,2) NOT NULL,
    damage VARCHAR(50) DEFAULT '',
    armor_rating INTEGER DEFAULT 0,
    ballistic_rating INTEGER DEFAULT 0,
    type VARCHAR(100) DEFAULT '',
    special TEXT DEFAULT '',
    parry BOOLEAN DEFAULT false,
    standard_mods TEXT DEFAULT '',
    description TEXT DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

-- Create indexes for common queries
CREATE INDEX idx_equipment_items_category ON equipment_items(category);
CREATE INDEX idx_equipment_items_subcategory ON equipment_items(subcategory);
CREATE INDEX idx_equipment_items_era ON equipment_items(era);
CREATE INDEX idx_equipment_items_rarity ON equipment_items(rarity);
CREATE INDEX idx_equipment_items_type ON equipment_items(type);
CREATE INDEX idx_equipment_items_cost ON equipment_items(cost);
CREATE INDEX idx_equipment_items_size ON equipment_items(size);

-- Create a composite index for weapon searches
CREATE INDEX idx_equipment_items_weapon_search ON equipment_items(category, subcategory, era, rarity);

-- Create a composite index for armor searches
CREATE INDEX idx_equipment_items_armor_search ON equipment_items(category, subcategory, era, armor_rating); 