import {extractMetadata, type Metadata} from "@/app/features/shared/files/metadata";
import {useAuthStore} from "@/security/auth/authstore/auth-store.ts";

type PayloadItem = {
    id: string
    metadata: Metadata
    file: File
}

// type Payload = Record<string, PayloadItem>

export class UploadForm {
    private readonly files: File[];
    private payload: PayloadItem[];
    private readonly formData: FormData;

    constructor(files: File[]) {
        this.files = files;
        this.formData = new FormData();
        this.payload = [];
    }

    public async prepare() {
        for (let i = 0; i < this.files.length; i++) {
            const metadata = await extractMetadata(this.files[i]);
            const file = this.files[i];
            const id = metadata.id;
            this.payload.push({ id, metadata, file })
        }

    }

    private buildFormData() {
        Object.values(this.payload).forEach(({ metadata, file }) => {
            this.formData.append(
                `metadata-${metadata.id}`,
                JSON.stringify(metadata)
            )

            this.formData.append(
                `file-${metadata.id}`,
                file,
                file.name
            )
        })
    }

    public async send(): Promise<Response> {
        this.buildFormData();
        const url = "http://localhost:8081/api/files/upload"



        const token = useAuthStore.getState().token;
        console.log(token);

        //const url = `${this.baseURL}/${endpoint}`;
        const options: RequestInit = {
            headers: {
                Authorization: `Bearer ${token}`,
            },
            method: "POST",
            body: this.formData};

        const response = await fetch(url, options);
        const contentType = response.headers.get("content-type");

        if (contentType && contentType.includes("application/json")) {
            return await response.json();
        }

        // Might be better off returning null here, need to rethink.
        await this.handleFailedUpload(response);
        return await response.json();
    }

    private async handleFailedUpload(response: Response): Promise<void> {
        if (!response.ok) {
            const errorText = await response.text();
            console.error(`Upload failed: ${response.status} ${response.statusText}`, errorText);
            throw new Error(`Upload failed with status ${response.status}`);
        }
    }

    //private async uploadIPFS(): Promise<Response> {
    //    this.formData.forEach()
    //}
}

