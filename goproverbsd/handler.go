package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type ProverbHandler struct {
	proverbs []string
}

func NewProverbHandler(proverbs []string) *ProverbHandler {
	return &ProverbHandler{
		proverbs: proverbs,
	}
}

// START OMIT
func (p *ProverbHandler) getProverb(preq *ProverbRequest) string {
	if preq.Offset == 0 {
		// if no offset specified, choose random
		preq.Offset = rand.Intn(len(p.proverbs))
	}
	if preq.Offset >= len(p.proverbs) {
		// if offset too large, wrap
		preq.Offset = preq.Offset % len(p.proverbs)
	}
	return p.proverbs[preq.Offset]
}

func (p *ProverbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var preq ProverbRequest
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&preq)
	var pres ProverbResponse
	pres.Proverb = p.getProverb(&preq)
	encoder := json.NewEncoder(w)
	encoder.Encode(&pres)
}

// END OMIT
