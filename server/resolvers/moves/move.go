package moves

import "strings"

type Move struct {
	Name        string
	Description string
	Tags        []string
}

func (m Move) NAME() string {
	return m.Name
}

func (m Move) DESCRIPTION() string {
	return m.Description
}

func (m Move) TAGS() []string {
	return m.Tags
}

func (m *Move) HasAnyTag(tags []string) bool {
	for _, argtag := range tags {
		for _, t := range m.Tags {
			if strings.EqualFold(t, argtag) {
				return true
			}
		}
	}
	return false
}

func (m *Move) HasAllTags(tags []string) bool {
	results := map[string]bool{}
	for _, argtag := range tags {
		results[argtag] = false
		for _, t := range m.Tags {
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
