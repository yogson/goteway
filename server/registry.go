package server

import "github.com/yogson/goteway/sdk"

var HandlerFunctions = map[string]sdk.IHandler{}


func Register(h sdk.IHandler) {
	HandlerFunctions[h.GetName()] = h
}