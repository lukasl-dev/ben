package step

type Command struct {
	// Command is the command to execute.
	Run string `json:"run,omitempty" hcl:"run"`

	// WorkDir is the working directory for the command. It overwrites the
	// working directory of the sheet.
	WorkDir string `json:"workDir,omitempty" hcl:"workDir,optional"`

	// ExitCodes is a slice of exit codes that are considered successful.
	ExitCodes []int `json:"exitCodes,omitempty" hcl:"exitCodes,optional"`
}
