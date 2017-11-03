package wotreplay

import (
	"encoding/json"
	"fmt"
)

// BattleResults is block data about "Battle Result"
type BattleResults struct {
	Results  battleResults
	Vehicles map[vehicleID]vehicle
	Data     map[vehicleID]struct {
		Flags int `json:"flags"`
	}
}

type battleResults struct {
	ArenaUniqueID int64 `json:"arenaUniqueID"`
	Personal      struct {
		PlayerID personalInfo   `json:"9473"` // Num9473
		Avatar   personalAvatar `json:"avatar"`
	}
	Vehicles map[vehicleID][]vehicle2 `json:"vehicles"`
	Avatars  map[vehicleID]avatar     `json:"avatars"`
	Players  map[vehicleID]player     `json:"players"`
	Common   struct {
		Division            interface{}  `json:"division"`
		FinishReason        finishReason `json:"finishReason"`
		GuiType             int          `json:"guiType"`
		ArenaCreateTime     int          `json:"arenaCreateTime"`
		ArenaTypeID         int          `json:"arenaTypeID"`
		Duration            int          `json:"duration"`
		GasAttackWinnerTeam int          `json:"gasAttackWinnerTeam"`
		WinnerTeam          int          `json:"winnerTeam"`
		VehLockMode         int          `json:"vehLockMode"`
		BonusType           int          `json:"bonusType"`
		Bots                struct{}     `json:"bots"`
	} `json:"common"`
}

type personalInfo struct {
	EventFortResourceList []interface{} `json:"eventFortResourceList"`
	VehTypeLockTime       int           `json:"vehTypeLockTime"`
	Stunned               int           `json:"stunned"`
	CreditsToDraw         int           `json:"creditsToDraw"`
	OrderFreeXPFactor100  int           `json:"orderFreeXPFactor100"`
	OrderXPFactor100      int           `json:"orderXPFactor100"`
	DamageAssistedRadio   int           `json:"damageAssistedRadio"`
	StunDuration          float64       `json:"stunDuration"`
	FreeXPReplay          interface{}   `json:"freeXPReplay"`
	WinPoints             int           `json:"winPoints"`
	QuestsProgress        struct {
		WGSSA18340Q00 []interface{} `json:"WGSSA-18340_q00"`
	} `json:"questsProgress"`
	StopRespawn                     bool                 `json:"stopRespawn"`
	CreditsContributionIn           int                  `json:"creditsContributionIn"`
	EventCredits                    int                  `json:"eventCredits"`
	EventXPList                     []interface{}        `json:"eventXPList"`
	XpReplay                        interface{}          `json:"xpReplay"`
	AchievementXP                   int                  `json:"achievementXP"`
	IgrXPFactor10                   int                  `json:"igrXPFactor10"`
	AogasFactor10                   int                  `json:"aogasFactor10"`
	OriginalCreditsContributionIn   int                  `json:"originalCreditsContributionIn"`
	OriginalCreditsPenalty          int                  `json:"originalCreditsPenalty"`
	DamagedWhileMoving              int                  `json:"damagedWhileMoving"`
	Kills                           int                  `json:"kills"`
	EventTMenXP                     int                  `json:"eventTMenXP"`
	PercentFromTotalTeamDamage      float64              `json:"percentFromTotalTeamDamage"`
	OriginalTMenXP                  int                  `json:"originalTMenXP"`
	MarkOfMastery                   int                  `json:"markOfMastery"`
	NoDamageDirectHitsReceived      int                  `json:"noDamageDirectHitsReceived"`
	BoosterCredits                  int                  `json:"boosterCredits"`
	OriginalGold                    int                  `json:"originalGold"`
	EventFreeXPList                 []interface{}        `json:"eventFreeXPList"`
	Tkills                          int                  `json:"tkills"`
	Index                           int                  `json:"index"`
	Shots                           int                  `json:"shots"`
	Team                            int                  `json:"team"`
	DeathCount                      int                  `json:"deathCount"`
	EventTMenXPFactor100List        [][]interface{}      `json:"eventTMenXPFactor100List"`
	DirectHits                      int                  `json:"directHits"`
	Spotted                         int                  `json:"spotted"`
	ExtPublic                       struct{}             `json:"extPublic"`
	DamageReceivedFromInvisibles    int                  `json:"damageReceivedFromInvisibles"`
	BoosterCreditsFactor100         int                  `json:"boosterCreditsFactor100"`
	PremiumCreditsFactor10          int                  `json:"premiumCreditsFactor10"`
	SoloFlagCapture                 int                  `json:"soloFlagCapture"`
	OrderFortResource               int                  `json:"orderFortResource"`
	MarksOnGun                      int                  `json:"marksOnGun"`
	PremiumVehicleXPFactor100       int                  `json:"premiumVehicleXPFactor100"`
	FactualXP                       int                  `json:"factualXP"`
	KilledAndDamagedByAllSquadmates int                  `json:"killedAndDamagedByAllSquadmates"`
	EventFreeXP                     int                  `json:"eventFreeXP"`
	EventGoldFactor100List          []interface{}        `json:"eventGoldFactor100List"`
	CreditsContributionOut          int                  `json:"creditsContributionOut"`
	DamageEventList                 interface{}          `json:"damageEventList"`
	Health                          int                  `json:"health"`
	Achievements                    []int                `json:"achievements"`
	OrderFreeXP                     int                  `json:"orderFreeXP"`
	EventGoldList                   []interface{}        `json:"eventGoldList"`
	BoosterTMenXPFactor100          int                  `json:"boosterTMenXPFactor100"`
	DossierPopUps                   [][]interface{}      `json:"dossierPopUps"`
	TdamageDealt                    int                  `json:"tdamageDealt"`
	ResourceAbsorbed                int                  `json:"resourceAbsorbed"`
	GoldReplay                      interface{}          `json:"goldReplay"`
	AutoEquipCost                   []int                `json:"autoEquipCost"`
	OriginalXP                      int                  `json:"originalXP"`
	Credits                         int                  `json:"credits"`
	DamagedWhileEnemyMoving         int                  `json:"damagedWhileEnemyMoving"`
	CreditsPenalty                  int                  `json:"creditsPenalty"`
	DamageDealt                     int                  `json:"damageDealt"`
	PercentFromSecondBestDamage     float64              `json:"percentFromSecondBestDamage"`
	CommittedSuicide                bool                 `json:"committedSuicide"`
	BoosterXP                       int                  `json:"boosterXP"`
	LifeTime                        int                  `json:"lifeTime"`
	EventTMenXPList                 []interface{}        `json:"eventTMenXPList"`
	DailyXPFactor10                 int                  `json:"dailyXPFactor10"`
	DamageRating                    int                  `json:"damageRating"`
	Repair                          int                  `json:"repair"`
	OriginalCredits                 int                  `json:"originalCredits"`
	DamageAssistedTrack             int                  `json:"damageAssistedTrack"`
	XpPenalty                       int                  `json:"xpPenalty"`
	XpByTmen                        [][]interface{}      `json:"xpByTmen"`
	SniperDamageDealt               int                  `json:"sniperDamageDealt"`
	FairplayFactor10                int                  `json:"fairplayFactor10"`
	OrderCreditsFactor100           int                  `json:"orderCreditsFactor100"`
	SubtotalTMenXP                  int                  `json:"subtotalTMenXP"`
	DamageBlockedByArmor            int                  `json:"damageBlockedByArmor"`
	Xp                              int                  `json:"xp"`
	BoosterXPFactor100              int                  `json:"boosterXPFactor100"`
	KillerID                        int                  `json:"killerID"`
	RefSystemXPFactor10             int                  `json:"refSystemXPFactor10"`
	OrderTMenXP                     int                  `json:"orderTMenXP"`
	FlagActions                     []int                `json:"flagActions"`
	OriginalXPPenalty               int                  `json:"originalXPPenalty"`
	OrderTMenXPFactor100            int                  `json:"orderTMenXPFactor100"`
	EventXPFactor100List            []interface{}        `json:"eventXPFactor100List"`
	OriginalFortResource            int                  `json:"originalFortResource"`
	SubtotalXP                      int                  `json:"subtotalXP"`
	SquadXP                         int                  `json:"squadXP"`
	OriginalCreditsContributionOut  int                  `json:"originalCreditsContributionOut"`
	OriginalFreeXP                  int                  `json:"originalFreeXP"`
	OrderCredits                    int                  `json:"orderCredits"`
	FreeXP                          int                  `json:"freeXP"`
	OrderXP                         int                  `json:"orderXP"`
	PremiumVehicleXP                int                  `json:"premiumVehicleXP"`
	FlagCapture                     int                  `json:"flagCapture"`
	EventCreditsList                []interface{}        `json:"eventCreditsList"`
	EventGold                       int                  `json:"eventGold"`
	Gold                            int                  `json:"gold"`
	FortResourceReplay              interface{}          `json:"fortResourceReplay"`
	EventXP                         int                  `json:"eventXP"`
	FactualCredits                  int                  `json:"factualCredits"`
	AutoLoadCost                    []int                `json:"autoLoadCost"`
	SubtotalFreeXP                  int                  `json:"subtotalFreeXP"`
	StunNum                         int                  `json:"stunNum"`
	AchievementFreeXP               int                  `json:"achievementFreeXP"`
	SubtotalCredits                 int                  `json:"subtotalCredits"`
	KillsBeforeTeamWasDamaged       int                  `json:"killsBeforeTeamWasDamaged"`
	BoosterTMenXP                   int                  `json:"boosterTMenXP"`
	PotentialDamageReceived         int                  `json:"potentialDamageReceived"`
	DirectTeamHits                  int                  `json:"directTeamHits"`
	DamageReceived                  int                  `json:"damageReceived"`
	PiercingsReceived               int                  `json:"piercingsReceived"`
	MovingAvgDamage                 int                  `json:"movingAvgDamage"`
	PremiumXPFactor10               int                  `json:"premiumXPFactor10"`
	CreditsReplay                   interface{}          `json:"creditsReplay"`
	Piercings                       int                  `json:"piercings"`
	PrevMarkOfMastery               int                  `json:"prevMarkOfMastery"`
	EventFreeXPFactor100List        []interface{}        `json:"eventFreeXPFactor100List"`
	ServiceProviderID               int                  `json:"serviceProviderID"`
	InfluencePoints                 int                  `json:"influencePoints"`
	DroppedCapturePoints            int                  `json:"droppedCapturePoints"`
	PersonalFortResource            int                  `json:"personalFortResource"`
	EventFortResourceFactor100List  []interface{}        `json:"eventFortResourceFactor100List"`
	DirectHitsReceived              int                  `json:"directHitsReceived"`
	TypeCompDescr                   int                  `json:"typeCompDescr"`
	DeathReason                     int                  `json:"deathReason"`
	CapturePoints                   int                  `json:"capturePoints"`
	DamageBeforeTeamWasDamaged      int                  `json:"damageBeforeTeamWasDamaged"`
	ExplosionHitsReceived           int                  `json:"explosionHitsReceived"`
	EventFortResource               int                  `json:"eventFortResource"`
	Details                         map[vehicleID]detail `json:"details"`
	SquadXPFactor100                int                  `json:"squadXPFactor100"`
	AchievementCredits              int                  `json:"achievementCredits"`
	OriginalCreditsToDraw           int                  `json:"originalCreditsToDraw"`
	IsPremium                       bool                 `json:"isPremium"`
	Mileage                         int                  `json:"mileage"`
	ExplosionHits                   int                  `json:"explosionHits"`
	RolloutsCount                   int                  `json:"rolloutsCount"`
	FortResource                    int                  `json:"fortResource"`
	AvatarDamageEventList           interface{}          `json:"avatarDamageEventList"`
	SubtotalGold                    int                  `json:"subtotalGold"`
	AppliedPremiumCreditsFactor10   int                  `json:"appliedPremiumCreditsFactor10"`
	Damaged                         int                  `json:"damaged"`
	AccountDBID                     int                  `json:"accountDBID"`
	OrderFortResourceFactor100      int                  `json:"orderFortResourceFactor100"`
	TmenXPReplay                    interface{}          `json:"tmenXPReplay"`
	AutoRepairCost                  int                  `json:"autoRepairCost"`
	EventCreditsFactor100List       []interface{}        `json:"eventCreditsFactor100List"`
	IsTeamKiller                    bool                 `json:"isTeamKiller"`
	TmenXP                          int                  `json:"tmenXP"`
	FactualFreeXP                   int                  `json:"factualFreeXP"`
	CapturingBase                   interface{}          `json:"capturingBase"`
	DamageAssistedStun              int                  `json:"damageAssistedStun"`
	AppliedPremiumXPFactor10        int                  `json:"appliedPremiumXPFactor10"`
	BoosterFreeXPFactor100          int                  `json:"boosterFreeXPFactor100"`
	SubtotalFortResource            int                  `json:"subtotalFortResource"`
	BoosterFreeXP                   int                  `json:"boosterFreeXP"`
	TdestroyedModules               int                  `json:"tdestroyedModules"`
	BattleNum                       int                  `json:"battleNum"`
}

type personalAvatar struct {
	CreditsReplay      interface{} `json:"creditsReplay"`
	FreeXPReplay       interface{} `json:"freeXPReplay"`
	Xp                 int         `json:"xp"`
	FairplayViolations []int       `json:"fairplayViolations"`
	EventFreeXP        int         `json:"eventFreeXP"`
	EventCredits       int         `json:"eventCredits"`
	XpReplay           interface{} `json:"xpReplay"`
	EnemyClub          struct{}    `json:"enemyClub"`
	DamageEventList    interface{} `json:"damageEventList"`
	EventFortResource  int         `json:"eventFortResource"`
	IsPrematureLeave   bool        `json:"isPrematureLeave"`
	SquadBonusInfo     interface{} `json:"squadBonusInfo"`
	WinnerIfDraw       int         `json:"winnerIfDraw"`
	FreeXP             int         `json:"freeXP"`
	AvatarKills        int         `json:"avatarKills"`
	EventTMenXP        int         `json:"eventTMenXP"`
	Club               struct {
		Achievements []interface{} `json:"achievements"`
	} `json:"club"`
	AvatarDamageEventList interface{}   `json:"avatarDamageEventList"`
	AvatarDamageDealt     int           `json:"avatarDamageDealt"`
	TotalDamaged          int           `json:"totalDamaged"`
	GoldReplay            interface{}   `json:"goldReplay"`
	EventGold             int           `json:"eventGold"`
	TmenXPReplay          interface{}   `json:"tmenXPReplay"`
	QuestsProgress        struct{}      `json:"questsProgress"`
	AccountDBID           int           `json:"accountDBID"`
	AvatarAmmo            []interface{} `json:"avatarAmmo"`
	Credits               int           `json:"credits"`
	FortResourceReplay    interface{}   `json:"fortResourceReplay"`
	EventXP               int           `json:"eventXP"`
	AvatarDamaged         int           `json:"avatarDamaged"`
	Team                  int           `json:"team"`
	ClanDBID              int           `json:"clanDBID"`
	FortClanDBIDs         []interface{} `json:"fortClanDBIDs"`
	FortBuilding          interface{}   `json:"fortBuilding"`
}

type avatar struct {
	FairplayViolations []int `json:"fairplayViolations"`
	TotalDamaged       int   `json:"totalDamaged"`
	AvatarKills        int   `json:"avatarKills"`
	AvatarDamaged      int   `json:"avatarDamaged"`
	AvatarDamageDealt  int   `json:"avatarDamageDealt"`
}

type player struct {
	Name        string `json:"name"`
	PrebattleID int    `json:"prebattleID"`
	IgrType     int    `json:"igrType"`
	ClanAbbrev  string `json:"clanAbbrev"`
	ClanDBID    int    `json:"clanDBID"`
	Team        int    `json:"team"`
}

type detail struct {
	Spotted                    int     `json:"spotted"`
	Crits                      int     `json:"crits"`
	DamageAssistedTrack        int     `json:"damageAssistedTrack"`
	DamageAssistedStun         int     `json:"damageAssistedStun"`
	Fire                       int     `json:"fire"`
	DeathReason                int     `json:"deathReason"`
	DamageReceived             int     `json:"damageReceived"`
	DamageDealt                int     `json:"damageDealt"`
	DamageAssistedRadio        int     `json:"damageAssistedRadio"`
	RickochetsReceived         int     `json:"rickochetsReceived"`
	Piercings                  int     `json:"piercings"`
	ExplosionHits              int     `json:"explosionHits"`
	DamageBlockedByArmor       int     `json:"damageBlockedByArmor"`
	NoDamageDirectHitsReceived int     `json:"noDamageDirectHitsReceived"`
	TargetKills                int     `json:"targetKills"`
	StunDuration               float64 `json:"stunDuration"`
	StunNum                    int     `json:"stunNum"`
	DirectHits                 int     `json:"directHits"`
}

// reason a match finished
type finishReason int

func (x finishReason) String() string {
	msg := []string{
		0: "Unknown",       // not used ?
		1: "Extermination", // 殲滅
		2: "Base Capture",  // 陣地占領
		3: "Timeout",       // 時間切れ
		4: "*Failure",      // *アリーナの初期化ミスあるいは内部エラー
		5: "*Technical",    // *サーバーの再起動あるいはキャンセル
		//9: "?",		    // Boot camp [8-9]
	}
	i := int(x)
	if !(0 <= i && i <= len(msg)) {
		return fmt.Sprintf("%d:???", i)
	}
	return fmt.Sprintf("%d:%s", int(x), msg[int(x)])
}

func (x finishReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
