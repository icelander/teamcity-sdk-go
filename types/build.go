package types

import (
	"fmt"
)

// Build represents a TeamCity build, along with its metadata.
type Build struct {
	ID          int64
	BuildTypeID string
	BuildType   struct {
		ID          string
		Name        string
		Description string
		ProjectName string
		ProjectID   string
		HREF        string
		WebURL      string
	}
	Triggered struct {
		Type string
		Date JSONTime
		User struct {
			Username string
		}
	}
	Changes struct {
		Change []Change
	}

	QueuedDate    JSONTime
	QueuePosition int64
	StartDate     JSONTime
	FinishDate    JSONTime
	Number        string
	Status        string
	StatusText    string
	State         string
	BranchName    string
	Personal      bool
	Running       bool
	Pinned        bool
	DefaultBranch bool
	HREF          string
	WebURL        string
	Agent         struct {
		ID     int64
		Name   string
		TypeID int64
		HREF   string
	}

	ProblemOccurrences struct {
		ProblemOccurrence []ProblemOccurrence
	}

	TestOccurrences struct {
		TestOccurrence []TestOccurrence
	}

	Tags []string `json:"tags.tag,omitempty"`

	Properties Properties `json:"properties"`
}


func (b *Build) String() string {
	return fmt.Sprintf("Build %d, %#v state=%s", b.ID, b.ComputedState(), b.State)
}

type State int

const (
	Unknown = State(iota)
	Queued
	Started
	Finished
)

func (b *Build) ComputedState() State {
	if b.QueuedDate == "" {
		return Unknown
	}
	if b.StartDate == "" {
		return Queued
	}
	if b.FinishDate == "" {
		return Started
	}
	return Finished
}