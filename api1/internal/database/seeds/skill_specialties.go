package seeds

import (
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
)

// GetSkillSpecialties returns comprehensive skill specialties from the RPG documentation
// Note: SkillID will be set during seeding based on skill names
func GetSkillSpecialties() []srd.SkillSpecialty {
	return []srd.SkillSpecialty{
		// Athletics Specialties
		{
			Name:        "Climbing",
			Type:        "typical",
			Description: "+2 bonus when climbing walls, trees, or other vertical surfaces",
			Bonus:       2,
		},
		{
			Name:        "Swimming",
			Type:        "typical",
			Description: "+2 bonus when swimming or moving through water",
			Bonus:       2,
		},
		{
			Name:        "Gymnastics",
			Type:        "typical",
			Description: "+2 bonus when performing acrobatic maneuvers, tumbling, or balancing",
			Bonus:       2,
		},
		{
			Name:        "Throwing",
			Type:        "typical",
			Description: "+2 bonus when throwing objects, weapons, or projectiles",
			Bonus:       2,
		},
		{
			Name:        "Dodge",
			Type:        "typical",
			Description: "+2 bonus when dodging attacks or avoiding danger",
			Bonus:       2,
		},
		{
			Name:        "Exotic Weapons",
			Type:        "exotic",
			Description: "Allows use of Athletics skill with weapons marked as 'Exotic'",
			Bonus:       0,
		},

		// Melee Specialties
		{
			Name:        "Swords",
			Type:        "typical",
			Description: "+2 bonus when using swords of any type",
			Bonus:       2,
		},
		{
			Name:        "Axes",
			Type:        "typical",
			Description: "+2 bonus when using axes and hatchets",
			Bonus:       2,
		},
		{
			Name:        "Polearms",
			Type:        "typical",
			Description: "+2 bonus when using spears, halberds, and other pole weapons",
			Bonus:       2,
		},
		{
			Name:        "Maces",
			Type:        "typical",
			Description: "+2 bonus when using maces, hammers, and blunt weapons",
			Bonus:       2,
		},
		{
			Name:        "Exotic Weapons",
			Type:        "exotic",
			Description: "Allows use of Melee skill with weapons marked as 'Exotic'",
			Bonus:       0,
		},

		// Pilot Specialties
		{
			Name:        "Ground Vehicles",
			Type:        "typical",
			Description: "+2 bonus when driving cars, carriages, and ground transport",
			Bonus:       2,
		},
		{
			Name:        "Riding",
			Type:        "typical",
			Description: "+2 bonus when riding horses and other mounts",
			Bonus:       2,
		},
		{
			Name:        "Watercraft",
			Type:        "typical",
			Description: "+2 bonus when piloting boats, ships, and water vessels",
			Bonus:       2,
		},
		{
			Name:        "Jet",
			Type:        "exotic",
			Description: "Allows use of Pilot skill with jet aircraft",
			Bonus:       0,
		},
		{
			Name:        "Rotorcraft",
			Type:        "exotic",
			Description: "Allows use of Pilot skill with helicopters and rotorcraft",
			Bonus:       0,
		},
		{
			Name:        "Sailing",
			Type:        "exotic",
			Description: "Allows use of Pilot skill with sailing vessels",
			Bonus:       0,
		},

		// Ranged Specialties
		{
			Name:        "Bows",
			Type:        "typical",
			Description: "+2 bonus when using bows and crossbows",
			Bonus:       2,
		},
		{
			Name:        "Firearms",
			Type:        "typical",
			Description: "+2 bonus when using guns and firearms",
			Bonus:       2,
		},
		{
			Name:        "Thrown",
			Type:        "typical",
			Description: "+2 bonus when throwing weapons like daggers, spears, or javelins",
			Bonus:       2,
		},
		{
			Name:        "Exotic Weapons",
			Type:        "exotic",
			Description: "Allows use of Ranged skill with weapons marked as 'Exotic'",
			Bonus:       0,
		},

		// Stealth Specialties
		{
			Name:        "Shadowing",
			Type:        "typical",
			Description: "+2 bonus when following someone without being detected",
			Bonus:       2,
		},
		{
			Name:        "Urban Stealth",
			Type:        "typical",
			Description: "+2 bonus when hiding in urban environments",
			Bonus:       2,
		},
		{
			Name:        "Wilderness Stealth",
			Type:        "typical",
			Description: "+2 bonus when hiding in natural environments",
			Bonus:       2,
		},

		// Thievery Specialties
		{
			Name:        "Lockpicking",
			Type:        "typical",
			Description: "+2 bonus when picking locks and bypassing mechanical security",
			Bonus:       2,
		},
		{
			Name:        "Pickpocketing",
			Type:        "typical",
			Description: "+2 bonus when stealing items from people without detection",
			Bonus:       2,
		},
		{
			Name:        "Security Systems",
			Type:        "typical",
			Description: "+2 bonus when disabling alarms and electronic security",
			Bonus:       2,
		},

		// Unarmed Specialties
		{
			Name:        "Grappling",
			Type:        "typical",
			Description: "+2 bonus when wrestling, grappling, or restraining opponents",
			Bonus:       2,
		},
		{
			Name:        "Striking",
			Type:        "typical",
			Description: "+2 bonus when punching, kicking, or striking with hands/feet",
			Bonus:       2,
		},
		{
			Name:        "Exotic Weapons",
			Type:        "exotic",
			Description: "Allows use of Unarmed skill with weapons marked as 'Exotic'",
			Bonus:       0,
		},

		// Academics Specialties
		{
			Name:        "Literature",
			Type:        "typical",
			Description: "+2 bonus when dealing with books, texts, and literary works",
			Bonus:       2,
		},
		{
			Name:        "History",
			Type:        "typical",
			Description: "+2 bonus when recalling historical events and knowledge",
			Bonus:       2,
		},
		{
			Name:        "Law",
			Type:        "typical",
			Description: "+2 bonus when dealing with legal matters and contracts",
			Bonus:       2,
		},
		{
			Name:        "Linguistics",
			Type:        "typical",
			Description: "+2 bonus when translating languages or analyzing text",
			Bonus:       2,
		},
		{
			Name:        "Research",
			Type:        "typical",
			Description: "+2 bonus when conducting academic research and investigation",
			Bonus:       2,
		},
		{
			Name:        "Bureaucracy",
			Type:        "typical",
			Description: "+2 bonus when navigating complex organizational systems",
			Bonus:       2,
		},
		{
			Name:        "Local Politics",
			Type:        "typical",
			Description: "+2 bonus when dealing with local government and politics",
			Bonus:       2,
		},
		{
			Name:        "National Politics",
			Type:        "typical",
			Description: "+2 bonus when dealing with national government and politics",
			Bonus:       2,
		},
		{
			Name:        "Scandals",
			Type:        "typical",
			Description: "+2 bonus when uncovering or dealing with political scandals",
			Bonus:       2,
		},
		{
			Name:        "Noble Houses",
			Type:        "typical",
			Description: "+2 bonus when dealing with noble families and aristocracy",
			Bonus:       2,
		},

		// Crafts Specialties
		{
			Name:        "Automotive",
			Type:        "typical",
			Description: "+2 bonus when working with vehicles and automotive systems",
			Bonus:       2,
		},
		{
			Name:        "Carpentry",
			Type:        "typical",
			Description: "+2 bonus when working with wood and wooden structures",
			Bonus:       2,
		},
		{
			Name:        "Jury Rigging",
			Type:        "typical",
			Description: "+2 bonus when making quick repairs with improvised materials",
			Bonus:       2,
		},
		{
			Name:        "Sculpting",
			Type:        "typical",
			Description: "+2 bonus when creating sculptures and artistic works",
			Bonus:       2,
		},
		{
			Name:        "Welding",
			Type:        "typical",
			Description: "+2 bonus when joining metals and working with heat",
			Bonus:       2,
		},
		{
			Name:        "Smithing",
			Type:        "typical",
			Description: "+2 bonus when forging weapons, armor, and metal objects",
			Bonus:       2,
		},
		{
			Name:        "Fletching",
			Type:        "typical",
			Description: "+2 bonus when making arrows and other projectile ammunition",
			Bonus:       2,
		},
		{
			Name:        "Cyberware",
			Type:        "exotic",
			Description: "Allows use of Crafts skill with cybernetic implants and technology",
			Bonus:       0,
		},

		// Investigation Specialties
		{
			Name:        "Crime Scenes",
			Type:        "typical",
			Description: "+2 bonus when investigating crime scenes and evidence",
			Bonus:       2,
		},
		{
			Name:        "Cryptography",
			Type:        "typical",
			Description: "+2 bonus when decoding messages and breaking codes",
			Bonus:       2,
		},
		{
			Name:        "Dreams",
			Type:        "typical",
			Description: "+2 bonus when analyzing dreams and subconscious patterns",
			Bonus:       2,
		},
		{
			Name:        "Forensic Accounting",
			Type:        "typical",
			Description: "+2 bonus when investigating financial records and fraud",
			Bonus:       2,
		},
		{
			Name:        "Riddles",
			Type:        "typical",
			Description: "+2 bonus when solving puzzles, riddles, and complex problems",
			Bonus:       2,
		},

		// Medicine Specialties
		{
			Name:        "Cardiology",
			Type:        "typical",
			Description: "+2 bonus when treating heart conditions and cardiovascular issues",
			Bonus:       2,
		},
		{
			Name:        "First Aid",
			Type:        "typical",
			Description: "+2 bonus when providing emergency medical treatment",
			Bonus:       2,
		},
		{
			Name:        "Pathology",
			Type:        "typical",
			Description: "+2 bonus when analyzing diseases and medical conditions",
			Bonus:       2,
		},
		{
			Name:        "Pharmacology",
			Type:        "typical",
			Description: "+2 bonus when working with drugs, medicines, and chemicals",
			Bonus:       2,
		},
		{
			Name:        "Surgery",
			Type:        "exotic",
			Description: "Allows use of Medicine skill for surgical procedures",
			Bonus:       0,
		},

		// Nature Lore Specialties
		{
			Name:        "Dogs",
			Type:        "typical",
			Description: "+2 bonus when working with dogs and canine animals",
			Bonus:       2,
		},
		{
			Name:        "Exotic Pets",
			Type:        "typical",
			Description: "+2 bonus when working with unusual or exotic animals",
			Bonus:       2,
		},
		{
			Name:        "Horses",
			Type:        "typical",
			Description: "+2 bonus when working with horses and equine animals",
			Bonus:       2,
		},
		{
			Name:        "Training",
			Type:        "typical",
			Description: "+2 bonus when training animals and teaching them commands",
			Bonus:       2,
		},
		{
			Name:        "Wild Animals",
			Type:        "typical",
			Description: "+2 bonus when dealing with wild and untamed animals",
			Bonus:       2,
		},
		{
			Name:        "Tracking",
			Type:        "typical",
			Description: "+2 bonus when following tracks and trails",
			Bonus:       2,
		},
		{
			Name:        "Hunting",
			Type:        "typical",
			Description: "+2 bonus when hunting animals for food or sport",
			Bonus:       2,
		},

		// Occult Specialties
		{
			Name:        "Ritual Magic",
			Type:        "typical",
			Description: "+2 bonus when performing magical rituals and ceremonies",
			Bonus:       2,
		},
		{
			Name:        "Spirit (type)",
			Type:        "typical",
			Description: "+2 bonus when dealing with specific types of spirits",
			Bonus:       2,
		},
		{
			Name:        "Offerings",
			Type:        "typical",
			Description: "+2 bonus when making offerings to spirits and deities",
			Bonus:       2,
		},
		{
			Name:        "Magical Theory",
			Type:        "typical",
			Description: "+2 bonus when understanding magical principles and theory",
			Bonus:       2,
		},
		{
			Name:        "Ghosts",
			Type:        "typical",
			Description: "+2 bonus when dealing with ghosts and spectral entities",
			Bonus:       2,
		},
		{
			Name:        "Magical Creatures",
			Type:        "typical",
			Description: "+2 bonus when dealing with supernatural and magical beasts",
			Bonus:       2,
		},
		{
			Name:        "Occult Curses",
			Type:        "typical",
			Description: "+2 bonus when dealing with curses and hexes",
			Bonus:       2,
		},
		{
			Name:        "Religion",
			Type:        "typical",
			Description: "+2 bonus when dealing with religious practices and beliefs",
			Bonus:       2,
		},
		{
			Name:        "Myth",
			Type:        "typical",
			Description: "+2 bonus when dealing with myths, legends, and folklore",
			Bonus:       2,
		},

		// Science Specialties
		{
			Name:        "Biology",
			Type:        "typical",
			Description: "+2 bonus when working with living organisms and biological systems",
			Bonus:       2,
		},
		{
			Name:        "Chemistry",
			Type:        "typical",
			Description: "+2 bonus when working with chemicals and chemical reactions",
			Bonus:       2,
		},
		{
			Name:        "Genetics",
			Type:        "typical",
			Description: "+2 bonus when working with genetic material and heredity",
			Bonus:       2,
		},
		{
			Name:        "Optics",
			Type:        "typical",
			Description: "+2 bonus when working with light, lenses, and optical systems",
			Bonus:       2,
		},
		{
			Name:        "Particle Physics",
			Type:        "typical",
			Description: "+2 bonus when working with subatomic particles and physics",
			Bonus:       2,
		},
		{
			Name:        "Data Retrieval",
			Type:        "typical",
			Description: "+2 bonus when accessing and retrieving digital information",
			Bonus:       2,
		},
		{
			Name:        "Digital Security",
			Type:        "typical",
			Description: "+2 bonus when working with computer security and encryption",
			Bonus:       2,
		},
		{
			Name:        "Programming",
			Type:        "typical",
			Description: "+2 bonus when writing and debugging computer code",
			Bonus:       2,
		},
		{
			Name:        "User Interface Design",
			Type:        "typical",
			Description: "+2 bonus when designing computer interfaces and user experiences",
			Bonus:       2,
		},
		{
			Name:        "Hacking",
			Type:        "exotic",
			Description: "Allows use of Science skill for computer hacking and intrusion",
			Bonus:       0,
		},
		{
			Name:        "EOD",
			Type:        "exotic",
			Description: "Allows use of Science skill for Explosive Ordnance Disposal",
			Bonus:       0,
		},
		{
			Name:        "Bioware",
			Type:        "exotic",
			Description: "Allows use of Science skill with biological implants and technology",
			Bonus:       0,
		},

		// Empathy Specialties
		{
			Name:        "Buried Feelings",
			Type:        "typical",
			Description: "+2 bonus when detecting hidden or repressed emotions",
			Bonus:       2,
		},
		{
			Name:        "Calming",
			Type:        "typical",
			Description: "+2 bonus when soothing and calming distressed individuals",
			Bonus:       2,
		},
		{
			Name:        "Emotions",
			Type:        "typical",
			Description: "+2 bonus when reading and understanding emotional states",
			Bonus:       2,
		},
		{
			Name:        "Lies",
			Type:        "typical",
			Description: "+2 bonus when detecting deception and falsehoods",
			Bonus:       2,
		},
		{
			Name:        "Motives",
			Type:        "typical",
			Description: "+2 bonus when understanding others' motivations and intentions",
			Bonus:       2,
		},

		// Performance Specialties
		{
			Name:        "Dance",
			Type:        "typical",
			Description: "+2 bonus when performing dance and movement arts",
			Bonus:       2,
		},
		{
			Name:        "Journalism",
			Type:        "typical",
			Description: "+2 bonus when writing and reporting news and stories",
			Bonus:       2,
		},
		{
			Name:        "Music Composition",
			Type:        "typical",
			Description: "+2 bonus when composing and arranging musical pieces",
			Bonus:       2,
		},
		{
			Name:        "Painting",
			Type:        "typical",
			Description: "+2 bonus when creating visual art and paintings",
			Bonus:       2,
		},
		{
			Name:        "Speeches",
			Type:        "typical",
			Description: "+2 bonus when delivering public speeches and orations",
			Bonus:       2,
		},
		{
			Name:        "String Instruments",
			Type:        "exotic",
			Description: "Allows use of Performance skill with string instruments",
			Bonus:       0,
		},
		{
			Name:        "Key Instruments",
			Type:        "exotic",
			Description: "Allows use of Performance skill with keyboard instruments",
			Bonus:       0,
		},
		{
			Name:        "Blow Instruments",
			Type:        "exotic",
			Description: "Allows use of Performance skill with wind instruments",
			Bonus:       0,
		},
		{
			Name:        "Other Instruments",
			Type:        "exotic",
			Description: "Allows use of Performance skill with other instrument types",
			Bonus:       0,
		},

		// Intimidation Specialties
		{
			Name:        "Direct Threats",
			Type:        "typical",
			Description: "+2 bonus when making direct and explicit threats",
			Bonus:       2,
		},
		{
			Name:        "Interrogation",
			Type:        "typical",
			Description: "+2 bonus when extracting information through intimidation",
			Bonus:       2,
		},
		{
			Name:        "Murderous Stare",
			Type:        "typical",
			Description: "+2 bonus when intimidating through silent, threatening presence",
			Bonus:       2,
		},
		{
			Name:        "Torture",
			Type:        "typical",
			Description: "+2 bonus when using physical intimidation and torture",
			Bonus:       2,
		},
		{
			Name:        "Veiled Threats",
			Type:        "typical",
			Description: "+2 bonus when making subtle or implied threats",
			Bonus:       2,
		},

		// Persuasion Specialties
		{
			Name:        "Fast Talking",
			Type:        "typical",
			Description: "+2 bonus when using quick, persuasive speech",
			Bonus:       2,
		},
		{
			Name:        "Inspiring",
			Type:        "typical",
			Description: "+2 bonus when motivating and inspiring others",
			Bonus:       2,
		},
		{
			Name:        "Sales Pitches",
			Type:        "typical",
			Description: "+2 bonus when selling products or services",
			Bonus:       2,
		},
		{
			Name:        "Seduction",
			Type:        "typical",
			Description: "+2 bonus when using charm and attraction to persuade",
			Bonus:       2,
		},
		{
			Name:        "Sermons",
			Type:        "typical",
			Description: "+2 bonus when delivering religious or inspirational speeches",
			Bonus:       2,
		},

		// Streetwise Specialties
		{
			Name:        "Black Market",
			Type:        "typical",
			Description: "+2 bonus when dealing with illegal goods and services",
			Bonus:       2,
		},
		{
			Name:        "Gangs",
			Type:        "typical",
			Description: "+2 bonus when dealing with street gangs and criminal organizations",
			Bonus:       2,
		},
		{
			Name:        "Navigation",
			Type:        "typical",
			Description: "+2 bonus when finding shortcuts and navigating urban areas",
			Bonus:       2,
		},
		{
			Name:        "Rumors",
			Type:        "typical",
			Description: "+2 bonus when gathering and spreading street rumors",
			Bonus:       2,
		},
		{
			Name:        "Undercover Work",
			Type:        "typical",
			Description: "+2 bonus when infiltrating criminal organizations",
			Bonus:       2,
		},

		// Style Specialties
		{
			Name:        "Clothing",
			Type:        "typical",
			Description: "+2 bonus when choosing and wearing appropriate clothing",
			Bonus:       2,
		},
		{
			Name:        "High Society",
			Type:        "typical",
			Description: "+2 bonus when fitting in with upper-class social circles",
			Bonus:       2,
		},
		{
			Name:        "Street",
			Type:        "typical",
			Description: "+2 bonus when fitting in with street-level social circles",
			Bonus:       2,
		},
		{
			Name:        "Stylish Recovery",
			Type:        "typical",
			Description: "+2 bonus when recovering gracefully from mistakes or accidents",
			Bonus:       2,
		},
		{
			Name:        "Making an Entrance",
			Type:        "typical",
			Description: "+2 bonus when creating dramatic and memorable entrances",
			Bonus:       2,
		},
		{
			Name:        "Gossip",
			Type:        "typical",
			Description: "+2 bonus when gathering and spreading social information",
			Bonus:       2,
		},

		// Subterfuge Specialties
		{
			Name:        "Detecting Lies",
			Type:        "typical",
			Description: "+2 bonus when spotting deception and falsehoods",
			Bonus:       2,
		},
		{
			Name:        "Hidden Meanings",
			Type:        "typical",
			Description: "+2 bonus when understanding subtle messages and implications",
			Bonus:       2,
		},
		{
			Name:        "Hiding Emotions",
			Type:        "typical",
			Description: "+2 bonus when concealing true feelings and reactions",
			Bonus:       2,
		},
		{
			Name:        "Long Cons",
			Type:        "typical",
			Description: "+2 bonus when executing complex, long-term deceptions",
			Bonus:       2,
		},
		{
			Name:        "Misdirection",
			Type:        "typical",
			Description: "+2 bonus when distracting and redirecting attention",
			Bonus:       2,
		},

		// Air Magic Specialties
		{
			Name:        "Electricity",
			Type:        "typical",
			Description: "+2 bonus when casting Air spells related to electricity",
			Bonus:       2,
		},
		{
			Name:        "Speed",
			Type:        "typical",
			Description: "+2 bonus when casting Air spells related to speed and movement",
			Bonus:       2,
		},
		{
			Name:        "Telekinesis",
			Type:        "typical",
			Description: "+2 bonus when casting Air spells related to telekinesis",
			Bonus:       2,
		},
		{
			Name:        "Thunder",
			Type:        "typical",
			Description: "+2 bonus when casting Air spells related to thunder and sound",
			Bonus:       2,
		},
		{
			Name:        "Wind",
			Type:        "typical",
			Description: "+2 bonus when casting Air spells related to wind and air",
			Bonus:       2,
		},

		// Dark Magic Specialties
		{
			Name:        "Cold",
			Type:        "typical",
			Description: "+2 bonus when casting Dark spells related to cold and ice",
			Bonus:       2,
		},
		{
			Name:        "Curses",
			Type:        "typical",
			Description: "+2 bonus when casting Dark spells related to curses",
			Bonus:       2,
		},
		{
			Name:        "Darkness",
			Type:        "typical",
			Description: "+2 bonus when casting Dark spells related to shadows and darkness",
			Bonus:       2,
		},
		{
			Name:        "Necromancy",
			Type:        "typical",
			Description: "+2 bonus when casting Dark spells related to death and undeath",
			Bonus:       2,
		},
		{
			Name:        "Negative Emotions",
			Type:        "typical",
			Description: "+2 bonus when casting Dark spells related to negative emotions",
			Bonus:       2,
		},

		// Earth Magic Specialties
		{
			Name:        "Alter Materials",
			Type:        "typical",
			Description: "+2 bonus when casting Earth spells related to material transformation",
			Bonus:       2,
		},
		{
			Name:        "Shaping",
			Type:        "typical",
			Description: "+2 bonus when casting Earth spells related to shaping earth and stone",
			Bonus:       2,
		},
		{
			Name:        "Stoicism",
			Type:        "typical",
			Description: "+2 bonus when casting Earth spells related to endurance and resilience",
			Bonus:       2,
		},

		// Light Magic Specialties
		{
			Name:        "Fire",
			Type:        "typical",
			Description: "+2 bonus when casting Light spells related to fire and heat",
			Bonus:       2,
		},
		{
			Name:        "Healing",
			Type:        "typical",
			Description: "+2 bonus when casting Light spells related to healing and restoration",
			Bonus:       2,
		},
		{
			Name:        "Light",
			Type:        "typical",
			Description: "+2 bonus when casting Light spells related to illumination",
			Bonus:       2,
		},
		{
			Name:        "Purity",
			Type:        "typical",
			Description: "+2 bonus when casting Light spells related to purification",
			Bonus:       2,
		},
		{
			Name:        "Revelation",
			Type:        "typical",
			Description: "+2 bonus when casting Light spells related to truth and revelation",
			Bonus:       2,
		},

		// Nature Magic Specialties
		{
			Name:        "Animals",
			Type:        "typical",
			Description: "+2 bonus when casting Nature spells related to animals",
			Bonus:       2,
		},
		{
			Name:        "Disease",
			Type:        "typical",
			Description: "+2 bonus when casting Nature spells related to disease and illness",
			Bonus:       2,
		},
		{
			Name:        "Instincts",
			Type:        "typical",
			Description: "+2 bonus when casting Nature spells related to instincts and intuition",
			Bonus:       2,
		},
		{
			Name:        "Plants",
			Type:        "typical",
			Description: "+2 bonus when casting Nature spells related to plants and vegetation",
			Bonus:       2,
		},
		{
			Name:        "Poison",
			Type:        "typical",
			Description: "+2 bonus when casting Nature spells related to poison and toxins",
			Bonus:       2,
		},
		{
			Name:        "Shapechange",
			Type:        "typical",
			Description: "+2 bonus when casting Nature spells related to transformation",
			Bonus:       2,
		},

		// Summoning Magic Specialties
		{
			Name:        "Affecting Magic",
			Type:        "typical",
			Description: "+2 bonus when casting Summoning spells that affect other magic",
			Bonus:       2,
		},
		{
			Name:        "Astral",
			Type:        "typical",
			Description: "+2 bonus when casting Summoning spells related to astral realms",
			Bonus:       2,
		},
		{
			Name:        "Banish",
			Type:        "typical",
			Description: "+2 bonus when casting Summoning spells to banish entities",
			Bonus:       2,
		},
		{
			Name:        "Bind",
			Type:        "typical",
			Description: "+2 bonus when casting Summoning spells to bind entities",
			Bonus:       2,
		},
		{
			Name:        "Exorcise",
			Type:        "typical",
			Description: "+2 bonus when casting Summoning spells to exorcise spirits",
			Bonus:       2,
		},
		{
			Name:        "Possession",
			Type:        "typical",
			Description: "+2 bonus when casting Summoning spells related to possession",
			Bonus:       2,
		},
		{
			Name:        "Summoning",
			Type:        "typical",
			Description: "+2 bonus when casting Summoning spells to call forth entities",
			Bonus:       2,
		},
	}
}
