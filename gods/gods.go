package main

import (
	"encoding/json"
	"fmt"
	"github.com/emirpasic/gods/lists/singlylinkedlist"
	"github.com/emirpasic/gods/queues/priorityqueue"
	"github.com/shopspring/decimal"
	"log"
)

type Element struct {
	Name     string                 `json:"name"`
	Orders   *singlylinkedlist.List `json:"orders"`
	Priority decimal.Decimal        `json:"priority"`
}

type Order struct {
	Id string `json:"id"`
}

func descPriority(a, b interface{}) int {
	priorityA := a.(Element).Priority
	priorityB := b.(Element).Priority
	return -decimalComparator(priorityA, priorityB) // "-" descending order
}

func ascPriority(a, b interface{}) int {
	priorityA := a.(Element).Priority
	priorityB := b.(Element).Priority
	return decimalComparator(priorityA, priorityB) // "-" descending order
}

func decimalComparator(a, b interface{}) int {
	aAsserted := a.(decimal.Decimal)
	bAsserted := b.(decimal.Decimal)
	return aAsserted.Cmp(bAsserted)
}

func main() {
	list := singlylinkedlist.New()
	list.Add(Order{Id: "123456"})
	list.Add(Order{Id: "1234567"})
	decimal0, err := decimal.NewFromString("123.33")
	if err != nil {
		log.Fatal(err)
	}
	decimal1, err := decimal.NewFromString("124.33")
	if err != nil {
		log.Fatal(err)
	}
	decimal2, err := decimal.NewFromString("125.33")
	if err != nil {
		log.Fatal(err)
	}
	element := Element{Name: "123.33", Orders: list, Priority: decimal0}
	element1 := Element{Name: "124.33", Orders: list, Priority: decimal1}
	element2 := Element{Name: "125.33", Orders: list, Priority: decimal2}
	descQueue := priorityqueue.NewWith(descPriority)
	descQueue.Enqueue(element1)
	descQueue.Enqueue(element)
	descQueue.Enqueue(element2)
	b := true
	for b {
		value, ok := descQueue.Dequeue()
		if ok {
			marshal, err := json.Marshal(value)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(marshal))
		}
		b = ok
	}
	ascQueue := priorityqueue.NewWith(ascPriority)
	ascQueue.Enqueue(element1)
	ascQueue.Enqueue(element)
	ascQueue.Enqueue(element2)
	b = true
	for b {
		value, ok := ascQueue.Dequeue()
		if ok {
			marshal, err := json.Marshal(value)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(marshal))
		}
		b = ok
	}
}
