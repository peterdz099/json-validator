import React from "react";
import Form from './Form'

function Content() {

    return(
        <div className="MainSquare">
            <img src={require('../assets/logo.png')} alt="Logo" />
            <div className="FromAndResult">
                <Form />
            </div>
        </div>
    )
}
export default Content;