package techniques

import (
	"embed"
	"encoding/json"
	"io/fs"

	"github.com/sirupsen/logrus"
)

//go:embed data/*.json
var techniquesJson embed.FS
var techniques map[string]*Technique = getTechniques()

type TechniqueResolver struct{}

func (tr *TechniqueResolver) TECHNIQUES() []*Technique {
	techniqueslice := []*Technique{}
	for _, t := range techniques {
		techniqueslice = append(techniqueslice, t)
	}
	return techniqueslice
}

type techniqueArgs struct {
	Name string
}

func (tr *TechniqueResolver) TECHNIQUE(args techniqueArgs) *Technique {
	return techniques[args.Name]
}

type techniqueTagArgs struct {
	Tags []string
}

func (tr *TechniqueResolver) TECHNIQUESWITHALLTAGS(args techniqueTagArgs) []*Technique {
	techniqueslice := []*Technique{}
	for _, t := range techniques {
		if t.HasAllTags(args.Tags) {
			techniqueslice = append(techniqueslice, t)
		}
	}
	return techniqueslice
}

func (tr *TechniqueResolver) TECHNIQUESWITHANYTAGS(args techniqueTagArgs) []*Technique {
	techniqueslice := []*Technique{}
	for _, t := range techniques {
		if t.HasAnyTag(args.Tags) {
			techniqueslice = append(techniqueslice, t)
		}
	}
	return techniqueslice
}

func getTechniques() map[string]*Technique {
	techniquemap := map[string]*Technique{}
	logrus.Info("loading technique data from files")
	walkErr := fs.WalkDir(techniquesJson, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rawTechnique, fileReadErr := techniquesJson.ReadFile(path)
		if fileReadErr != nil {
			return fileReadErr
		}
		newTechnique := &Technique{}
		jsonErr := json.Unmarshal(rawTechnique, newTechnique)
		if jsonErr != nil {
			logrus.Fatal(jsonErr)
		}
		techniquemap[newTechnique.Name] = newTechnique
		return nil
	})
	if walkErr != nil {
		logrus.Fatal(walkErr)
	}
	logrus.Infof("loaded %d techniques", len(techniquemap))
	return techniquemap
}
