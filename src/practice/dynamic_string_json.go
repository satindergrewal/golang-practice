package main

import (
	"encoding/json"
	"fmt"
)

type meminfo struct {
	string
}

type memfalse struct {
	Result map[string]memtrue    `json:"result"`
	Error  interface{} `json:"error"`
	ID     string      `json:"id"`
}

type memtrue struct {
	Size             int           `json:"size"`
	Fee              float64       `json:"fee"`
	Time             int           `json:"time"`
	Height           int           `json:"height"`
	Startingpriority float64       `json:"startingpriority"`
	Currentpriority  float64       `json:"currentpriority"`
	Depends          []interface{} `json:"depends"`
}

func (s meminfo) greeting() {
	fmt.Println("Hello", s.string)
}

func main() {
	a := meminfo{"World"}
	a.greeting()



  mfJson := `{"result":[
  "7921fbd3ad4d1593aa580833946cb123260f81480069a37f649d2c8cc10d770b",
  "7a1cae3b6e7dcaf94e2422b75aa4eb13fd54de1094e23a2fc46346a4d7d40448",
  "adc6251d4efbc3409e4c615f1ef298d7c8805519efe15b7f009c910e34cc2f56",
  "35306da9b38b89df74386723607d2f7fe8364cf4e6fa6b0c656d18c4e236f962",
  "1aff2a2ef54283d4622084f69424f64f1afbd7c2de9b0c67e50ac2a2b6e54175",
  "385e697b5884179ac85e7bdcb2909834edfeb17774b51c749cdd1327e064009a",
  "f9b4e1667bf01e9fc1507f35b77a90ef6b70508f6c5ce7a814238a2e0c893dc0",
  "e837e17ef3141aae95f7963030356976706951ce547f50eb9557ef309a2c79c1",
  "19056d2181304e68188d0dbe43cc1ee545e81f147ff889600827b016b4f20cd2",
  "25ebbd51408d7d5fb8aede3a79492be14f1c5b1823315d12a8b564b920e7a0e4",
  "0c4162f75c28682bbe5cbc518a116f689612a24871d0f30450c3fcac578825f0"
],"error":null,"id":"curltest"}`

	mtJson := `{"result":{
  "7921fbd3ad4d1593aa580833946cb123260f81480069a37f649d2c8cc10d770b": {
    "size": 4611,
    "fee": 0.00110000,
    "time": 1547431224,
    "height": 1183180,
    "startingpriority": 8135355.08111136,
    "currentpriority": 8135355.08111136,
    "depends": [
    ]
  },
  "7a1cae3b6e7dcaf94e2422b75aa4eb13fd54de1094e23a2fc46346a4d7d40448": {
    "size": 244,
    "fee": 0.00005000,
    "time": 1547431225,
    "height": 1183180,
    "startingpriority": 0,
    "currentpriority": 0,
    "depends": [
      "f9b4e1667bf01e9fc1507f35b77a90ef6b70508f6c5ce7a814238a2e0c893dc0"
    ]
  },
  "adc6251d4efbc3409e4c615f1ef298d7c8805519efe15b7f009c910e34cc2f56": {
    "size": 1398,
    "fee": 0.00000000,
    "time": 1547431224,
    "height": 1183179,
    "startingpriority": 12206881.05515588,
    "currentpriority": 12410329.07274181,
    "depends": [
    ]
  },
  "35306da9b38b89df74386723607d2f7fe8364cf4e6fa6b0c656d18c4e236f962": {
    "size": 1633,
    "fee": 0.00031200,
    "time": 1547431169,
    "height": 1183179,
    "startingpriority": 144161.4906832298,
    "currentpriority": 144968.9440993789,
    "depends": [
    ]
  },
  "1aff2a2ef54283d4622084f69424f64f1afbd7c2de9b0c67e50ac2a2b6e54175": {
    "size": 2411,
    "fee": 0.00060000,
    "time": 1547431224,
    "height": 1183179,
    "startingpriority": 1201919575.784357,
    "currentpriority": 1213938820.69863,
    "depends": [
    ]
  },
  "385e697b5884179ac85e7bdcb2909834edfeb17774b51c749cdd1327e064009a": {
    "size": 245,
    "fee": 0.00005000,
    "time": 1547431225,
    "height": 1183180,
    "startingpriority": 0,
    "currentpriority": 0,
    "depends": [
      "7a1cae3b6e7dcaf94e2422b75aa4eb13fd54de1094e23a2fc46346a4d7d40448"
    ]
  },
  "f9b4e1667bf01e9fc1507f35b77a90ef6b70508f6c5ce7a814238a2e0c893dc0": {
    "size": 244,
    "fee": 0.00005000,
    "time": 1547431224,
    "height": 1183180,
    "startingpriority": 0,
    "currentpriority": 0,
    "depends": [
      "19056d2181304e68188d0dbe43cc1ee545e81f147ff889600827b016b4f20cd2"
    ]
  },
  "e837e17ef3141aae95f7963030356976706951ce547f50eb9557ef309a2c79c1": {
    "size": 3599,
    "fee": 0.00087000,
    "time": 1547431224,
    "height": 1183179,
    "startingpriority": 9724241.379310345,
    "currentpriority": 9805957.693422196,
    "depends": [
    ]
  },
  "19056d2181304e68188d0dbe43cc1ee545e81f147ff889600827b016b4f20cd2": {
    "size": 244,
    "fee": 0.00005000,
    "time": 1547431224,
    "height": 1183180,
    "startingpriority": 3236483169.896907,
    "currentpriority": 3236483169.896907,
    "depends": [
    ]
  },
  "25ebbd51408d7d5fb8aede3a79492be14f1c5b1823315d12a8b564b920e7a0e4": {
    "size": 244,
    "fee": 0.00062045,
    "time": 1547431169,
    "height": 1183179,
    "startingpriority": 4663894206185.567,
    "currentpriority": 4664307019462.062,
    "depends": [
    ]
  },
  "0c4162f75c28682bbe5cbc518a116f689612a24871d0f30450c3fcac578825f0": {
    "size": 2014,
    "fee": 0.00051000,
    "time": 1547431224,
    "height": 1183179,
    "startingpriority": 19170552.08677022,
    "currentpriority": 19616378.8794858,
    "depends": [
    ]
  }
},"error":null,"id":"curltest"}`

  fmt.Println(mfJson)
  //fmt.Println(mfJson)

	//result := make(map[string]memtrue)
	var result memfalse
	json.Unmarshal([]byte(mtJson), &result)

	fmt.Println(result.Result)
}
