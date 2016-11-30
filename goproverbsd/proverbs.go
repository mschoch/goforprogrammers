package main

import (
	"bufio"
	"io"
	"strings"
)

// START OMIT
func loadProverbs(r io.Reader) ([]string, error) {
	var rv []string
	br := bufio.NewReader(r)
	line, err := br.ReadString('\n')
	for err == nil {
		rv = append(rv, strings.TrimSpace(line))
		line, err = br.ReadString('\n')
	}
	if err != nil && err != io.EOF {
		return nil, err
	}
	return rv, nil
}

// END OMIT
