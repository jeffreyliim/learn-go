/*	LinkedList	*/
package main

import (
	"fmt"
	"math"
)

type person struct {
	Age       int
	Weight    int
	FirstName string
	Next      *person
}

//	Mark -> Samuel -> Jeff
func main() {

	pl := &person{23, 55, "Mark", nil}
	pl = pl.addNodeEnd(&person{24, 57, "Samuel", nil})
	pl = pl.addNodeEnd(&person{25, 25, "Jeff", nil})
	pl = pl.addNodeEnd(&person{20, 55, "Kim", nil})
	pl = pl.addNodeEnd(&person{10, 25, "Lee", nil})
	//pl = pl.addNodeFront(&person{44, 22, "Mich", nil})

	//pl = pl.deleteNodeBack()
	//pl = pl.deleteNodeFront()
	pl = pl.reverseRecursive()
	printList(pl)

	//fmt.Println(pl)
	//	print 1
	//getMiddleNode := pl.getMiddleNode()
	//fmt.Println(getMiddleNode, "Get middle node")

	//	 print 2
	//n := 6
	//getNthNode := pl.getNthNode(n)
	//fmt.Printf("Getting %v th node: %v", n, getNthNode)
	//
	////	print 3
	//printList(pl)

}

func printList(pl *person) {
	for l := pl; l != nil; l = l.Next {
		fmt.Println(l)
	}
}

func (pl *person) addNodeEnd(newPerson *person) *person {
	for p := pl; p != nil; p = p.Next {
		if p.Next == nil {
			p.Next = newPerson
			return pl
		}
	}
	return pl
}
func (pl *person) addNodeFront(newPerson *person) *person {
	currentPlist := pl
	newPlist := newPerson
	newPlist.Next = currentPlist
	return newPlist
}

func (pl *person) deleteNodeBack() *person {
	for p := pl; p != nil; p = p.Next {
		// we will filter out the node that has next != nil and next.next as nil. #returns Samuel
		if p.Next != nil && p.Next.Next == nil {
			p.Next = nil
		}
	}
	return pl
}

func (pl *person) deleteNodeFront() *person {
	return pl.Next
}

func (pl *person) getMiddleNode() *person {
	counter := 0
	list := make([]person, counter)
	for p := pl; p != nil; p = p.Next {
		counter++
		list = append(list, *p)
	}

	countNode := int(math.Round(float64((counter - 1) / 2)))
	return &list[countNode]
}

func (pl *person) reverse() *person {
	var prev *person
	if pl.Next == nil {
		return pl
	} else {
		for {
			if pl == nil {
				break
			}
			next := pl.Next
			pl.Next = prev
			prev = pl
			pl = next
		}
		return prev
	}
}

func (pl *person) reverseRecursive() *person {
	if pl.Next == nil {
		return pl
	} else {
		next := pl.Next.reverseRecursive()
		pl.Next.Next = pl
		pl.Next = nil
		return next
	}
}

func (pl *person) getNthNode(n int) *person {
	var m []*person
	for p := pl; p != nil; p = p.Next {
		m = append(m, p)
	}

	if n < 1 {
		return nil
	}
	if n > len(m) {
		return nil
	}
	return m[n-1]
}
