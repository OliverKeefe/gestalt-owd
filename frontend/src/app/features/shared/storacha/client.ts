import * as Client from "@storacha/client";
import cred from "./cred.json";
import type { StorachaUserInfo } from "@/app/features/shared/storacha/types.ts";

export const StorachaClient = await Client.create();

export async function StorachaLogin(u: unknown): Promise<void> {
    if (!isUser(u)) return;
    await StorachaClient.login(u.account);
    await StorachaClient.setCurrentSpace(u.did);
    return;
}

function isUser(u: StorachaUserInfo): u is StorachaUserInfo {
    return (
        typeof u?.did === "string" &&
        u.did.startsWith("did:key") &&
        typeof u?.account === "string" &&
        u.account.includes("@")
    );
}

async function getAccount(claim: any): Promise<string> {
    return "";
}

async function getDID(): Promise<string> {
    return "";
}