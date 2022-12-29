/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	tcell "github.com/gdamore/tcell/v2"
	tview "github.com/rivo/tview"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		chatGPT3()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Message struct {
	Timestamp time.Time
	Username  string
	Text      string
}

func chatGPT3() {
	// c := gogpt.NewClient("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")
	// ctx := context.Background()

	// result := ""

	// for result != ":exit" {
	// 	label := emoji.WhiteQuestionMark
	// 	prompt := promptui.Prompt{
	// 		Label: label,
	// 	}
	// 	res, err := prompt.Run()

	// 	if err != nil {
	// 		fmt.Printf("Prompt failed %v\n", err)
	// 		return
	// 	}
	// 	result = res
	// 	fmt.Printf("You answered %s\n", result)
	// }

	// req := gogpt.CompletionRequest{
	// 	Model:     "text-ada-001",
	// 	MaxTokens: 500,
	// 	Prompt:    "Hello, how are you?",
	// }

	// resp, err := c.CreateCompletion(ctx, req)
	// if err != nil {
	// 	return
	// }
	// fmt.Println(resp.Choices[0].Text)

	app := tview.NewApplication()

	// Create a new flex layout that divides the screen into two columns.
	flex := tview.NewFlex().SetDirection(tview.FlexRow)

	// Create a chat history box on the left.
	history := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true)
	history.SetBorder(true).SetTitle("Chat History")
	flex.AddItem(history, 0, 1, true)

	var input *tview.InputField
	// Create an input field on the bottom.
	input = tview.NewInputField().
		SetLabel("Enter message: ").
		SetFieldWidth(30).
		SetDoneFunc(func(key tcell.Key) {
			// When the user hits enter, add their message to the chat history.
			history.Write([]byte(input.GetText()))
			input.SetText("")
		})
	flex.AddItem(input, 1, 1, false)

	// Start the application.
	if err := app.SetRoot(flex, true).Run(); err != nil {
		fmt.Println(err)
	}

}

// app := tview.NewApplication()

// 	messages := tview.NewList()
// 	inputField := tview.NewInputField()
// 	usernameField := tview.NewInputField().SetLabel("Username: ")

// 	// Set the default username.
// 	usernameField.SetText("User")

// 	// Create a layout that divides the screen into three rows.
// 	// Create a layout that divides the screen into three rows.
// 	layout := tview.NewFlex().SetDirection(tview.FlexRow).
// 		AddItem(tview.NewBox().SetBorder(true).SetTitle("Messages").AddItem(messages, 0, 1, false), 0, 1, false).
// 		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
// 			AddItem(tview.NewBox().SetBorder(true).AddItem(inputField, 1, 0, false), 1, 0, false).
// 			AddItem(tview.NewBox().SetBorder(true).AddItem(inputField, 1, 0, false), 1, 0, false))

// 	// Handle user input.
// 	inputField.SetDoneFunc(func(key tcell.Key) {
// 		if key == tcell.KeyEnter {
// 			// Get the current timestamp and username.
// 			timestampStr := time.Now().Format(time.RFC1123)
// 			timestamp, err := time.Parse(time.RFC1123, timestampStr)
// 			if err != nil {
// 				fmt.Println(err)
// 				return
// 			}
// 			username := usernameField.GetText()

// 			// Create a new message.
// 			message := &Message{
// 				Timestamp: timestamp,
// 				Username:  username,
// 				Text:      inputField.GetText(),
// 			}

// 			// Clear the input field.
// 			inputField.SetText("")

// 			// Add the message to the list.
// 			messages.AddItem(fmt.Sprintf("[%s] %s: %s", message.Timestamp, message.Username, message.Text), "", 0, nil)
// 			app.SetFocus(messages)
// 		}
// 	})

// 	// Start the application.
// 	if err := app.SetRoot(layout, true).SetFocus(inputField).Run(); err != nil {
// 		fmt.Println(err)
// 	}
