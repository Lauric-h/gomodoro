Delete when project done
## Specs and todos
- [X] Afficher les chiffres qui décomptent
- [X] Gérer les minutes 
- [X] Gestion des long break  
- [X] ajout pause
- [X] Régler départ horloge
- [X] Régler les long break tous les multiples de 4
- [X] Ajout temps maximum (60 minutes)rd 
- [X] Checker si trop d'arguments
- [X] logger dans un fichier txt la session
- [] Add total time dans le logger  
- [] Ajout style avec Pterm
- [X] Ajout touche pour quitter  
- [] Ajout touche pour arreter  ?
- [] Ajout touche pour mettre en pause/restart  
- [] add flags/subcommand pour temps total
- [X] add flags/subcommand pour temps de pause
- [X] add flags/subcommand pour temps de travail
- [] add README
- [] quitter proprement le programme -> ecrire la duree, sessions etc.
- [] Refactor - utiliser stdlib time ?
- [] add godoc

### Première itération : 
au lancement, worktimer de 25minutes et pause de 5 min  
Ajout des long break
Pas de choix de temps  
Pas de touche pour arrêter  
Pas de log du nombre de session

### Deuxieme itération
Arreter a tout moment via touches pour quitter / pause  
Changer le temps des pauses courtes/pauses longues/worktimer au démarrage

### Troisième itération
Logger les sessions dans un fichier  
Définir un temps total au démarrage  
Possibilité d'afficher les logs/stats  

### Quatrième itération
Ajout alertes/notif + possibilité de disables
Style avec Pterm
Ajout help

