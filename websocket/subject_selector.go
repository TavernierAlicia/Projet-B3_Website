package main

func SelectSubj(pro bool, subjectNum string) string {
	subject := ""
	if pro == true {
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
	} else {
		//3 cases
		switch subjectNum {
		case "1":
			subject = "J'ai un problème avec ma commande"
		case "2":
			subject = "J'ai trouvé un bug!"
		case "3":
			subject = "Autre"
		default:
			subject = "Autre"
		}
	}
	return subject
}
