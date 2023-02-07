package main

import (
	"context"
	"example/redis"
	"fmt"
	"strconv"
	"time"
)

const defaultPort = "8080"

func main() {

	rcl := redis.NewRedis()
	go func() {
		i := 0
		for {
			err := rcl.Publish(strconv.Itoa(i))
			if err != nil {
				panic(err)
			}
			i++
		}
	}()
	ctx := context.Background()
	ch := rcl.SubscribeCh(ctx)
	tick := time.Tick(time.Second)
	i := 0
	for {
		select {
		case msg := <-ch:
			fmt.Printf("%s\r", msg)
			i++
		case <-tick:
			fmt.Printf("1sec: %d\n", i)
			i = 0
		}
	}

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = defaultPort
	// }
	// rcl := redis.NewRedis()
	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Redis: rcl}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe("localhost:"+port, nil))
}
