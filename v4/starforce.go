package v4

import (
	myItem "github.com/Geonhyeong/go_test_lib/proto"
)

func StarForce(myItem *myItem.Item) *myItem.Item {
	myItem.Id = 2
	myItem.Value = "Lee GeonHyeong"

	myItem.ItemEventSourcing.CHUC = 4
	myItem.ItemEventSourcing.Option1 = 6

	return myItem
}
