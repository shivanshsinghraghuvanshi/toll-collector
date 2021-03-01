const rootUrl = `http://localhost:8000/graphql`;




// var graphql = JSON.stringify({
//   query: "query generate {\r\n  generateMatrix(num: 51) {\r\n    special\r\n    matrix\r\n  }\r\n}\r\n",
//   variables: {}
// })
// .then(response => response.text())
//   .then(result => console.log(result))
//   .catch(error => console.log('error', error));

export const fetchResult =  function(graphql){
var myHeaders = new Headers();
myHeaders.append("Accept-Encoding", "gzip, deflate, br");
myHeaders.append("Content-Type", "application/json");
myHeaders.append("Accept", "application/json");
myHeaders.append("Connection", "keep-alive");
myHeaders.append("DNT", "1");
myHeaders.append("Origin", "http://localhost:3000");
var requestOptions = {
  method: 'POST',
  headers: myHeaders,
  body: graphql,
  redirect: 'follow'
};

return fetch(rootUrl, requestOptions);

}
