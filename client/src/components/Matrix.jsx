
import {
	BrowserRouter as Router,
	Switch,
	Route,
	Link
} from "react-router-dom";

import { ApolloProvider } from '@apollo/react-hooks';

import { useQuery, useMutation } from '@apollo/react-hooks';
import { gql } from 'apollo-boost';
import ApolloClient from 'apollo-boost'
import { useEffect, useState } from "react";


import { fetchResult } from './../service/Service'
function Matrix() {

	const [param, setParam] = useState(1)

	const [matrixData, setMatrixData] = useState(null)
	const [specialNumber, setSpecialNumber] = useState(null)


	const graphql = JSON.stringify({
		query: `query generate {\r\n  generateMatrix(num: ${param}) {\r\n    special\r\n    matrix\r\n  }\r\n}\r\n`,
		variables: {}
	})
	useEffect(() => {

		fetchResult(graphql).then(response => response.text())
			.then((result) => {
				result = JSON.parse(result);
				console.log(result)
				setMatrixData(result.data.generateMatrix.matrix);
				setSpecialNumber(result.data.generateMatrix.special);

			})
			.catch(error => console.log('error', error));

	}, [])
	const generate = () => {
		fetchResult(graphql).then(response => response.text())
			.then((result) => {
				result = JSON.parse(result);

				const newArr = [];
				while (result.data.generateMatrix.matrix.length) newArr.push(result.data.generateMatrix.matrix.splice(0, param));
				console.log(newArr)
				setMatrixData(newArr);
				setSpecialNumber(result.data.generateMatrix.special);

			})
			.catch(error => console.log('error', error));
	}
	const changeInput = (e) => {
		setParam(e.target.value);
	}

	const renderMatrix = () => {
		if (matrixData !== null) {
			return matrixData.map((pair) => {

				return (
					<tr>
						{pair}
					</tr>
				)
			})
		}
	}



	return (
		<div>
			<h2>Set the number here</h2>
			<input type="text" name="" placeholder="Enter number" onChange={changeInput} /><button onClick={generate}>Generate</button>
			<h3>Special Number :{specialNumber}</h3>
			<h3>Spiral Matrix: {renderMatrix()}</h3>

		</div>
	);
}

export default Matrix;
