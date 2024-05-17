package search

import (
	"bufio"
	"bytes"
	"chercher/utils/config"
	"encoding/json"
	"errors"
	"fmt"
	"os/exec"
)

// Todo: Implement inherited type ObsidianSearcher
// Todo: Implement bleve search instead of using ripgrep

type LocalSearcher struct {
	dir         string
	configDir   string
	ripgrepPath string
}

type RipgrepOutput struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}
type RipgrepMatch struct {
	Path struct {
		Text string `json:"text"`
	} `json:"path"`
	Lines struct {
		Text string `json:"text"`
	} `json:"lines"`
	LineNumber     int `json:"line_number"`
	AbsoluteOffset int `json:"absolute_offset"`
	Submatches     []struct {
		Match struct {
			Text string `json:"text"`
		} `json:"match"`
		Start int `json:"start"`
		End   int `json:"end"`
	} `json:"submatches"`
}

func (l *LocalSearcher) Search(query string) ([]Result, error) {
	var result []Result
	cmd := exec.Command(l.ripgrepPath, "-i", query, "--json", l.dir)
	output, _ := cmd.Output()
	s := bufio.NewScanner(bytes.NewReader(output))
	fmt.Println(cmd.String(), string(output))
	var errs []error
	for s.Scan() {
		var v RipgrepOutput
		if err := json.Unmarshal(s.Bytes(), &v); err != nil {
			errs = append(errs, fmt.Errorf("failed to unmarshal ripgrep output: %w", err))
		}
		if v.Type == "match" {
			var match RipgrepMatch
			if err := json.Unmarshal(v.Data, &match); err != nil {
				errs = append(errs, fmt.Errorf("failed to unmarshal ripgrep match: %w", err))
			}
			result = append(result, Result{
				Href:    match.Path.Text,
				Title:   match.Lines.Text,
				Source:  "Local",
				Context: match.Lines.Text,
			})
		}
	}

	if len(errs) > 0 {
		return result, errors.Join(errs...)
	}

	return result, nil
}

func MakeLocalSearcher(searcherConfig config.SearcherConfig) (*LocalSearcher, error) {
	if SearcherType(searcherConfig.Type) != SearcherTypeLocal {
		return nil, fmt.Errorf("searcher type mismatch: %s", searcherConfig.Type)
	}

	dir := ""
	ok := true

	if iDir, found := searcherConfig.Config["dir"]; found {
		if dir, ok = iDir.(string); !ok {
			return nil, fmt.Errorf("config.dir is not a string")
		}
	} else {
		return nil, fmt.Errorf("missing config.dir")
	}

	ripgrepPath := ""
	if iRipgrepPath, found := searcherConfig.Config["ripgrep_path"]; found {
		if ripgrepPath, ok = iRipgrepPath.(string); !ok {
			return nil, fmt.Errorf("config.ripgrep_path is not a string")
		}
	} else {
		return nil, fmt.Errorf("missing config.ripgrep_path")
	}

	return &LocalSearcher{
		dir:         dir,
		configDir:   searcherConfig.GetDataDir(),
		ripgrepPath: ripgrepPath,
	}, nil
}
