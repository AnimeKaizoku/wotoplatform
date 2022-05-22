package wotoGroups_test

import (
	"testing"
	"wp-server/wotoPacks/database/groupsDatabase"
)

func TestGenerateGroupId(t *testing.T) {
	for i := 0; i < 10; i++ {
		id := groupsDatabase.GetNewGroupId()
		if id == "" {
			t.Error("GenerateGroupId() returned empty string")
			return
		}
	}
}
