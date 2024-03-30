import axios from "axios";
import { API_URL } from "../config/urls";

import messages from "../assets/texts/messages.json"

interface ValidationRespone {
    valid: boolean;
    message: string;
}

export const validateJson = async (file: File) : Promise<ValidationRespone> => {
    try {
        const formData = new FormData();
        formData.append('file', file);

        const { data, status } = await axios.post(
            API_URL,
            formData,
            {
                headers: {
                    'Content-Type': 'multipart/form-data',
                },
            },
        );

        return data;

    } catch (error) {
        if (axios.isAxiosError(error)) {
            if (error.request.responseText != ""){
                return {valid: false, message: error.request.responseText}
            }
            return {valid: false, message: error.message}
        } else {
            return {valid: false, message: messages.unexpectedError }
        }
    }
};