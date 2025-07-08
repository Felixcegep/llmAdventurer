
# 📦 Anime Battle CLI - UI Enhancement Guide

---

## 🧠 Section 1 : System Prompt (pour assistant ou développeur AI)

```
Tu es un développeur Go travaillant sur un jeu CLI basé sur Cobra.

Tu dois :
- Garder un code propre, modulaire et idiomatique (cmd/, internal/, etc.)
- Offrir une expérience utilisateur fluide dans le terminal (UX CLI)
- Rendre l'interface vivante avec couleurs, pauses, emojis et menus clairs
- Réagir aux erreurs utilisateur et donner du feedback immédiat
- Ne jamais lancer le programme sans que la variable d’environnement `GEMINI_API_KEY` soit définie
```

---

## 🧑‍💻 Section 2 : Step-by-Step — Améliorer l’UI en CLI

### ⚠️ AVANT DE COMMENCER
> 🚫 **NE PAS LANCER LE PROGRAMME sans définir la clé API** :
```bash
# Bash / Zsh / Git Bash
export GEMINI_API_KEY="ta_clé"

# PowerShell
$env:GEMINI_API_KEY="ta_clé"
```

---

### 1️⃣ Ajouter des couleurs dans le terminal

```bash
go get github.com/fatih/color
```

```go
import "github.com/fatih/color"

color.Green("✅ Attaque réussie !")
color.Red("❌ Vous avez subi des dégâts.")
color.Cyan("🎒 Inventaire ouvert.")
```

---

### 2️⃣ Afficher des menus clairs

```go
fmt.Println("========== COMBAT ==========")
fmt.Printf("👤 %s (PV: %d)  🆚  🐲 %s (PV: %d)\n", p.HeroName, p.Health, b.Name, b.Health)
fmt.Println("Actions disponibles :")
fmt.Println("  [1] ⚔️  Attaquer")
fmt.Println("  [2] 🧪  Se soigner")
fmt.Println("  [3] 🎒  Utiliser un objet")
fmt.Print("👉 Entrez votre choix > ")
```

---

### 3️⃣ Ajouter du rythme avec `time.Sleep`

```go
fmt.Println("L'ennemi vous fixe...")
time.Sleep(2 * time.Second)
```

---

### 4️⃣ Afficher les effets des actions

```go
fmt.Printf("🎯 Vous infligez %d dégâts à %s !\n", damage, b.Name)
fmt.Printf("💖 Vous récupérez %d PV.\n", healAmount)
```

---

### 5️⃣ Nettoyer l’écran entre les tours (facultatif)

```go
fmt.Print("\033[H\033[2J") // Efface l'écran dans la plupart des terminaux
```

---

### 6️⃣ Utiliser des emojis pour dynamiser

- ⚔️ Attaque
- 🧪 Potion
- 🛡️ Défense
- 🎒 Inventaire
- 💥 Dégâts
- 💖 Soins

---

### 7️⃣ Créer des barres de vie

```go
func renderHealthBar(name string, current, max int) {
    bar := strings.Repeat("█", current*10/max)
    fmt.Printf("%s [%s] %d/%d PV\n", name, bar, current, max)
}
```

---

### 8️⃣ Centraliser les menus

```go
func showMainMenu() {
    fmt.Println("🎮 MENU PRINCIPAL")
    fmt.Println("[1] Nouvelle Partie")
    fmt.Println("[2] Charger Partie")
    fmt.Println("[3] Quitter")
}
```

---

### ✅ Résultat attendu

- Interface conviviale, claire et engageante
- Plus immersive même sans interface graphique
- Rythme plus fluide et lisible
