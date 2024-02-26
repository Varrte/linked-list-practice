package main

import "fmt"

// DO NOT EDIT ---------------------------------------------------------------------------------------------------------
type Node struct {
	value int
	next  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

func printList(list *Node) {
	curNode := list
	for i := 0; curNode != nil; i++ {
		fmt.Printf("%d: %d\n", i, curNode.Value())
		curNode = curNode.Next()
	}
}

// ---------------------------------------------------------------------------------------------------------------------

pushFront добавляет новый элемент в начало списка
func pushFront(list *Node, value int) {
			item:=&Node{
			value: list.value,
			next: list.next,
			}
			list.value=value
			list.next=item
	//panic("not implemented")
}

// pushBack добавляет новый элемент в конец списка
func pushBack(list *Node, value int) {
			thirdItem:=&Node{
				value: value,
				next: nil,
			}
			secondItem:=list
			for secondItem.next!=nil{
				secondItem=secondItem.next
			}
			secondItem.next=thirdItem

	// panic("not implemented")
}

// // count возвращает кол-во элементов в списке
	 func count(list *Node) int {
		count:=0
		item:=list
		for item.next!=nil{
			item=item.next
			count++
		}
		return count

	// 	panic("not implemented")
 }

// popFront возвращает значение первого элемента и удаляет его из списка
func popFront(list *Node) int {
	item:=list
	firstValue:=list.value
	list.value=item.next.value
	list.next=item.next.next
	

	return firstValue 
	//panic("not implemented")
}

// // popBack возвращает значение последнего элемента и удаляет его из списка
func popBack(list *Node) int {
		secondItem:=list
		//for secondItem.next!=nil{
			for i:=0;i<count(list)-1;i++{
			secondItem=secondItem.next
		}
		lastValue:=secondItem.next.value
		secondItem.next=nil
		return lastValue
	//panic("not implemented")
}

// isValueInList ищет значение в списке и возвращает true, если оно найдено, в ином случае - false
func isValueInList(list *Node, value int) bool {
	secondItem:=list

	for secondItem.next!=nil{
		if secondItem.value==value{
			return true
		}
		secondItem=secondItem.next
	}
	//panic("not implemented")
	return false
}

// // getValueByIndex возвращает значение из списка по индексу, если оно есть, в ином случае - error("index out of range")
// func getValueByIndex(list *Node, index int) (int, error) {
// 	panic("not implemented")
// }

// // insert добавляет элемент в список в соответствующий индекс
// func insert(list *Node, index, value int) {
// 	panic("not implemented")
// }

// // deleteByIndex удаляет элемент из списка по индексу и возвращает его значение. Если индекс неправильный - возвращает error("index out of range")
// func deleteByIndex(list *Node, index int) (int, error) {
// 	panic("not implemented")
// }

// // sort сортирует список (*)
// func sort(list *Node) {
// 	panic("not implemented")
// }

func main() {
	linkedList := &Node{
		value: 0,
		next:  nil,
	}
	listValue:=[]int{1,2,3,4,5}

	for _,v:= range listValue{
		pushFront(linkedList,v)
	}
	for _,v:= range listValue{
			pushBack(linkedList,v)
		 }
	fmt.Println(popFront(linkedList))
	fmt.Println(popBack(linkedList))
	fmt.Println(isValueInList(linkedList,3))
	printList(linkedList)
}
