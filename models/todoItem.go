package models

import (
	"time"
)

type Todo struct {
	ID		int32
	Title		string
	Created		time.Time
	CreatedBy	int32
	Assigned	[]int32
	Status		Status
	Deleted		bool
	List		int32
}

type State int 

const (
	Incomplete State = iota
	In_progress
	Completed
)

type Status struct {
	State 	State
	Prior	Status
	Time	time.Time
	User	int32
}
