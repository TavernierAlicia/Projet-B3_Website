package main

func SelectSubj(subjectNum string) string {
	subject := ""
	//7 cases
	switch subjectNum {
	case "1":
		subject = "Je souhaite en savoir plus concernant le service"
	case "2":
		subject = "J'aimerais prendre rendez-vous avec l'équipe Click'N Drink"
	case "3":
		subject = "Je souhaite m'inscrire au service"
	case "4":
		subject = "J'ai trouvé un bug!"
	case "5":
		subject = "J'ai rencontré un problème avec une commande"
	case "6":
		subject = "Je souhaite fermer mon compte"
	case "7":
		subject = "Autre"
	default:
		subject = "Autre"
	}
	return subject
}
