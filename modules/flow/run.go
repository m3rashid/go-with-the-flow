package flow

import (
	"context"
	"log"

	"github.com/m3rashid/go-with-the-flow/db"
	flow "github.com/m3rashid/go-with-the-flow/modules/flow/schema"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	v8 "rogchap.com/v8go"
)

var DefaultFunctions = map[string]string{
	"add":      `const add = (...args) => args.reduce((a, b) => a + b, 0);`,
	"multiply": `const multiply = (...args) => args.reduce((a, b) => a * b, 1);`,
	"divide":   `const divide = (a, b) => a / b;`,
	"subtract": `const subtract = (a, b) => a - b;`,
	"average":  `const average = (...args) => sum(...args) / args.length;`,
}

func RunFunc(function string, globalFunctions ...string) interface{} {
	contextFunctions := []string{}
	contextFunctions = append(contextFunctions, globalFunctions...)
	for _, fun := range DefaultFunctions {
		contextFunctions = append(contextFunctions, fun)
	}

	ctx := v8.NewContext()
	for _, fun := range contextFunctions {
		ctx.RunScript(fun, "main.js")
	}
	ctx.RunScript(function, "main.js")

	data, err := ctx.RunScript("data", "main.js")
	if err != nil {
		log.Println("Error in running data script", err)
		return nil
	}

	log.Println(data)
	return data
}

func Run(data map[string]interface{}, operationType interface{}, collectionName string, client *mongo.Client) {
	var workflow flow.Workflow

	collection := db.OpenCollection(client, flow.WORKFLOWS_MODEL_NAME)
	err := collection.FindOne(context.Background(), bson.M{
		"currentTrigger": operationType,
		"collections":    bson.A{collectionName},
	}).Decode(&workflow)
	if err != nil {
		return
	}

	if workflow.CurrentNode >= len(workflow.Nodes) {
		return
	}

	RunFunc((workflow.Nodes[workflow.CurrentNode]).Handler, workflow.GlobalFunctions)
}

func RunDeleteFlow(documentKey string, w *flow.Workflow, client *mongo.Client) {
	//
}
