package runsummary

import (
	"os"

	"github.com/vercel/turbo/cli/internal/cache"
	"github.com/vercel/turbo/cli/internal/fs"
	"github.com/vercel/turbo/cli/internal/turbopath"
	"github.com/vercel/turbo/cli/internal/util"
)

// TaskCacheSummary is an extended version of cache.ItemStatus
// that includes TimeSaved and some better data.
type TaskCacheSummary struct {
	Local     bool   `json:"local"`            // Deprecated, but keeping around for --dry=json
	Remote    bool   `json:"remote"`           // Deprecated, but keeping around for --dry=json
	Status    string `json:"status"`           // should always be there
	Source    string `json:"source,omitempty"` // can be empty on status:miss
	TimeSaved int    `json:"timeSaved"`        // always include, but can be 0
}

// NewTaskCacheSummary decorates a cache.ItemStatus into a TaskCacheSummary
// Importantly, it adds the derived keys of `source` and `status` based on
// the local/remote booleans. It would be nice if these were just included
// from upstream, but that is a more invasive change.
func NewTaskCacheSummary(itemStatus cache.ItemStatus) TaskCacheSummary {
	status := cache.CacheEventMiss
	if itemStatus.Hit {
		status = cache.CacheEventHit
	}
	var source string
	if itemStatus.Hit {
		source = itemStatus.Source
	}

	cs := TaskCacheSummary{
		Status:    status,
		Source:    source,
		TimeSaved: itemStatus.TimeSaved,
	}

	// Assign these deprecated fields Local and Remote based on the information available
	// in the itemStatus. Note that these fields are problematic, because an ItemStatus isn't always
	// the composite of both local and remote caches. That means that an ItemStatus might say it
	// was a local cache hit, and we return remote: false here. That's misleading because it does
	// not mean that there is no remote cache hit, it _could_ mean that we never checked the remote
	// cache. These fields are being deprecated for this reason.
	cs.Local = itemStatus.Hit && itemStatus.Source == cache.CacheSourceFS
	cs.Remote = itemStatus.Hit && itemStatus.Source == cache.CacheSourceRemote

	return cs
}

// TaskSummary contains information about the task that was about to run
// TODO(mehulkar): `Outputs` and `ExcludedOutputs` are slightly redundant
// as the information is also available in ResolvedTaskDefinition. We could remove them
// and favor a version of Outputs that is the fully expanded list of files.
type TaskSummary struct {
	TaskID                 string                                `json:"taskId,omitempty"`
	Task                   string                                `json:"task"`
	Package                string                                `json:"package,omitempty"`
	Hash                   string                                `json:"hash"`
	ExpandedInputs         map[turbopath.AnchoredUnixPath]string `json:"inputs"`
	ExternalDepsHash       string                                `json:"hashOfExternalDependencies"`
	CacheSummary           TaskCacheSummary                      `json:"cache"`
	Command                string                                `json:"command"`
	CommandArguments       []string                              `json:"cliArguments"`
	Outputs                []string                              `json:"outputs"`
	ExcludedOutputs        []string                              `json:"excludedOutputs"`
	LogFile                string                                `json:"logFile"`
	Dir                    string                                `json:"directory,omitempty"`
	Dependencies           []string                              `json:"dependencies"`
	Dependents             []string                              `json:"dependents"`
	ResolvedTaskDefinition *fs.TaskDefinition                    `json:"resolvedTaskDefinition"`
	ExpandedOutputs        []turbopath.AnchoredSystemPath        `json:"expandedOutputs"`
	Framework              string                                `json:"framework"`
	EnvMode                util.EnvMode                          `json:"envMode"`
	EnvVars                TaskEnvVarSummary                     `json:"environmentVariables"`
	DotEnv                 turbopath.AnchoredUnixPathArray       `json:"dotEnv"`
	Execution              *TaskExecutionSummary                 `json:"execution,omitempty"` // omit when it's not set
}

// GetLogs reads the Logfile and returns the data
func (ts *TaskSummary) GetLogs() []byte {
	bytes, err := os.ReadFile(ts.LogFile)
	if err != nil {
		return []byte{}
	}
	return bytes
}

// TaskEnvConfiguration contains the environment variable inputs for a task
type TaskEnvConfiguration struct {
	Env            []string `json:"env"`
	PassThroughEnv []string `json:"passThroughEnv"`
}

// TaskEnvVarSummary contains the environment variables that impacted a task's hash
type TaskEnvVarSummary struct {
	Specified TaskEnvConfiguration `json:"specified"`

	Configured  []string `json:"configured"`
	Inferred    []string `json:"inferred"`
	PassThrough []string `json:"passthrough"`
}

// cleanForSinglePackage converts a TaskSummary to remove references to workspaces
func (ts *TaskSummary) cleanForSinglePackage() {
	dependencies := make([]string, len(ts.Dependencies))
	for i, dependency := range ts.Dependencies {
		dependencies[i] = util.StripPackageName(dependency)
	}
	dependents := make([]string, len(ts.Dependents))
	for i, dependent := range ts.Dependents {
		dependents[i] = util.StripPackageName(dependent)
	}
	task := util.StripPackageName(ts.TaskID)

	ts.TaskID = task
	ts.Task = task
	ts.Dependencies = dependencies
	ts.Dependents = dependents
	ts.Dir = ""
	ts.Package = ""
}
