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
	"strings"

	"github.com/rogerscuall/gotm/adapters/db"
	"github.com/rogerscuall/gotm/packages/types/tasks"
	"github.com/rogerscuall/gotm/ports"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var updateField string
var updateValue string

// - represents the udpate command
var update = &cobra.Command{
	Use:   "update",
	Short: "Update a field of a task",
	Long: `It is used to update a field of a task.
The field needs to be specified in the command line.
gotm update <TASK_ID> -f <TASK_FIELD> -v "UPDATE"`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {

		dbName = viper.GetString("db_name")
		var dbAdapter ports.DbPort
		var err error
		dbAdapter, err = db.NewAdapter(dbName)
		defer dbAdapter.CloseDbConnection()

		taskID := strings.TrimSpace(args[0])
		fmt.Println("Looking for task ID: ", taskID)
		bTask, err := dbAdapter.GetVal(taskID)
		if err != nil {
			log.Fatal("The task was not found: ", err)
		}

		if !tasks.IsFieldUpdatable(updateField) {
			log.Fatal("The field is not updatable: ", updateField)
		}

		task := tasks.Task{}
		err = task.FromBytes(bTask)
		if err != nil {
			log.Fatal("We could not parse the task: ", err)
		}

		switch updateField {
		case "name":
			task.Name = updateValue
		case "description":
			task.Description = updateValue
			bTask = tasks.ToBytes(task)
			err = dbAdapter.SetVal(taskID, bTask)
			if err != nil {
				log.Fatal("Task was not updated: ", err)
			}
			fmt.Println("Task was updated")
		case "related":
			// missing the validation of the task that is going to be added to the updateValue
			//task.TaskRelations = append(task.TaskRelations, updateValue)
			fmt.Println("Not implemented yet")
		case "parentProject":
			fmt.Println("Not implemented yet")
		}
	},
}

func init() {

	update.Flags().StringVarP(&updateField, "field", "f", "description", "field to be updated")
	update.Flags().StringVarP(&updateValue, "value", "v", "", "updates a field of a task")
	rootCmd.AddCommand(update)
	//update.AddCommand(updateFieldCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// update.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// update.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
