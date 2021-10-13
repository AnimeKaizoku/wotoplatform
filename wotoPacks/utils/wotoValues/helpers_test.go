package wotoValues_test

import (
	"log"
	"strconv"
	"strings"
	"testing"
	"wp-server/wotoPacks/utils/wotoValues"
)

func TestMakeSure1(t *testing.T) {
	b := []byte{140, 100, 20, 50, 60}
	b = wotoValues.MakeSureByte(b, 2)
	log.Println(b)
	if len(b) != 2 {
		t.Errorf("length b was %d, expected %d", len(b), 2)
	}
}

func TestMakeSure2(t *testing.T) {
	b := []byte("005")
	b = wotoValues.MakeSureByte(b, 8)
	log.Println("\"" + string(b) + "\"")

	myI, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		t.Errorf("got an error while trying to convert %v to int: %v",
			string(b), err)
	}
	log.Println(myI)

	if len(b) != 8 {
		t.Errorf("length b was %d, expected %d", len(b), 8)
	}
}

func TestMakeSure3(t *testing.T) {
	b := []byte(wotoValues.MakeSureNum(5, 8))
	b = wotoValues.MakeSureByte(b, 8)
	log.Println("\"" + string(b) + "\"")
	if len(b) != 8 {
		t.Errorf("length b was %d, expected %d", len(b), 8)
	}
}

func TestMakeSure4(t *testing.T) {
	b := []byte("0000000000000000000000000000000005")
	b = wotoValues.MakeSureByte(b, 8)
	log.Println("\"" + string(b) + "\"")

	myI, err := strconv.Atoi(strings.TrimSpace(string(b)))
	if err != nil {
		t.Errorf("got an error while trying to convert %v to int: %v",
			string(b), err)
	}
	log.Println(myI)

	if len(b) != 8 {
		t.Errorf("length b was %d, expected %d", len(b), 8)
	}
}
