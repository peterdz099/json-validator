import React, { useState } from "react";
import { Form, JsonValidationCallback } from '../form/Form';
import './Content.css';

import logo from "../../assets/images/logo.png"


export const Content: React.FC = () => {
    const [show, setShow] = useState<boolean>(false);
    const [render, setRender] = useState<boolean>(false);
    const [message, setMessage] = useState<string>();
    const [promptType, setPromptType] = useState<string>();

    function showPrompt(): void {
        setRender(true);
        setTimeout(() => {
            setShow(true);
        }, 10);

        setTimeout(() => {
            setShow(false);
            setTimeout(() => {
                setRender(false);
            }, 500);
            
        }, 3000); 
    }

    const handleChildValue: JsonValidationCallback = (isJsonValid, message, info = false) => {
        setMessage(message);
        isJsonValid ? setPromptType("success") : (info ? setPromptType("primary"): setPromptType("danger"));
        showPrompt();
    };

    return (
        <div className="mainSquare">
            <div className="logoWrapper">
                <img src={logo} alt="Logo" />
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
