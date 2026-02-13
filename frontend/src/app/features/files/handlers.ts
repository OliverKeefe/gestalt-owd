import {RestHandler} from "@/app/features/shared/api/rest/rest-handler.ts";
import type { Metadata } from "@/app/features/home/hooks/data";

export interface CursorReq {
    modified_at: string | null;
    uuid: string | null;
}

export interface GetAllMetadataReq {
    user_id: string | null;
    cursor: CursorReq;
    limit: number;
}

interface GetAllMetadataResp {
    status: string;
    metadata: Metadata[];
}

const api = new RestHandler(`http://127.0.0.1:8081`);

export async function getAllMetadata(request: GetAllMetadataReq): Promise<GetAllMetadataResp> {
        const resp: any = await api.handlePost<GetAllMetadataReq, any>(`api/files/get-all`, request);
        console.log("RAW RESPONSE:", resp);
        return {
            status: resp.status || resp.Status || "fetched",
            metadata: Array.isArray(resp.data) ? resp.data : (Array.isArray(resp.Data) ? resp.Data : []),
        };
}

export async function uploadFiles(files: File[]) {
    const form = new UploadForm(files);
    await form.prepare();
    return await form.send();
}
