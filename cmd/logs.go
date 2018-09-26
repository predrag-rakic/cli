package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	client "github.com/semaphoreci/cli/api/client"
	"github.com/semaphoreci/cli/cmd/utils"
	"github.com/semaphoreci/cli/config"
	"github.com/spf13/cobra"
)

type Event struct {
	Timestamp int32  `json:"timestamp"`
	Type      string `json:"event"`
	Output    string `json:"output"`
}

type Events struct {
	Events []Event `json:"events"`
}

var logsCmd = &cobra.Command{
	Use:   "logs [JOB ID]",
	Short: "Display logs generated by a job.",
	Long:  ``,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]

		c := client.NewJobsV1AlphaApi()
		job, err := c.GetJob(id)

		utils.Check(err)

		url := fmt.Sprintf("https://%s/jobs/%s/raw_logs.json", config.GetHost(), job.Metadata.Id)

		req, err := http.NewRequest("GET", url, nil)

		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Authorization", fmt.Sprintf("Token %s", config.GetAuth()))

		client := &http.Client{}
		resp, err := client.Do(req)

		utils.Check(err)

		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)

		events := Events{}
		json.Unmarshal(body, &events)

		for _, e := range events.Events {
			if e.Type == "cmd_output" {
				fmt.Println(e.Output)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(logsCmd)
}
