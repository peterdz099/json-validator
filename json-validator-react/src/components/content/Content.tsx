import './Content.css';
import React, { useState } from "react";
import Form from "../form/Form";


const Content: React.FC = () => {
    const [show, setShow] = useState<boolean>(false);
    const [render, setrender] = useState<boolean>(false);
    const [message, setMessage] = useState<string>();
    const [promptType, setPromptType] = useState<string>();

    function showPrompt() {
        setrender(true);
        setTimeout(() => {
            setShow(true);
        }, 10);

        setTimeout(() => {
            setShow(false);
            setTimeout(() => {
                setrender(false);
            }, 500);
            
        }, 3000); 
    }

    const handleChildValue = (isJsonValid: boolean, message: string, info ?: boolean ) => {
        setMessage(message);
        isJsonValid ? setPromptType("success") : (info ? setPromptType("primary"): setPromptType("danger"));
        showPrompt();
      };

    return (
        <div className="mainSquare">
            <div className="logoWrapper">
                <img src={require('../../assets/logo.png')} alt="Logo" />
            </div>
            <div className="contentWrapper">
                {render && <div className="alert-message" >
                    <div 
                        className={show ? `alert alert-${promptType} fade-enter` : `alert alert-${promptType} fade-exit`} 
                        role="alert"
                        style={{padding: 0}}
                    >
                        {message}
                    </div>
                </div>}
                <Form onJsonValidation={handleChildValue}></Form>
            </div>
        </div>
    );
}

export default Content;
