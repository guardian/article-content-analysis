const namesJsonUrl = 'https://s3-eu-west-1.amazonaws.com/bechdel-test-names/names.json'
const nlp = require('compromise');
const fetch = require("node-fetch");


const getNamesForPath = (name) =>
fetch(namesJsonUrl).then(function(response){
  return response.json()
}).then(function(names) {
  const a = nlp(name, names).people().data();
  return a;
});


exports.handler = async (event) => {
  let body = JSON.parse(event.body);
  let res =  await getNamesForPath(body.name);
  return res;
};
