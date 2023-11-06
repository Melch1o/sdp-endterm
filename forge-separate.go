package main

import "fmt"

// Factory pattern - making weapons
// Interface for weapon
type IWeapon interface {
	getType() string
	getName() string
	getAtk() int
	getMp() int
}

// Concrete weapon
type Weapon struct {
	weaponType string
	name       string
	atk        int
	mp         int
}

func (w Weapon) getType() string {
	return w.weaponType
}

func (w Weapon) getName() string {
	return w.name
}

func (w Weapon) getAtk() int {
	return w.atk
}

func (w Weapon) getMp() int {
	return w.mp
}

// Concrete products
type Mace struct {
	Weapon
}

func makeMace() IWeapon {
	var n string
	fmt.Print("Choose name for your weapon: ")
	fmt.Scan(&n)
	return &Mace{
		Weapon{
			weaponType: "Mace",
			name:       n,
			atk:        20,
			mp:         0,
		},
	}
}

type Sword struct {
	Weapon
}

func makeSword() IWeapon {
	var n string
	fmt.Print("Choose name for your weapon: ")
	fmt.Scan(&n)
	return &Sword{
		Weapon{
			weaponType: "Sword",
			name:       n,
			atk:        10,
			mp:         5,
		},
	}
}

type Dagger struct {
	Weapon
}

func makeDagger() IWeapon {
	var n string
	fmt.Print("Choose name for your weapon: ")
	fmt.Scan(&n)
	return &Dagger{
		Weapon{
			weaponType: "Dagger",
			name:       n,
			atk:        5,
			mp:         15,
		},
	}
}

// Decorator pattern - upgrading weapons
// Decorator for atkack
type Sharper struct {
	w IWeapon
}

func (s Sharper) getType() string {
	return s.w.getType()
}

func (s Sharper) getName() string {
	return s.w.getName()
}

func (s Sharper) getAtk() int {
	return s.w.getAtk() + 5
}

func (s Sharper) getMp() int {
	return s.w.getMp()
}

// Decorator for Magic power
type Enchanter struct {
	w IWeapon
}

func (e Enchanter) getType() string {
	return e.w.getType()
}

func (e Enchanter) getName() string {
	return e.w.getName()
}

func (e Enchanter) getAtk() int {
	return e.w.getAtk()
}

func (e Enchanter) getMp() int {
	return e.w.getMp() + 5
}
