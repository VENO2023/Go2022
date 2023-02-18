package main

import (
	"fmt"
	"log"
	"net/http"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2023-02-01 18:22:04
 */

func sayHi(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "{\"code\":0,\"message\":\"OK\",\"data\":{\"straight\":{\"nodes\":[{\"id\":2,\"thirdCode\":\"ty735985\",\"bayonetType\":0,\"acType\":1,\"pileNo\":735985},{\"id\":1,\"thirdCode\":\"ty737050\",\"bayonetType\":0,\"acType\":2,\"pileNo\":737050}],\"edges\":[{\"from\":2,\"to\":1,\"distance\":1065}]},\"reverse\":{\"nodes\":[{\"id\":4,\"thirdCode\":\"sjz737000\",\"bayonetType\":0,\"acType\":2,\"pileNo\":737000},{\"id\":3,\"thirdCode\":\"sjz735940\",\"bayonetType\":0,\"acType\":1,\"pileNo\":735940}],\"edges\":[{\"from\":4,\"to\":3,\"distance\":1060}]}}}")
	fmt.Fprint(w, "{\n    \"code\": 0,\n    \"message\": \"OK\",\n    \"data\": {\n        \"straight\": {\n            \"nodes\": [\n                {\n                    \"id\": 1000051,\n                    \"thirdCode\": \"1\",\n                    \"bayonetType\": 0,\n                    \"acType\": 1,\n                    \"pileNo\": 200\n                },\n                {\n                    \"id\": 1000052,\n                    \"thirdCode\": \"2\",\n                    \"bayonetType\": 0,\n                    \"acType\": 0,\n                    \"pileNo\": 500\n                },\n                {\n                    \"id\": 1000085,\n                    \"thirdCode\": \"zhi1\",\n                    \"bayonetType\": 1,\n                    \"acType\": 1,\n                    \"pileNo\": 700\n                },\n                {\n                    \"id\": 1000053,\n                    \"thirdCode\": \"3\",\n                    \"bayonetType\": 0,\n                    \"acType\": 0,\n                    \"pileNo\": 1000\n                },\n                {\n                    \"id\": 1000086,\n                    \"thirdCode\": \"zhi2\",\n                    \"bayonetType\": 1,\n                    \"acType\": 2,\n                    \"pileNo\": 1300\n                },\n                {\n                    \"id\": 1000054,\n                    \"thirdCode\": \"4\",\n                    \"bayonetType\": 0,\n                    \"acType\": 2,\n                    \"pileNo\": 1500\n                },\n                {\n                    \"id\": 1000055,\n                    \"thirdCode\": \"5\",\n                    \"bayonetType\": 0,\n                    \"acType\": 1,\n                    \"pileNo\": 2000\n                },\n                {\n                    \"id\": 1000056,\n                    \"thirdCode\": \"6\",\n                    \"bayonetType\": 0,\n                    \"acType\": 0,\n                    \"pileNo\": 5000\n                },\n                {\n                    \"id\": 1000057,\n                    \"thirdCode\": \"7\",\n                    \"bayonetType\": 0,\n                    \"acType\": 2,\n                    \"pileNo\": 7000\n                }\n            ],\n            \"edges\": [\n                {\n                    \"from\": 1000051,\n                    \"to\": 1000052,\n                    \"distance\": 300\n                },\n                {\n                    \"from\": 1000052,\n                    \"to\": 1000053,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 1000053,\n                    \"to\": 1000054,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 1000054,\n                    \"to\": 1000055,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 1000055,\n                    \"to\": 1000056,\n                    \"distance\": 3000\n                },\n                {\n                    \"from\": 1000056,\n                    \"to\": 1000057,\n                    \"distance\": 2000\n                },\n                {\n                    \"from\": 1000085,\n                    \"to\": 1000053,\n                    \"distance\": 300\n                },\n                {\n                    \"from\": 1000054,\n                    \"to\": 1000086,\n                    \"distance\": 200\n                }\n            ]\n        },\n        \"reverse\": {\n            \"nodes\": [\n                {\n                    \"id\": 1000058,\n                    \"thirdCode\": \"11\",\n                    \"bayonetType\": 0,\n                    \"acType\": 1,\n                    \"pileNo\": 5000\n                },\n                {\n                    \"id\": 1000059,\n                    \"thirdCode\": \"22\",\n                    \"bayonetType\": 0,\n                    \"acType\": 0,\n                    \"pileNo\": 4500\n                },\n                {\n                    \"id\": 1000087,\n                    \"thirdCode\": \"zhi3\",\n                    \"bayonetType\": 1,\n                    \"acType\": 2,\n                    \"pileNo\": 4300\n                },\n                {\n                    \"id\": 1000060,\n                    \"thirdCode\": \"33\",\n                    \"bayonetType\": 0,\n                    \"acType\": 0,\n                    \"pileNo\": 4000\n                },\n                {\n                    \"id\": 1000061,\n                    \"thirdCode\": \"44\",\n                    \"bayonetType\": 0,\n                    \"acType\": 0,\n                    \"pileNo\": 3500\n                },\n                {\n                    \"id\": 1000062,\n                    \"thirdCode\": \"55\",\n                    \"bayonetType\": 0,\n                    \"acType\": 2,\n                    \"pileNo\": 3000\n                },\n                {\n                    \"id\": 1000088,\n                    \"thirdCode\": \"zhi4\",\n                    \"bayonetType\": 1,\n                    \"acType\": 1,\n                    \"pileNo\": 2600\n                }\n            ],\n            \"edges\": [\n                {\n                    \"from\": 1000058,\n                    \"to\": 1000059,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 1000059,\n                    \"to\": 1000060,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 1000060,\n                    \"to\": 1000061,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 1000061,\n                    \"to\": 1000062,\n                    \"distance\": 500\n                },\n                {\n                    \"from\": 1000059,\n                    \"to\": 1000087,\n                    \"distance\": 200\n                },\n                {\n                    \"from\": 1000088,\n                    \"to\": 1000060,\n                    \"distance\": 200\n                }\n            ]\n        }\n    }\n}")
}
func main() {
	http.HandleFunc("/v1/facilities-manage-service/access_control/digraph", sayHi)
	log.Fatal(http.ListenAndServe("localhost:18083", nil))
}
