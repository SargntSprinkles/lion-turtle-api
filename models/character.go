package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	CharacterName        string
	Playbook             string
	TrainingWater        bool
	TrainingEarth        bool
	TrainingFire         bool
	TrainingAir          bool
	TrainingWeapons      bool
	TrainingTech         bool
	FightingStyle        string
	BackgroundMilitary   bool
	BackgroundMonastic   bool
	BackgroundOutlaw     bool
	BackgroundPrivileged bool
	BackgroundUrban      bool
	BackgroundWilderness bool
	Hometown             string
	// demeanorOne    string
	// demeanorTwo    string
	// demeanorThree  string
	// demeanorFour   string
	// demeanorFive   string
	// appearance     string
	// historyOne     string
	// historyTwo     string
	// historyThree   string
	// historyFour    string
	// historyFive    string
	// connectionsOne string
	// connectionsTwo string

	// STATS
	// creativity int
	// focus      int
	// harmony    int
	// passion    int

	// CONDITIONS
	// fatigue  int
	// afraid   bool
	// angry    bool
	// guilty   bool
	// insecure bool
	// troubled bool

	// BALANCE
	// center                   int
	// balance                  int
	// momentOfBalanceAvailable bool

	// feature???

	// GROWTH
	// growth                      int
	// growthYourPlaybook          int
	// growthAnotherPlaybook       int
	// growthRaiseStat             int
	// growthShiftCenter           int
	// growthUnlockMomentOfBalance int

	// ABILITIES
	// moves      []playbooks.Move
	// techniques []techniques.Technique
}
