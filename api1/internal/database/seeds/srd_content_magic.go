package seeds

import (
	"github.com/google/uuid"
)

// Helper function to get SRDEntry ID by title
// This is a mapping of titles to UUIDs that will be resolved during seeding
func GetSRDEntryIDByTitle(title string) uuid.UUID {
	// This function is used during seeding to create a mapping
	// The actual UUIDs will be resolved when the entries are created
	return uuid.Nil
}

// MagicSRDContent represents content for magic SRD entries
type MagicSRDContent struct {
	Title   string
	Content string
}

func GetMagicSRDContent() []MagicSRDContent {
	return []MagicSRDContent{
		{
			Title:   "What is Magic?",
			Content: `Magic is the ability to alter reality to one's will by channeling forces from other dimensions known as Astral Realms. "The Gift" is the label used for the ability to channel magic. Most Gifted are born with the gift, but it can also be attained through rituals, bargains, or relics. The most common gift is "Partial Gift" (one realm), while "Full Gift" allows channeling all realms. Summoning is not an Astral Realm but refers to the space between all realms.`,
		},
		{
			Title:   "The Gift and Astral Realms",
			Content: `The Gift is the innate or acquired ability to channel magic from Astral Realms: Air, Dark, Earth, Illusion, Light, Nature, or Summoning. Each realm governs different magical skills and effects. Full Gifted can learn all, Partial Gifted only one. Summoning is unique, drawing from the space between realms.`,
		},
		{
			Title:   "Limits of Magic",
			Content: `Magic cannot create something from nothing, nor bestow permanent bonuses without sacrifice. It cannot directly affect anything outside the caster's line of sight, and concentration cannot be maintained while unconscious. Magic requires a strong connection to the Astral Realm, a clear sense of purpose, and focus (often aided by Godspeech, gestures, or foci).`,
		},
		{
			Title:   "Godspeech and Gestures",
			Content: `Godspeech is the language of the gods and Astral Entities. To non-Gifted, it is incomprehensible; to the Gifted, it conveys images and feelings. Using Godspeech and gestures is standard for spellcasting, helping to focus and direct magical energies.`,
		},
		{
			Title:   "Spellcasting Rules",
			Content: `To cast a spell, roll Magic Skill + Attribute. The skill depends on the magical school; the attribute varies by effect. Spell level determines power and backlash. See "Spell Levels and Backlash" for details. Spells may have thresholds for resistance or backlash.`,
		},
		{
			Title:   "Spell Levels and Backlash",
			Content: `Spells are ranked 1-10 by effect. Each level has a Backlash Threshold. Failing to meet the threshold causes stun damage but does not necessarily mean failure. Example: Level 3 spell has threshold 2; if you roll 1 success, you take 1 stun damage. See the SRD for level-by-level effects.`,
		},
		{
			Title:   "Spell Terms and Attributes",
			Content: `Key terms: Sustained (requires concentration, -2 all actions per spell), Target (element, object, being, effect), Area (possible, no, yes), Tolerance (does it affect tolerance?), Attribute (used for the roll), Resistance (how the target resists: none, size, defense, physical, mental). Attributes: Brawn (impact), Agility (precision), Logic (composition), Wits (perception), Power (mind), Cool (focus).`,
		},
		{
			Title:   "Targeting and Range",
			Content: `Self/touch: affect self or touched target. Range: line of sight, -2 penalty per range band beyond engaged. Multiple targets: split dice pool, all penalties/bonuses applied before split. Area: -4 penalty per area size, -3 to defense per area size. Area spells that do damage go against general armor and soak.`,
		},
		{
			Title:   "Ritual Spellcasting",
			Content: `Ritual casting allows extra time for bonuses (up to Cool/2). Base time is 1 hour, increasing for each bonus. Ritual preparation (Occult + Logic, 1 hour) can reduce area penalties. Rituals often involve drawing Godspeech runes and preparing the area.`,
		},
		{
			Title:   "Anchored Magic",
			Content: `Anchoring a spell makes it permanent, tied to a physical object or creature marked with Godspeech runes. Anchors must be appropriately sized. Destroying the anchor ends the spell. Anchoring requires a sacrifice of reagents (see "Reagents and Foci.`,
		},
		{
			Title:   "Conjunctive Magic",
			Content: `Some effects require combining multiple magic skills. The caster must have sufficient ranks in all required skills; the lowest is used for the dice pool. Example: Transforming earth to water (Air + Earth).`,
		},
		{
			Title:   "Reagents and Foci",
			Content: `Reagents are consumed to reduce backlash. Foci (staves, wands, orbs) help focus magic, increasing range or sustaining spells. Only one focus can benefit a spell at a time. Sustaining foci can maintain spells after casting.`,
		},
		{
			Title: "Air Magic - Spy the Unseen (Level 1)",
			Content: `Spy the Unseen
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: Defense (Dodge only)

This spell allows the caster to feel things that cannot be seen by sending out gusts of air and feeling whom or what it impacts. Including that which cannot be seen.

Successes over the unseen beings defense gives the caster knowledge of their presence as long as it is inside the sustained spells area. Further successes gives the caster knowledge about the beings; size, nature, and whether it is unarmed. Unattended yet invisible objects are automatically detected.`,
		},
		{
			Title: "Air Magic - Sense Electricity (Level 1)",
			Content: `Sense Electricity
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the target to sense electric currents in the area around them.

Success on this spell informs the target of electrical currents around the area. Impressive success gives greater information such as what is powered or what produces the electricity.`,
		},
		{
			Title: "Air Magic - Understand Physics (Level 2)",
			Content: `Understand Physics
Level: 2 (Knowing)
Backlash DT: 1
Sustained: Yes
Target: Object
Area: Possible
Tolerance: No
Attribute: Logic
Resistance: Object size

This spell allows the caster to intuitively understand items that make use of kinetic or electrical energies. This could be a computer, electric light or crossbow.

Success on this spell removes untrained penalties when using the target object. Additional successes can be used as teamwork bonus to science or craft rolls altering the function of the object (but does not give any bonus to its normal use).`,
		},
		{
			Title: "Air Magic - Cut through hardship (Level 3)",
			Content: `Cut through hardship
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Being
Area: possible
Tolerance: No
Attribute: Cool
Resistance: No

This spell uses the air to make a cutting shield around the target(s) allowing them to move unhindered by weather or other environmental hindrances to their movement.

Successes on this spell can decrease the penalties or effects that hinder movement in some way, as long as it is sustained. It has no effect on solid objects and cannot grant bonuses.`,
		},
		{
			Title: "Air Magic - Air is no resistance (Level 3)",
			Content: `Air is no resistance
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Object (weapon)
Area: Possible
Tolerance: No
Attribute: Agility
Resistance: Object Size

This spell alters the air resistance around weapons allowing them to be brought to bear faster. It cannot increase the weapons initiative to a bonus.

Net Successes on this spell decreases initiative penalties from weapon size.
Exceptional success can be used to give the weapon Armor Piercing 1.`,
		},
		{
			Title: "Air Magic - Mage hand (Level 4)",
			Content: `Mage hand
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Object
Area: No
Tolerance: No
Attribute: Agility
Resistance: Object Size

This spell allows the caster to move unattended objects through the air using Telekinesis. Or close doors, pull levers etc.

Successes on this spell can move the object as one move action. Critical success allows for a distance equal to two move actions. The mage cannot inflict damage using this spell but can for example move a live grenade or trigger a damaging trap.`,
		},
		{
			Title: "Air Magic - Shift Current (Level 4)",
			Content: `Shift Current
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Element/object
Area: Possible
Tolerance: No
Attribute: Logic
Resistance: Object Size

This spell allows the caster to manipulate electric currents. A prepared mage may be able to redirect incoming lightning but mostly this can be used to shift currents inside an object. Any mage can use this to damage electronics but those skilled in science or thievery may be able to open electrical locks or make a metal detector give the wrong signal.

Net successes can inflict an equal number of damage to an electronic object. Alternatively it can redirect a ranged electric attack or lower the damage of an electric weapon by net successes. A complex science or thievery check may allow more intricate manipulation.
Impressive success: electronic items can be turned off.`,
		},
		{
			Title: "Air Magic - Kinetic Shield (Level 5)",
			Content: `Kinetic Shield
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Wits OR Cool
Resistance: No

This spell allows the caster to shield targets against any physical or elemental attack, but has no effect on direct magic such as curses.

The spell replaces the (willing) target's defense with the casters (Wits OR Cool + Air) with appropriate penalties for distance to the target or for shielding multiple targets. However the shield suffers no defense penalty from Auto-fire, reach or area attacks.`,
		},
		{
			Title: "Air Magic - Wind Wall (Level 5)",
			Content: `Wind Wall
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Brawn
Resistance: No

This spell creates a massive draft of air that can go in any direction. This throws off any projectile entering or moving through the area.

Successes on this spell gives an equal penalty to any ranged attack traveling into, out of, or through the area.`,
		},
		{
			Title: "Air Magic - Slow Fall (Level 5)",
			Content: `Slow Fall
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being/object
Area: Possible
Tolerance: No
Attribute: Agility
Resistance: Size
Reaction: This spell can be cast as a swift action outside the casters turn as long as they are not surprised and have a Swift action to spend.

This Spell makes the target land safely from any fall.

Success on this spell slows the target to fall at a safe speed.
Impressive success: The caster can choose to let the target fall slowly or simply slow them just before impact.`,
		},
		{
			Title: "Air Magic - Alter Speed (Level 6)",
			Content: `Alter Speed
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility
Resistance: Physical

This spell allows the caster to increase and decrease the target's speed.

Successes on this spell alters the target's Agility (Down to minimum 1, up to augmented maximum.) Impressive success can be used to increase or decrease the target's movement by 1.`,
		},
		{
			Title: "Air Magic - Light Body (Level 6)",
			Content: `Light Body
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: Possible
Tolerance: Yes (half)
Attribute: Agility or Brawn
Resistance: No

This spell uses telekinesis to assist the target to make great leaps and take less damage from falls.

Successes on this spell add automatic successes to Athletics rolls to Jump or land safely.
Impressive success: Successes also provide bonus dice to Climbing and gymnastics (but not dodge). A roll cannot receive both successes and a dice bonus.`,
		},
		{
			Title: "Air Magic - Alter Current (Level 6)",
			Content: `Alter Current
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Object/Element
Area: Possible
Tolerance: No
Attribute: Logic
Resistance: Object Size

This spell allows the caster to increase or decrease electrical currents. This means disrupting electrical lights, altering electrical weapons output and draining batteries.

Net successes on this spell can increase or decrease electrical damage.
Impressive success allows the spell to drain batteries or short out systems.`,
		},
		{
			Title: "Air Magic - Telekinetic Push (Level 7)",
			Content: `Telekinetic Push
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being/object
Area: Possible
Tolerance: No
Attribute: Brawn
Resistance: Physical Resistance/object size

This spell allows the caster to push objects or people from afar.

Each net success of this spell allows the caster to force an object or being a horizontal distance equal to one move. Impressive success makes the target go prone at the end of the movement. If the target impacts something while there are still "unused" moves they take 5 damage per unused move.`,
		},
		{
			Title: "Air Magic - Lightning (Level 7)",
			Content: `Lightning
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being/object
Area: Possible
Tolerance: No
Attribute: Agility
Resistance: Defense

This spell allows the caster to send bolts of lightning forward from their hands electrocuting targets.

On success this spell inflicts damage equal to Logic + net hits. (If the target is wearing metal armor the damage is increased by 2) On Impressive success the target suffers the Shocked condition.`,
		},
		{
			Title: "Air Magic - Flash Forward (Level 8)",
			Content: `Flash Forward
Level: 8 (Transformation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility
Resistance: No

This spell allows the caster to become a blur of speed, flashing forward rapidly. The caster can move up or down as long as there is no solid object in between.

Success on this spell moves the caster along a straight line to the target destination as long as it is within their line of sight (normal penalties for distance apply). They are not subject to things such as elemental hazards or trap triggers in between.`,
		},
		{
			Title: "Air Magic - Wind Walking (Level 8)",
			Content: `Wind Walking
Level: 8 (Transformation)
Backlash DT: 4
Sustained: Yes
Target: Being
Area: Possible
Tolerance: No
Attribute: Logic
Resistance: Physical/Size

This spell allows the caster to turn the target into a human kite, making them float on the air-currents.

Success on this spell makes the target float be carried by wind. Each level of wind will carry the target one move in the wind direction per turn.
Net successes allows for the target to avoid obstacles and greater precision in the movement. Every net success after the first makes the target able to ignore one level of wind so that they can stop or slow down.`,
		},
		{
			Title: "Air Magic - Telekinetic Strength (Level 9)",
			Content: `Telekinetic Strength
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Brawn
Resistance: Physical

The mage layers telekinesis on the target to help or hinder. The telekinesis can add its strength to blows and slow incoming ones. Or it can make a weight bearing down upon the target making them weaker and help incoming blows push in harder.

Successes on this spell increases and decreases the targets Brawn. It affects everything from dice pools to damage and soak.`,
		},
		{
			Title: "Air Magic - Wind Blades (Level 9)",
			Content: `Wind Blades
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Half
Attribute: Agility
Resistance: No

The caster can summon forth wind sharp as blades from the target's arms.

Successes on this spell determine size and damage of the wind blade, should the caster wish to limit size the damage cannot exceed size +2. The caster may split successes between each of the subject's arms to create multiple blades.`,
		},
		{
			Title: "Air Magic - Telekinetic Grasp (Level 10)",
			Content: `Telekinetic Grasp
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Brawn
Resistance: Defense (Grapple check)

This spell allows the mage to grasp something as if in a giant hand. Crushing, smashing, moving or simply holding the target.

Successes on this spell allows the caster to establish a grapple at a distance using the spellcasting dice pool as a grapple check.`,
		},
		{
			Title: "Air Magic - Fly (Level 10)",
			Content: `Fly
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility or Brawn
Resistance: Size

This spell allows the mage to slip the bonds of earth and fly.

Each net hit allows the target to travel one move in any direction per move action.`,
		},
		// Dark Magic Spells
		{
			Title: "Dark Magic - Sense Death (Level 1)",
			Content: `Sense Death
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense death and decay in the area around them.

Success on this spell informs the caster of the presence of death, decay, or undead in the area. Impressive success gives greater information such as the type of death, how long ago it occurred, or the nature of the undead.`,
		},
		{
			Title: "Dark Magic - Sense Fear (Level 1)",
			Content: `Sense Fear
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense fear and terror in the area around them.

Success on this spell informs the caster of the presence of fear or terror in the area. Impressive success gives greater information such as the source of the fear or the intensity of the emotion.`,
		},
		{
			Title: "Dark Magic - Understand Pain (Level 2)",
			Content: `Understand Pain
Level: 2 (Knowing)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to understand the pain and suffering of others.

Success on this spell gives the caster insight into the target's pain and suffering. Additional successes can be used as teamwork bonus to medicine or psychology rolls related to healing or understanding the target's condition.`,
		},
		{
			Title: "Dark Magic - Shadow Step (Level 3)",
			Content: `Shadow Step
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Agility
Resistance: No

This spell allows the caster to step through shadows to move short distances.

Success on this spell allows the caster to move through shadows to a nearby location within line of sight. The distance is limited by the amount of shadow available and the caster's skill.`,
		},
		{
			Title: "Dark Magic - Dark Vision (Level 3)",
			Content: `Dark Vision
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the target to see clearly in complete darkness.

Success on this spell grants the target the ability to see in darkness as if it were daylight. The spell does not work in magical darkness or areas where light is completely absent due to supernatural means.`,
		},
		{
			Title: "Dark Magic - Shadow Hand (Level 4)",
			Content: `Shadow Hand
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Object
Area: No
Tolerance: No
Attribute: Agility
Resistance: Object Size

This spell allows the caster to manipulate objects using shadows as if they were hands.

Successes on this spell allow the caster to move or manipulate objects through shadows. The shadows can perform simple tasks like opening doors, picking up small objects, or triggering mechanisms.`,
		},
		{
			Title: "Dark Magic - Fear Touch (Level 4)",
			Content: `Fear Touch
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to instill fear in a target through touch.

Success on this spell causes the target to experience intense fear. The target may flee, freeze, or suffer penalties to their actions depending on the intensity of the fear and their mental resistance.`,
		},
		{
			Title: "Dark Magic - Shadow Shield (Level 5)",
			Content: `Shadow Shield
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Cool
Resistance: No

This spell creates a shield of darkness that protects the target from harm.

Successes on this spell create a shield that provides protection against physical and magical attacks. The shield is most effective in darkness and may be weakened in bright light.`,
		},
		{
			Title: "Dark Magic - Death Ward (Level 5)",
			Content: `Death Ward
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell protects the target from death magic and necromantic effects.

Successes on this spell provide protection against death magic, necromantic spells, and effects that would cause death or severe harm. The ward may also provide some protection against disease and poison.`,
		},
		{
			Title: "Dark Magic - Shadow Cloak (Level 5)",
			Content: `Shadow Cloak
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Agility
Resistance: No

This spell wraps the target in shadows, making them difficult to see and track.

Successes on this spell make the target harder to see and track, especially in darkness or shadowy areas. The cloak may also provide some protection against detection spells and abilities.`,
		},
		{
			Title: "Dark Magic - Alter Pain (Level 6)",
			Content: `Alter Pain
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Mental

This spell allows the caster to increase or decrease the target's sensitivity to pain.

Successes on this spell can either increase the target's pain sensitivity (making them more vulnerable to damage) or decrease it (providing some protection against pain-based effects).`,
		},
		{
			Title: "Dark Magic - Shadow Form (Level 6)",
			Content: `Shadow Form
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility
Resistance: Physical

This spell allows the target to become partially incorporeal, passing through solid objects.

Successes on this spell allow the target to pass through solid objects and become more difficult to harm with physical attacks. The target may also be able to move through walls and other barriers.`,
		},
		{
			Title: "Dark Magic - Death Touch (Level 6)",
			Content: `Death Touch
Level: 6 (Alteration)
Backlash DT: 3
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to inflict death magic through touch.

Success on this spell inflicts necromantic damage on the target. The damage bypasses normal armor and may cause additional effects such as disease, paralysis, or death depending on the target's resistance and the spell's power.`,
		},
		{
			Title: "Dark Magic - Shadow Strike (Level 7)",
			Content: `Shadow Strike
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Agility
Resistance: Defense

This spell allows the caster to attack from shadows, striking with darkness itself.

Success on this spell inflicts damage equal to Power + net hits. The attack may bypass some forms of armor and can be made from concealment or through shadows.`,
		},
		{
			Title: "Dark Magic - Fear Wave (Level 7)",
			Content: `Fear Wave
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a wave of fear that affects all beings in an area.

Successes on this spell cause all targets in the area to experience fear. The intensity and effects depend on the number of successes and the targets' mental resistance. Some targets may flee, freeze, or suffer penalties to their actions.`,
		},
		{
			Title: "Dark Magic - Shadow Walk (Level 8)",
			Content: `Shadow Walk
Level: 8 (Transformation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility
Resistance: No

This spell allows the caster to travel through shadows to distant locations.

Success on this spell allows the caster to travel through shadows to a distant location. The distance is limited by the caster's skill and the availability of shadows along the path.`,
		},
		{
			Title: "Dark Magic - Death Cloud (Level 8)",
			Content: `Death Cloud
Level: 8 (Transformation)
Backlash DT: 4
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a cloud of death magic that affects all beings in an area.

Successes on this spell create a cloud that inflicts necromantic damage on all targets in the area. The cloud may also cause disease, paralysis, or other death-related effects.`,
		},
		{
			Title: "Dark Magic - Shadow Mastery (Level 9)",
			Content: `Shadow Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility
Resistance: Physical

This spell grants the target complete mastery over shadows and darkness.

Successes on this spell allow the target to control shadows, create darkness, and manipulate shadow-based magic with great skill. The target may also gain enhanced abilities in darkness and shadow.`,
		},
		{
			Title: "Dark Magic - Death Mastery (Level 9)",
			Content: `Death Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Mental

This spell grants the target mastery over death magic and necromantic forces.

Successes on this spell allow the target to control death magic, manipulate life force, and command undead with great skill. The target may also gain enhanced abilities related to death and necromancy.`,
		},
		{
			Title: "Dark Magic - Shadow Realm (Level 10)",
			Content: `Shadow Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a realm of pure shadow and darkness.

Successes on this spell create an area where shadows and darkness reign supreme. Within this realm, shadow magic is enhanced, light is diminished, and the caster has great control over the environment.`,
		},
		{
			Title: "Dark Magic - Death Realm (Level 10)",
			Content: `Death Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a realm where death magic and necromantic forces are dominant.

Successes on this spell create an area where death magic is enhanced, life force is diminished, and the caster has great control over death and necromantic forces.`,
		},
		// Earth Magic Spells
		{
			Title: "Earth Magic - Sense Stone (Level 1)",
			Content: `Sense Stone
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense stone, metal, and earth in the area around them.

Success on this spell informs the caster of the presence of stone, metal, or earth in the area. Impressive success gives greater information such as the type of material, its composition, or hidden deposits.`,
		},
		{
			Title: "Earth Magic - Sense Structure (Level 1)",
			Content: `Sense Structure
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense the structural integrity of buildings and objects.

Success on this spell informs the caster of the structural integrity of buildings, walls, or objects in the area. Impressive success gives greater information such as weak points, load-bearing elements, or hidden passages.`,
		},
		{
			Title: "Earth Magic - Understand Metal (Level 2)",
			Content: `Understand Metal
Level: 2 (Knowing)
Backlash DT: 1
Sustained: Yes
Target: Object
Area: No
Tolerance: No
Attribute: Logic
Resistance: Object Size

This spell allows the caster to intuitively understand metal objects and their properties.

Success on this spell removes untrained penalties when using metal objects. Additional successes can be used as teamwork bonus to craft rolls related to metalworking or engineering.`,
		},
		{
			Title: "Earth Magic - Stone Skin (Level 3)",
			Content: `Stone Skin
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Brawn
Resistance: No

This spell toughens the target's skin, making it more resistant to damage.

Successes on this spell provide additional armor or damage resistance to the target. The effect is most noticeable against physical attacks and may also provide some protection against environmental hazards.`,
		},
		{
			Title: "Earth Magic - Metal Bond (Level 3)",
			Content: `Metal Bond
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Object
Area: No
Tolerance: No
Attribute: Logic
Resistance: Object Size

This spell strengthens the bonds between metal atoms, making objects more durable.

Successes on this spell make metal objects more resistant to damage, wear, and corrosion. The effect may also improve the object's structural integrity and performance.`,
		},
		{
			Title: "Earth Magic - Stone Hand (Level 4)",
			Content: `Stone Hand
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Object
Area: No
Tolerance: No
Attribute: Brawn
Resistance: Object Size

This spell allows the caster to manipulate stone and earth as if it were clay.

Successes on this spell allow the caster to shape, move, or manipulate stone and earth. The caster can create simple structures, move rocks, or reshape terrain within the spell's limits.`,
		},
		{
			Title: "Earth Magic - Metal Shape (Level 4)",
			Content: `Metal Shape
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Object
Area: No
Tolerance: No
Attribute: Logic
Resistance: Object Size

This spell allows the caster to shape and manipulate metal objects.

Successes on this spell allow the caster to reshape, repair, or modify metal objects. The caster can create simple metal tools, repair damaged items, or alter the shape of existing objects.`,
		},
		{
			Title: "Earth Magic - Stone Shield (Level 5)",
			Content: `Stone Shield
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Brawn
Resistance: No

This spell creates a shield of stone that protects the target from harm.

Successes on this spell create a shield that provides excellent protection against physical attacks. The shield is particularly effective against projectiles and melee weapons.`,
		},
		{
			Title: "Earth Magic - Metal Ward (Level 5)",
			Content: `Metal Ward
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Logic
Resistance: No

This spell creates a ward of metal that protects against various forms of attack.

Successes on this spell create a ward that provides protection against physical, magical, and elemental attacks. The ward may also provide some protection against disease and poison.`,
		},
		{
			Title: "Earth Magic - Earth Armor (Level 5)",
			Content: `Earth Armor
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Brawn
Resistance: No

This spell encases the target in armor made of earth and stone.

Successes on this spell create armor that provides excellent protection against physical attacks. The armor may also provide some protection against environmental hazards and extreme temperatures.`,
		},
		{
			Title: "Earth Magic - Alter Density (Level 6)",
			Content: `Alter Density
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Object
Area: No
Tolerance: No
Attribute: Logic
Resistance: Object Size

This spell allows the caster to increase or decrease the density of materials.

Successes on this spell can make objects heavier or lighter, stronger or weaker. The effect may also alter the object's physical properties such as conductivity, hardness, or flexibility.`,
		},
		{
			Title: "Earth Magic - Stone Form (Level 6)",
			Content: `Stone Form
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Brawn
Resistance: Physical

This spell allows the target to take on properties of stone, becoming more durable and resistant.

Successes on this spell make the target more resistant to damage, wear, and environmental hazards. The target may also gain enhanced strength and durability while the spell is active.`,
		},
		{
			Title: "Earth Magic - Metal Form (Level 6)",
			Content: `Metal Form
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Logic
Resistance: Physical

This spell allows the target to take on properties of metal, becoming more conductive and reflective.

Successes on this spell make the target more conductive to electricity, heat, and other forms of energy. The target may also gain enhanced resistance to certain types of damage and environmental effects.`,
		},
		{
			Title: "Earth Magic - Stone Strike (Level 7)",
			Content: `Stone Strike
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Brawn
Resistance: Defense

This spell allows the caster to attack with the force of stone and earth.

Success on this spell inflicts damage equal to Brawn + net hits. The attack may cause additional effects such as stunning, knocking prone, or creating difficult terrain.`,
		},
		{
			Title: "Earth Magic - Metal Strike (Level 7)",
			Content: `Metal Strike
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Logic
Resistance: Defense

This spell allows the caster to attack with the precision and conductivity of metal.

Success on this spell inflicts damage equal to Logic + net hits. The attack may cause additional effects such as electrical damage, metal poisoning, or disruption of electronic devices.`,
		},
		{
			Title: "Earth Magic - Earth Walk (Level 8)",
			Content: `Earth Walk
Level: 8 (Transformation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: Yes
Attribute: Brawn
Resistance: No

This spell allows the caster to travel through earth and stone as if it were air.

Success on this spell allows the caster to move through solid earth and stone. The caster can tunnel, create passages, or move through walls and other solid barriers.`,
		},
		{
			Title: "Earth Magic - Metal Walk (Level 8)",
			Content: `Metal Walk
Level: 8 (Transformation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: Yes
Attribute: Logic
Resistance: No

This spell allows the caster to travel through metal objects and structures.

Success on this spell allows the caster to move through metal objects, walls, and structures. The caster can pass through metal barriers, travel along metal surfaces, or move through metal-based technology.`,
		},
		{
			Title: "Earth Magic - Stone Mastery (Level 9)",
			Content: `Stone Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Brawn
Resistance: Physical

This spell grants the target complete mastery over stone and earth.

Successes on this spell allow the target to control stone, earth, and related materials with great skill. The target may also gain enhanced abilities related to stone and earth magic.`,
		},
		{
			Title: "Earth Magic - Metal Mastery (Level 9)",
			Content: `Metal Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Logic
Resistance: Physical

This spell grants the target complete mastery over metal and related materials.

Successes on this spell allow the target to control metal, alloys, and related materials with great skill. The target may also gain enhanced abilities related to metal magic and technology.`,
		},
		{
			Title: "Earth Magic - Stone Realm (Level 10)",
			Content: `Stone Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Brawn
Resistance: Physical

This spell creates a realm where stone and earth reign supreme.

Successes on this spell create an area where stone and earth magic is enhanced, other elements are diminished, and the caster has great control over the environment.`,
		},
		{
			Title: "Earth Magic - Metal Realm (Level 10)",
			Content: `Metal Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: Physical

This spell creates a realm where metal and technology reign supreme.

Successes on this spell create an area where metal magic is enhanced, other elements are diminished, and the caster has great control over metal and technology.`,
		},
		// Illusion Magic Spells
		{
			Title: "Illusion Magic - Sense Illusion (Level 1)",
			Content: `Sense Illusion
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense illusions and magical deceptions in the area around them.

Success on this spell informs the caster of the presence of illusions or magical deceptions in the area. Impressive success gives greater information such as the type of illusion, its source, or how to dispel it.`,
		},
		{
			Title: "Illusion Magic - Sense Truth (Level 1)",
			Content: `Sense Truth
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense truth and lies in the area around them.

Success on this spell informs the caster of the presence of truth or deception in the area. Impressive success gives greater information such as the nature of the deception or the source of the truth.`,
		},
		{
			Title: "Illusion Magic - Understand Deception (Level 2)",
			Content: `Understand Deception
Level: 2 (Knowing)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Logic
Resistance: Mental

This spell allows the caster to understand the nature of deception and illusion.

Success on this spell gives the caster insight into the nature of deception and illusion. Additional successes can be used as teamwork bonus to rolls related to detecting or creating illusions.`,
		},
		{
			Title: "Illusion Magic - Minor Illusion (Level 3)",
			Content: `Minor Illusion
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: No

This spell creates a minor illusion that can deceive the senses.

Successes on this spell create a simple illusion that can deceive sight, sound, or other senses. The illusion is limited in complexity and may be dispelled by disbelief or physical interaction.`,
		},
		{
			Title: "Illusion Magic - Disguise Self (Level 3)",
			Content: `Disguise Self
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Logic
Resistance: No

This spell allows the caster to alter their appearance to appear as someone else.

Successes on this spell allow the caster to change their appearance to resemble another person or creature. The disguise may be dispelled by disbelief or magical detection.`,
		},
		{
			Title: "Illusion Magic - Silent Image (Level 4)",
			Content: `Silent Image
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: No

This spell creates a silent visual illusion that can deceive the eyes.

Successes on this spell create a visual illusion that appears real but makes no sound. The illusion may be dispelled by disbelief or physical interaction.`,
		},
		{
			Title: "Illusion Magic - Ventriloquism (Level 4)",
			Content: `Ventriloquism
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: No

This spell allows the caster to make sounds appear to come from another location.

Successes on this spell allow the caster to create sounds that appear to originate from a different location. The sounds may be voices, music, or other audio effects.`,
		},
		{
			Title: "Illusion Magic - Mirror Image (Level 5)",
			Content: `Mirror Image
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Logic
Resistance: No

This spell creates illusory duplicates of the caster that can confuse attackers.

Successes on this spell create illusory duplicates of the caster. Attackers may target the illusions instead of the real caster, providing protection through misdirection.`,
		},
		{
			Title: "Illusion Magic - Invisibility (Level 5)",
			Content: `Invisibility
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Logic
Resistance: No

This spell makes the target invisible to normal sight.

Successes on this spell make the target invisible to normal sight. The invisibility may be dispelled by magical detection or if the target takes certain actions.`,
		},
		{
			Title: "Illusion Magic - Blur (Level 5)",
			Content: `Blur
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Logic
Resistance: No

This spell makes the target appear blurred and difficult to target.

Successes on this spell make the target appear blurred and difficult to target with attacks. The effect may also provide some protection against detection spells.`,
		},
		{
			Title: "Illusion Magic - Alter Appearance (Level 6)",
			Content: `Alter Appearance
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Logic
Resistance: Mental

This spell allows the caster to alter the appearance of themselves or others.

Successes on this spell allow the caster to change the appearance of the target. The changes may be subtle or dramatic, and may affect how others perceive and interact with the target.`,
		},
		{
			Title: "Illusion Magic - Phantasmal Force (Level 6)",
			Content: `Phantasmal Force
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Logic
Resistance: Mental

This spell creates a phantasmal force that can interact with the target's mind and senses.

Successes on this spell create a phantasmal force that can affect the target's perception and behavior. The force may cause fear, confusion, or other mental effects.`,
		},
		{
			Title: "Illusion Magic - Disguise Other (Level 6)",
			Content: `Disguise Other
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Logic
Resistance: Mental

This spell allows the caster to alter the appearance of another person or creature.

Successes on this spell allow the caster to change the appearance of the target to resemble another person or creature. The disguise may be dispelled by disbelief or magical detection.`,
		},
		{
			Title: "Illusion Magic - Phantasmal Killer (Level 7)",
			Content: `Phantasmal Killer
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Logic
Resistance: Mental

This spell creates a phantasmal killer that can harm the target through fear and illusion.

Success on this spell creates a phantasmal killer that can inflict damage on the target through fear and mental trauma. The damage may be physical or mental in nature.`,
		},
		{
			Title: "Illusion Magic - Mass Suggestion (Level 7)",
			Content: `Mass Suggestion
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: Mental

This spell allows the caster to suggest actions to multiple targets simultaneously.

Successes on this spell allow the caster to suggest actions to all targets in the area. The suggestions must be reasonable and may be resisted by strong-willed individuals.`,
		},
		{
			Title: "Illusion Magic - Etherealness (Level 8)",
			Content: `Etherealness
Level: 8 (Transformation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: Yes
Attribute: Logic
Resistance: No

This spell allows the caster to become ethereal, passing through solid objects and becoming immune to most physical effects.

Success on this spell allows the caster to become ethereal and pass through solid objects. The caster may also become immune to most physical attacks and effects while ethereal.`,
		},
		{
			Title: "Illusion Magic - Mass Invisibility (Level 8)",
			Content: `Mass Invisibility
Level: 8 (Transformation)
Backlash DT: 4
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: No

This spell makes all targets in an area invisible to normal sight.

Successes on this spell make all targets in the area invisible to normal sight. The invisibility may be dispelled by magical detection or if the targets take certain actions.`,
		},
		{
			Title: "Illusion Magic - Illusion Mastery (Level 9)",
			Content: `Illusion Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Logic
Resistance: Mental

This spell grants the target complete mastery over illusion and deception.

Successes on this spell allow the target to create and control illusions with great skill. The target may also gain enhanced abilities related to illusion magic and deception.`,
		},
		{
			Title: "Illusion Magic - Reality Warp (Level 9)",
			Content: `Reality Warp
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: Mental

This spell allows the caster to warp reality within a limited area.

Successes on this spell allow the caster to alter the nature of reality within the area. The changes may affect physical laws, perception, or the behavior of objects and beings within the area.`,
		},
		{
			Title: "Illusion Magic - Illusion Realm (Level 10)",
			Content: `Illusion Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: Mental

This spell creates a realm where illusion and deception reign supreme.

Successes on this spell create an area where illusion magic is enhanced, reality is malleable, and the caster has great control over perception and deception.`,
		},
		{
			Title: "Illusion Magic - Reality Realm (Level 10)",
			Content: `Reality Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Logic
Resistance: Mental

This spell creates a realm where the caster has complete control over reality itself.

Successes on this spell create an area where the caster has complete control over reality, perception, and the nature of existence within the realm.`,
		},
		// Light Magic Spells
		{
			Title: "Light Magic - Sense Light (Level 1)",
			Content: `Sense Light
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense light and illumination in the area around them.

Success on this spell informs the caster of the presence and intensity of light in the area. Impressive success gives greater information such as the source of the light, its color, or magical properties.`,
		},
		{
			Title: "Light Magic - Sense Good (Level 1)",
			Content: `Sense Good
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense goodness and purity in the area around them.

Success on this spell informs the caster of the presence of good or pure beings in the area. Impressive success gives greater information such as the nature of the goodness or the source of the purity.`,
		},
		{
			Title: "Light Magic - Understand Purity (Level 2)",
			Content: `Understand Purity
Level: 2 (Knowing)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to understand the nature of purity and goodness.

Success on this spell gives the caster insight into the nature of purity and goodness. Additional successes can be used as teamwork bonus to rolls related to healing or purification.`,
		},
		{
			Title: "Light Magic - Light Step (Level 3)",
			Content: `Light Step
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Agility
Resistance: No

This spell allows the caster to step through light to move short distances.

Success on this spell allows the caster to move through light to a nearby location within line of sight. The distance is limited by the amount of light available and the caster's skill.`,
		},
		{
			Title: "Light Magic - Light Vision (Level 3)",
			Content: `Light Vision
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the target to see clearly in any light conditions.

Success on this spell grants the target the ability to see clearly in any light conditions, including darkness, bright light, or magical illumination.`,
		},
		{
			Title: "Light Magic - Light Hand (Level 4)",
			Content: `Light Hand
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Object
Area: No
Tolerance: No
Attribute: Agility
Resistance: Object Size

This spell allows the caster to manipulate objects using light as if they were hands.

Successes on this spell allow the caster to move or manipulate objects through light. The light can perform simple tasks like opening doors, picking up small objects, or triggering mechanisms.`,
		},
		{
			Title: "Light Magic - Healing Touch (Level 4)",
			Content: `Healing Touch
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell allows the caster to heal wounds through touch.

Success on this spell heals damage and may cure minor ailments. The healing is most effective on living beings and may be less effective on undead or constructs.`,
		},
		{
			Title: "Light Magic - Light Shield (Level 5)",
			Content: `Light Shield
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Cool
Resistance: No

This spell creates a shield of light that protects the target from harm.

Successes on this spell create a shield that provides protection against physical and magical attacks. The shield is most effective against darkness and evil, and may be weakened in areas of strong darkness.`,
		},
		{
			Title: "Light Magic - Protection from Evil (Level 5)",
			Content: `Protection from Evil
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell protects the target from evil and malevolent forces.

Successes on this spell provide protection against evil magic, malevolent beings, and effects that would cause harm through evil means. The ward may also provide some protection against corruption and disease.`,
		},
		{
			Title: "Light Magic - Light Cloak (Level 5)",
			Content: `Light Cloak
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Agility
Resistance: No

This spell wraps the target in light, making them difficult to harm and easy to see.

Successes on this spell make the target more resistant to harm and easier to see in darkness. The cloak may also provide some protection against detection spells and abilities.`,
		},
		{
			Title: "Light Magic - Alter Purity (Level 6)",
			Content: `Alter Purity
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Mental

This spell allows the caster to increase or decrease the target's purity and goodness.

Successes on this spell can either increase the target's purity (making them more resistant to evil) or decrease it (making them more vulnerable to corruption).`,
		},
		{
			Title: "Light Magic - Light Form (Level 6)",
			Content: `Light Form
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility
Resistance: Physical

This spell allows the target to take on properties of light, becoming more radiant and pure.

Successes on this spell make the target more radiant and pure, gaining enhanced abilities related to light and goodness. The target may also become more resistant to evil and corruption.`,
		},
		{
			Title: "Light Magic - Purification Touch (Level 6)",
			Content: `Purification Touch
Level: 6 (Alteration)
Backlash DT: 3
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to purify the target through touch.

Success on this spell purifies the target, removing corruption, disease, or evil influences. The purification may also heal damage and restore the target's natural state.`,
		},
		{
			Title: "Light Magic - Light Strike (Level 7)",
			Content: `Light Strike
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Agility
Resistance: Defense

This spell allows the caster to attack with the power of light and goodness.

Success on this spell inflicts damage equal to Power + net hits. The attack is most effective against evil and undead, and may cause additional effects such as blinding or purification.`,
		},
		{
			Title: "Light Magic - Healing Wave (Level 7)",
			Content: `Healing Wave
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: No

This spell creates a wave of healing energy that affects all beings in an area.

Successes on this spell heal damage and may cure ailments for all targets in the area. The healing is most effective on living beings and may be less effective on undead or constructs.`,
		},
		{
			Title: "Light Magic - Light Walk (Level 8)",
			Content: `Light Walk
Level: 8 (Transformation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility
Resistance: No

This spell allows the caster to travel through light to distant locations.

Success on this spell allows the caster to travel through light to a distant location. The distance is limited by the caster's skill and the availability of light along the path.`,
		},
		{
			Title: "Light Magic - Purification Cloud (Level 8)",
			Content: `Purification Cloud
Level: 8 (Transformation)
Backlash DT: 4
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a cloud of purification that affects all beings in an area.

Successes on this spell create a cloud that purifies all targets in the area, removing corruption, disease, or evil influences. The cloud may also heal damage and restore natural states.`,
		},
		{
			Title: "Light Magic - Light Mastery (Level 9)",
			Content: `Light Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Agility
Resistance: Physical

This spell grants the target complete mastery over light and goodness.

Successes on this spell allow the target to control light, create illumination, and manipulate light-based magic with great skill. The target may also gain enhanced abilities related to light and goodness.`,
		},
		{
			Title: "Light Magic - Purity Mastery (Level 9)",
			Content: `Purity Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Mental

This spell grants the target mastery over purity and goodness.

Successes on this spell allow the target to control purity, manipulate goodness, and command benevolent forces with great skill. The target may also gain enhanced abilities related to purity and healing.`,
		},
		{
			Title: "Light Magic - Light Realm (Level 10)",
			Content: `Light Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a realm of pure light and goodness.

Successes on this spell create an area where light and goodness reign supreme. Within this realm, light magic is enhanced, darkness is diminished, and the caster has great control over the environment.`,
		},
		{
			Title: "Light Magic - Purity Realm (Level 10)",
			Content: `Purity Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a realm where purity and goodness are dominant.

Successes on this spell create an area where purity and goodness are enhanced, evil is diminished, and the caster has great control over purification and healing forces.`,
		},
		// Nature Magic Spells
		{
			Title: "Nature Magic - Sense Life (Level 1)",
			Content: `Sense Life
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense life and living things in the area around them.

Success on this spell informs the caster of the presence of living beings in the area. Impressive success gives greater information such as the type of life, its health, or its emotional state.`,
		},
		{
			Title: "Nature Magic - Sense Growth (Level 1)",
			Content: `Sense Growth
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense growth and natural processes in the area around them.

Success on this spell informs the caster of the presence of growing things, natural processes, or environmental changes in the area. Impressive success gives greater information such as the rate of growth, the health of plants, or environmental conditions.`,
		},
		{
			Title: "Nature Magic - Understand Life (Level 2)",
			Content: `Understand Life
Level: 2 (Knowing)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to understand the nature of life and living things.

Success on this spell gives the caster insight into the nature of life and living things. Additional successes can be used as teamwork bonus to rolls related to healing, medicine, or biology.`,
		},
		{
			Title: "Nature Magic - Plant Growth (Level 3)",
			Content: `Plant Growth
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: No

This spell accelerates the growth of plants in an area.

Successes on this spell cause plants in the area to grow rapidly. The growth may create difficult terrain, provide cover, or produce useful plant materials.`,
		},
		{
			Title: "Nature Magic - Animal Bond (Level 3)",
			Content: `Animal Bond
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a bond between the caster and an animal.

Successes on this spell create a bond that allows the caster to communicate with and influence the animal. The bond may also provide some protection or benefits to both the caster and the animal.`,
		},
		{
			Title: "Nature Magic - Plant Control (Level 4)",
			Content: `Plant Control
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: No

This spell allows the caster to control and manipulate plants in an area.

Successes on this spell allow the caster to control the movement, growth, and behavior of plants in the area. The caster can create barriers, clear paths, or use plants for various purposes.`,
		},
		{
			Title: "Nature Magic - Animal Control (Level 4)",
			Content: `Animal Control
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to control and influence animals.

Successes on this spell allow the caster to control the behavior and actions of animals. The control may be subtle influence or direct command, depending on the animal's intelligence and the caster's skill.`,
		},
		{
			Title: "Nature Magic - Natural Shield (Level 5)",
			Content: `Natural Shield
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell creates a shield of natural materials that protects the target from harm.

Successes on this spell create a shield that provides protection against physical and magical attacks. The shield is most effective in natural environments and may be enhanced by nearby plants or animals.`,
		},
		{
			Title: "Nature Magic - Life Ward (Level 5)",
			Content: `Life Ward
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell protects the target from harm by enhancing their natural life force.

Successes on this spell provide protection against damage, disease, and harmful effects by enhancing the target's natural life force and healing abilities.`,
		},
		{
			Title: "Nature Magic - Natural Armor (Level 5)",
			Content: `Natural Armor
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell encases the target in armor made of natural materials.

Successes on this spell create armor that provides excellent protection against physical attacks. The armor may also provide some protection against environmental hazards and extreme temperatures.`,
		},
		{
			Title: "Nature Magic - Alter Growth (Level 6)",
			Content: `Alter Growth
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Mental

This spell allows the caster to increase or decrease the growth rate of living things.

Successes on this spell can accelerate or slow the growth of plants, animals, or other living things. The effect may also affect healing rates, aging, or other natural processes.`,
		},
		{
			Title: "Nature Magic - Plant Form (Level 6)",
			Content: `Plant Form
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Physical

This spell allows the target to take on properties of plants, becoming more resilient and connected to nature.

Successes on this spell make the target more resilient to damage, environmental hazards, and disease. The target may also gain enhanced abilities related to plants and natural environments.`,
		},
		{
			Title: "Nature Magic - Animal Form (Level 6)",
			Content: `Animal Form
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Physical

This spell allows the target to take on properties of animals, gaining enhanced senses and abilities.

Successes on this spell give the target enhanced senses, strength, or other animal abilities. The target may also gain the ability to communicate with animals or take on animal-like behaviors.`,
		},
		{
			Title: "Nature Magic - Natural Strike (Level 7)",
			Content: `Natural Strike
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Defense

This spell allows the caster to attack with the power of nature itself.

Success on this spell inflicts damage equal to Power + net hits. The attack may cause additional effects such as disease, poison, or environmental damage.`,
		},
		{
			Title: "Nature Magic - Life Wave (Level 7)",
			Content: `Life Wave
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: No

This spell creates a wave of life energy that affects all beings in an area.

Successes on this spell heal damage and may cure ailments for all living targets in the area. The wave may also enhance natural healing and growth processes.`,
		},
		{
			Title: "Nature Magic - Natural Walk (Level 8)",
			Content: `Natural Walk
Level: 8 (Transformation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: No

This spell allows the caster to travel through natural environments with great speed and ease.

Success on this spell allows the caster to move through natural environments as if they were part of the landscape. The caster can move through plants, over difficult terrain, or along natural paths with great efficiency.`,
		},
		{
			Title: "Nature Magic - Life Cloud (Level 8)",
			Content: `Life Cloud
Level: 8 (Transformation)
Backlash DT: 4
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: No

This spell creates a cloud of life energy that affects all beings in an area.

Successes on this spell create a cloud that enhances life, healing, and growth for all living targets in the area. The cloud may also provide protection against disease and environmental hazards.`,
		},
		{
			Title: "Nature Magic - Nature Mastery (Level 9)",
			Content: `Nature Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Physical

This spell grants the target complete mastery over nature and natural forces.

Successes on this spell allow the target to control nature, manipulate natural forces, and command living things with great skill. The target may also gain enhanced abilities related to nature and life.`,
		},
		{
			Title: "Nature Magic - Life Mastery (Level 9)",
			Content: `Life Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Mental

This spell grants the target mastery over life and living forces.

Successes on this spell allow the target to control life, manipulate living forces, and command life energy with great skill. The target may also gain enhanced abilities related to life and healing.`,
		},
		{
			Title: "Nature Magic - Nature Realm (Level 10)",
			Content: `Nature Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Physical

This spell creates a realm where nature and natural forces reign supreme.

Successes on this spell create an area where nature magic is enhanced, artificial forces are diminished, and the caster has great control over natural environments and living things.`,
		},
		{
			Title: "Nature Magic - Life Realm (Level 10)",
			Content: `Life Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a realm where life and living forces are dominant.

Successes on this spell create an area where life magic is enhanced, death is diminished, and the caster has great control over life energy and living forces.`,
		},
		// Summoning Magic Spells
		{
			Title: "Summoning Magic - Sense Portal (Level 1)",
			Content: `Sense Portal
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense portals and dimensional rifts in the area around them.

Success on this spell informs the caster of the presence of portals, dimensional rifts, or other planar connections in the area. Impressive success gives greater information such as the destination, stability, or nature of the portal.`,
		},
		{
			Title: "Summoning Magic - Sense Entity (Level 1)",
			Content: `Sense Entity
Level: 1 (Detection)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: Yes
Tolerance: No
Attribute: Wits
Resistance: No

This spell allows the caster to sense otherworldly entities and summoned beings in the area around them.

Success on this spell informs the caster of the presence of otherworldly entities, summoned beings, or extraplanar creatures in the area. Impressive success gives greater information such as the type of entity, its power, or its origin.`,
		},
		{
			Title: "Summoning Magic - Understand Summoning (Level 2)",
			Content: `Understand Summoning
Level: 2 (Knowing)
Backlash DT: 1
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to understand the nature of summoning and extraplanar magic.

Success on this spell gives the caster insight into the nature of summoning magic and extraplanar forces. Additional successes can be used as teamwork bonus to rolls related to summoning or planar travel.`,
		},
		{
			Title: "Summoning Magic - Minor Summon (Level 3)",
			Content: `Minor Summon
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: No

This spell allows the caster to summon minor extraplanar entities or create small portals.

Successes on this spell allow the caster to summon minor entities or create small portals to other planes. The summoned entities are typically weak and may only be able to perform simple tasks.`,
		},
		{
			Title: "Summoning Magic - Planar Step (Level 3)",
			Content: `Planar Step
Level: 3 (Minor Alteration)
Backlash DT: 2
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell allows the caster to take a single step into another plane of existence.

Success on this spell allows the caster to briefly step into another plane. The caster may be able to see, hear, or interact with the other plane, but cannot remain there for long.`,
		},
		{
			Title: "Summoning Magic - Summon Familiar (Level 4)",
			Content: `Summon Familiar
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Area
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell allows the caster to summon a familiar - a minor extraplanar entity bound to serve them.

Successes on this spell allow the caster to summon and maintain a familiar. The familiar can perform various tasks, provide assistance, or serve as a companion to the caster.`,
		},
		{
			Title: "Summoning Magic - Planar Hand (Level 4)",
			Content: `Planar Hand
Level: 4 (Minor Manipulation)
Backlash DT: 2
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: No

This spell allows the caster to reach through planar boundaries to manipulate objects or interact with other planes.

Successes on this spell allow the caster to reach through planar boundaries to interact with objects or beings on other planes. The caster may be able to retrieve items, communicate, or perform other actions across planar boundaries.`,
		},
		{
			Title: "Summoning Magic - Planar Shield (Level 5)",
			Content: `Planar Shield
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell creates a shield that protects against extraplanar threats and planar magic.

Successes on this spell create a shield that provides protection against extraplanar attacks, planar magic, and otherworldly threats. The shield may also provide some protection against dimensional effects.`,
		},
		{
			Title: "Summoning Magic - Binding Ward (Level 5)",
			Content: `Binding Ward
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell creates a ward that prevents extraplanar entities from entering or affecting the protected area.

Successes on this spell create a ward that prevents extraplanar entities from entering the area or affecting those within it. The ward may also provide some protection against summoning magic.`,
		},
		{
			Title: "Summoning Magic - Planar Cloak (Level 5)",
			Content: `Planar Cloak
Level: 5 (Shielding)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: No

This spell wraps the target in planar energy, making them difficult to detect or affect across planar boundaries.

Successes on this spell make the target more difficult to detect or affect across planar boundaries. The cloak may also provide some protection against extraplanar detection and scrying.`,
		},
		{
			Title: "Summoning Magic - Alter Binding (Level 6)",
			Content: `Alter Binding
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Mental

This spell allows the caster to strengthen or weaken the binding of summoned entities.

Successes on this spell can strengthen the binding of summoned entities (making them more loyal and powerful) or weaken it (making them more independent or rebellious).`,
		},
		{
			Title: "Summoning Magic - Planar Form (Level 6)",
			Content: `Planar Form
Level: 6 (Alteration)
Backlash DT: 3
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Physical

This spell allows the target to take on properties of extraplanar entities, becoming more connected to other planes.

Successes on this spell make the target more connected to other planes, gaining enhanced abilities related to extraplanar magic and planar travel. The target may also become more resistant to planar effects.`,
		},
		{
			Title: "Summoning Magic - Banish Entity (Level 6)",
			Content: `Banish Entity
Level: 6 (Alteration)
Backlash DT: 3
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Mental

This spell allows the caster to banish extraplanar entities back to their home plane.

Success on this spell can banish extraplanar entities back to their home plane. The effectiveness depends on the entity's power and the caster's skill.`,
		},
		{
			Title: "Summoning Magic - Summon Entity (Level 7)",
			Content: `Summon Entity
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: No

This spell allows the caster to summon powerful extraplanar entities to serve them.

Successes on this spell allow the caster to summon and maintain powerful extraplanar entities. The summoned entities can perform complex tasks, provide powerful assistance, or serve as formidable allies.`,
		},
		{
			Title: "Summoning Magic - Planar Strike (Level 7)",
			Content: `Planar Strike
Level: 7 (Manipulation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: No
Attribute: Power
Resistance: Defense

This spell allows the caster to attack with the power of extraplanar forces.

Success on this spell inflicts damage equal to Power + net hits. The attack may cause additional effects such as planar damage, dimensional effects, or extraplanar afflictions.`,
		},
		{
			Title: "Summoning Magic - Planar Walk (Level 8)",
			Content: `Planar Walk
Level: 8 (Transformation)
Backlash DT: 4
Sustained: No
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: No

This spell allows the caster to travel between planes of existence.

Success on this spell allows the caster to travel to another plane of existence. The destination and duration depend on the caster's skill and the nature of the target plane.`,
		},
		{
			Title: "Summoning Magic - Planar Cloud (Level 8)",
			Content: `Planar Cloud
Level: 8 (Transformation)
Backlash DT: 4
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a cloud of planar energy that affects all beings in an area.

Successes on this spell create a cloud that affects all targets in the area with planar effects. The cloud may cause dimensional damage, planar afflictions, or other extraplanar effects.`,
		},
		{
			Title: "Summoning Magic - Summoning Mastery (Level 9)",
			Content: `Summoning Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Physical

This spell grants the target complete mastery over summoning and extraplanar magic.

Successes on this spell allow the target to control summoning magic, manipulate extraplanar forces, and command summoned entities with great skill. The target may also gain enhanced abilities related to summoning and planar travel.`,
		},
		{
			Title: "Summoning Magic - Planar Mastery (Level 9)",
			Content: `Planar Mastery
Level: 9 (Major Alteration)
Backlash DT: 5
Sustained: Yes
Target: Being
Area: No
Tolerance: Yes
Attribute: Power
Resistance: Mental

This spell grants the target mastery over planar forces and dimensional magic.

Successes on this spell allow the target to control planar forces, manipulate dimensional magic, and command extraplanar entities with great skill. The target may also gain enhanced abilities related to planar travel and dimensional manipulation.`,
		},
		{
			Title: "Summoning Magic - Summoning Realm (Level 10)",
			Content: `Summoning Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a realm where summoning and extraplanar forces reign supreme.

Successes on this spell create an area where summoning magic is enhanced, extraplanar forces are dominant, and the caster has great control over summoned entities and planar effects.`,
		},
		{
			Title: "Summoning Magic - Planar Realm (Level 10)",
			Content: `Planar Realm
Level: 10 (Major Manipulation)
Backlash DT: 5
Sustained: Yes
Target: Area
Area: Yes
Tolerance: No
Attribute: Power
Resistance: Mental

This spell creates a realm where planar forces and dimensional magic are dominant.

Successes on this spell create an area where planar magic is enhanced, dimensional forces are dominant, and the caster has great control over extraplanar entities and dimensional effects.`,
		},
	}
}
