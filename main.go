package main

import (
	"fmt"
	"net/url"
	"practice/ccAvenueEncryption"
	"strings"
)

func main() {
	//goroutines.BasicsOfGoroutines()
	//goroutines.PrintPattern1(5)
	//goroutines.PrintPattern1usingRecursion(3, 1, 1)
	//goroutines.DiamondPattern(9)
	// code for binary search
	//arr := []int{1, 2, 3, 4, 5}
	//println(goroutines.BinarySearch([]int{1, 2, 3, 4, 5}, 3))
	//println(goroutines.BinarySearchUsingRecursion(arr, 7, 0, len(arr)-1))
	//goroutines.PrintPatternPyramidUsingRecursion(9, 0, 0, true)
	//weeklyChallenge.CreateCode("Ashutosh Pandey")

	/**
	CCAVENUE_TESTING
	**/

	workingKey := "1106DB35C320B2AD8C5EFE664DFD1A07" // Replace with your actual key

	//arrayString := []string{
	//	"265201",
	//	strconv.FormatInt(time.Now().Unix(), 10),
	//	"http://k8s-neuro-ingressd-68355dad55-1868134886.ap-south-1.elb.amazonaws.com/payment/accepted",
	//	"http://k8s-neuro-ingressd-68355dad55-1868134886.ap-south-1.elb.amazonaws.com/payment/rejected",
	//	"1000",
	//}
	//
	//encryptedText := ccAvenueEncryption.Encrypt(arrayString, workingKey)
	//fmt.Println("Encrypted:", encryptedText)
	//decryptedText := ccAvenueEncryption.Decrypt(
	//	"ded923e9778ec628edefdc5cb38a4e06b24ac3932aa1a6c7e997091b78209b2840a2cf8a0f9034afed219afa4270aa67ea4dd1c354f2f0734d98690b557fb5afac3a796c8c00aea78bb4a7e619e200bb29be9808fb74b82584c3a76b71212f27d10cf4ceb329ff6e600bd63d5d2f9a3ccb2a0dfd7ffd92d876438c23edc705e4bb3a8deb7269687d1ef797a983254f00c4f9824eaaa1f552b47180fb74706feb5b82a3e1a2108dbf266b43c1a3bd0863ffb034aae71ac071a5d394f908e0d396ead808235e8d0ef641474699460d2499",
	//	workingKey)
	//fmt.Println("Decrypted:", string(decryptedText))
	//ccAvenueEncryption.DoRequest()
	//cronJob.RunCronJobs()

	res, _ := ccAvenueEncryption.DecryptEncVal(
		"9654a1702038d251ec6952ed5e3929476e592d59f92b134eac5ec306aa38b52265f0cd5b521bdf6260eb5dbc2241c10b5f5bfa3f5e48a5b8da555ecb564b4c7f10d1a244eebdb2fcef1c0360489842516d90e0f69b25792e8f86d411874cc9b686541dd423056092fe15193ed407df27d020b6c810aa91019fe75f22a70d5a8685310c7110e0ab47551ced40720c8f64f6273a75a0d2598b9c82f812d41878549c1167ca3aa0a4f976c578981be5d77814e68faf52a511613a3a7bfb9a8f028444a446e7d278b21fb2fbc1647144e898ab9b85ff4e4b50bd74b2ffa3421066cfabb3432010f87ef1e6d94f9096869874f5cf6445874e886faf1b5c41ced03b18792701574ea83b5ef8c06504f860bda53c5e41adaaa25d63163003e9327f40e21eeb2cd34caf039dfd28ecff30df454a964e1d2a0d025a67ad2d969382f016174d008cf89ec790fc306bacf71a46535a7e5f4b43cd28a5709a58b8a9d6676c424182fd43eac4143c0e4c13a7e86bfe5017cda37659a707c72189e0faffa6d771942a4078beeed2d5194d5e6d59423021319803798e0edc07e2aee3e4832f5520793a5b61aa5c1fb73cc317949e57518077534ac2499d65727289600c5b83a301a46fde70e090f3f550102f13e558d9dc47205a825abf42a366de5882980f930bc2699857526bab53408b208ecc80ced1d522cef191f2e9a3b52769fc81db7cad5a97618b686baa82a77310cd6c82e555842c88b41691c29a309cb8e569d691012d5dee9603e7af9339221df25824d2bb55acaa013e6847639ace476688a0d0695e99f6687270474390764cfa37c44bfbe5bb695065c85d3c43f7654d944afe076c2695d17ece88d1b72dd61c008ebb061f2ce3c9a8678bc2494e88b03d70c05d77cdddc98d7193475a52302cb2231ad067bffa1abdefeb6d81625402c4c1e9a1168a9746556fa2ac8d5f8671344fdfc1084f6f50d8696aaac6523b04b66acc8862fed3fa2b5950a7155aead314c703a9",
		workingKey)

	//"order_id=1692210783&tracking_id=312010178136&bank_ref_no=1692214073910&order_status=Success&failure_message=&payment_mode=Net Banking&card_name=AvenuesTest&status_code=null&status_message=Y&currency=INR&amount=5000000.00&billing_name=&billing_address=&billing_city=&billing_state=&billing_zip=&billing_country=&billing_tel=&billing_email=&delivery_name=&delivery_address=&delivery_city=&delivery_state=&delivery_zip=&delivery_country=&delivery_tel=&merchant_param1=&merchant_param2=&merchant_param3=&merchant_param4=&merchant_param5=&vault=N&offer_type=null&offer_code=null&discount_value=0.0&mer_amount=5000000.00&eci_value=null&retry=N&response_code=0&billing_notes=&trans_date=17%2F08%2F2023%2000%3A57%3A55&bin_country=%08%08%08%08%08%08%08%08%08%08%08%08%08%08"

	// Parse the URL-encoded string
	values, err := url.ParseQuery(strings.ReplaceAll(res, "\x08", "")) // Removing invalid characters
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	println(res)

	// Extract and print each value
	for key, val := range values {
		fmt.Printf("%s: %s\n", key, val[0])
	}

}

//package main
//
//import (
//	"fmt"
//	"net/http"
//	"practice/ccAvenueEncryption"
//	"strconv"
//	"strings"
//	"time"
//)
//
//const htmlContent = `\
//<html>
//<head>
//	<title>Sub-merchant checkout page</title>
//	<script src="http://ajax.googleapis.com/ajax/libs/jquery/1.10.2/jquery.min.js"></script>
//</head>
//<body>
//<form id="nonseamless" method="post" name="redirect" action="https://test.ccavenue.com/transaction/transaction.do?command=initiateTransaction" >
//		<input type="hidden" id="encRequest" name="encRequest" value="$encReq">
//		<input type="hidden" name="access_code" id="access_code" value="$xscode">
//		<script language='javascript'>document.redirect.submit();</script>
//</form>
//</body>
//</html>
//`
//
//func main() {
//	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//
//		arrayString := []string{
//			"265201",
//			strconv.FormatInt(time.Now().Unix(), 10),
//			"http://k8s-neuro-ingressd-68355dad55-1868134886.ap-south-1.elb.amazonaws.com/payment/accepted",
//			"http://k8s-neuro-ingressd-68355dad55-1868134886.ap-south-1.elb.amazonaws.com/payment/rejected",
//			"1000",
//		}
//
//		encReq := ccAvenueEncryption.Encrypt(arrayString, "1106DB35C320B2AD8C5EFE664DFD1A07") // Replace with actual value
//		xscode := "AVUM05KH10AQ49MUQA"                                                        // Replace with actual value
//
//		html := replaceVariables(htmlContent, map[string]string{
//			"$encReq": encReq,
//			"$xscode": xscode,
//		})
//
//		w.Header().Set("Content-Type", "text/html")
//		fmt.Fprint(w, html)
//	})
//
//	http.HandleFunc("/response", func(w http.ResponseWriter, r *http.Request) {
//		r.ParseForm()
//		response := r.Form.Get("response") // Assuming the response parameter contains the transaction response
//
//		// Process the response as needed
//		fmt.Println("Transaction Response:", response)
//	})
//
//	fmt.Println("Server listening on :8087")
//	http.ListenAndServe(":8087", nil)
//}
//
//func replaceVariables(input string, variables map[string]string) string {
//	output := input
//	for key, value := range variables {
//		output = strings.Replace(output, key, value, -1)
//	}
//	return output
//}
