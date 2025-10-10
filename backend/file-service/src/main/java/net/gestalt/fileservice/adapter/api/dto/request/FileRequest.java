package net.gestalt.fileservice.adapter.api.dto.request;

import java.time.Instant;
import java.util.UUID;

public record FileRequest (
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
