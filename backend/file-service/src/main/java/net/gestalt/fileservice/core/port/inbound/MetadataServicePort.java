package net.gestalt.fileservice.core.port.inbound;

import net.gestalt.fileservice.adapter.api.dto.response.FileResponse;
import net.gestalt.fileservice.core.model.FileModel;

public interface MetadataServicePort {
    boolean createFileInfo(FileModel fileModel);
    FileModel getFileInfo(FileModel fileModel);
    boolean updateFileInfo(FileModel fileModel);
    boolean deleteFileInfo(FileModel fileModel);
}
