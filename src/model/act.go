/******************************************************************************
    Informations about acts (birth, marriage, death certificates).

    @license    GPL
    @history    2022-02-16 01:51:37+01:00, Thierry Graff : Creation
********************************************************************************/
package model

type Act interface {}

type BC struct { // implements Act
    // Header   interface{}
    Document interface{}
    Person   Person
}

/** 
    Returns the birth
**/
func ComputeBC(p *Person, dirActs string) (interface{}) {
/* 
    file := dirActs
	//y, err := ioutil.ReadFile(dirActs)
	if err != nil {
	    if os.IsNotExist(err){
            return nil; // p.Acts remains empty
	    }
		panic(err)
	}
//	err = yaml.Unmarshal(y, &config)
	if err != nil {
		panic(err)
	}
*/
    return nil
}
