import {extractMetadata, type Metadata} from "@/app/features/shared/files/metadata";

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
                "metadata",
                new Blob([JSON.stringify(metadata)], { type: "application/html" })
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

        //const url = `${this.baseURL}/${endpoint}`;
        const options: RequestInit = {method: "POST", body: this.formData};

        const response = await fetch(url, options);
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
}

