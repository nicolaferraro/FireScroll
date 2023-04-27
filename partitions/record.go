package partitions

import "time"

type RecordKey struct {
	Pk string `validate:"require"`
	Sk string `validate:"require"`
}

type Record struct {
	Pk   string
	Sk   string
	Data map[string]any

	CreatedAt time.Time
	UpdatedAt time.Time
}

type RecordMutation struct {
	Pk       string `validate:"require"`
	Sk       string `validate:"require"`
	Data     *map[string]any
	Mutation Operation
}