// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package fpath

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strings"
)

// FPath creates an OS independently Path Handling Structure
type FPath struct {
	local  bool   //set if file is local
	scheme string //url scheme
	bucket string //set, when remote scheme
	path   string //filepath OR remote path (without bucket)
}

// New creates Struct from handed over URL
func New(url string) (p FPath, err error) {

	// Check for Schema
	sepstring := strings.SplitN(url, "://", 2)

	switch len(sepstring) {
	case 2: // Has Schema
		p.scheme = sepstring[0]
		if p.scheme != "sj" {
			return FPath{}, errors.New("unsupported URL scheme")
		}
		//Trim initial / of the path and clean it, afterwards split on first /
		split := strings.SplitN(path.Clean(strings.TrimLeft(sepstring[1], "/")), "/", 2)
		p.bucket = split[0]
		if len(split) == 2 {
			p.path = filepath.ToSlash(split[1])
		}
	case 1: // No Scheme
		p.local = true
		p.path = sepstring[0]
	default: // Everything else is malformed
		return FPath{}, fmt.Errorf("malformed URL: %s", url)
	}

	// Check for Windows Special Handling Prefix
	cprefix, err := regexp.Compile(`^\\\\\?\\(UNC\\)?`)
	if err != nil {
		return FPath{}, err
	}

	// when Prefix present, omit further changes to the path
	if prefix := cprefix.FindString(p.path); prefix != "" {

		p.scheme = prefix
		p.path = strings.Replace(p.path, prefix, "", 1) //Strip Prefix
		return p, nil
	}

	// when file is local, ensure path absolute
	if p.IsLocal() && !filepath.IsAbs(p.path) {

		fullpath, err := filepath.Abs(p.path)
		if err != nil {
			return FPath{}, fmt.Errorf("unable to create absolute path for: %s", p.path)
		}

		p.path = fullpath
	}
	return p, nil
}

// Join is merging segment to the path
func (p *FPath) Join(segment string) *FPath {
	p.path = filepath.Join(p.path, segment)
	if !p.local {
		p.path = filepath.ToSlash(p.path)
	}
	return p
}

// Folder returns the parent folder of path
func (p FPath) Folder() string {
	return filepath.Dir(p.path)
}

// IsFolder returns if path is a folder
func (p FPath) IsFolder() bool {
	fileInfo, err := os.Stat(p.path)
	if err != nil {
		//fmt.Println(err)
		return false
	}
	return fileInfo.IsDir()
}

// Base returns Base Segment of the path
func (p FPath) Base() string {
	return filepath.Base(p.path)
}

// Bucket returns first segment of path
func (p FPath) Bucket() string {
	return p.bucket
}

// BucketPath returns Path including prefixed bucket
func (p FPath) BucketPath() string {
	if !p.local && p.bucket != "" {
		return p.bucket + "/" + p.path
	}
	return ""
}

// Path returns the URL path without schema
func (p FPath) Path() string {
	return p.path
}

// IsLocal returns whether URL refers to local or remote location
func (p FPath) IsLocal() bool {
	return p.local
}

// HasScheme returns if URL had a schema
func (p FPath) HasScheme() bool {
	return p.scheme != ""
}

// Scheme returns the scheme for URL-s
func (p FPath) Scheme() string {
	return p.scheme
}

// String returns entire URL
func (p FPath) String() string {
	if p.HasScheme() {
		return p.scheme + "://" + p.BucketPath()
	}
	return p.path
}
