package itto

import (
	"fmt"

	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/player"
)

var highPlungeFrames []int
var lowPlungeFrames []int

const lowPlungeHitmark = 37
const highPlungeHitmark = 40
const collisionHitmark = lowPlungeHitmark - 6

const lowPlungePoiseDMG = 150.0
const lowPlungeRadius = 3.0

const lowPlungePoiseDMGB = 150.0
const lowPlungeRadiusB = 4.0

const highPlungePoiseDMG = 200.0
const highPlungeRadius = 5.0

const highPlungePoiseDMGB = 200.0
const highPlungeRadiusB = 6.0

func init() {
	// low_plunge -> x
	lowPlungeFrames = frames.InitAbilSlice(83)
	lowPlungeFrames[action.ActionAttack] = 49
	lowPlungeFrames[action.ActionCharge] = 49
	lowPlungeFrames[action.ActionSkill] = 48
	lowPlungeFrames[action.ActionBurst] = 49
	lowPlungeFrames[action.ActionDash] = 37
	lowPlungeFrames[action.ActionSwap] = 60

	// high_plunge -> x
	highPlungeFrames = frames.InitAbilSlice(87)
	highPlungeFrames[action.ActionAttack] = 53
	lowPlungeFrames[action.ActionCharge] = 52
	highPlungeFrames[action.ActionSkill] = 51
	highPlungeFrames[action.ActionBurst] = 51
	highPlungeFrames[action.ActionDash] = 40
	highPlungeFrames[action.ActionWalk] = 85
	highPlungeFrames[action.ActionSwap] = 62
}

// Low Plunge attack damage queue generator
// Use the "collision" optional argument if you want to do a falling hit on the way down
// Default = 0
func (c *char) LowPlungeAttack(p map[string]int) (action.Info, error) {
	defer c.Core.Player.SetAirborne(player.Grounded)
	switch c.Core.Player.Airborne() {
	case player.AirborneXianyun:
		return c.lowPlungeXY(p)
	default:
		return action.Info{}, fmt.Errorf("%s low_plunge can only be used while airborne", c.Base.Key.String())
	}
}

func (c *char) lowPlungeXY(p map[string]int) (action.Info, error) {
	collision, ok := p["collision"]
	if !ok {
		collision = 0 // Whether or not collision hit
	}

	if collision > 0 {
		c.plungeCollision(collisionHitmark)
	}

	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Low Plunge",
		AttackTag:  attacks.AttackTagPlunge,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeBlunt,
		PoiseDMG:   lowPlungePoiseDMG,
		Element:    attributes.Physical,
		Durability: 25,
		Mult:       lowPlunge[c.TalentLvlAttack()],
	}
	lowPlungeRadius := lowPlungeRadius
	if c.StatusIsActive(burstBuffKey) {
		ai.Element = attributes.Geo
		ai.IgnoreInfusion = true
		ai.PoiseDMG = lowPlungePoiseDMGB
		lowPlungeRadius = lowPlungeRadiusB
	}

	c.Core.QueueAttack(
		ai,
		combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, lowPlungeRadius),
		lowPlungeHitmark,
		lowPlungeHitmark,
	)

	return action.Info{
		Frames:          frames.NewAbilFunc(lowPlungeFrames),
		AnimationLength: lowPlungeFrames[action.InvalidAction],
		CanQueueAfter:   lowPlungeFrames[action.ActionDash],
		State:           action.PlungeAttackState,
	}, nil
}

// High Plunge attack damage queue generator
// Use the "collision" optional argument if you want to do a falling hit on the way down
// Default = 0
func (c *char) HighPlungeAttack(p map[string]int) (action.Info, error) {
	defer c.Core.Player.SetAirborne(player.Grounded)
	switch c.Core.Player.Airborne() {
	case player.AirborneXianyun:
		return c.highPlungeXY(p)
	default:
		return action.Info{}, fmt.Errorf("%s high_plunge can only be used while airborne", c.Base.Key.String())
	}
}

func (c *char) highPlungeXY(p map[string]int) (action.Info, error) {
	collision, ok := p["collision"]
	if !ok {
		collision = 0 // Whether or not collision hit
	}

	if collision > 0 {
		c.plungeCollision(collisionHitmark)
	}

	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "High Plunge",
		AttackTag:  attacks.AttackTagPlunge,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeBlunt,
		PoiseDMG:   highPlungePoiseDMG,
		Element:    attributes.Physical,
		Durability: 25,
		Mult:       highPlunge[c.TalentLvlAttack()],
	}
	highPlungeRadius := highPlungeRadius
	if c.StatusIsActive(burstBuffKey) {
		ai.Element = attributes.Geo
		ai.IgnoreInfusion = true
		ai.PoiseDMG = highPlungePoiseDMGB
		highPlungeRadius = highPlungeRadiusB
	}

	c.Core.QueueAttack(
		ai,
		combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, highPlungeRadius),
		highPlungeHitmark,
		highPlungeHitmark,
	)

	return action.Info{
		Frames:          frames.NewAbilFunc(highPlungeFrames),
		AnimationLength: highPlungeFrames[action.InvalidAction],
		CanQueueAfter:   highPlungeFrames[action.ActionDash],
		State:           action.PlungeAttackState,
	}, nil
}

// Plunge normal falling attack damage queue generator
// Standard - Always part of high/low plunge attacks
func (c *char) plungeCollision(delay int) {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Plunge Collision",
		AttackTag:  attacks.AttackTagPlunge,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeSlash,
		Element:    attributes.Physical,
		Durability: 0,
		Mult:       collision[c.TalentLvlAttack()],
	}

	// TODO: Check if C6 can expire during plunge
	if c.StatusIsActive(burstBuffKey) {
		ai.Element = attributes.Geo
		ai.IgnoreInfusion = true
	}

	c.Core.QueueAttack(ai, combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, 1), delay, delay)
}
