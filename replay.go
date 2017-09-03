// Package wotreplay provides World_of_Tanks Replay Parser.
package wotreplay

import (
	"encoding/json"
	"log"
)

const magic = 288633362

// Replay stores an entire battle in a replay file
type Replay struct {
	File struct {
		Header struct {
			Magic      uint32
			BlockCount uint32
		}
		Blocks    []dataBlock // len(Blocks) == Header.BlockCount
		ReplayLen uint32
		XXX       [4]byte
		Replay    []byte // binary replay data. len(Replay) == ReplayLen
	}
	MatchStart    MatchStart
	MatchEnd      MatchEnd
	BattleResults BattleResults
	*log.Logger
}

type dataBlock struct {
	Length uint32
	Data   json.RawMessage
}

func (r Replay) String() string {
	b, _ := json.MarshalIndent(struct {
		Date     string
		Map      string
		Gamemode string
		// Result   string
		Reason  string
		Tank    string
		Player  string
		Version string
	}{
		Date:     r.MatchStart.DateTime,
		Map:      r.MatchStart.MapDisplayName,
		Tank:     r.MatchStart.PlayerVehicle,
		Player:   r.MatchStart.PlayerName,
		Gamemode: r.MatchStart.BattleType.String(),
		Reason:   r.BattleResults.Results.Common.FinishReason.String(),
		Version:  r.MatchStart.ClientVersionFromXML,
	}, "", "\t")
	return string(b)
}

type vehicleID json.Number

type vehicle struct {
	ForbidInBattleInvitations bool     `json:"forbidInBattleInvitations"`
	VehicleType               string   `json:"vehicleType"`
	IsAlive                   bool     `json:"isAlive"`
	Name                      string   `json:"name"`
	IgrType                   int      `json:"igrType"`
	PotapovQuestIDs           []int    `json:"potapovQuestIDs"`
	ClanAbbrev                string   `json:"clanAbbrev"`
	Team                      int      `json:"team"`
	Events                    struct{} `json:"events"`
	IsTeamKiller              bool     `json:"isTeamKiller"`
}

type vehicle2 struct {
	Spotted                      int         `json:"spotted"`
	ExtPublic                    struct{}    `json:"extPublic"`
	DamageAssistedTrack          int         `json:"damageAssistedTrack"`
	DamageReceivedFromInvisibles int         `json:"damageReceivedFromInvisibles"`
	DirectTeamHits               int         `json:"directTeamHits"`
	DamageDealt                  int         `json:"damageDealt"`
	PiercingsReceived            int         `json:"piercingsReceived"`
	SniperDamageDealt            int         `json:"sniperDamageDealt"`
	SoloFlagCapture              int         `json:"soloFlagCapture"`
	DamageAssistedRadio          int         `json:"damageAssistedRadio"`
	StunDuration                 float64     `json:"stunDuration"`
	Piercings                    int         `json:"piercings"`
	DamageBlockedByArmor         int         `json:"damageBlockedByArmor"`
	Xp                           int         `json:"xp"`
	InfluencePoints              int         `json:"influencePoints"`
	DroppedCapturePoints         int         `json:"droppedCapturePoints"`
	StopRespawn                  bool        `json:"stopRespawn"`
	DeathCount                   int         `json:"deathCount"`
	Index                        int         `json:"index"`
	DirectHitsReceived           int         `json:"directHitsReceived"`
	KillerID                     int         `json:"killerID"`
	ExplosionHitsReceived        int         `json:"explosionHitsReceived"`
	AchievementXP                int         `json:"achievementXP"`
	DeathReason                  int         `json:"deathReason"`
	CapturePoints                int         `json:"capturePoints"`
	DamageEventList              interface{} `json:"damageEventList"`
	Health                       int         `json:"health"`
	AchievementCredits           int         `json:"achievementCredits"`
	Achievements                 []int       `json:"achievements"`
	Mileage                      int         `json:"mileage"`
	Shots                        int         `json:"shots"`
	Kills                        int         `json:"kills"`
	FortResource                 int         `json:"fortResource"`
	FlagCapture                  int         `json:"flagCapture"`
	Damaged                      int         `json:"damaged"`
	TdamageDealt                 int         `json:"tdamageDealt"`
	ResourceAbsorbed             int         `json:"resourceAbsorbed"`
	Credits                      int         `json:"credits"`
	LifeTime                     int         `json:"lifeTime"`
	NoDamageDirectHitsReceived   int         `json:"noDamageDirectHitsReceived"`
	Stunned                      int         `json:"stunned"`
	AccountDBID                  int         `json:"accountDBID"`
	IsTeamKiller                 bool        `json:"isTeamKiller"`
	TypeCompDescr                int         `json:"typeCompDescr"`
	CapturingBase                interface{} `json:"capturingBase"`
	DamageAssistedStun           int         `json:"damageAssistedStun"`
	RolloutsCount                int         `json:"rolloutsCount"`
	Tkills                       int         `json:"tkills"`
	PotentialDamageReceived      int         `json:"potentialDamageReceived"`
	DamageReceived               int         `json:"damageReceived"`
	FlagActions                  []int       `json:"flagActions"`
	WinPoints                    int         `json:"winPoints"`
	ExplosionHits                int         `json:"explosionHits"`
	Team                         int         `json:"team"`
	TdestroyedModules            interface{} `json:"tdestroyedModules"` // int or bool
	StunNum                      int         `json:"stunNum"`
	AchievementFreeXP            int         `json:"achievementFreeXP"`
	DirectHits                   int         `json:"directHits"`
}
