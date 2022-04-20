package playbooks

type Playbook struct {
	Name               string
	Source             string
	Principles         []string
	StartingCreativity int32
	StartingFocus      int32
	StartingHarmony    int32
	StartingPassion    int32
	DemeanorOptions    []string
	HistoryQuestions   []string
	FeatureName        string
	FeatureDescription string
	Connections        []string
	MomentOfBalance    string
	GrowthQuestion     string
	Moves              []string
	Technique          string
}

func (p Playbook) NAME() string {
	return p.Name
}

func (p Playbook) SOURCE() string {
	return p.Source
}

func (p Playbook) PRINCIPLES() []string {
	return p.Principles
}

func (p Playbook) STARTINGCREATIVITY() int32 {
	return p.StartingCreativity
}

func (p Playbook) STARTINGFOCUS() int32 {
	return p.StartingFocus
}

func (p Playbook) STARTINGHARMONY() int32 {
	return p.StartingHarmony
}

func (p Playbook) STARTINGPASSION() int32 {
	return p.StartingPassion
}

func (p Playbook) DEMEANOROPTIONS() []string {
	return p.DemeanorOptions
}

func (p Playbook) HISTORYQUESTIONS() []string {
	return p.HistoryQuestions
}

func (p Playbook) FEATURENAME() string {
	return p.FeatureName
}

func (p Playbook) FEATUREDESCRIPTION() string {
	return p.FeatureDescription
}

func (p Playbook) CONNECTIONS() []string {
	return p.Connections
}

func (p Playbook) MOMENTOFBALANCE() string {
	return p.MomentOfBalance
}

func (p Playbook) GROWTHQUESTION() string {
	return p.GrowthQuestion
}

func (p Playbook) MOVES() []string {
	return p.Moves
}

func (p Playbook) TECHNIQUE() string {
	return p.Technique
}
