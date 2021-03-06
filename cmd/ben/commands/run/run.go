package run

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/lukasl-dev/ben/cmd/ben/handler"
	"github.com/lukasl-dev/ben/internal/spinner"
	"github.com/lukasl-dev/ben/loader"
	"github.com/lukasl-dev/ben/runner"
	"github.com/lukasl-dev/ben/sheet"
	"github.com/lukasl-dev/ben/sheet/job"
	"github.com/lukasl-dev/ben/sheet/step"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"os"
)

// run contains flags that have been passed to the command.
type run struct {
	// sheetURI is the URI of the sheet to run.
	sheetURI string

	// skipConfirm indicates that the confirmation prompt should be skipped.
	skipConfirm bool
}

// Command constructs the 'run' command.
func Command() *cobra.Command {
	run := new(run)
	cmd := &cobra.Command{
		Use:     "run",
		Short:   "Runs a sheet in the given working directory.",
		RunE:    run.run,
		PreRunE: run.pre,
	}
	run.flags(cmd.Flags())
	return cmd
}

// flags binds the flags to r.
func (r *run) flags(set *pflag.FlagSet) {
	set.StringVarP(&r.sheetURI, "sheet", "s", "ben.yml", "Path or URL to the sheet configuration file.")
}

// pre prepares r for running.
func (r *run) pre(cmd *cobra.Command, _ []string) error {
	if r.sheetURI == "" {
		return cmd.Usage()
	}
	return nil
}

// run runs the command.
func (r *run) run(*cobra.Command, []string) error {
	s, err := loadSheet(r.sheetURI)
	if err != nil {
		return loadingFailed(err)
	}
	if err := s.Validate(); err != nil {
		return invalidSheet(err)
	}

	if !r.skipConfirm {
		_, err := promptConfirmation(s)
		if err != nil {
			return confirmationFailed(err)
		}
	}

	if err := runJobs(*s); err != nil {
		return err
	}
	return nil
}

// loadSheet loads a sheet from the given URI and shows a spinner while it is
// loading.
func loadSheet(uri string) (*sheet.Sheet, error) {
	s := spinner.New("Loading sheet...", spinner.Options{})
	s.Start()

	loaded, err := loader.Load(uri, loader.NoOptions)
	if err != nil {
		s.Error(fmt.Sprintf("Failed to load sheet: %s", err))
	}
	s.Success("Sheet has been loaded.")
	return loaded, err
}

// loadingFailed wraps err in a handler.Error.
func loadingFailed(err error) *handler.Error {
	return &handler.Error{
		Err:   err,
		Title: "Failed to load sheet",
		Suggestions: []string{
			"Make sure the sheet exists and is valid.",
		},
	}
}

// invalidSheet wraps err in a handler.Error.
func invalidSheet(err error) *handler.Error {
	return &handler.Error{
		Err:   err,
		Title: "The provided sheet contains invalid data",
	}
}

// promptConfirmation prompts the user to confirm to run the sheet.
func promptConfirmation(sheet *sheet.Sheet) (b bool, err error) {
	confirm := &survey.Confirm{
		Message: fmt.Sprintf("Are you sure that you want to run the sheet '%s'?", sheet.Name),
		Default: false,
	}
	err = survey.AskOne(confirm, &b, survey.WithValidator(validateConfirmation))
	return b, err
}

// validateConfirmation validates ans to be true.
func validateConfirmation(ans interface{}) error {
	if ans == true {
		return nil
	}
	return errors.New("you must confirm to run the sheet")
}

// confirmationFailed wraps err in a handler.Error.
func confirmationFailed(err error) *handler.Error {
	return &handler.Error{
		Err:   err,
		Title: "Failed to show confirmation prompt",
	}
}

// runJobs runs the jobs in workDir.
func runJobs(s sheet.Sheet) error {
	for i, j := range s.Jobs {
		if err := runJob(s, j, i+1, len(s.Jobs)); err != nil {
			return err
		}
	}
	return nil
}

// runJob runs j in workDir.
func runJob(s sheet.Sheet, j job.Job, pos, size int) error {
	spin := createJobSpinner(j, pos, size)
	spin.Start()

	for i, stp := range j.Steps {
		spin.Update(fmt.Sprintf("%s (%d/%d): %s (%d/%d)", j.Name, pos, size, stp.Name, i+1, len(j.Steps)))
		err := runStep(j, stp)
		if err != nil {
			spin.Error(fmt.Sprintf("%s (%d/%d): Failed on step '%s'", j.Name, pos, size, stp.Name))
			return stepFailed(stp, err)
		}
	}
	spin.Success(fmt.Sprintf("%s (%d/%d): Completed", j.Name, pos, size))
	return nil
}

// createJobSpinner creates a spinner for the given job.
func createJobSpinner(j job.Job, pos, size int) *spinner.Spinner {
	text := fmt.Sprintf("%s (%d/%d): Pending...", j.Name, pos, size)
	return spinner.New(text, spinner.Options{})
}

// runStep runs st.
func runStep(j job.Job, stp step.Step) error {
	logs, _ := os.OpenFile("ben.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	logrus.SetOutput(logs)

	w := logrus.StandardLogger().WithFields(logrus.Fields{
		"job":  j.Name,
		"step": stp.Name,
	}).Writer()

	return runner.Run(stp, runner.Options{Stdout: w, Stderr: w})
}

// stepFailed wraps err in a handler.Error.
func stepFailed(st step.Step, err error) *handler.Error {
	return &handler.Error{
		Err:   err,
		Title: fmt.Sprintf("Failed to run step '%s'", st.Name),
	}
}
