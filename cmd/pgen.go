package cmd

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	characters = "[]()/.,*&"
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()),
)

var pgen = &cobra.Command{
	Use:   "pgen",
	Short: "Password generator",

	Run: genPass,
}

func init() {
	rootCmd.AddCommand(pgen)
	pgen.Flags().IntP("characters", "c", 2, "Use special characters")
	pgen.Flags().IntP("numbers", "n", 2, "Use numbers")
	pgen.Flags().IntP("length", "l", 8, "Password length")
}

func genPass(cmd *cobra.Command, args []string) {

	// flags
	c, err := cmd.Flags().GetInt("characters")
	nu, err := cmd.Flags().GetInt("numbers")
	nl, err := cmd.Flags().GetInt("length")

	if err != nil {
		color.Red("Invalid flags")
	}

	tl := nl - nu - c
	if tl <= 0 {
		color.Red("Invalid length")
		return
	}

	// fabric pass
	pNu := strconv.Itoa(rand.Int())
	pCh := getRandCharacters(c)
	pL := getRandChars(tl)

	psw := pL + pCh + pNu[:nu]
	clipboard.WriteAll(psw)

	fmt.Printf("New password: %s\n", psw)

	sPass := savePassMsg(3)
	if sPass {

		uPsw := usageMsg()
		savePass(psw, uPsw)
	}
}
func savePass(p string, u string) {
	d := color.New(color.FgCyan)
	d.Printf("Password saved for %s ! \n", u)
}

func usageMsg() string {

	d := color.New(color.FgYellow)
	d.Print("Type the password usage\n")
	buf := bufio.NewReader(os.Stdin)
	res, err := buf.ReadString('\n')
	if err != nil {
		color.Red("Invalid alias")
	}

	return strings.Replace(res[:len(res)-1], " ", "", -1)
}

func savePassMsg(t int) bool {
	for i := 0; i < t; i++ {
		d := color.New(color.FgYellow)
		d.Print("Save password? (yes|no)\n")

		buf := bufio.NewReader(os.Stdin)
		res, err := buf.ReadString('\n')
		res = strings.Replace(res[:len(res)-1], " ", "", -1)
		if err != nil ||
			(!strings.EqualFold(res, "yes") && !strings.EqualFold(res, "no")) {
			color.Red("Invalid input")
		} else {
			switch res {
			case "yes":
				return true
			case "no":
				return false
			}
		}
	}
	return false
}

func getRandChars(l int) string {
	s := make([]byte, l)
	for i := range s {
		s[i] = charset[rand.Intn(len(charset))]
	}
	return string(s)
}

func getRandCharacters(l int) string {
	s := make([]byte, l)
	for i := range s {
		s[i] = characters[rand.Intn(len(characters))]
	}
	return string(s)
}
