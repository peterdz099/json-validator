import { useRef, useState } from "react";
import { HttpHandler } from "../handlers/HttpHandler";

function Form() {
    
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
          http.sendFile(file);
          if (formRef.current) {
            formRef.current.reset();
          }
          setFile(null);
          
        } else {
          console.log("No file selected.");
        }
      }

      return (
        <form ref={formRef} onSubmit={handleSubmit} className="fileForm">
          <input 
            id="jsonFile" 
            type="file" 
            className="form-control form-control-lg" 
            onChange={handleFileChange} 
            accept=".json" 
            placeholder="Choose a JSON file"
          />
          <button type="submit" className="submitButton">Submit</button>
        </form>
      );
}

export default Form;