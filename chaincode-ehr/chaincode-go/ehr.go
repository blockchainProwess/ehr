package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	
	"github.com/hyperledger/fabric-chaincode-go/shim"
    sc "github.com/hyperledger/fabric-protos-go/peer"
)

// SmartContract struct
type SmartContract struct {
}


// PatientDetails struct
type PatientDetails struct {
	Name 		 		string `json:"name"`
	DOB 		 		string `json:"date_of_birth"`
	Gender 				string `json:"gender"`
	AdhaarID	 		string `json:"adhaar_id"`
	DoorNo				string `json:"door_number"`
	StreetName  		string `json:"street_number"`
	City        		string `json:"city"`
	District    		string `json:"district"`
	State       		string `json:"state"`	
	JobType				string `json:"job_type"`
	Occupation			string `json:"occupation"`
	Password     		string `json:"password"`
	MobileNo			string `json:"mobile_no"`
	JoinDate			string `json:"join_date"`
}


// TechnicianDetails struct 
type TechnicianDetails struct {
	Name 		 		string `json:"name"`
	DOB 		 		string `json:"date_of_birth"`
	Gender 				string `json:"gender"`
	License 			string `json:"license"`
	AdhaarID	 		string `json:"adhaar_id"`
	DoorNo				string `json:"door_number"`
	StreetName  		string `json:"street_number"`
	City        		string `json:"city"`
	District    		string `json:"district"`
	State       		string `json:"state"`	
	UGDegree 			string `json:"ug_degree"`
	UGUniversity 		string `json:"ug_univesity"`
	PGDegree 			string `json:"pg_degree"`
	PGUniversity 		string `json:"pg_univesity"`
	CenterName 		 	string `json:"center_name"`
	HDoorNo				string `json:"door_number"`
	HStreetName  		string `json:"street_number"`
	HCity        		string `json:"city"`
	HDistrict    		string `json:"district"`
	HState       		string `json:"state"`	
	Password     		string `json:"password"`
	JoinDate			string `json:"join_date"`
}

// DoctorDetails struct 
type DoctorDetails struct {
	Name 		 		string `json:"name"`
	DOB 		 		string `json:"date_of_birth"`
	Gender 				string `json:"gender"`
	License 			string `json:"license"`
	AdhaarID	 		string `json:"adhaar_id"`
	DoorNo				string `json:"door_number"`
	StreetName  		string `json:"street_number"`
	City        		string `json:"city"`
	District    		string `json:"district"`
	State       		string `json:"state"`	
	UGDegree 			string `json:"ug_degree"`
	UGUniversity 		string `json:"ug_univesity"`
	PGDegree 			string `json:"pg_degree"`
	PGUniversity 		string `json:"pg_univesity"`
	CenterName 		 	string `json:"center_name"`
	HDoorNo				string `json:"door_number"`
	HStreetName  		string `json:"street_number"`
	HCity        		string `json:"city"`
	HDistrict    		string `json:"district"`
	HState       		string `json:"state"`	
	Password     		string `json:"password"`
	JoinDate			string `json:"join_date"`
}


// DoctorVisitaion struct 
type DoctorVisitaion struct {
	InstanceID    string `json:"instance_id"`     // unix timestamp
	License  	  string `json:"license"`
	Problem       string `json:"problem` 
	Observations  string `json:"observations"`    // should be comma seperated
	Tests         string `json:"tests"`
	BillDetails   string `json:"bill_details"`    // comma seperated Total amount, generated on, bill breakdown 
	Timestamp     string `json:"timestamp"`
}


// TechnicianVisitation struct
type TechnicianVisitation struct {
	InstanceID          string `json:"instance_id"`
	License				string `json:"license"`
	Tests 				string `json:"tests"`
	DiagnosticReports	string `json:"diagnostic_reports"`
	BillDetails			string `json:"bill_details"`
	Timestamp     		string `json:"timestamp"`
}

// DoctorSummaryVisitaion struct 
type DoctorSummaryVisitaion struct {
	InstanceID    string `json:"instance_id"`     // unix timestamp
	License  	  string `json:"license"`
	DoctorTestObservations  string `json:"doctor_test_observations"`
	Prescription  string `json:"prescription"`
	BillDetails   string `json:"bill_details"`    // comma seperated Total amount, generated on, bill breakdown 
	Timestamp     string `json:"timestamp"`
}



// Init function
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

// Invoke function
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "initLedger" {
		return s.initLedger(APIstub)
	} else if function == "recordDoctorDetails" {
		return s.recordDoctorDetails(APIstub, args)
	} else if function == "recordPatientDetails" {
		return s.recordPatientDetails(APIstub, args)
	} else if function == "recordTechnicianDetails" {
		return s.recordTechnicianDetails(APIstub, args)
	} else if function == "recordDoctorVisitationTransaction" {
		return s.recordDoctorVisitationTransaction(APIstub, args)
	} else if function == "recordTechnicianVisitationTransaction" {
		return s.recordTechnicianVisitationTransaction(APIstub, args)
	} else if function == "recordDoctorSummaryVisitationTransaction" {
		return s.recordDoctorSummaryVisitationTransaction(APIstub, args)
	} 	
	 
	return shim.Error("Invalid Smart Contract function name.")
}

/*
 * The initLedger method *
Will add test data (10 asset catches)to our network
*/
func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {	
	return shim.Success(nil)
}


/*
 * The recordPatientDetails method *
 */
 func (s *SmartContract) recordPatientDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 14 {
		return shim.Error("Incorrect number of arguments. Expecting 14")
	}
 


	var patientdetailstx = PatientDetails{
											Name:args[0],
											DOB:args[1],
											Gender:args[2],
											AdhaarID:args[3],
											DoorNo:args[4],
											StreetName:args[5],
											City:args[6],
											District:args[7],
											State:args[8],
											JobType:args[9],
											Occupation:args[10],
											Password:args[11],
											MobileNo:args[12],
											JoinDate:args[13],
							
	}

	patientdetailstxAsBytes, _ := json.Marshal(patientdetailstx)
	
	err := APIstub.PutState(args[3], patientdetailstxAsBytes)
	
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record patient details: %s", args[3]))
	}

	return shim.Success(nil)

}


/*
 * The recordDoctorDetails method *
 */
 func (s *SmartContract) recordDoctorDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 22 {
		return shim.Error("Incorrect number of arguments. Expecting 22")
	}
	
	var doctordetailstx = DoctorDetails{
											Name 		 :args[0],		
											DOB 		 :args[1],		
											Gender 		 :args[2],		
											License 	 :args[3],		
											AdhaarID	 :args[4],		
											DoorNo		 :args[5],		
											StreetName   :args[6],		
											City         :args[7],		
											District     :args[8],		
											State        :args[9],		
											UGDegree 	 :args[10],		
											UGUniversity :args[11],		
											PGDegree 	 :args[12],		
											PGUniversity :args[13],		
											CenterName 	 :args[14],	 	
											HDoorNo		 :args[15],		
											HStreetName  :args[16],		
											HCity        :args[17],		
											HDistrict    :args[18],		
											HState       :args[19],				
											Password     :args[20],
											JoinDate	 :args[21],							
									
										}

	doctordetailstxAsBytes, _ := json.Marshal(doctordetailstx)
	
	err := APIstub.PutState(args[3], doctordetailstxAsBytes)
	
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record doctor details: %s", args[3]))
	}

	return shim.Success(nil)
}




/*
 * The recordTechnicianDetails method *
 */
 func (s *SmartContract) recordTechnicianDetails(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 22 {
		return shim.Error("Incorrect number of arguments. Expecting 22")
	}
	
	var techniciandetailstx = TechnicianDetails{
												Name 		 :args[0],		
												DOB 		 :args[1],		
												Gender 		 :args[2],		
												License 	 :args[3],		
												AdhaarID	 :args[4],		
												DoorNo		 :args[5],		
												StreetName   :args[6],		
												City         :args[7],		
												District     :args[8],		
												State        :args[9],		
												UGDegree 	 :args[10],		
												UGUniversity :args[11],		
												PGDegree 	 :args[12],		
												PGUniversity :args[13],		
												CenterName 	 :args[14],	 	
												HDoorNo		 :args[15],		
												HStreetName  :args[16],		
												HCity        :args[17],		
												HDistrict    :args[18],		
												HState       :args[19],				
												Password     :args[20],
												JoinDate	 :args[21],							
												
										}

	techniciandetailstxAsBytes, _ := json.Marshal(techniciandetailstx)
	
	err := APIstub.PutState(args[3], techniciandetailstxAsBytes)
	
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record technician details: %s", args[3]))
	}

	return shim.Success(nil)
}



/*
 * The recordDoctorVisitationTransaction method *
 */
 func (s *SmartContract) recordDoctorVisitationTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 8 {
		return shim.Error("Incorrect number of arguments. Expecting 8")
	}
	
	var dctrvsttx = DoctorVisitaion{
											InstanceID    :args[1],
											License  	  :args[2],
											Problem       :args[3], 
											Observations  :args[4],
											Tests         :args[5],
											BillDetails   :args[6],
											Timestamp	  :args[7], 
												
										}

	dctrvsttxAsBytes, _ := json.Marshal(dctrvsttx)
	// key = "DIV_"+args[0]
	err := APIstub.PutState(args[0], dctrvsttxAsBytes)
	
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record technician details: %s", args[0]))
	}

	return shim.Success(nil)
}


/*
 * The recordTechnicianVisitationTransaction method *
 */
 func (s *SmartContract) recordTechnicianVisitationTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	
	var tchnsnvsttx = TechnicianVisitation{
											InstanceID    		:args[1],	
											License				:args[2], 	 	
											BillDetails 	 	:args[3],
											Tests				:args[4],
											DiagnosticReports 	:args[5],
											Timestamp	  		:args[6],
												
										}

	tchnsnvsttxAsBytes, _ := json.Marshal(tchnsnvsttx)
	// key = "DGV_"+args[0]
	err := APIstub.PutState(args[0], tchnsnvsttxAsBytes)
	
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record technician details: %s", args[0]))
	}

	return shim.Success(nil)
}


/*
 * The recordDoctorSummaryVisitationTransaction method *
 */
 func (s *SmartContract) recordDoctorSummaryVisitationTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 7 {
		return shim.Error("Incorrect number of arguments. Expecting 7")
	}
	
	var dctrsmryvsttx = DoctorSummaryVisitaion{
											InstanceID    :args[1],
											License  	  :args[2],
											DoctorTestObservations :args[3],
											Prescription  :args[4],
											BillDetails   :args[5],
											Timestamp	  :args[6], 
											}

	dctrsmryvsttxAsBytes, _ := json.Marshal(dctrsmryvsttx)
	// key = "DSV_"+args[0]
	err := APIstub.PutState(args[0], dctrsmryvsttxAsBytes)
	
	if err != nil {
		return shim.Error(fmt.Sprintf("Failed to record technician details: %s", args[0]))
	}

	return shim.Success(nil)
}




/*
 * The query method *
 */
func (s *SmartContract) queryTransaction(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	assetAsBytes, _ := APIstub.GetState(args[0])
	if assetAsBytes == nil {
		return shim.Error("Could not locate asset")
	}
	return shim.Success(assetAsBytes)
}



/*
 * The queryTransactionHistory method *
 */
 func (s *SmartContract) queryTransactionHistory(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	recordID := args[0]

	fmt.Printf("- start getAsset History: %s\n", recordID)

	resultsIterator, err := APIstub.GetHistoryForKey(recordID)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing historic values for the asset
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(response.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Value\":")
		// if it was a delete operation on given key, then we need to set the
		//corresponding value null. Else, we will write the response.Value
		//as-is (as the Value itself a JSON asset)
		if response.IsDelete {
			buffer.WriteString("null")
		} else {
			buffer.WriteString(string(response.Value))
		}

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(response.Timestamp.Seconds, int64(response.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"IsDelete\":")
		buffer.WriteString("\"")
		buffer.WriteString(strconv.FormatBool(response.IsDelete))
		buffer.WriteString("\"")

		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- get Asset History returning:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())

}





func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

