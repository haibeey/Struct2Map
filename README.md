## Struct2Map

Struct2Map coverts go structs types to an instance of `map[string]interface{}`



    Usage: go get github.com/haibeey/struct2Map

    val,err:=struct2Map.Struct2Map(struct)
    if err != nil {
        t.Errorf(err.Error())
    }

