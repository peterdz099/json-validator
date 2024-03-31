# json-validator

A json validator-project is build for verifying JSON data defined as AWS:IAM:Role
Policy. It contains either frontend application build using React and backend application build in Go.

## Excercise interpretation

The task did not specifically specify the project requirements, which left room for my own interpretation. I decided to create a front-end and a back-end application, which communicate with each other using the HTTP protocol. Using the application accessible via localhost:3000, the user can specify a *.json file, which is then sent via HTTP Request to port :8080, where its content is verified. The verification begins by checking if the JSON format is correct, if it includes all necessary fields, if the fields are of the specified type, and if they are not empty. If it successfully passes this verification, the Resource field is examined. If not, the application sends an HTTP Response with code 400 with a suitable message to the user. The Resource field, on the other hand, is checked by a method that returns true if the JSON is valid (any value of the Resource field does not contain an asterisk) or false if it is not valid (any value of the Resource field contains an asterisk). After the verification is completed, an HTTP response is returned with the code 200 and the appropriate information for the user. I assumed that the correct form of JSON is the form in the task body, in addition, checking the AWS documentation, I decided that its form should look as follows:
```sh
{
	"PolicyName": string, required
	"PolicyDocument": {
	     "Version": string, required
	     "Statment": [
	     	{
	     		"Sid": string, optional,
	     		"Effect": string ("Allow" | "Deny"), required
	     		"Action": string | []string, required
	     		"Resource": string | []string, required
	     	}, ...
	     ]
	}
}
```

## How to run
**Prequisities**:
1. GIT
2. Node.js - latest Version
3. Go, latest Version
&nbsp;


**Running instruction**

Clone git repository.
  ```sh
    git clone https://github.com/peterdz099/json-validator.git
  ```

Change the current working directory.
  ```sh
    cd json-validator
  ```

Run Go application
  ```sh
    cd json-validator-go/cmd/json-validator
  ```
  ```sh
    go run .
  ```
In other terminal run React application
  ```sh
    cd /json-validator/json-validator-react/
  ```
  ```sh
    npm install
  ```
  ```sh
    npm start
  ```
You can now use application on
 ```sh
    https://localhost:3000
  ```
You can use *.json files from **json-validator/json-validator-go/test/test-data** to play with the application.  
&nbsp;

**Running Tests**
Change the current working directory.
  ```sh
    cd json-validator/json-validator-go/tests
  ```
  ```sh
    go test -v
  ```
