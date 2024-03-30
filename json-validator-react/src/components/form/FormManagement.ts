import React from 'react';

export function resetForm(formRef: React.RefObject<HTMLFormElement>): void{
    if (formRef.current) {
        formRef.current.reset();
    }
}

export function disableButtonForLimitedTime(buttonRef: React.RefObject<HTMLButtonElement>): void{
    if(buttonRef.current){
        buttonRef.current.disabled = true;
        setTimeout(() => {
          if(buttonRef.current){
            buttonRef.current.disabled = false;
          }
        }, 4000); 
    }   
}

export{}