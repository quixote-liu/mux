package mux

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
)

type Group struct {
	parts           map[string]http.HandlerFunc
	partsForWild    map[string]http.HandlerFunc
	partsForAllWild map[string]http.HandlerFunc
}

func newGroup() *Group {
	return &Group{
		parts:           make(map[string]http.HandlerFunc),
		partsForWild:    make(map[string]http.HandlerFunc),
		partsForAllWild: make(map[string]http.HandlerFunc),
	}
}

func (g *Group) register(pattern string, handler http.HandlerFunc) error {
	pat := filepath.Clean(pattern)

	match, err := g.parsePattern(pat)
	if err != nil {
		return err
	}

	var errPathConflict = fmt.Errorf("the router path(%s) is conflict", pattern)
	switch match {
	case matchNil:
		if g.parts[pat] != nil {
			return errPathConflict
		}
		g.parts[pat] = handler
	case matchPart:
		if g.partsForWild[pat] != nil {
			return errPathConflict
		}
		g.partsForWild[pat] = handler
	case matchAll:
		if g.partsForAllWild[pat] != nil {
			return errPathConflict
		}
		g.partsForAllWild[pat] = handler
	}

	return nil
}

const (
	matchNil = iota + 1
	matchPart
	matchAll
)

func (g *Group) parsePattern(pattern string) (match int, err error) {
	parts := strings.Split(pattern, "/")
	for i, p := range parts {
		if strings.HasPrefix(p, ":") {
			match = matchPart
			return
		}
		if strings.HasPrefix(p, "*") {
			if i != len(parts)-1 {
				err = fmt.Errorf("the router path(%s) with wildcard[*] format error")
				return
			}
			match = matchAll
			return
		}
	}
	match = matchNil
	return
}

func (g *Group) find(path string) http.HandlerFunc {
	
}
