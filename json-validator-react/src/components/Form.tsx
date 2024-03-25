import { useRef, useState } from "react";
import { HttpHandler } from "../handlers/HttpHandler";

interface FormProps {
  onJsonValidation: (jsJsonValid: boolean, message: string, ) => void;
}

const Form: React.FC<FormProps> = (props) => {
    
    const http: HttpHandler = new HttpHandler('http://127.0.0.1:8080/validate');

    const [file, setFile] = useState<File | null>(null);
    const formRef = useRef<HTMLFormElement>(null);

    function handleFileChange(e: React.ChangeEvent<HTMLInputElement>) {
        const fileList = e.target.files;
        if (fileList && fileList.length > 0) {
          const selectedFile = fileList[0];
          if (selectedFile.type === "application/json") {
            setFile(selectedFile);
          } else {
            if (formRef.current) {
                formRef.current.reset();
              }
          }
        }
      }

    function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
        e.preventDefault();
        if (file) {
          let message: string = "Internal Error"
          let jsJsonValid: boolean = false;
          
          try{
            http.validateJson(file).then(data => {
                jsJsonValid = data.valid;
                message = data.filename; //change immiediately
                props.onJsonValidation(jsJsonValid, message);
                if (formRef.current) {
                  formRef.current.reset();
                }
                setFile(null);
              }
            )
          }
          catch(error){
            props.onJsonValidation(jsJsonValid, message);
          }

        } else {
          console.log("No file selected.");
        }
      }

      return (
        <form ref={formRef} onSubmit={handleSubmit}>
          <input
            className="form-control"
            id="jsonFile"
            type="file"
            accept=".json"
            placeholder="Choose a JSON file"
            onChange={handleFileChange}
            lang="en-US"
          />
          <button type="submit" className="submitButton">Submit</button>
        </form>
      );
}

export default Form;