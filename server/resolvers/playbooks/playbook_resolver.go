package playbooks

import (
	"embed"
	"encoding/json"
	"io/fs"

	"github.com/sirupsen/logrus"
)

//go:embed data/*.json
var playbooksJson embed.FS
var playbooks map[string]*Playbook = getPlaybooks()

type PlaybookResolver struct{}

func (pr *PlaybookResolver) PLAYBOOKS() []*Playbook {
	playbookslice := []*Playbook{}
	for _, pb := range playbooks {
		playbookslice = append(playbookslice, pb)
	}
	return playbookslice
}

func Playbooks() map[string]*Playbook {
	return playbooks
}

type playbookArgs struct {
	Name string
}

func (pr *PlaybookResolver) PLAYBOOK(args playbookArgs) *Playbook {
	return playbooks[args.Name]
}

func getPlaybooks() map[string]*Playbook {
	playbookmap := map[string]*Playbook{}
	logrus.Info("loading playbook data from files")
	walkErr := fs.WalkDir(playbooksJson, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rawPlaybook, fileReadErr := playbooksJson.ReadFile(path)
		if fileReadErr != nil {
			return fileReadErr
		}
		newPlaybook := &Playbook{}
		jsonErr := json.Unmarshal(rawPlaybook, newPlaybook)
		if jsonErr != nil {
			logrus.Fatal(jsonErr)
		}
		playbookmap[newPlaybook.Name] = newPlaybook
		return nil
	})
	if walkErr != nil {
		logrus.Fatal(walkErr)
	}
	logrus.Infof("loaded %d playbooks", len(playbookmap))
	return playbookmap
}
