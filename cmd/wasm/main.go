package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"syscall/js"

// 	"github.com/darchlabs/infra/pkg/project"
// )

// func main() {
// 	done := make(chan struct{}, 0)

// 	js.Global().Set("provideInfra", js.FuncOf(provideInfra))

// 	<-done
// }

// func provideInfra(this js.Value, args []js.Value) interface{} {
// 	// define context for tracking pulumi proccess
// 	// ctx := context.Background()

// 	// parse project values
// 	p, err := project.ToProjectWasm(args)
// 	if err != nil {
// 		// temporal
// 		p.Provisioning = false
// 		p.Error = err.Error()
// 		b, err := json.Marshal(p)
// 		if err != nil {
// 			return err.Error()
// 		}

// 		return string(b)
// 	}

// 	fmt.Printf("%+v \n", p)

// 	// run pulumi stack
// 	p.Provisioning = true
// 	// err = pulumiinfra.Pulumi(ctx, *p)
// 	// if err != nil {
// 	// 	// temporal
// 	// 	p.Provisioning = false
// 	// 	p.Error = err.Error()
// 	// 	b, err := json.Marshal(p)
// 	// 	if err != nil {
// 	// 		return err.Error()
// 	// 	}

// 	// 	return string(b)
// 	// }

// 	// parse project struct to json string
// 	p.Provisioning = false
// 	p.Ready = true
// 	b, err := json.Marshal(p)
// 	if err != nil {
// 		return err.Error()
// 	}

// 	// return project json to frontend
// 	return string(b)
// }
