package main

import "fmt"

// Facade pattern
// Facade object
type forgeFacade struct {
	inv []IWeapon
}

// Singleton of Facade
var singleInstance *forgeFacade

func getInstance() *forgeFacade {
	if singleInstance == nil {
		singleInstance = &forgeFacade{}
	}
	return singleInstance
}

// 1. Accessing inventory
func (f *forgeFacade) inventory() bool {
	if len(f.inv) != 0 {
		for i, w := range f.inv {
			fmt.Printf(`------ %d ------
						Type: %s
						Name: %s
						Attack: %d
						MP: %d\n`, i, w.getType(), w.getName(), w.getAtk(), w.getMp())
		}
		return true
	} else {
		fmt.Println("There's nothing in your inventory. Maybe consider ordering something?")
		return false
	}
}

// 2. Making a weapon
func (f *forgeFacade) makeWeapon() {
	fmt.Println(`Sir, what weapon you want?
				1. Mace
				2. Sword
				3. Dagger`)
	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		f.inv = append(f.inv, makeMace())
	case 2:
		f.inv = append(f.inv, makeSword())
	case 3:
		f.inv = append(f.inv, makeDagger())
	default:
		fmt.Println("Please, choose from our assortment")
		f.makeWeapon()
	}
}

// 3. Upgrading a weapon
func (f *forgeFacade) upgradeWeapon() {
	emptyInv := f.inventory()
	if !emptyInv {
		fmt.Println("What weapon needs upgrade?")
		var i int
		fmt.Scan(&i)
		fmt.Println(`What kind of upgrade?
					1. Sharpen 
					2. Enchant`)
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			f.inv[i-1] = Sharper{f.inv[i-1]}
		case 2:
			f.inv[i-1] = Enchanter{f.inv[i-1]}
		default:
			fmt.Println("Please, choose from our assortment")
			f.upgradeWeapon()
		}
	}
}

// Facade itself
func (f *forgeFacade) blacksmithFacade() {
LOOP:
	for {
		fmt.Println(`Welcome to our blacksmith! What can i help you with?
					1. My inventory
					2. Make me a weapon
					3. Upgrade me a weapon
					4. Leave\n`)
		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			f.inventory()
		case 2:
			f.makeWeapon()
		case 3:
			f.upgradeWeapon()
		case 4:
			fmt.Println("Good luck! Come back later")
			break LOOP
		default:
			fmt.Println("We don't offer such services")
		}
	}
}
