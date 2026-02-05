import {ObjectType} from "@/app/features/shared/storacha/types.ts";
import type {Metadata} from "@/app/features/shared/files/metadata.ts";
import {StorachaClient, StorachaLogin} from "@/app/features/shared/storacha/client.ts";
import cred from "@/app/features/shared/storacha/cred.json";


export async function uploadObject(
    //spaceDID: string,
    //delegation,
    filedata: Uint8Array[],
    metadata: Metadata[],
    objectType: ObjectType
): Promise<string | null> {


    //const account = await StorachaClient.login("test@test.galacy");
    //const space = await StorachaClient.createSpace("test-space", {account})
    //await StorachaClient.setCurrentSpace(`did:key:${did}`);
    await StorachaLogin(cred).then();

    switch (objectType) {
        case ObjectType.FILE: {
            const file = makeDataBlobLike(filedata[0], metadata[0]);
            try {
                const cid = await StorachaClient.uploadFile(file);

                return cid.toString();
            } catch (err) {
                console.log("unable to upload to storacha network: ", err);
                return null;
            }
        }

        //case ObjectType.DIRECTORY: {}

        //case ObjectType.CAR_FILE_SHARDS: {}
    }
    return null;
}

function makeDataBlobLike(filedata: Uint8Array, metadata: Metadata): File {
    try {
        return new File([filedata], metadata.path);
    } catch (err) {
        console.log("unable to make file blob like: ", err)
        return new File(0, null); // fix this abomination.
    }
}