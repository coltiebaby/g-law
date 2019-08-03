package matches

import (
	"context"
	"encoding/json"

	"github.com/coltiebaby/g-law/riot"
	"github.com/coltiebaby/g-law/riot/v4"
)

type Match struct {
	Frames []Frame `json:frames`
}

// Get only the events you want.
func (m *Match) Filter(by ...string) {
	for i, frame := range m.Frames {
		events := filter(frame.Events, by)
		frame.Events = events
		m.Frames[i] = frame
	}
}

type Frame struct {
	Events []Event `json:"events"`
}

type Event struct {
	Timestamp int    `json:"timestamp"`
	ItemID    int    `json:"itemId,omitempty"`
	KillerID  int    `json:"killerId,omitempty"`
	VictimID  int    `json:"victimId,omitempty"`
	Type      string `json:"type"`
}

func GetTimeline(ctx context.Context, match_id string) (*Match, error) {
	match := &Match{}

	req := riot.RiotRequest{
		Type:    "match",
		Uri:     "timelines/by-match/" + match_id,
		Version: v4.VERSION,
	}

	resp, _ := req.GetData()
	err := json.NewDecoder(resp.Body).Decode(match)
	if err != nil {
		return match, err
	}

	return match, nil
}