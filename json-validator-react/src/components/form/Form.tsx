import './Form.css';
import { useRef, useState } from "react";
import { validateJson } from "../../handlers/HttpHandler";

import messages from "../../assets/texts/messages.json"
import { disableButtonForLimitedTime, resetForm } from './FormManagement';


export interface JsonValidationCallback {
  (jsJsonValid: boolean, message: string, info ?: boolean): void;
}

interface FormProps {
  onJsonValidation: JsonValidationCallback;
}

export const Form: React.FC<FormProps> = (props) => {
    const [file, setFile] = useState<File | null>(null);
    const formRef = useRef<HTMLFormElement>(null);
    const buttonRef = useRef<HTMLButtonElement>(null);

    function handleFileChange(e: React.ChangeEvent<HTMLInputElement>): void {
        const fileList = e.target.files;
        if (fileList && fileList.length > 0) {
          const selectedFile = fileList[0];
          if (selectedFile.type === "application/json") {
            setFile(selectedFile);
          } else {
              resetForm(formRef);
          }
        }else{
          resetForm(formRef);
          setFile(null);
        }
      }

    function handleSubmit(e: React.FormEvent<HTMLFormElement>) : void {
        e.preventDefault();
        if (file && !buttonRef.current?.disabled) {
          try{
            validateJson(file).then(data => {
                const isJsonValid = data.valid;
                const message = data.message; 
                props.onJsonValidation(isJsonValid, message);
                resetForm(formRef);
                disableButtonForLimitedTime(buttonRef); 
              }
            )
          }
          catch(error){
            props.onJsonValidation(false, messages.InternalError);
          }
        } else if (!buttonRef.current?.disabled) {
          props.onJsonValidation(false, messages.selectFile, true);
          disableButtonForLimitedTime(buttonRef);
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
