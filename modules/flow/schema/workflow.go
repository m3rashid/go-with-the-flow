package flow

const WORKFLOWS_MODEL_NAME = "allWorkflows"

const RUNNING_WORKFLOWS_MODEL_NAME = "runningWorkflows"

type Workflow struct {
	Nodes           []Node   `json:"nodes" bson:"nodes" validate:"required"`
	CurrentNode     int      `json:"currentNode" bson:"currentNode" validate:"required"`
	Collections     []string `json:"collections" bson:"collections" validate:"required"`
	CurrentTrigger  string   `json:"currentTrigger" bson:"currentTrigger" validate:""`
	GlobalFunctions string   `json:"globalFunctions" bson:"globalFunctions" validate:""`
}
