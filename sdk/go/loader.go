package main

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
)

type dictName string

const (
	dictBaseURI = "https://cdn.jsdelivr.net/gh/Reallife/test-mono"

	DictDir1 = "dir_1"
	DictDir2 = "dir_2"
)

type Dict struct {
	Meta struct {
		Name        string `json:"name"`
		Version     string `json:"version"`
		Description string `json:"description"`
	} `json:"meta"`
	Guide []struct {
		ID          string `json:"id"`
		Description string `json:"description"`
	} `json:"guide"`
}

type Loader struct {
	cache map[string]*Dict
}

func (l *Loader) GetDict(ctx context.Context, name dictName, version string) (*Dict, error) {
	cacheKey := string(name) + "@" + version

	if dict, ok := l.cache[cacheKey]; ok {
		return dict, nil
	}

	url := dictBaseURI + "@" + string(name) + "/file/" + version + "/" + string(name) + "/file.json"

	c := http.Client{}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, http.NoBody)

	resp, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	dict := &Dict{}

	if err = json.Unmarshal(content, dict); err != nil {
		return nil, err
	}

	return dict, nil
}
