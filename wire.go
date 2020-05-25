// Copyright 2020 Seong Yong-ju. All rights reserved.

package main

import "github.com/google/wire"

func InitializeEvent() Event {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}
}
