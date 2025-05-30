package main

import (
	"fmt"
	linkedlist "linked_list/linked_list"
)

func test_erros() {
	println("Teste 1 --------------------------")
	ll := linkedlist.NewLinkedList()

	_, err := ll.Get(0)
	if err != nil {
		fmt.Println("Error ao resgatar")
	}

	_, err = ll.Get(3)
	if err != nil {
		fmt.Println("Error ao resgatar")
	}
}

func test_add_node() {
	println("Teste 2 --------------------------")
	ll := linkedlist.NewLinkedList()

	ll.AddNode("linha 1")
	ll.AddNode("linha 2")
	ll.AddNode("linha 3")

	fmt.Println(ll.Len)
	v, err := ll.Get(0)
	if err == nil {
		fmt.Println("Valor: ", v)
	}

	v, err = ll.Get(1)
	if err == nil {
		fmt.Println("Valor: ", v)
	}

	v, err = ll.Get(2)
	if err == nil {
		fmt.Println("Valor: ", v)
	}

	err = ll.Set(5, "Valor teste 1")
	if err != nil {
		fmt.Println(err, "- ok")
	}
}

func test_set() {
	println("Teste 3 --------------------------")
	ll := linkedlist.NewLinkedList()

	ll.AddNode("linha 1")
	ll.AddNode("linha 2")
	ll.AddNode("linha 3")

	err := ll.Set(5, "Valor teste 1")
	if err != nil {
		fmt.Println(err, "ok")
	}

	err = ll.Set(0, "Valor teste 1")
	if err == nil {
		fmt.Println("ok")
	}

	err = ll.Set(1, "Valor teste 2")
	if err == nil {
		fmt.Println("ok")
	}

	err = ll.Set(2, "Valor teste 3")
	if err == nil {
		fmt.Println("ok")
	}
}

func test_remove() {
	println("Teste 4 --------------------------")
	ll := linkedlist.NewLinkedList()

	err := ll.Remove(0)
	if err != nil {
		fmt.Println(err, " - ok")
	}
	ll.AddNode("linha 1")

	err = ll.Remove(0)
	if err != nil {
		fmt.Println(err, " - not ok")
	}

	if ll.Len != 0 {
		fmt.Println("not ok len 1")
	}

	v, err := ll.Get(0)
	if err == nil {
		fmt.Println("not ok get depois de remover", v)
	}

	ll.AddNode("linha 1")
	ll.AddNode("linha 2")

	err = ll.Remove(1)
	if err != nil {
		fmt.Println(err, " - not ok")
	}

	if ll.Len != 1 {
		fmt.Println("not ok len 2, len: ")
	}

	v, err = ll.Get(1)
	if err == nil {
		fmt.Println("not ok get depois de remover", v)
	}

	ll.AddNode("linha 2")
	ll.AddNode("linha 3")

	err = ll.Remove(1)
	if err != nil {
		fmt.Println(err, " - not ok")
	}

	if ll.Len != 2 {
		fmt.Println("not ok len 3")
	}

	v, _ = ll.Get(1)
	if v != "linha 3" {
		fmt.Println("not ok get depois de remover", v)
	}

	pivot := ll.Start
	for pivot != nil {
		fmt.Println(pivot.Value)
		pivot = pivot.Next
	}
}

func test_insert() {
	println("Teste 5 --------------------------")

	ll := linkedlist.NewLinkedList()

	err := ll.Insert(1, "tes")
	if err == nil {
		println("deu errado")
	}

	ll.AddNode("linha 1")
	err = ll.Insert(0, "linha 0")
	if err != nil {
		println(err, "ao inserir linha 0")
	}

	if ll.Len != 2 {
		println("len deveria ser 2")
	}

	ll.Insert(1, "linha 0.5")
	v, _ := ll.Get(1)
	if v != "linha 0.5" {
		println("index 1 deveria ser linha 0.5")
		println("index 1: ", v)
	}

	println("ok")
}

func main() {
	test_erros()
	test_add_node()
	test_set()
	test_remove()
	test_insert()
}
