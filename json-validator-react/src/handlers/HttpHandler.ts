import axios from "axios";

export class HttpHandler {
    private url: string;

    constructor(url :string) {
        this.url = url;
    }
    
    async validateJson(file: File) {
        try {
            const formData = new FormData();
            formData.append('file', file);

            const { data, status } = await axios.post(
                this.url,
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
                return {valid: false, message: "An unexpected error occurred"}
            }
        }
    }
}    
