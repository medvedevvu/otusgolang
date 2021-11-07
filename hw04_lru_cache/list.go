package hw04lrucache

import "fmt"

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
	PrintAll()
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	countItems int
	head       *ListItem
	tail       *ListItem
}

func (l list) Len() int {
	return l.countItems
}

func (l list) IsEmpty() bool {
	return l.countItems == 0
}

func (l *list) PushFront(v interface{}) *ListItem {
	newli := &ListItem{Value: v, Prev: nil, Next: nil}
	if l.Len() == 0 {
		l.head = newli
		l.tail = newli
	} else {
		l.head.Prev = newli
		newli.Next = l.head
		l.head = newli
	}
	l.countItems++
	return newli
}

func (l *list) PushBack(v interface{}) *ListItem {
	newli := &ListItem{Value: v, Prev: nil, Next: nil}
	if l.Len() == 0 {
		l.head = newli
		l.tail = newli
	} else {
		newli.Prev = l.tail
		l.tail.Next = newli
		l.tail = newli
	}
	l.countItems++
	return newli
}

func (l list) Front() *ListItem {
	return l.head
}

func (l list) Back() *ListItem {
	return l.tail
}

func (l *list) Remove(i *ListItem) {
	curr := l.head
	if curr == nil {
		return
	}
	if curr.Value == i.Value { // head is the node with value key.
		curr = curr.Next
		l.countItems--
		if curr != nil {
			l.head = curr
			l.head.Prev = nil
		} else {
			l.tail = nil // only one element in list.
		}
		return
	}
	for curr.Next != nil {
		if curr.Next.Value == i.Value {
			curr.Next = curr.Next.Next
			if curr.Next == nil { // last element case.
				l.tail = curr
			} else {
				curr.Next.Prev = curr
			}
			l.countItems--
			return
		}
		curr = curr.Next
	}
}

func (l *list) PrintAll() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.Value, " ")
		temp = temp.Next
	}
}

func (l *list) MoveToFront(i *ListItem) {
	l.Remove(i)
	l.PushFront(i.Value)
}

func NewList() List {
	return &list{}
}
