package plan

import (
	"regexp"

	"github.com/collection-of-inconceivable-repos/terrafy/internal/app/terrafy/util"
)

type PlaylistDefinition struct {
	Name        string                    `yaml:"name"`
	Description string                    `yaml:"description"`
	Image       string                    `yaml:"image"`
	Tracks      []PlaylistTrackDefinition `yaml:"tracks"`
}

type PlaylistTrackDefinition struct {
	Link   string `yaml:"link"`
	Name   string `yaml:"name"`
	Artist string `yaml:"artist"`
	Album  string `yaml:"album"`
}

var spotifyLinkPattern *regexp.Regexp = regexp.MustCompile(`^https?://open\.spotify\.com/track/(?P<trackId>[a-zA-Z0-9]{22})(?:\?.*)?$`)

func (ptd *PlaylistTrackDefinition) HasSpotifyLink() bool {
	if util.IsBlank(ptd.Link) {
		return false
	}

	return spotifyLinkPattern.MatchString(ptd.Link)
}

// func (ptd *PlaylistTrackDefinition) HasSearchMetadata() bool {

// }
