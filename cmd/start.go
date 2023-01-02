/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"chatgpt-cli/util"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/fatih/color"
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	gogpt "github.com/sashabaranov/go-gpt3"
	"github.com/spf13/cobra"
)

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the Chat with ChatGPT",
	Long: `To start the chat box, please use the "start" command,
you will need to open a command prompt or terminal window on your computer.`,
	Run: func(cmd *cobra.Command, args []string) {
		username, _ := cmd.Flags().GetString("user")
		displayChat(username)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
	var Username string
	startCmd.Flags().StringVarP(&Username, "user", "u", "", "Your Username")
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
	// Username  string
	Text string
}

func chatGPT3(prompt string) string {
	envConfig, err := util.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	c := gogpt.NewClient(envConfig.ChatGPTToken)
	ctx := context.Background()

	req := gogpt.CompletionRequest{
		Model:     "text-ada-001",
		MaxTokens: 500,
		Prompt:    prompt,
	}
	resp, err := c.CreateCompletion(ctx, req)
	if err != nil {
		return ""
	}
	return resp.Choices[0].Text
}

func displayChat(username string) {

	// Colorize the string using the Color object
	colorFuncGreen := color.New(color.FgGreen).SprintFunc()
	colorFuncRed := color.New(color.FgRed).SprintFunc()

	app := tview.NewApplication()
	inputField := tview.NewInputField()
	layout := tview.NewFlex().SetDirection(tview.FlexRow)

	introOutput := "Welcome to the ChatGPT " + colorFuncGreen(username) + "! Have fun and find something helpful with " + colorFuncRed("Alfred")

	// Create a text view for the introduction.
	intro := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true)

	writer := tview.ANSIWriter(intro)
	writer.Write([]byte(introOutput))
	intro.SetBorder(true)
	layout.AddItem(intro, 3, 1, false)
	// Create a layout that divides the screen into three rows.

	history := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWrap(true)
	history.SetBorder(true).SetTitle("Chat History")
	layout.AddItem(history, 0, 1, true)

	// AddItem(messages, 0, 1, false).
	// AddItem(usernameField, 1, 0, false).

	// Handle user input.
	inputField.SetLabel("Enter message: ").
		SetDoneFunc(func(key tcell.Key) {
			if key == tcell.KeyEnter {
				// Get the current timestamp and username.
				timestampStr := time.Now().Format(time.RFC1123)
				timestamp, err := time.Parse(time.RFC1123, timestampStr)
				if err != nil {
					fmt.Println(err)
					return
				}

				// Create a new message.
				message := &Message{
					Timestamp: timestamp,
					Text:      inputField.GetText(),
				}

				// the chat will quit when user type :exit
				if message.Text == ":exit" {
					app.Stop()
				}

				// Input
				writerHistory := tview.ANSIWriter(history)
				output := fmt.Sprintf("[%s] %s: %s", message.Timestamp, colorFuncGreen(username), message.Text)
				outputBytes := []byte(output + "\n")
				writerHistory.Write(outputBytes)
				history.ScrollToEnd()
				inputField.SetText("")

				//Response message from chatGPT
				resp := chatGPT3(message.Text)
				respFromChatGPT := fmt.Sprintf("[%s] %s: %s", message.Timestamp, colorFuncRed("Alfred"), resp)
				outputBytes_1 := []byte(respFromChatGPT + "\n\n")
				writerHistory.Write(outputBytes_1)
				history.ScrollToEnd()

				inputField.SetText("")
			}
		})
	inputField.SetBorder(true)
	layout.AddItem(inputField, 3, 1, true)
	// Start the application.
	if err := app.SetRoot(layout, true).SetFocus(inputField).Run(); err != nil {
		fmt.Println(err)
	}

}
