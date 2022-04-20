package techniques

import "strings"

type Technique struct {
	Name        string
	Approach    string
	Description string
	Tags        []string
}

func (t Technique) NAME() string {
	return t.Name
}

func (t Technique) APPROACH() string {
	return t.Approach
}

func (t Technique) DESCRIPTION() string {
	return t.Description
}

func (t Technique) TAGS() []string {
	return t.Tags
}

func (t *Technique) HasAnyTag(tags []string) bool {
	for _, argtag := range tags {
		for _, t := range t.Tags {
			if strings.EqualFold(t, argtag) {
				return true
			}
		}
	}
	return false
}

func (t *Technique) HasAllTags(tags []string) bool {
	results := map[string]bool{}
	for _, argtag := range tags {
		results[argtag] = false
		for _, t := range t.Tags {
			if strings.EqualFold(t, argtag) {
				results[argtag] = true
			}
		}
	}
	for _, has := range results {
		if !has {
			return false
		}
	}
	return true
}
