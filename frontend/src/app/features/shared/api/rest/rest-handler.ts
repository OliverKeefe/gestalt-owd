
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

}