package run

import (
	"errors"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/lukasl-dev/ben/cmd/ben/handler"
	"github.com/lukasl-dev/ben/internal/spinner"
	"github.com/lukasl-dev/ben/sheet"
	"github.com/lukasl-dev/ben/sheet/job"
	"github.com/lukasl-dev/ben/sheet/step"
	"github.com/lukasl-dev/ben/sheet/steprunner"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// run contains flags that have been passed to the command.
type run struct {
	// sheetURI is the URI of the sheet to run.
	sheetURI string

	// skipConfirm indicates that the confirmation prompt should be skipped.
	skipConfirm bool

	// workDir is the working directory to use.
	workDir string
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
	set.StringVarP(&r.workDir, "working-directory", "w", "", "Working directory to run the sheet in. By default, the current directory is used.")
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

	if err := runJobs(s.Jobs, r.workDir); err != nil {
		return err
	}
	return nil
}

// loadSheet loads a sheet from the given URI and shows a spinner while it is
// loading.
func loadSheet(uri string) (*sheet.Sheet, error) {
	s := spinner.New(spinner.Options{
		Pending: "Sheet is loading...",
		Success: "Sheet has been loaded.",
		Error:   "Failed to load sheet.",
	})
	s.Start()

	loaded, err := sheet.Load(uri)
	if err != nil {
		s.Error()
	}
	s.Success()
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
func runJobs(jobs []job.Job, workDir string) error {
	for _, j := range jobs {
		if err := runJob(j, workDir); err != nil {
			return err
		}
	}
	return nil
}

// runJob runs j in workDir.
func runJob(j job.Job, workDir string) error {
	spin := createJobSpinner(j)
	spin.Start()

	if err := runSteps(spin, j, workDir); err != nil {
		return err
	}
	spin.Success()
	return nil
}

// createJobSpinner creates a spinner for the given job.
func createJobSpinner(j job.Job) *spinner.Spinner {
	return spinner.New(spinner.Options{
		Pending: fmt.Sprintf("Job %s: 0/%d", j.Name, len(j.Steps)),
		Success: fmt.Sprintf("Job %s: Completed", j.Name),
		Error:   fmt.Sprintf("Job: %s: Failed to complete", j.Name),
	})
}

// runSteps runs the steps in workDir. The spin is updated after each step.
func runSteps(spin *spinner.Spinner, j job.Job, workDir string) error {
	for i, st := range j.Steps {
		spin.Update(fmt.Sprintf("Job %s: %d/%d: %s", j.Name, i+1, len(j.Steps), st.Name))

		err := runStep(st, workDir)
		if err != nil {
			spin.Error()
			return stepFailed(st, err)
		}
	}
	return nil
}

// runStep runs st.
func runStep(st step.Step, workDir string) error {
	return steprunner.Step(st, steprunner.StepOptions{
		Command: steprunner.CommandOptions{
			WorkDir: workDir,
		},
	})
}

// stepFailed wraps err in a handler.Error.
func stepFailed(st step.Step, err error) *handler.Error {
	return &handler.Error{
		Err:   err,
		Title: fmt.Sprintf("Failed to run step '%s'", st.Name),
	}
}
