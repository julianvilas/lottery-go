package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/julianvilas/gmailer"
	lottery "github.com/julianvilas/lottery-go"
	"github.com/spf13/cobra"
)

var region string

// monitor represents the monitor command
var monitorCmd = &cobra.Command{
	Use:   "monitor <email> <number>...",
	Short: "monitors your lottery numbers and notifies you",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("incorrect number of args, want 2 or more, got %v", len(args))
		}

		email := args[0]
		numbers := args[1:]

		n, err := stringToInt(numbers...)
		if err != nil {
			return err
		}

		return monitor(email, n)
	},
}

func init() {
	monitorCmd.Flags().StringVarP(&region, "region", "r", "eu-west-1", "sets the aws-region to send the email using ses")

	RootCmd.AddCommand(monitorCmd)
}

func monitor(email string, numbers []int) error {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(region),
	}))
	svc := ses.New(sess)

	gm := gmailer.New(svc)

	var err error
	var done bool

	if verbose {
		log.Printf("checking %v numbers, notifying %v", numbers, email)
	}

	c := time.Tick(1 * time.Minute)
	for _ = range c {
		if verbose {
			log.Println("checking numbers again")
		}

		numbers, done, err = checkNumbers(gm, email, numbers)
		if err != nil {
			return err
		}
		if done {
			break
		}
	}

	return nil
}

func checkNumbers(gm *gmailer.Mailer, email string, numbers []int) (nums []int, done bool, err error) {
	res, err := lottery.CheckNumbers(numbers...)
	if err != nil {
		return nil, false, err
	}

	if verbose {
		log.Printf("%+v", res)
	}

loop:
	for _, sr := range res {
		switch {
		case sr.Error != 0:
			return nil, false, fmt.Errorf("api returned error: %v", sr.Error)
		case sr.Status < 0 || sr.Status > 4:
			return nil, false, fmt.Errorf("api returned unknown status: %v", sr.Status)
		case sr.Status == 0:
			if verbose {
				log.Println("raffle has not started yet. Skipping...")
			}
			break loop
		case sr.Prize > 0:
			if err := gm.SendRaw(gmailer.Email{
				Subject: "Number has been awarded",
				Body:    fmt.Sprintf("The number %v has been awarded with %v. The status of the raffle is %v.", sr.Num, sr.Prize, sr.Status),
				From:    email,
				Dest:    []string{email},
			}); err != nil {
				return nil, false, err
			}
			numbers = deleteNumber(numbers, sr.Num)

			if sr.Status == 4 {
				done = true
			}
		}
	}

	return numbers, done, nil
}

func deleteNumber(numbers []int, number int) []int {
	var pos int
	found := false

	for i, val := range numbers {
		if val == number {
			pos = i
			found = true
			break
		}
	}

	if found {
		numbers = append(numbers[:pos], numbers[pos+1:]...)
	}

	if verbose {
		log.Printf("Deleting %v from %+v, it was found: %v", number, numbers, found)
	}

	return numbers
}
