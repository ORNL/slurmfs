package slurmfs

//////////////// Generics \\\\\\\\\\\\\

type SlurmError struct {
	Error string
	Error_Number int
}

type SlurmResponse struct {
	Errors []SlurmError
}

////////////// List jobs API \\\\\\\\\\\\

type JobsResponse struct {
	SlurmResponse
	Jobs []JobMeta
}

// V0.0.38_JOB_RESOURCES
type Resources struct {
	Nodes string
	Allocated_CPUs int
	Allocated_Hosts int
    // Allocated_Nodes []NodeAllocation # why?
}

// V0.0.39_JOB_INFO subset
type JobMeta struct {
	Name string
	Partition string
	Account string

	User_Name string
	Job_State string

	Tasks int
	Max_Nodes int
	Node_Count int

	Dependency string
	Requeue bool
	Restart_Cnt int

	Command string
	Environment map[string]string

	Time_Limit int64
	Deadline int64
	End_Time int64
	Eligible_Time int64
	/*Resize_Time int64 srsly
	Accrue_Time int64
	Pre_Sus_Time int64
	Suspend_Time int64
	Accrue_Time int64*/

	Required_Nodes string
	Excluded_Nodes string
	Current_Working_Directory string
	Start_Time int64
	Job_Resources Resources
	Standard_Output string
	Standard_Error string
	Exit_Code int
	//Script: string
}

////////////// Create jobs API \\\\\\\\\\\\

// V0.0.38_JOB_SUBMISSION
type JobSubmission struct {
	Script string   `json:"script"`
	Job JobProps    `json:"job"`
	Jobs []JobProps `json:"jobs"`
}

// V0.0.38_JOB_PROPERTIES 
type JobProps struct {
        Name string          `json:"name,omitempty"`
	Partition string     `json:"partition,omitempty"`
	Account string       `json:"account,omitempty"`

	Nodes int            `json:"nodes,omitempty"`
        // min,max node request
        Tasks int            `json:"tasks,omitempty"`
	Tasks_Per_Node int   `json:"tasks_per_node,omitempty"`
	GPU_Binding string   `json:"gpu_binding,omitempty"`
	GPU_Frequency string `json:"gpu_frequency,omitempty"`
	GPUs_Per_Task string `json:"gpus_per_task,omitempty"`
	No_Kill bool         `json:"no_kill,omitempty"`
        // do not auto-kill if a node fails

	Time_Limit int       `json:"time_limit,omitempty"`
        // minutes
	Deadline string       `json:"deadline,omitempty"`
        // remove if "est start" > deadline-min_walltime
	Time_Minimum int      `json:"time_minimum,omitempty"`
        // minutes
	CPUs_Per_Task int    `json:"cpus_per_task,omitempty"`
	Begin_Time int64     `json:"begin_time,omitempty"`
	
	Argv []string                 `json:"argv,omitempty"`
	Environment map[string]string `json:"environment"`

	Dependency string    `json:"dependency,omitempty"`
	Requeue bool         `json:"requeue,omitempty"`

	Reservation string   `json:"reservation,omitempty"`
	Container string     `json:"container,omitempty"`

	Current_Working_Directory string `json:"current_working_directory,omitempty"`
	Distribution string              `json:"distribution,omitempty"`

	Signal int           `json:"signal,omitempty"`
	Sig_Time int64       `json:"sig_time,omitempty"`
        // file names
	Standard_Input string  `json:"Standard_Input,omitempty"`
	Standard_Output string `json:"standard_output,omitempty"`
}

type SubmitResponse struct {
	SlurmResponse
	Job_Id int
	Step_Id string // new job step ID
	Job_Submit_User_Msg string
}

