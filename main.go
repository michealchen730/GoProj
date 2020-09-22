package main

import "GoProj1.0/routers"

func main(){
	r := routers.SetupRouter()
	_ = r.Run(":8081")
}
