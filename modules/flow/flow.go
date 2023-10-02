package flow

type Step struct {
	Name    string
	Handler func(data interface{}) bool
}

type Flow struct {
	Name       string
	Trigger    string
	Collection string
	Steps      []Step
}

var RegisteredFlows []Flow

func RunFlow(collectionName string, operationType interface{}, data interface{}) {
	for _, flow := range RegisteredFlows {
		if flow.Collection == collectionName && flow.Trigger == operationType {
			// handle each step sequentially in the flow
			flowRunner := make(chan int)
			for {
				step := <-flowRunner
				if step == len(flow.Steps) {
					break
				}
				flow.Steps[step].Handler(data)
				flowRunner <- step + 1
			}

			// for _, step := range flow.Steps {
			// 	step.Handler(data)
			// }
		}
	}
}
