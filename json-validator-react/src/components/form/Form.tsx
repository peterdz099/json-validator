import './Form.css';
import { useRef, useState } from "react";
import { HttpHandler } from "../../handlers/HttpHandler";

interface FormProps {
  onJsonValidation: (jsJsonValid: boolean, message: string, info ?: boolean ) => void;
}

const Form: React.FC<FormProps> = (props) => {
    
    const http: HttpHandler = new HttpHandler('http://127.0.0.1:8080/validate');

    const [file, setFile] = useState<File | null>(null);
    const formRef = useRef<HTMLFormElement>(null);
    const buttonRef = useRef<HTMLButtonElement>(null);

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
        if (file && !buttonRef.current?.disabled) {
          let message: string = "Internal Error"
          let isJsonValid: boolean = false;
          
          try{
            http.validateJson(file).then(data => {
                isJsonValid = data.valid;
                message = data.message; 
                props.onJsonValidation(isJsonValid, message);
                if (formRef.current) {
                  formRef.current.reset();
                }
                if(buttonRef.current){
                  buttonRef.current.disabled = true;
                  setTimeout(() => {
                    if(buttonRef.current){
                      buttonRef.current.disabled = false;
                    }
                  }, 4000); 
                }    
                setFile(null);
              }
            )
          }
          catch(error){
            props.onJsonValidation(isJsonValid, message);
          }

        } else {
          if(!buttonRef.current?.disabled){
            props.onJsonValidation(false, "Select a file", true);
            if(buttonRef.current){
                buttonRef.current.disabled = true;
                setTimeout(() => {
                  if(buttonRef.current){
                    buttonRef.current.disabled = false;
                  }
                }, 4000); 
            }
          }
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
          <button ref={buttonRef}type="submit" className="submitButton">Submit</button>
        </form>
      );
}

export default Form;