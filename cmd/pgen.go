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
	"github.com/dacors-m/fingerp/utils"
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
	Run:   genPass,
}

func init() {
	rootCmd.AddCommand(pgen)
	pgen.Flags().IntP("characters", "c", 2, "amount of special characters")
	pgen.Flags().IntP("numbers", "n", 2, "amount of numbers")
	pgen.Flags().IntP("length", "l", 8, "password length")
	pgen.Flags().BoolP("store", "s", true, "store password")
	pgen.Flags().StringP("reference", "r", "", "password reference")
	pgen.MarkFlagRequired("reference")
}

func genPass(cmd *cobra.Command, args []string) {

	// flags
	tl, err := cmd.Flags().GetInt("length")
	nu, err := cmd.Flags().GetInt("numbers")
	c, err := cmd.Flags().GetInt("characters")
	s, err := cmd.Flags().GetBool("store")
	ref, err := cmd.Flags().GetString("reference")

	if err != nil {
		color.Red("Invalid flags")
	}

	passLenght := tl - nu - c
	if passLenght <= 0 {
		color.Red("Invalid length")
		return
	}

	// generatge pass
	pL := utils.GetRandChars(passLenght)
	pCh := utils.GetRandCharacters(c)
	pNu := strconv.Itoa(rand.Int())

	psw := pL + pCh + pNu[:nu]

	clipboard.WriteAll(psw)

	if s {
		fmt.Println("store")
		fmt.Println(ref)
	}

	fmt.Printf("New password: %s\n", psw)
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
