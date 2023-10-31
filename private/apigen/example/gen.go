// Copyright (C) 2022 Storj Labs, Inc.
// See LICENSE for copying information.

//go:build ignore
// +build ignore

package main

import (
	"time"

	"storj.io/common/uuid"

	"storj.io/storj/private/apigen"
	"storj.io/storj/private/apigen/example/myapi"
)

func main() {
	a := &apigen.API{PackageName: "example", Version: "v0", BasePath: "/api"}

	g := a.Group("Documents", "docs")

	g.Get("/", &apigen.Endpoint{
		Name:           "Get Documents",
		Description:    "Get the paths to all the documents under the specified paths",
		GoName:         "Get",
		TypeScriptName: "get",
		Response: []struct {
			ID             uuid.UUID      `json:"id"`
			Path           string         `json:"path"`
			Date           time.Time      `json:"date"`
			Metadata       myapi.Metadata `json:"metadata"`
			LastRetrievals []struct {
				User string    `json:"user"`
				When time.Time `json:"when"`
			} `json:"last_retrievals"`
		}{},
		ResponseMock: []struct {
			ID             uuid.UUID      `json:"id"`
			Path           string         `json:"path"`
			Date           time.Time      `json:"date"`
			Metadata       myapi.Metadata `json:"metadata"`
			LastRetrievals []struct {
				User string    `json:"user"`
				When time.Time `json:"when"`
			} `json:"last_retrievals"`
		}{{
			ID:   uuid.UUID{},
			Path: "/workspace/notes.md",
			Metadata: myapi.Metadata{
				Owner: "Storj",
				Tags:  [][2]string{{"category", "general"}},
			},
			LastRetrievals: []struct {
				User string    `json:"user"`
				When time.Time `json:"when"`
			}{{
				User: "Storj",
				When: time.Now().Add(-time.Hour),
			}},
		}},
	})

	g.Get("/{path}", &apigen.Endpoint{
		Name:           "Get One",
		Description:    "Get the document in the specified path",
		GoName:         "GetOne",
		TypeScriptName: "getOne",
		Response:       myapi.Document{},
		PathParams: []apigen.Param{
			apigen.NewParam("path", ""),
		},
		ResponseMock: myapi.Document{
			ID:        uuid.UUID{},
			Date:      time.Now().Add(-24 * time.Hour),
			PathParam: "ID",
			Body:      "## Notes",
			Version: myapi.Version{
				Date:   time.Now().Add(-30 * time.Minute),
				Number: 1,
			},
		},
	})

	g.Get("/{path}/tag/{tagName}", &apigen.Endpoint{
		Name:           "Get a tag",
		Description:    "Get the tag of the document in the specified path and tag label ",
		GoName:         "GetTag",
		TypeScriptName: "getTag",
		Response:       [2]string{},
		PathParams: []apigen.Param{
			apigen.NewParam("path", ""),
			apigen.NewParam("tagName", ""),
		},
		ResponseMock: [2]string{"category", "notes"},
	})

	g.Get("/{path}/versions", &apigen.Endpoint{
		Name:           "Get Version",
		Description:    "Get all the version of the document in the specified path",
		GoName:         "GetVersions",
		TypeScriptName: "getVersions",
		Response:       []myapi.Version{},
		PathParams: []apigen.Param{
			apigen.NewParam("path", ""),
		},
		ResponseMock: []myapi.Version{
			{Date: time.Now().Add(-360 * time.Hour), Number: 1},
			{Date: time.Now().Add(-5 * time.Hour), Number: 2},
		},
	})

	g.Post("/{path}", &apigen.Endpoint{
		Name:           "Update Content",
		Description:    "Update the content of the document with the specified path and ID if the last update is before the indicated date",
		GoName:         "UpdateContent",
		TypeScriptName: "updateContent",
		Response: struct {
			ID        uuid.UUID `json:"id"`
			Date      time.Time `json:"date"`
			PathParam string    `json:"pathParam"`
			Body      string    `json:"body"`
		}{},
		Request: struct {
			Content string `json:"content"`
		}{},
		QueryParams: []apigen.Param{
			apigen.NewParam("id", uuid.UUID{}),
			apigen.NewParam("date", time.Time{}),
		},
		PathParams: []apigen.Param{
			apigen.NewParam("path", ""),
		},
		ResponseMock: struct {
			ID        uuid.UUID `json:"id"`
			Date      time.Time `json:"date"`
			PathParam string    `json:"pathParam"`
			Body      string    `json:"body"`
		}{
			ID:        uuid.UUID{},
			Date:      time.Now(),
			PathParam: "ID",
			Body:      "## Notes\n### General",
		},
	})

	g = a.Group("Users", "users")

	g.Get("/", &apigen.Endpoint{
		Name:           "Get Users",
		Description:    "Get the list of registered users",
		GoName:         "Get",
		TypeScriptName: "get",
		Response: []struct {
			Name    string `json:"name"`
			Surname string `json:"surname"`
			Email   string `json:"email"`
		}{},
		ResponseMock: []struct {
			Name    string `json:"name"`
			Surname string `json:"surname"`
			Email   string `json:"email"`
		}{
			{Name: "Storj", Surname: "Labs", Email: "storj@storj.test"},
			{Name: "Test1", Surname: "Testing", Email: "test1@example.test"},
			{Name: "Test2", Surname: "Testing", Email: "test2@example.test"},
		},
	})

	g.Post("/", &apigen.Endpoint{
		Name:           "Create User",
		Description:    "Create a user",
		GoName:         "Create",
		TypeScriptName: "create",
		Request: []struct {
			Name    string `json:"name"`
			Surname string `json:"surname"`
			Email   string `json:"email"`
		}{},
	})

	a.MustWriteGo("api.gen.go")
	a.MustWriteTS("client-api.gen.ts")
	a.MustWriteTSMock("client-api-mock.gen.ts")
	a.MustWriteDocs("apidocs.gen.md")
}
