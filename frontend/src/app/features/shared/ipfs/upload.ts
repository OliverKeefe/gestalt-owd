import {type FileData, ObjectType} from "@/app/features/shared/ipfs/types.ts";
import type {Metadata} from "@/app/features/shared/files/metadata.ts";
import {StorachaClient, StorachaLogin} from "@/app/features/shared/ipfs/client.ts";
import cred from "@/app/features/shared/ipfs/cred.json";

// Needs interfaces etc... duh...
export async function uploadObject(
    //spaceDID: string,
    //delegation,
    id: string,
    fileData: File,
    metadata: Metadata,
    objectType: ObjectType
): Promise<Map<string, string[]> | null> {


    //const account = await StorachaClient.login("test@test.galacy");
    //const space = await StorachaClient.createSpace("test-space", {account})
    //await StorachaClient.setCurrentSpace(`did:key:${did}`);
    await StorachaLogin(cred).then();

    const cids = new Map<string, string[]>();

        switch (objectType) {
            case ObjectType.FILE: {
                const file = makeDataBlobLike(fileData, metadata);
                try {
                    const cid = await StorachaClient.uploadFile(file);
                    cids.set(id, [cid.toString()]);
                } catch (err) {
                    console.log("unable to upload to ipfs network: ", err);
                    return null;
                }
            }
        }
        //case ObjectType.DIRECTORY: {}

        //case ObjectType.CAR_FILE_SHARDS: {}
    return cids;
    }



function makeDataBlobLike(fileData: File, metadata: Metadata): File {
    try {
        return new File([fileData], metadata.path);
    } catch (err) {
        console.log("unable to make file blob like: ", err)
        return new File([], "invalid");
    }
}