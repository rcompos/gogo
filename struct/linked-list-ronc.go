package main

import(
	"fmt"
)

type link struct {
	value	int32
	next	*link
}

func main() {
	fmt.Println("Linked List")

	uno := &link{1, nil}
	list := uno
	dos := &link{2, nil}
	tres := &link{3, nil}

	printList(list)
	addList(dos, list)
	addList(tres, list)

	printList(list)

	list = reverseList(list)
	printList(list)

}

func printList(plist *link) {
	for p := plist; p != nil; p = p.next {
		fmt.Println(p)
	}
}

func addList(newLink, plist *link) *link {
	if plist == nil {
		return plist
	}
	for p := plist; p != nil; p = p.next {
		if p.next == nil {
			p.next = newLink
			return plist
		}
	}
	return plist
}

func reverseList(plist *link) *link {
	if plist == nil {
		return plist
	}
	p := plist
	if p.next == nil {
		return p
	} else {
		newHead := reverseList(p.next)
		p.next.next = p
		p.next = nil
		return newHead
	}
}
