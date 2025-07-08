package gamelogic

import (
	"fmt"
	"github.com/fatih/color"
	"strings"
	"time"
)

func calculateDamage(attack, defense int) int {
	damage := attack - (defense / 2)
	return damage
}
func Combat(p *Player, b *Boss, dialogue string) bool {
	clearScreen()
	color.Cyan("📝 Dialogue: %s", dialogue)
	time.Sleep(5 * time.Second)
	for {
		clearScreen()
		color.Yellow("========== COMBAT ==========")
		renderHealthBar(p.HeroName, p.Health, p.MaxHealth)
		renderHealthBar(b.Name, b.Health, b.MaxHealth)
		color.Yellow("============================")

		if b.Health <= 0 || p.Health <= 0 {
			if b.Health <= 0 {
				color.Green("🎉 You defeated %s!", b.Name)
				clearScreen()
				return true
			} else {
				defeatScreen(b)

				return false
			}
		} else {

			var choice int
			color.Cyan("Actions disponibles :")
			color.Cyan("  [1] ⚔️  Attaquer")
			color.Cyan("  [2] 🧪  Se soigner")
			color.Cyan("  [3] 🎒  Utiliser un objet")
			color.Cyan("  [4] ✨  Superpouvoir")
			fmt.Print("👉 Entrez votre choix > ")
			fmt.Scan(&choice)
			if 0 < choice && choice <= 4 {
				playerTurn(p, b, choice)
				if b.Health > 0 {
					bossTurn(p, b)
					if p.Health <= 0 {
						defeatScreen(b)
						return false

					}
				} else {

					return true
				}

			} else {
				color.Red("Choix invalide. Veuillez réessayer.")
				time.Sleep(1 * time.Second)
				continue
			}
		}

	}
}
func EndScreen(game Game) {
	clearScreen()

	player := game.Players
	victoryMessage := game.Dialogues[len(game.Dialogues)-1].Content

	art := []string{
		"",
		"🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌",
		"🌟                                           🌟",
		"🌟         🏆 MISSION ACCOMPLISHED 🏆         🌟",
		"🌟___________________________________________🌟",
		"",
		fmt.Sprintf("  Hero: %s", player.HeroName),
		"  Final Stats:",
		fmt.Sprintf("    ✦ Attack: %d", player.Attack),
		fmt.Sprintf("    ✦ Defense: %d", player.Defense),
		fmt.Sprintf("    ✦ Super Power: %s", player.SuperPower.NameItem),
		"",
		"  Bosses Defeated:",
	}
	for _, boss := range game.Bosses {
		art = append(art, fmt.Sprintf("    ☠️  %s", boss.Name))
	}
	art = append(art, "")
	art = append(art, "  Locations Liberated:")
	for _, level := range game.Levels {
		art = append(art, fmt.Sprintf("    🗺️  %s", level))
	}
	art = append(art, "")
	art = append(art, "🎖️ Final Mission Report:")
	art = append(art, fmt.Sprintf("   \"%s\"", victoryMessage))
	art = append(art, "")
	art = append(art, "🌟 Your legend as an exorcist will echo through the Order.")
	art = append(art, "   Press ENTER to return to reality.")
	art = append(art, "🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌🌌")

	for _, line := range art {
		color.Magenta(line)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Scanln()
}

func defeatScreen(b *Boss) {
	clearScreen()
	art := []string{
		"=========================================",
		"               💀 DEFEAT 💀              ",
		"=========================================",
		fmt.Sprintf("   You were defeated by %s...", b.Name),
		"",
		"          ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒           ",
		"         ▒▒░░░░░░░░░░░░░░░░▒▒          ",
		"         ▒▒░░   RIP Hero   ░░▒▒        ",
		"         ▒▒░░   ⚔️👺⚔️     ░░▒▒        ",
		"          ▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒▒           ",
		"",
		"   Live to fight another day, warrior.",
		"=========================================",
	}

	for _, line := range art {
		color.Red(line)
		time.Sleep(200 * time.Millisecond)
	}
}

func renderHealthBar(name string, current, max int) {
	barLength := 20
	// Ensure current health is not negative and not greater than max health for rendering purposes
	displayCurrent := current
	if displayCurrent < 0 {
		displayCurrent = 0
	}
	if displayCurrent > max {
		displayCurrent = max
	}

	filled := int(float64(displayCurrent) / float64(max) * float64(barLength))
	if filled < 0 { // Should not happen with displayCurrent clamping, but as a safeguard
		filled = 0
	}
	if filled > barLength { // Should not happen with displayCurrent clamping, but as a safeguard
		filled = barLength
	}

	empty := barLength - filled
	if empty < 0 { // Should not happen if filled is correctly clamped
		empty = 0
	}

	bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
	color.White("%s [%s] %d/%d PV", name, bar, current, max)
}

func playerTurn(p *Player, b *Boss, choices int) {
	switch choices {
	case 1:
		color.Cyan("🗡️  Vous attaquez %s!", b.Name)
		damage := calculateDamage(p.Attack, b.Defense)
		b.Health -= damage
		color.Green("🎯 Vous infligez %d dégâts à %s !\n", damage, b.Name)
		time.Sleep(1 * time.Second)

	case 2:
		p.healingpotion()
	case 3:
		PlayerUseItem(p, b)

	case 4:
		color.Yellow("💥 Super Pouvoir Activé!")
		time.Sleep(1 * time.Second)
		p.Health += p.SuperPower.Health
		damage := calculateDamage(p.SuperPower.Attack, b.Defense)
		b.Health -= damage
		p.Defense += p.SuperPower.Defense
		color.Green("💖 Vous récupérez %d PV et infligez %d dégâts !\n", p.SuperPower.Health, damage)
		time.Sleep(1 * time.Second)

	}
}

func bossTurn(p *Player, b *Boss) {
	color.Red("👹 %s vous attaque!", b.Name)
	time.Sleep(1 * time.Second)
	damage := calculateDamage(b.Attack, p.Defense)
	p.Health -= damage
	color.Red("💔 Vous avez pris %d dégâts! Il vous reste %d PV.\n", damage, p.Health)
	time.Sleep(1 * time.Second)
}
func PlayerUseItem(p *Player, b *Boss) {
	if len(p.Inventory) == 0 {
		color.Red("📦 Inventaire vide !")
		return
	}
	color.Cyan("🎒 Vos objets :")
	for i, item := range p.Inventory {
		color.Cyan("  [%d] %s (Soins: %d, Atk: %d, Def: %d)", i, item.NameItem, item.Health, item.Attack, item.Defense)
	}

	var choiceItem int
	color.Cyan("Choisissez un objet > ")
	fmt.Scan(&choiceItem)

	if choiceItem < 0 || choiceItem >= len(p.Inventory) {
		color.Red("Choix d'objet invalide.")
		return
	}

	item := p.Inventory[choiceItem]

	p.Health += item.Health
	b.Health -= calculateDamage(item.Attack, b.Defense)
	p.Defense += item.Defense

	color.Green("✨ Vous utilisez %s. PV: %d, Dégâts infligés: %d, Défense: %d\n", item.NameItem, item.Health, calculateDamage(item.Attack, b.Defense), item.Defense)
	time.Sleep(1 * time.Second)

	// Remove item
	p.Inventory = append(p.Inventory[:choiceItem], p.Inventory[choiceItem+1:]...)
}
func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
