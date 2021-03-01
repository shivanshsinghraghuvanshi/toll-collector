
import { useState } from "react";
import {
	BrowserRouter as Router,
	Switch,
	Route,
	Link
} from "react-router-dom";

import { fetchResult } from './../service/Service'



function TollBooth() {


	const [refID, setRfid] = useState(null)
	const [ownerId, setOwnerId] = useState(0)
	const [carnumber, setCarNumber] = useState("")
	const [ncarnumber, setnCarNumber] = useState("")
	const [nrefID, nsetRfid] = useState("")
	const [valid, isValid] = useState(null)




	const [status, setStatus] = useState(null)
	const [tollid, setTollID] = useState(null)
	const [amount, setAmount] = useState(null)
	const [remarks, setRemarks] = useState(null)
	const [payRfid, setPayrfid] = useState(null)
	const graphql = JSON.stringify({
		query: "mutation newrfid($e:NewRFID){\r\n  createRFID(input:$e)\r\n}\r\n",
		variables: { "e": { "ownerid": ownerId, "carnumber": carnumber } }
	})
	const payTolltax = JSON.stringify({
		query: "mutation mypayment($e:PayTollTax){\r\n  payTollTax(input:$e)\r\n}",
		variables: { "e": { "rfid": payRfid, "tollid": tollid, "amount": amount, "remarks": remarks } }
	})

	const graphqlverify = JSON.stringify({
		query: "mutation validate($e:ValidateRFID!){\r\n  validateRFID(input:$e)\r\n}",
		variables: { "e": { "rfid": nrefID, "carnumber": ncarnumber } }
	})
	const generateRFID = () => {
		fetchResult(graphql).then(response => response.text())
			.then((result) => {
				result = JSON.parse(result);
				console.log(result)
				setRfid(result.data.createRFID);

			})
			.catch(error => console.log('error', error));
	}

	const validateRFID = () => {
		fetchResult(graphqlverify).then(response => response.text())
			.then((result) => {
				result = JSON.parse(result);
				console.log(result)
				if (result.data.validateRFID) {
					isValid("true")
				} else {
					isValid("false")
				}

			})
			.catch(error => console.log('error', error));
	}

	const changeOwner = (e) => {
		setOwnerId(e.target.value);
	}
	const changeCarNumber = (e) => {
		setCarNumber(e.target.value);
	}
	const changecarnumverify = (e) => {
		setnCarNumber(e.target.value);
	}
	const changerfid = (e) => {
		nsetRfid(e.target.value);
	}

	const ctollid = (e) => {
		setTollID(e.target.value)

	}
	const cprfid = (e) => {
		setPayrfid(e.target.value)
	}
	const camount = (e) => {
		setAmount(e.target.value)
	}
	const cremarks = (e) => {
		setRemarks(e.target.value)
	}

	const paytolltax = () => {
		fetchResult(payTolltax).then(response => response.text())
			.then((result) => {
				result = JSON.parse(result);
				console.log(result)
				if (result.data.payTollTax) {
					setStatus("true")
				} else {
					setStatus("false")
				}

			})
			.catch(error => console.log('error', error));
	}
	return (
		<div>

			<h2>Generate RFID</h2>
			<input type="text" name="" placeholder="Enter OwnerID" onChange={changeOwner} />
			<input type="text" name="" placeholder="Enter Car Number" onChange={changeCarNumber} />

			<button onClick={generateRFID}>Generate RFID</button>

			<h3>The generated RFID is : {refID}</h3>

			<br />
			<br />
			<div>
				<h2>Validate RFID</h2>
				<input type="text" name="" placeholder="Enter Car Number" onChange={changecarnumverify} />
				<input type="text" name="" placeholder="Enter RFID" onChange={changerfid} />

				<button onClick={validateRFID}>Validate RFID</button>
				<h3>This rfid validity is  : {valid}</h3>
			</div>



			<br />
			<br />
			<div>
				<h2>Pay Toll Tax</h2>
				<input type="text" name="" placeholder="Enter toll id" onChange={ctollid} />
				<input type="text" name="" placeholder="Enter RFID" onChange={cprfid} />
				<input type="text" name="" placeholder="Enter amount" onChange={camount} />
				<input type="text" name="" placeholder="Enter remarks" onChange={cremarks} />
				<button onClick={paytolltax}>Pay Toll Tax</button>
				<h3>trasaction status is  : {status}</h3>
			</div>
		</div>

	);
}

export default TollBooth;
