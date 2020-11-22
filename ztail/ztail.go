package main

import (
	"os"

	"github.com/01-edu/z01"

	"fmt"
)

func main() {

	var nombre int64
	var fichier string
	var optionC string
	checkFile := 0
	NOTFILE := ""

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

			if stockArgs[k] != '.' {
				notFile = true
				NOTFILE = args[i]
			} else if stockArgs[k] == '.' {
				checkFile++
				if fichier == "" {
					_, err := os.Open(args[i])

					if err != nil {
						fmt.Printf("tail: %s: No such file or directory\n", args[i])
						return
					} else {
						fichier = args[i]
						notFile = false
						break
					}
				}
			}
		}

		if checkFile == 0 {
			fmt.Printf("tail: invalid number of bytes: %q\n", stockArgs)
			return
		}
	}

	if checkFile >= 1 && notFile == true {

	}

	if checkFile >= 1 {
		fichierOpen, _ := os.Open(fichier)

		if optionC == "-c" && nombre == 0 {
			fmt.Printf("tail: option requires an argument -- 'c'\n")
			return
		}

		if optionC == "-c" { //Avec le "-c"
			sizeFileWithOption, _ := fichierOpen.Stat()

			if checkFile >= 1 && notFile == true {
				if sizeFileWithOption.Size() <= nombre {
					sizeFile, _ := fichierOpen.Stat()
					data := make([]byte, sizeFile.Size())
					count, _ := fichierOpen.Read(data)
					for i := 0; i < checkFile; i++ {
						if i == checkFile-1 {
							fmt.Printf("===> %s <===\n", fichier)
							fmt.Printf("%s", data[:count])
							fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", NOTFILE)
							return
						}
						fmt.Printf("===> %s <===\n", fichier)
						fmt.Printf("%s\n", data[:count])
					}
				} else {

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
					for i := 0; i < checkFile; i++ {
						if i == checkFile-1 {
							fmt.Printf("===> %s <===\n", fichier)
							for j := len(tab) - 1; j >= 0; j-- {
								fmt.Printf(string(tab[j]))
							}
							fmt.Printf("tail: cannot open '%s' for reading: No such file or directory\n", NOTFILE) //je suis ici
							return
						}
						fmt.Printf("===> %s <===\n", fichier)
						for j := len(tab) - 1; j >= 0; j-- {
							fmt.Printf(string(tab[j]))
						}
						z01.PrintRune('\n')
					}

				}
			}

			if sizeFileWithOption.Size() <= nombre { //Print tt le fichier

				if checkFile == 1 {
					sizeFile, _ := fichierOpen.Stat()     //Récupère des infos sur le fichier
					data := make([]byte, sizeFile.Size()) //sizeFile.Size() récupère la taille du fichier et crée un tableau de cette taille
					count, _ := fichierOpen.Read(data)
					fmt.Printf("%s\n", data[:count])
				} else {
					sizeFile, _ := fichierOpen.Stat()
					data := make([]byte, sizeFile.Size())
					count, _ := fichierOpen.Read(data)
					for i := 0; i < checkFile; i++ {
						if i == checkFile-1 {
							fmt.Printf("===> %s <===\n", fichier)
							fmt.Printf("%s", data[:count])
							break
						}
						fmt.Printf("===> %s <===\n", fichier)
						fmt.Printf("%s\n", data[:count])
					}
				}
			} else { // Print le tableau
				if checkFile == 1 {
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
				} else {
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
					for i := 0; i < checkFile; i++ {
						if i == checkFile-1 {
							fmt.Printf("===> %s <===\n", fichier)
							for j := len(tab) - 1; j >= 0; j-- {
								fmt.Printf(string(tab[j]))
							}
							break
						}
						fmt.Printf("===> %s <===\n", fichier)
						for j := len(tab) - 1; j >= 0; j-- {
							fmt.Printf(string(tab[j]))
						}
						z01.PrintRune('\n')
					}
				}
			}

		} else { //Sans le "-c"  Print tous le fichier, GOOD
			if checkFile == 1 {
				sizeFile, _ := fichierOpen.Stat()     //Récupère des infos sur le fichier
				data := make([]byte, sizeFile.Size()) //sizeFile.Size() récupère la taille du fichier et crée un tableau de cette taille
				count, _ := fichierOpen.Read(data)
				fmt.Printf("%s\n", data[:count])
			} else {
				sizeFile, _ := fichierOpen.Stat()     //Récupère des infos sur le fichier
				data := make([]byte, sizeFile.Size()) //sizeFile.Size() récupère la taille du fichier et crée un tableau de cette taille
				count, _ := fichierOpen.Read(data)
				for i := 0; i < checkFile; i++ {
					if i == checkFile-1 {
						fmt.Printf("===> %s <===\n", fichier)
						fmt.Printf("%s", data[:count])
						break
					}
					fmt.Printf("===> %s <===\n", fichier)
					fmt.Printf("%s\n", data[:count])
				}
			}
		}
	}
}

var (
	args       = os.Args[1:]
	tailleArgs = len(args)
	notFile    = true
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

//FINI
