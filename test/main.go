package main

import (
	"encoding/json"
	"fmt"
	"github.com/kcswag/kcgods/lists/doublylinkedlist"
)

type Node struct {
	Identifier  string
	Description string
}

func main() {
	heyN := &Node{Identifier: "hey"}
	hoN := &Node{Identifier: "ho"}
	fuckN := &Node{Identifier: "fuck"}
	itN := &Node{Identifier: "it"}
	damnN := &Node{Identifier: "damn"}
	shitN := &Node{Identifier: "shit"}

	dl := doublylinkedlist.New[*Node]()
	dl.Add(heyN, hoN, fuckN, itN)
	dl.Prepend(damnN)
	dl.Append(shitN)
	dlJson, _ := json.Marshal(dl)

	dl2 := doublylinkedlist.New[*Node]()
	json.Unmarshal(dlJson, &dl2)
	fmt.Println(dl2)

}
