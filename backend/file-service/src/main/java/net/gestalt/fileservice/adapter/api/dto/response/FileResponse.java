package net.gestalt.fileservice.adapter.api.dto.response;

import java.time.Instant;
import java.util.UUID;

public record FileResponse(
        UUID id,
        UUID orgId,
        UUID ownerId,
        String storageType,
        String storageKey,
        Long fileSize,
        String checksum,
        Instant createdAt,
        Instant lastModifiedAt
) {

}
