package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	userdomain "github.com/dacors-m/fingerp/domain/user"
	"github.com/dacors-m/fingerp/passwordrepository"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var puser = &cobra.Command{
	Use:   "puser",
	Short: "User Setup",
	Run:   setupUser,
}

func init() {
	rootCmd.AddCommand(puser)
	puser.Flags().StringP("user", "u", "", "user name")
	puser.MarkFlagRequired("user")
}

func setupUser(cmd *cobra.Command, args []string) {
	u, err := cmd.Flags().GetString("user")

	if err != nil {
		color.Red("Invalid flags")
	}

	d := color.New(color.FgCyan)
	d.Print("Type the new user password \n")

	buf := bufio.NewReader(os.Stdin)
	res, err := buf.ReadString('\n')
	if err != nil {
		color.Red("Invalid password")
	}

	p := strings.Replace(res[:len(res)-1], " ", "", -1)

	fmt.Println(u, p)

	passRepo := passwordrepository.NewPasswordRepository()
	userDomain := userdomain.NewUserDomain(passRepo)

	err = userDomain.SetupUser(u, p)
	if err != nil {
		fmt.Println("fail")
	}

}
