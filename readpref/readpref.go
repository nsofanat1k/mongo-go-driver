package readpref

import (
	"errors"
	"time"

	"github.com/10gen/mongo-go-driver/server"
)

var (
	errInvalidReadPreference = errors.New("can not specify tags or max staleness on primary")
)

// Primary constructs a read preference with a PrimaryMode.
func Primary() *ReadPref {
	return &ReadPref{mode: PrimaryMode}
}

// PrimaryPreferred constructs a read preference with a PrimaryPreferredMode.
func PrimaryPreferred(opts ...Option) *ReadPref {
	rp, _ := New(PrimaryPreferredMode, opts...)
	return rp
}

// SecondaryPreferred constructs a read preference with a SecondaryPreferredMode.
func SecondaryPreferred(opts ...Option) *ReadPref {
	rp, _ := New(SecondaryPreferredMode, opts...)
	return rp
}

// Secondary constructs a read preference with a SecondaryMode.
func Secondary(opts ...Option) *ReadPref {
	rp, _ := New(SecondaryMode, opts...)
	return rp
}

// Nearest constructs a read preference with a NearestMode.
func Nearest(opts ...Option) *ReadPref {
	rp, _ := New(NearestMode, opts...)
	return rp
}

func New(mode Mode, opts ...Option) (*ReadPref, error) {
	rp := &ReadPref{
		mode: mode,
	}

	if mode == PrimaryMode && len(opts) != 0 {
		return nil, errInvalidReadPreference
	}

	for _, opt := range opts {
		opt(rp)
	}

	return rp, nil
}

// ReadPref determines which servers are considered suitable for read operations.
type ReadPref struct {
	maxStaleness    time.Duration
	maxStalenessSet bool
	mode            Mode
	tagSets         []server.TagSet
}

// MaxStaleness is the maximum amount of time to allow
// a server to be considered eligible for selection. The
// second return value indicates if this value has been set.
func (r *ReadPref) MaxStaleness() (time.Duration, bool) {
	return r.maxStaleness, r.maxStalenessSet
}

// Mode indicates the mode of the read preference.
func (r *ReadPref) Mode() Mode {
	return r.mode
}

// TagSets are multiple tag sets indicating
// which servers should be considered.
func (r *ReadPref) TagSets() []server.TagSet {
	return r.tagSets
}