package telegram

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"

	"gopkg.in/telebot.v3"
)

func TestSendString(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Hello",
			args: args{
				message: "Hello, world!",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendString(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			bs, err := json.MarshalIndent(got, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(bs))
		})
	}
}

func TestSendHtml(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name    string
		args    args
		want    *telebot.Message
		wantErr bool
	}{
		{
			name: "Sample",
			args: args{
				message: "This is <b>bold</b>, <i>italic</i>, <u>underlined</u>, and <s>strikethrough</s> text. " +
					"Here is a <code>single line of code</code> and a <pre>block of code</pre>. " +
					"Also, here is a link: <a href='https://example.com'>Example.com</a> and a <tg-spoiler>spoiler</tg-spoiler>.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SendHtml(tt.args.message)
			if (err != nil) != tt.wantErr {
				t.Errorf("SendMarkdownV2() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			bs, err := json.MarshalIndent(got, "", "  ")
			if err != nil {
				log.Fatalln(err)
			}
			fmt.Println(string(bs))
		})
	}
}
