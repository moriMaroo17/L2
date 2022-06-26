package main

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(int) error
	dispenseItem() error
}
