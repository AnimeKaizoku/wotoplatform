package wotoUsers_test

import (
	"encoding/json"
	"log"
	"testing"
	"wp-server/wotoPacks/core/wotoValues"
	"wp-server/wotoPacks/wotoActions/wotoUsers"

	ws "github.com/ALiwoto/StrongStringGo/strongStringGo"
)

func TestGetUserLikedListResult(t *testing.T) {
	var allList []*wotoValues.LikedListElement
	var i int64

	for i = 0; i < 10; i++ {
		allList = append(allList, &wotoValues.LikedListElement{
			UniqueId: "uniqueID(testing)-" + ws.ToBase10(i),
			OwnerId:  3498435894358435989,
		})
	}

	result := &wotoUsers.GetUserLikedListResult{
		LikedList: allList,
	}
	b, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.Error(err)
		return
	}

	log.Println(string(b))

	log.Println("")
}
