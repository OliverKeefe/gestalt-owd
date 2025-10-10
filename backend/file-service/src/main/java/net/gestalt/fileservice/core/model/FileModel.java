package net.gestalt.fileservice.core.model;

import java.time.Instant;
import java.util.ArrayList;
import java.util.LinkedList;
import java.util.UUID;

public class FileModel {

    public record File (
            UUID id,
            UUID orgId,
            UUID ownerId,
            String storageType,
            String storageKey,
            Long fileSize,
            String checksum,
            Instant createdAt,
            Instant lastModifiedAt
    ) {}

    private final File file;

    private ArrayList<File> versions;

    public FileModel(UUID id, UUID orgId, UUID ownerId, String storageType, String storageKey,
                     Long fileSize, String checksum, Instant createdAt, Instant lastModifiedAt) {

        this.file = new File(id, orgId, ownerId, storageType, storageKey,
                             fileSize, checksum, createdAt, lastModifiedAt);
    }

    public File getFileInfo() {
        return file;
    }

    public ArrayList<File> getFullVersionHistory() {
        return this.versions;
    }

    public void setVersionHistory(File newVersion) {
        this.versions.add(newVersion);
    }
}




