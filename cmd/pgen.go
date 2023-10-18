package cmd

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

const (
	characters = "[]()/.,"
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
}

func genPass(cmd *cobra.Command, args []string) {

	// flags
	c, _ := cmd.Flags().GetInt("characters")
	nu, _ := cmd.Flags().GetInt("numbers")

	// fabric pass
	pNu := strconv.Itoa(rand.Int())
	pCh := getRandCharacters(c)

	psw := pCh + pNu[:nu]
	clipboard.WriteAll(psw)

	color.Cyan(fmt.Sprint("Gen pass:"))
	color.White(psw)

}

func getRandCharacters(l int) string {
	s := make([]byte, l)
	for i := range s {
		s[i] = characters[rand.Intn(len(characters))]
	}
	return string(s)
}
