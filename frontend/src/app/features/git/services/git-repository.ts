import type {Commit} from "@/app/features/git/types.ts";

class GitRepository {
    private commitHistory: Commit[];
    private name: string;
    private size: number;

    constructor(name: string, size: number, commits: Commit[]) {
    }
}

export default GitRepository