package main

import (
	"os"

	"fmt"
)

func main() {

	var nombre int64
	var fichier string
	var optionC string
	checkFile := 0

	//mettre les arguments dans les variables

	for i := range args {
		stockArgs := args[i]

		if args[i] == "-c" && stockArgs[0] == '-' {
			optionC = args[i]
			continue
		}

		if stockArgs[0] >= '0' && stockArgs[0] <= '9' {
			for j := range args[i] {
				Rune := []rune(args[i])

				if Rune[j] < '0' || Rune[j] > '9' {
					break
				}
				nombre = Atoi(args[i])
			}
			continue
		}

		for k := range stockArgs {
			if stockArgs[k] == '.' {
				checkFile++
				if fichier == "" {
					_, err := os.Open(args[i])

					if err != nil {
						fmt.Printf("tail: %s: No such file or directory\n", args[i])
						return
					} else {
						fichier = args[i]

					}
				}
			}
		}
		if checkFile == 0 {
			fmt.Printf("tail: invalid number of bytes: %q\n", stockArgs)
			return
		}
	}
	if checkFile == 1 {
		fichierOpen, _ := os.Open(fichier)

		if optionC == "-c" && nombre == 0 {
			fmt.Printf("tail: option requires an argument -- 'c'\n")
			return
		}

		if optionC == "-c" { //Avec le "-c"
			sizeFileWithOption, _ := fichierOpen.Stat()

			if sizeFileWithOption.Size() <= nombre { //Print tt le fichier
				sizeFile, _ := fichierOpen.Stat()     //Récupère des infos sur le fichier
				data := make([]byte, sizeFile.Size()) //sizeFile.Size() récupère la taille du fichier et crée un tableau de cette taille
				count, _ := fichierOpen.Read(data)
				fmt.Printf("%q\n", data[:count])
			} else { // Print le tableau

				sizeFile, _ := fichierOpen.Stat()
				data := make([]byte, sizeFile.Size())
				var tab []byte
				count, _ := fichierOpen.Read(data)
				count = count

				compteur := 1
				for i := int64(len(data) - 1); i >= 0; i-- {

					if int64(compteur) <= nombre {
						tab = append(tab, data[i])
					}
					compteur++
				}

				for j := len(tab) - 1; j >= 0; j-- {
					fmt.Printf(string(tab[j]))
				}
			}

		} else { //Sans le "-c"  Print tous le fichier, GOOD

			sizeFile, _ := fichierOpen.Stat()     //Récupère des infos sur le fichier
			data := make([]byte, sizeFile.Size()) //sizeFile.Size() récupère la taille du fichier et crée un tableau de cette taille
			count, _ := fichierOpen.Read(data)
			fmt.Printf("%q\n", data[:count])
		}
	}
}

var (
	args       = os.Args[1:]
	tailleArgs = len(args)
)

func Atoi(s string) int64 {
	var res int64 = 0

	for i, val := range s {
		if s[0] >= '0' && s[0] <= '9' {
			var dig int64 = 0

			for j := '1'; j <= val; j++ {
				dig++
			}
			res = res*10 + dig
		}

		if i > 0 {
			if !(val >= '0' && val <= '9') {
				res = 0
				return res
			}
		}
	}
	return res
}

/*********************************************A FINIR*********************************************/

//"tail: illegal offset"
//"tail: %s: No such file or directory"

/*	tail fichier.txt : Print tous le fichier, GOOD
	tail -c 15 fichier.txt : Print les 15 derniers bytes du fichier.txt

	Conditions :
	SI le INT est plus grand que la taille entiere du fichier tu print le fichier
	SI le nom du fichier en args est différent du nom de fichier, print erreur :tail: %s: No such file or directory (file.Name), GOOD
	SI l'argument du nombre est autre qu'un nombre renvoyer "tail: invalid number of bytes: ‘’"
	SI l'argument du fichier est en plusieurs fois, il faut print plusieurs fois

	Commandes utiles :
	os.OpenFile() : Ouvrir le fichier
	file.Name() : Renvoie le nom du file
	os.Exit() : Sortir de la sortie standard
	os.Stdin() : Équivalent à cat. Prenez l'entrée standard et imprimez-la sur la sortie standard
	os.Stdout() : fichier qui est immédiatement imprimé à l'écran par la console

	data := make([]byte, 100)
	count, err := file.Read(data) : lire un fichier
*/

//FINI
