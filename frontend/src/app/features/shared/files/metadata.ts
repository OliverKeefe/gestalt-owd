import MediaInfo from "mediainfo.js";


/**
 * Metadata interface maps extracted metadata to object.
 * Fields `path`, `relativePath`, `lastModified`, `lastModifiedDate`
 * `size` and `fileType` are set using MediaInfo.js metadata
 * extraction. The rest, `id`, `ownerId` and `checkSum` are
 * injected once extraction is complete. `id` serves as a key to
 * associate each metadata object with its respective ByteData.
 * */
export type Metadata = ExtractedMetadata& {
    id: string
    ownerId: string
    checkSum: string
}

type ExtractedMetadata = {
    path: string
    relativePath: string
    lastModified: number
    lastModifiedDate: string
    size: number
    fileType: string
}

/**
 * <p>
 * extractMetadata uses MediaInfo to extract metadata from a file uploaded
 * via an UploadDialog's Dropzone component.
 * </p>
 * <p>
 * Metadata is extracted client-side and sent to backend via REST api due to
 * metadata loss that occurs with multipart form uploads.
 * </p>
 *
 * @param file - a File, uploaded via UploadDialog's Dropzone component.
 * @returns result - extracted metadata parsed to JSON.
 * */
export async function extractMetadata(file: File): Promise<Metadata> {
    const mediaInfo = await MediaInfo({
        format: "JSON",
        locateFile: (path) => `/mediainfo/dist/${path}`,
    })

    const result = await mediaInfo.analyzeData(
        () => file.size,
        async (chunkSize, offset) => {
            const buffer = await file.slice(offset, offset + chunkSize).arrayBuffer()
            return new Uint8Array(buffer)
        }
    )

    mediaInfo.close()

    if (!result) {
        throw new Error("MediaInfo returned void")
    }

    const extracted = JSON.parse(result) as ExtractedMetadata

    return {
        ...extracted,
        id: crypto.randomUUID(),
        ownerId: crypto.randomUUID(),
        checkSum: "a".repeat(256),
    }
}