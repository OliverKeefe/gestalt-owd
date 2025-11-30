
/*
* Handles client-side REST requests and responses (POST, GET, PUT, CREATE, DELETE).
* */
export class RestHandler {
    private readonly baseURL: string;

    /**
     * Constructor.
     * @param baseURL The base URL for REST api calls to the Gestalto Control Plane Backend.
     * */
    constructor(baseURL: string) {
        this.baseURL = baseURL;
    }

    private async handleFailedRequest(response: Response): Promise<void> {
        if (!response.ok) {
            const errorText = await response.text();
            console.error(`Request failed: ${response.status} ${response.statusText}`, errorText);
            throw new Error(`Request failed with status ${response.status}`);
        }
    }

    public async handleGet<R = unknown>(endpoint: string): Promise<R> {
        const url = `${this.baseURL}/${endpoint}`;
        const options: RequestInit = { method: "GET" };
        const response = await fetch(url, options);
        await this.handleFailedRequest(response);
        return await response.json();
    }

    public async handlePost<T, R = unknown>(endpoint: string, payload: T): Promise<R> {
        const url = `${this.baseURL}/${endpoint}`;
        const options: RequestInit = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify(payload)
        };

        const response = await fetch(url, options);
        await this.handleFailedRequest(response);
        return await response.json();
    }

}