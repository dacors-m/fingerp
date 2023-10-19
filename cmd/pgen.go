package cmd

import (
	"math/rand"
	"strconv"
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

	d := color.New(color.FgCyan, color.Bold)
	d.Printf("Gen pass: %s\n", psw)

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
