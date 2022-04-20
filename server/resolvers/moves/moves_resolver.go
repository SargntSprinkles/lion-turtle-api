package moves

import (
	"embed"
	"encoding/json"
	"io/fs"

	"github.com/sirupsen/logrus"
)

//go:embed data/*.json
var movesJson embed.FS
var moves map[string]*Move = getMoves()

type MoveResolver struct{}

func (mr *MoveResolver) MOVES() []*Move {
	moveslice := []*Move{}
	for _, t := range moves {
		moveslice = append(moveslice, t)
	}
	return moveslice
}

type moveArgs struct {
	Name string
}

func (mr *MoveResolver) MOVE(args moveArgs) *Move {
	return moves[args.Name]
}

type moveTagArgs struct {
	Tags []string
}

func (mr *MoveResolver) MOVESWITHALLTAGS(args moveTagArgs) []*Move {
	moveslice := []*Move{}
	for _, t := range moves {
		if t.HasAllTags(args.Tags) {
			moveslice = append(moveslice, t)
		}
	}
	return moveslice
}

func (mr *MoveResolver) MOVESWITHANYTAGS(args moveTagArgs) []*Move {
	moveslice := []*Move{}
	for _, t := range moves {
		if t.HasAnyTag(args.Tags) {
			moveslice = append(moveslice, t)
		}
	}
	return moveslice
}

func getMoves() map[string]*Move {
	movemap := map[string]*Move{}
	logrus.Info("loading move data from files")
	walkErr := fs.WalkDir(movesJson, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		rawMove, fileReadErr := movesJson.ReadFile(path)
		if fileReadErr != nil {
			return fileReadErr
		}
		newMove := &Move{}
		jsonErr := json.Unmarshal(rawMove, newMove)
		if jsonErr != nil {
			logrus.Fatal(jsonErr)
		}
		movemap[newMove.Name] = newMove
		return nil
	})
	if walkErr != nil {
		logrus.Fatal(walkErr)
	}
	logrus.Infof("loaded %d moves", len(movemap))
	return movemap
}
