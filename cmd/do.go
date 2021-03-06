/*
Copyright © 2022 Roger Gomez rogerscuall@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/rogerscuall/gotm/adapters/db"
	"github.com/rogerscuall/gotm/packages/types/tasks"
	"github.com/rogerscuall/gotm/ports"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// completeCmd represents the complete command
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Do will mark a task as completed",
	Args:  cobra.MinimumNArgs(1),
	Long: `Do will mark a task as completed.
it needs a ID in the format of: gotm do <ID>`,
	Run: func(cmd *cobra.Command, args []string) {
		dbName = viper.GetString("db_name")
		var dbAdapter ports.DbPort
		var err error
		dbAdapter, err = db.NewAdapter(dbName)
		defer dbAdapter.CloseDbConnection()

		id := args[0]
		bTask, err := dbAdapter.GetVal(id)
		if err != nil {
			log.Fatal("The task was not found: ", err)
		}
		var task tasks.Task
		task.FromBytes(bTask)
		fmt.Println(task)
		task.Complete()
	},
}

func init() {
	rootCmd.AddCommand(doCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
