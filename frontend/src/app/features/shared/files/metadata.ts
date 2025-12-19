import MediaInfo from "mediainfo.js";

let mediaInfoPromise: Promise<any> | null = null;

/**
* getMediaInfo initializes MediaInfo.js module
* locateFile gets the MediaInfo precompiled wasm binary from public/mediainfo.
* @returns mediaInfoPromise - Promise containing MediaInfo format and binary path.
* */
function getMediaInfo() {
    if (!mediaInfoPromise) {
        mediaInfoPromise = MediaInfo({
            format: "JSON",
            locateFile: (file) => `/mediainfo/${file}`,
        });
    }
    return mediaInfoPromise;
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
export async function extractMetadata(file: File) {
    console.log("extractMetadata called", file);
    const mediaInfo = await getMediaInfo();
     const getSize = () => file.size;

     const readChunk = (chunkSize: number, offset: number) =>
         file.slice(offset, offset + chunkSize).arrayBuffer();
     const result = await mediaInfo.analyzeData(getSize, readChunk);
     return JSON.parse(result);
}