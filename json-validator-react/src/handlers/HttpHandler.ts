import axios from "axios";

export class HttpHandler {
    private url: string;

    constructor(url :string) {
        this.url = url;
    }
    
    async sendFile(file: File) {
        console.log(file);
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

            console.log(JSON.stringify(data, null, 4));
            console.log(status);

            return data;

        } catch (error) {
            if (axios.isAxiosError(error)) {
                console.log('error message: ', error.message);
                return error.message;
            } else {
                console.log('unexpected error: ', error);
                return 'An unexpected error occurred';
            }
        }
    }
}    
