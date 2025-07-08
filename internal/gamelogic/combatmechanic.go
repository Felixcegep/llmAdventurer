package gamelogic

import (
	"fmt"
	"time"
)

func calculateDamage(attack, defense int) int {
	damage := attack - (defense / 2)
	return damage
}
func Combat(p *Player, b *Boss, dialogue string) bool {
	clearScreen()
	fmt.Printf("\n%s ❤️ %d  |  %s 💀 %d\n", p.HeroName, p.Health, b.Name, b.Health)
	fmt.Println("📝 Dialogue:", dialogue)
	time.Sleep(5 * time.Second)
	for {
		fmt.Printf("\n%s ❤️ %d  |  %s 💀 %d\n", p.HeroName, p.Health, b.Name, b.Health)

		if b.Health <= 0 || p.Health <= 0 {
			if b.Health <= 0 {
				fmt.Println("🎉 You defeated the boss!")
				return true
			} else {
				fmt.Println("💀 You were defeated...")
				return false
			}
		} else {

			var choice int
			fmt.Println("Choose your action:")
			fmt.Println("  [1] Attack")
			fmt.Println("  [2] Heal")
			fmt.Println("  [3] Use Item")
			fmt.Println("  [4] Superpower")
			fmt.Print("Enter your choice: ")
			fmt.Scan(&choice)
			if 0 < choice && choice <= 4 {
				playerTurn(p, b, choice)
				if b.Health > 0 {
					bossTurn(p, b)
					if p.Health <= 0 {
						return false
					}
				} else {
					return true
				}

			} else {
				fmt.Println("choice dont exist")
				continue
			}
		}

	}
}
func playerTurn(p *Player, b *Boss, choices int) {
	switch choices {
	case 1:
		fmt.Println("🗡️  You attack!")
		b.Health -= calculateDamage(p.Attack, b.Defense)

	case 2:
		p.healingpotion()
	case 3:
		PlayerUseItem(p, b)

	case 4:
		fmt.Println("💥 Super Power Activated!")

		p.Health += p.SuperPower.Health
		b.Health -= calculateDamage(p.SuperPower.Attack, b.Defense)
		p.Defense += p.SuperPower.Defense

	}
}

func bossTurn(p *Player, b *Boss) {
	fmt.Printf("👹 %s attacks you!\n", b.Name)
	damage := calculateDamage(b.Attack, p.Defense)
	p.Health -= damage
	fmt.Printf("You took %d damage! Your health is now %d\n", damage, p.Health)
}
func PlayerUseItem(p *Player, b *Boss) {
	if len(p.Inventory) == 0 {
		fmt.Println("📦 Inventory empty!")
		return
	}
	fmt.Println("🎒 Your items:")
	for i, item := range p.Inventory {
		fmt.Printf("  [%d] %s (Heal: %d, Atk: %d, Def: %d)\n", i, item.NameItem, item.Health, item.Attack, item.Defense)
	}

	var choiceItem int
	fmt.Print("Choose item: ")
	fmt.Scan(&choiceItem)

	item := p.Inventory[choiceItem]

	p.Health += item.Health
	b.Health -= calculateDamage(item.Attack, b.Defense)
	p.Defense += item.Defense

	// Remove item
	p.Inventory = append(p.Inventory[:choiceItem], p.Inventory[choiceItem+1:]...)
}
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
