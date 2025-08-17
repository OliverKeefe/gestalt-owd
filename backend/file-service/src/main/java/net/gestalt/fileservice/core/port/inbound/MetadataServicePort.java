package net.gestalt.fileservice.core.port.inbound;

import net.gestalt.fileservice.adapter.api.dto.response.FileResponse;
import net.gestalt.fileservice.core.model.FileModel;

public interface MetadataServicePort {
    FileResponse getFile(FileModel fileModel);
}
