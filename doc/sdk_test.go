package doc

import (
	"fmt"
	"testing"
	"time"

	"github.com/unified-tool/go-zinc/doc/schemas"
)

var (
	index = "test-index"
)

func Test_zincDocImpl_DeleteDocument(t *testing.T) {
	sdk, err := NewSDK("http://localhost:4080", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	if err := sdk.DeleteDocument(index, "11111111111111"); err != nil {
		t.Fatal(err)
	}
}

func Test_zincDocImpl_InsertDocument(t *testing.T) {
	sdk, err := NewSDK("http://localhost:4080", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	err = sdk.InsertDocument(index, map[string]string{
		"title": "钢铁侠美国队长复仇者联盟",
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_zincDocImpl_SearchDocuments(t *testing.T) {
	sdk, err := NewSDK("http://localhost:4080", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	resp, err := sdk.SearchDocuments(index, &schemas.SearchRequest{
		SearchType: "matchphrase",
		Query: struct {
			Term      string    `json:"term"`
			StartTime time.Time `json:"start_time"`
			EndTime   time.Time `json:"end_time"`
		}{
			Term:      "美国",
			StartTime: time.Time{},
			EndTime:   time.Time{},
		},
		SortFields: nil,
		From:       0,
		MaxResults: 1000,
		Source:     nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	for _, hit := range resp.Hits.Hits {
		fmt.Println(hit.Source)
	}
}

func Test_zincDocImpl_UpdateDocument(t *testing.T) {
	sdk, err := NewSDK("http://localhost:4080", "admin", "admin")
	if err != nil {
		t.Fatal(err)
	}
	if err := sdk.UpdateDocument(index, "286b4037-9462-40d7-95ed-ea3ddbfaa80b", map[string]string{
		"title": "钢铁侠美国队长复仇者联盟",
	}); err != nil {
		t.Fatal(err)
	}
}
