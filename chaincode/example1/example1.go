package main
 
import "fmt"
import "github.com/hyperledger/fabric/core/chaincode/shim"
import "github.com/hyperledger/fabric/protos/peer"
 
type SampleChaincode struct {
}

type PersonalInfo struct {
	Firstname	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
	DOB			string	`json:"DOB"`
	Email		string	`json:"email"`
	Mobile		string	`json:"mobile"`
}

type FinancialInfo struct {
    MonthlySalary      int `json:"monthlySalary"`
    MonthlyRent        int `json:"monthlyRent"`
    OtherExpenditure   int `json:"otherExpenditure"`
    MonthlyLoanPayment int `json:"monthlyLoanPayment"`
}

type LoanApplication struct {
    ID                     string        `json:"id"`
    PropertyId             string        `json:"propertyId"`
    LandId                 string        `json:"landId"`
    PermitId               string        `json:"permitId"`
    BuyerId                string        `json:"buyerId"`
    SalesContractId        string        `json:"salesContractId"`
    PersonalInfo           PersonalInfo  `json:"personalInfo"`
    FinancialInfo          FinancialInfo `json:"financialInfo"`
    Status                 string        `json:"status"`
    RequestedAmount        int           `json:"requestedAmount"`
    FairMarketValue        int           `json:"fairMarketValue"`
    ApprovedAmount         int           `json:"approvedAmount"`
    ReviewerId             string        `json:"reviewerId"`
    LastModifiedDate       string        `json:"lastModifiedDate"`
}

func (t *SampleChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
    return shim.Success(nil)
}
 
func (t *SampleChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
    function, args := stub.GetFunctionAndParameters()

    fmt.Println("Invoke is processing" + function )

    // validate function name
    if function == "CreateLoanApplication" {	// create loan application
    	return	t.CreateLoanApplication(stub, args)
    }

    return shim.Success(nil)
}

// create loan application
func (t *SampleChaincode) CreateLoanApplication(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	// check application data
	if len(args) <= 0 {
		return shim.Error("Incorrect number of arguments")
	}



	return shim.Success(nil)
} 

func main() {
    err := shim.Start(new(SampleChaincode))
    if err != nil {
        fmt.Println("Could not start SampleChaincode")
    } else {
        fmt.Println("SampleChaincode successfully started")
    }
 
}