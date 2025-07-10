package seeds

import (
	"github.com/mremperor-atwork/rpg1/api1/internal/features/srd"
)

// GetVirtues returns comprehensive virtues from the RPG documentation
func GetVirtues() []srd.VirtueVice {
	return []srd.VirtueVice{
		{
			Name:        "Charitable",
			Type:        "Virtue",
			Description: "You are generous and giving, always willing to help those in need.",
			Effect:      "Regain 1 Willpower when making a sacrifice for charity. Regain all Willpower when putting yourself in mortal danger for charity.",
		},
		{
			Name:        "Kind",
			Type:        "Virtue",
			Description: "You are gentle and compassionate, treating others with warmth and understanding.",
			Effect:      "Regain 1 Willpower when making a sacrifice for kindness. Regain all Willpower when putting yourself in mortal danger for kindness.",
		},
		{
			Name:        "Generous",
			Type:        "Virtue",
			Description: "You are giving and open-handed, sharing what you have with others.",
			Effect:      "Regain 1 Willpower when making a sacrifice for generosity. Regain all Willpower when putting yourself in mortal danger for generosity.",
		},
		{
			Name:        "Helpful",
			Type:        "Virtue",
			Description: "You are always ready to assist others, going out of your way to be useful.",
			Effect:      "Regain 1 Willpower when making a sacrifice to help others. Regain all Willpower when putting yourself in mortal danger to help others.",
		},
		{
			Name:        "Loyal",
			Type:        "Virtue",
			Description: "You are faithful and devoted to your friends, family, and causes.",
			Effect:      "Regain 1 Willpower when making a sacrifice for loyalty. Regain all Willpower when putting yourself in mortal danger for loyalty.",
		},
		{
			Name:        "Faithful",
			Type:        "Virtue",
			Description: "You are deeply religious or spiritual, guided by strong beliefs.",
			Effect:      "Regain 1 Willpower when making a sacrifice for your faith. Regain all Willpower when putting yourself in mortal danger for your faith.",
		},
		{
			Name:        "Protective",
			Type:        "Virtue",
			Description: "You are a guardian, always ready to defend those who cannot defend themselves.",
			Effect:      "Regain 1 Willpower when making a sacrifice to protect others. Regain all Willpower when putting yourself in mortal danger to protect others.",
		},
		{
			Name:        "Honest",
			Type:        "Virtue",
			Description: "You are truthful and sincere, always speaking the truth even when it's difficult.",
			Effect:      "Regain 1 Willpower when making a sacrifice for honesty. Regain all Willpower when putting yourself in mortal danger for honesty.",
		},
		{
			Name:        "Just",
			Type:        "Virtue",
			Description: "You believe in fairness and justice, always seeking to do what is right.",
			Effect:      "Regain 1 Willpower when making a sacrifice for justice. Regain all Willpower when putting yourself in mortal danger for justice.",
		},
		{
			Name:        "Courageous",
			Type:        "Virtue",
			Description: "You are brave and fearless, facing danger with determination.",
			Effect:      "Regain 1 Willpower when making a sacrifice for courage. Regain all Willpower when putting yourself in mortal danger for courage.",
		},
		{
			Name:        "Wise",
			Type:        "Virtue",
			Description: "You are thoughtful and insightful, making decisions based on deep understanding.",
			Effect:      "Regain 1 Willpower when making a sacrifice for wisdom. Regain all Willpower when putting yourself in mortal danger for wisdom.",
		},
		{
			Name:        "Patient",
			Type:        "Virtue",
			Description: "You are calm and enduring, able to wait for the right moment.",
			Effect:      "Regain 1 Willpower when making a sacrifice for patience. Regain all Willpower when putting yourself in mortal danger for patience.",
		},
		{
			Name:        "Humble",
			Type:        "Virtue",
			Description: "You are modest and unpretentious, not seeking recognition for your deeds.",
			Effect:      "Regain 1 Willpower when making a sacrifice for humility. Regain all Willpower when putting yourself in mortal danger for humility.",
		},
		{
			Name:        "Temperate",
			Type:        "Virtue",
			Description: "You are moderate and self-controlled, avoiding excess in all things.",
			Effect:      "Regain 1 Willpower when making a sacrifice for temperance. Regain all Willpower when putting yourself in mortal danger for temperance.",
		},
		{
			Name:        "Diligent",
			Type:        "Virtue",
			Description: "You are hardworking and persistent, never giving up on your goals.",
			Effect:      "Regain 1 Willpower when making a sacrifice for diligence. Regain all Willpower when putting yourself in mortal danger for diligence.",
		},
	}
}

// GetVices returns comprehensive vices from the RPG documentation
func GetVices() []srd.VirtueVice {
	return []srd.VirtueVice{
		{
			Name:        "Wrathful",
			Type:        "Vice",
			Description: "You are quick to anger and violence, often losing control when provoked.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to wrath. Regain all Willpower when your wrath causes potentially lethal harm.",
		},
		{
			Name:        "Gluttonous",
			Type:        "Vice",
			Description: "You are excessive in your consumption, whether of food, drink, or other pleasures.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to gluttony. Regain all Willpower when your gluttony causes potentially lethal harm.",
		},
		{
			Name:        "Deceptive",
			Type:        "Vice",
			Description: "You are dishonest and manipulative, often lying to get what you want.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to deception. Regain all Willpower when your deception causes potentially lethal harm.",
		},
		{
			Name:        "Greedy",
			Type:        "Vice",
			Description: "You are obsessed with wealth and possessions, always wanting more.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to greed. Regain all Willpower when your greed causes potentially lethal harm.",
		},
		{
			Name:        "Selfish",
			Type:        "Vice",
			Description: "You are concerned only with your own interests, often at the expense of others.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to selfishness. Regain all Willpower when your selfishness causes potentially lethal harm.",
		},
		{
			Name:        "Lustful",
			Type:        "Vice",
			Description: "You are driven by physical desires and pleasures, often to excess.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to lust. Regain all Willpower when your lust causes potentially lethal harm.",
		},
		{
			Name:        "Envious",
			Type:        "Vice",
			Description: "You are jealous of what others have, often resenting their success.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to envy. Regain all Willpower when your envy causes potentially lethal harm.",
		},
		{
			Name:        "Prideful",
			Type:        "Vice",
			Description: "You are arrogant and vain, believing yourself superior to others.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to pride. Regain all Willpower when your pride causes potentially lethal harm.",
		},
		{
			Name:        "Slothful",
			Type:        "Vice",
			Description: "You are lazy and indolent, avoiding work and responsibility.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to sloth. Regain all Willpower when your sloth causes potentially lethal harm.",
		},
		{
			Name:        "Cowardly",
			Type:        "Vice",
			Description: "You are fearful and timid, often running from danger or difficult situations.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to cowardice. Regain all Willpower when your cowardice causes potentially lethal harm.",
		},
		{
			Name:        "Cruel",
			Type:        "Vice",
			Description: "You enjoy causing pain and suffering to others, finding pleasure in their misery.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to cruelty. Regain all Willpower when your cruelty causes potentially lethal harm.",
		},
		{
			Name:        "Reckless",
			Type:        "Vice",
			Description: "You act without thinking, often putting yourself and others in danger.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to recklessness. Regain all Willpower when your recklessness causes potentially lethal harm.",
		},
		{
			Name:        "Paranoid",
			Type:        "Vice",
			Description: "You are suspicious and distrustful, always expecting betrayal.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to paranoia. Regain all Willpower when your paranoia causes potentially lethal harm.",
		},
		{
			Name:        "Addictive",
			Type:        "Vice",
			Description: "You are prone to addiction, whether to substances, gambling, or other vices.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to addiction. Regain all Willpower when your addiction causes potentially lethal harm.",
		},
		{
			Name:        "Impulsive",
			Type:        "Vice",
			Description: "You act on sudden urges without considering consequences.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to impulsiveness. Regain all Willpower when your impulsiveness causes potentially lethal harm.",
		},
		{
			Name:        "Vengeful",
			Type:        "Vice",
			Description: "You hold grudges and seek revenge for real or perceived slights.",
			Effect:      "Regain 1 Willpower when getting someone in trouble due to vengeance. Regain all Willpower when your vengeance causes potentially lethal harm.",
		},
	}
}
