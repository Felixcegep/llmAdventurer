
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
> 🚫 **NE PAS LANCER LE PROGRAMME 

```

---

il y a un bug avec le healthbar
loom Soul [████████████████████] 250/250 PV
============================
Actions disponibles :
[1] ⚔️  Attaquer
[2] 🧪  Se soigner
[3] 🎒  Utiliser un objet
[4] ✨  Superpouvoir
👉 Entrez votre choix > 4
💥 Super Pouvoir Activé!
💖 Vous récupérez 50 PV et infligez -7 dégâts !
👹 Gloom Soul vous attaque!
💔 Vous avez pris 10 dégâts! Il vous reste 190 PV.
========== COMBAT ==========
panic: strings: negative Repeat count

goroutine 1 [running]:
strings.Repeat({0x1095d01?, 0x3?}, 0x1d?)
C:/Program Files/Go/src/strings/strings.go:624 +0x585
awesomeProject1/internal/gamelogic.renderHealthBar({0xc000321b6c, 0x4}, 0xbe, 0x96)
C:/Users/fella/GolandProjects/awesomeProject1/internal/gamelogic/combatmechanic.go:68 +0x9c
awesomeProject1/internal/gamelogic.Combat(0xc000099a78, 0xc000099b48, {0xc00037c000, 0x5c})
C:/Users/fella/GolandProjects/awesomeProject1/internal/gamelogic/combatmechanic.go:21 +0x158
awesomeProject1/cmd.init.func1(0xc000196c00?, {0x10960bf?, 0x4?, 0x10960c3?})
C:/Users/fella/GolandProjects/awesomeProject1/cmd/play.go:37 +0x3d9
github.com/spf13/cobra.(*Command).execute(0x1656d20, {0x16aa100, 0x0, 0x0})
C:/Users/fella/go/pkg/mod/github.com/spf13/cobra@v1.9.1/command.go:1019 +0xa91
github.com/spf13/cobra.(*Command).ExecuteC(0x1656fe0)
C:/Users/fella/go/pkg/mod/github.com/spf13/cobra@v1.9.1/command.go:1148 +0x46f
github.com/spf13/cobra.(*Command).Execute(...)
C:/Users/fella/go/pkg/mod/github.com/spf13/cobra@v1.9.1/command.go:1071
awesomeProject1/cmd.Execute()
C:/Users/fella/GolandProjects/awesomeProject1/cmd/root.go:30 +0x1a
main.main()
C:/Users/fella/GolandProjects/awesomeProject1/main.go:9 +0xf
exit status 2
et je veux que la health barre affiche zero si le hero perd 