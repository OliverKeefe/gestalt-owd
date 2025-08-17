package net.gestalt.fileservice.bootstrap.mapper;

import net.gestalt.fileservice.adapter.api.dto.request.FileRequest;
import net.gestalt.fileservice.core.model.FileModel;

public class FileMapper {

    public FileMapper() {}

    public FileModel toModel(FileRequest fileRequest) {
        FileModel fileModel = new FileModel(
                fileRequest.id(),
                fileRequest.orgId(),
                fileRequest.ownerId(),
                fileRequest.storageType(),
                fileRequest.storageKey(),
                fileRequest.fileSize(),
                fileRequest.checksum(),
                fileRequest.createdAt(),
                fileRequest.lastModifiedAt()
        );

        return fileModel;
    }

    public FileModel toModel(FileEntity fileEntity) {}
}
